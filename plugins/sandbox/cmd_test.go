// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	plugin "unikraft.com/cloud/plugins/sandbox"
	"unikraft.com/cloud/sdk/platform"
)

// mockPlugin is a sandbox plugin that serves one command: enough of the API
// for a Cmd to run against, and a record of what it was asked.
type mockPlugin struct {
	mu sync.Mutex

	stdout, stderr []byte
	more           chan struct{} // closed and replaced whenever there is more output

	exited   chan struct{}
	exitCode int32
	dead     bool

	run      plugin.RunCommandRequest
	stdin    []byte
	stdinEOF bool
	resizes  [][2]uint16

	// refuse and drop are set before the plugin serves anything, so they need
	// no lock of their own: refuse answers the socket with that status, and drop
	// closes it as the plugin does a command forgotten under it.
	refuse int
	drop   string

	signals []int
	forgot  bool
}

func newFakePlugin() *mockPlugin {
	return &mockPlugin{more: make(chan struct{}), exited: make(chan struct{})}
}

func (f *mockPlugin) write(stdout, stderr string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.stdout = append(f.stdout, stdout...)
	f.stderr = append(f.stderr, stderr...)
	close(f.more)
	f.more = make(chan struct{})
}

// exit ends the command once; a later call, as from a SIGKILL after the
// command already died, changes nothing.
func (f *mockPlugin) exit(code int32) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.dead {
		return
	}
	f.dead = true
	f.exitCode = code
	close(f.exited)
}

// started is the request the command was run from.
func (f *mockPlugin) started() plugin.RunCommandRequest {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.run
}

func (f *mockPlugin) forgotten() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.forgot
}

func (f *mockPlugin) sentSignals() []int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]int(nil), f.signals...)
}

func (f *mockPlugin) stdinSeen() (string, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return string(f.stdin), f.stdinEOF
}

func (f *mockPlugin) resized() [][2]uint16 {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([][2]uint16(nil), f.resizes...)
}

func (f *mockPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const prefix = "/v1/instances/inst-1/plugins/sandbox"
	path := strings.TrimPrefix(r.URL.Path, prefix)

	reply := func(data any) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": data})
	}

	switch {
	case path == "/commands" && r.Method == http.MethodPost:
		f.mu.Lock()
		_ = json.NewDecoder(r.Body).Decode(&f.run)
		f.mu.Unlock()
		reply(plugin.RunCommandData{Uuid: "cmd-1"})

	case path == "/commands/cmd-1" && r.Method == http.MethodDelete:
		f.mu.Lock()
		f.forgot = true
		f.mu.Unlock()
		reply(nil)

	case path == "/commands/cmd-1/ws" && f.refuse != 0:
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(f.refuse)
		fmt.Fprint(w, "<html>\n<head><title>Bad Gateway</title></head>\n<body>openresty</body>\n</html>")

	case path == "/commands/cmd-1/ws":
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer func() { _ = conn.CloseNow() }()
		if f.drop != "" {
			_ = conn.Close(websocket.StatusPolicyViolation, f.drop)
			return
		}

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()
		go f.take(ctx, cancel, conn)
		f.follow(ctx, conn)

	case path == "/commands/cmd-1/signal" && r.Method == http.MethodPost:
		var req struct {
			Signal int `json:"signal"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)

		f.mu.Lock()
		f.signals = append(f.signals, req.Signal)
		f.mu.Unlock()
		// SIGKILL cannot be ignored: the command dies of it, as a real one does.
		if req.Signal == int(sigKILL) {
			f.exit(137)
		}
		reply(nil)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// follow sends the output as it comes, a frame per stream, then the exit.
func (f *mockPlugin) follow(ctx context.Context, conn *websocket.Conn) {
	var sentOut, sentErr int
	for {
		f.mu.Lock()
		out, errOut := f.stdout[sentOut:], f.stderr[sentErr:]
		sentOut, sentErr = len(f.stdout), len(f.stderr)
		more, dead, code := f.more, f.dead, f.exitCode
		f.mu.Unlock()

		for _, frame := range [][]byte{append([]byte{1}, out...), append([]byte{2}, errOut...)} {
			if len(frame) > 1 && conn.Write(ctx, websocket.MessageBinary, frame) != nil {
				return
			}
		}
		if dead {
			_ = conn.Write(ctx, websocket.MessageText, fmt.Appendf(nil, `{"event":"exit","exitcode":%d}`, code))
			_ = conn.Close(websocket.StatusNormalClosure, "")
			return
		}

		select {
		case <-more:
		case <-f.exited:
		case <-ctx.Done():
			return
		}
	}
}

// take records the input and control messages the client sends.
func (f *mockPlugin) take(ctx context.Context, cancel context.CancelFunc, conn *websocket.Conn) {
	defer cancel()
	for {
		typ, msg, err := conn.Read(ctx)
		if err != nil {
			return
		}

		var control struct {
			Type       string
			Cols, Rows uint16
		}
		f.mu.Lock()
		switch {
		case typ == websocket.MessageBinary:
			f.stdin = append(f.stdin, msg...)
		case json.Unmarshal(msg, &control) != nil:
		case control.Type == "eof":
			f.stdinEOF = true
		case control.Type == "resize":
			f.resizes = append(f.resizes, [2]uint16{control.Cols, control.Rows})
		}
		f.mu.Unlock()
	}
}

// newTarget serves fake over a test server and returns a Target addressing
// it.
func newTarget(t *testing.T, fake *mockPlugin) Target {
	t.Helper()
	return newTargetServing(t, fake)
}

func newTargetServing(t *testing.T, h http.Handler) Target {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)

	return Target{
		Client:   plugin.NewClient(),
		Instance: platform.Instance{Uuid: "inst-1"},
		Plugin:   "sandbox",
		Opts: []plugin.Option{
			plugin.WithEndpoint(srv.URL),
			plugin.WithPluginName("sandbox"),
			plugin.WithToken("token"),
			plugin.WithHTTPClient(srv.Client()),
		},
	}
}

// TestCmdRun pins the whole of a successful run: the command line the plugin
// is asked to run and the output both streams deliver.
func TestCmdRun(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.write("hello\n", "warning\n")
	fake.exit(0)

	var stdout, stderr bytes.Buffer
	cmd := target.Command(t.Context(), "echo", "hello")
	cmd.Dir = "/srv"
	cmd.Env = map[string]string{"DEBUG": "true"}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	require.NoError(t, cmd.Run())

	assert.Equal(t, "hello\n", stdout.String())
	assert.Equal(t, "warning\n", stderr.String())
	assert.Equal(t, "cmd-1", cmd.UUID)

	fake.mu.Lock()
	defer fake.mu.Unlock()
	assert.Equal(t, plugin.CommandLineArgs{"echo", "hello"}, fake.run.Cmd,
		"the words go as words, with no quoting for a shell to undo")
	require.NotNil(t, fake.run.Cwd)
	assert.Equal(t, "/srv", *fake.run.Cwd)
	require.NotNil(t, fake.run.Env)
	assert.Equal(t, map[string]string{"DEBUG": "true"}, *fake.run.Env)
}

// TestCmdRunCommandLine pins that a whole shell line reaches the plugin as it
// was written, with nothing quoted on its behalf.
func TestCmdRunCommandLine(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.write("a b\n", "")
	fake.exit(0)

	var stdout bytes.Buffer
	cmd := target.CommandLine(t.Context(), "echo a b > /dev/stderr; echo a b")
	cmd.Stdout = &stdout

	require.NoError(t, cmd.Run())
	assert.Equal(t, "a b\n", stdout.String())

	fake.mu.Lock()
	defer fake.mu.Unlock()
	assert.Equal(t, plugin.CommandLineShell("echo a b > /dev/stderr; echo a b"), fake.run.Cmd,
		"a line the user typed is still a line for the instance's shell to read")
}

// TestCmdCommandForms pins the two forms a command reaches the plugin in:
// words it runs without a shell, and a line its shell reads.
func TestCmdCommandForms(t *testing.T) {
	for _, tt := range []struct {
		name    string
		cmd     Cmd
		want    plugin.CommandLineUnion
		wantErr string
	}{
		{name: "args", cmd: Cmd{Args: []string{"echo", "a b"}}, want: plugin.CommandLineArgs{"echo", "a b"}},
		{name: "cmdline", cmd: Cmd{Cmdline: "echo a b"}, want: plugin.CommandLineShell("echo a b")},
		{name: "neither", cmd: Cmd{}, wantErr: "no command given"},
		{name: "both", cmd: Cmd{Args: []string{"echo"}, Cmdline: "echo"}, wantErr: "only one of Args and Cmdline"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cmd.commandLine()
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestCmdStreamsStayApart pins that the two streams are never folded into
// one: a nil writer discards its stream rather than sending it to the other.
func TestCmdStreamsStayApart(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.write("out", "err")
	fake.exit(0)

	var stdout, stderr bytes.Buffer
	cmd := target.Command(t.Context(), "sh", "-c", "...")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	require.NoError(t, cmd.Run())
	assert.Equal(t, "out", stdout.String())
	assert.Equal(t, "err", stderr.String())
}

// TestCmdNilStderrDiscards pins that a caller that set only Stdout sees only
// the command's standard output, as os/exec does.
func TestCmdNilStderrDiscards(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.write("out", "err")
	fake.exit(0)

	var stdout bytes.Buffer
	cmd := target.Command(t.Context(), "sh", "-c", "...")
	cmd.Stdout = &stdout

	require.NoError(t, cmd.Run())
	assert.Equal(t, "out", stdout.String())
}

// TestCmdStdin pins that a reader is forwarded in full and that the command's
// standard input is closed after it.
func TestCmdStdin(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	cmd := target.Command(t.Context(), "cat")
	cmd.Stdin = strings.NewReader("one\ntwo\n")
	cmd.Stdout = io.Discard

	require.NoError(t, cmd.Start())

	// The command only ends once its input has been read to the end, as cat
	// does.
	require.Eventually(t, func() bool {
		_, eof := fake.stdinSeen()
		return eof
	}, 5*time.Second, 10*time.Millisecond)
	fake.exit(0)

	require.NoError(t, cmd.Wait())

	seen, eof := fake.stdinSeen()
	assert.Equal(t, "one\ntwo\n", seen)
	assert.True(t, eof)
}

// TestCmdStdinChunked pins that a reader larger than one chunk still arrives
// whole.
func TestCmdStdinChunked(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	want := strings.Repeat("x", stdinChunkSize*2+7)

	cmd := target.Command(t.Context(), "cat")
	cmd.Stdin = strings.NewReader(want)
	cmd.Stdout = io.Discard

	require.NoError(t, cmd.Start())
	require.Eventually(t, func() bool {
		_, eof := fake.stdinSeen()
		return eof
	}, 5*time.Second, 10*time.Millisecond)
	fake.exit(0)
	require.NoError(t, cmd.Wait())

	seen, _ := fake.stdinSeen()
	assert.Equal(t, want, seen)
}

// TestCmdExitStatus pins that a command that exits non-zero is reported as
// such, with the status it exited with and everything it wrote.
func TestCmdExitStatus(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.write("", "no such file\n")
	fake.exit(2)

	var stdout, stderr bytes.Buffer
	cmd := target.Command(t.Context(), "ls", "/nope")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	var exit *ExitError
	err := cmd.Run()
	require.ErrorAs(t, err, &exit)
	assert.Equal(t, 2, exit.Code)
	assert.Equal(t, 2, exit.ExitCode())
	assert.Equal(t, "cmd-1", exit.UUID)
	assert.Equal(t, 2, cmd.ExitCode)
	assert.Equal(t, "no such file\n", stderr.String())
}

// TestCmdExitStatusZero pins that a command that exits zero is no error, and
// that the status is readable all the same.
func TestCmdExitStatusZero(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)
	fake.exit(0)

	cmd := target.Command(t.Context(), "true")
	cmd.Stdout = io.Discard

	require.NoError(t, cmd.Run())
	assert.Equal(t, 0, cmd.ExitCode)
}

// TestCmdCancelWithinWaitDelay pins that a command that dies of the interrupt
// within the delay is reported as having ended, with its output and the
// status it died with rather than the cancellation.
func TestCmdCancelWithinWaitDelay(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	ctx, cancel := context.WithCancel(t.Context())
	var stdout bytes.Buffer
	cmd := target.Command(ctx, "sleep", "300")
	cmd.Stdout = &stdout
	cmd.WaitDelay = 10 * time.Second
	cmd.Cancel = func() error {
		// A command that takes the interrupt: it writes a last line and dies
		// of it, as one killed by SIGINT does.
		fake.write("interrupted\n", "")
		fake.exit(130)
		return nil
	}

	require.NoError(t, cmd.Start())
	cancel()

	var exit *ExitError
	require.ErrorAs(t, cmd.Wait(), &exit)
	assert.Equal(t, 130, exit.Code)
	assert.Equal(t, "interrupted\n", stdout.String())
}

// TestCmdCancelWaitsByDefault pins that an unset WaitDelay waits for the
// interrupted command however long it takes, as the zero value of the field
// of the same name of os/exec.Cmd does.
func TestCmdCancelWaitsByDefault(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	ctx, cancel := context.WithCancel(t.Context())
	var stdout bytes.Buffer
	cmd := target.Command(ctx, "sleep", "300")
	cmd.Stdout = &stdout
	cmd.Cancel = func() error {
		fake.write("interrupted\n", "")
		fake.exit(130)
		return nil
	}

	require.NoError(t, cmd.Start())
	cancel()

	var exit *ExitError
	require.ErrorAs(t, cmd.Wait(), &exit)
	assert.Equal(t, 130, exit.Code)
	assert.Equal(t, "interrupted\n", stdout.String())
}

// TestCmdCancelWaitDelayExpires pins that a command that ignores the interrupt
// is killed once a delay that was asked for is out, is reported as
// interrupted, and leaves no record on the instance.
func TestCmdCancelWaitDelayExpires(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	ctx, cancel := context.WithCancel(t.Context())
	cmd := target.Command(ctx, "sleep", "300")
	cmd.Stdout = io.Discard
	cmd.WaitDelay = 50 * time.Millisecond

	require.NoError(t, cmd.Start())
	cancel()

	err := cmd.Wait()
	require.ErrorIs(t, err, context.Canceled)
	require.ErrorContains(t, err, "cmd-1")
	assert.Equal(t, []int{int(sigINT), int(sigKILL)}, fake.sentSignals())
	assert.True(t, fake.forgotten(), "the killed command's record was dropped")
}

// TestCmdCancelError pins that an error from Cancel ends the wait with it.
func TestCmdCancelError(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	ctx, cancel := context.WithCancel(t.Context())
	cmd := target.Command(ctx, "sleep", "300")
	cmd.Stdout = io.Discard

	sentinel := errors.New("could not interrupt")
	cmd.Cancel = func() error { return sentinel }

	require.NoError(t, cmd.Start())
	cancel()

	assert.ErrorIs(t, cmd.Wait(), sentinel)
}

// TestCmdOutputStreamedWhileRunning pins that output is delivered as the
// command produces it, not only once it has ended.
func TestCmdOutputStreamedWhileRunning(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	seen := make(chan struct{})
	cmd := target.Command(t.Context(), "sh", "-c", "...")
	cmd.Stdout = writerFunc(func(p []byte) (int, error) {
		if strings.Contains(string(p), "early") {
			select {
			case <-seen:
			default:
				close(seen)
			}
		}
		return len(p), nil
	})

	require.NoError(t, cmd.Start())
	fake.write("early\n", "")

	select {
	case <-seen:
	case <-time.After(5 * time.Second):
		require.FailNow(t, "output was not delivered while the command was still running")
	}

	fake.exit(0)
	require.NoError(t, cmd.Wait())
}

// TestCmdOutputPanicEndsTheCommand pins that a writer that panics ends the
// command with an error, and takes nothing else down with it.
func TestCmdOutputPanicEndsTheCommand(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	cmd := target.Command(t.Context(), "sh", "-c", "...")
	cmd.Stdout = writerFunc(func(p []byte) (int, error) { panic("lone ESC") })

	require.NoError(t, cmd.Start())
	fake.write("\x1b", "")

	err := cmd.Wait()
	require.ErrorContains(t, err, "panicked")
	assert.ErrorContains(t, err, "cmd-1")
}

// TestCmdMisuse pins what a Cmd refuses: no command, a second Start, a Wait
// without a Start, and a second Wait.
func TestCmdMisuse(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)
	fake.exit(0)

	t.Run("no-command", func(t *testing.T) {
		cmd := target.CommandArgs(t.Context(), nil)
		require.Error(t, cmd.Err)
		assert.ErrorContains(t, cmd.Run(), "no command given")
	})

	t.Run("no-command-line", func(t *testing.T) {
		cmd := target.CommandLine(t.Context(), "")
		require.Error(t, cmd.Err)
		assert.ErrorContains(t, cmd.Run(), "no command given")
	})

	t.Run("wait-without-start", func(t *testing.T) {
		cmd := target.Command(t.Context(), "true")
		assert.ErrorContains(t, cmd.Wait(), "not started")
	})

	t.Run("double-start", func(t *testing.T) {
		cmd := target.Command(t.Context(), "true")
		cmd.Stdout = io.Discard
		require.NoError(t, cmd.Start())
		require.ErrorContains(t, cmd.Start(), "already started")
		require.NoError(t, cmd.Wait())
	})

	t.Run("double-wait", func(t *testing.T) {
		cmd := target.Command(t.Context(), "true")
		cmd.Stdout = io.Discard
		require.NoError(t, cmd.Run())
		assert.ErrorContains(t, cmd.Wait(), "already called")
	})
}

// TestCmdStartFailure pins that a plugin that does not report a UUID is a
// failure to start, named as such.
func TestCmdStartFailure(t *testing.T) {
	target := newTargetServing(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"success","data":{"uuid":""}}`)
	}))

	err := target.Command(t.Context(), "true").Run()
	require.ErrorContains(t, err, "did not report a command UUID")
	assert.ErrorContains(t, err, `"sandbox"`)
}

// TestACommandNobodyCanFollowIsKilled pins that a command whose socket is
// refused is killed and forgotten rather than left running unwatched, and that
// the refusal is said in one line.
func TestACommandNobodyCanFollowIsKilled(t *testing.T) {
	for _, tt := range []struct {
		status int
		want   string
	}{
		{http.StatusBadGateway, `failed to follow command: the "sandbox" plugin answered 502 Bad Gateway`},
		{http.StatusNotFound, `failed to follow command: the "sandbox" plugin cannot stream commands; update its image`},
	} {
		t.Run(http.StatusText(tt.status), func(t *testing.T) {
			fake := newFakePlugin()
			fake.refuse = tt.status
			target := newTarget(t, fake)

			cmd := target.Command(t.Context(), "ls")
			err := cmd.Start()

			require.Error(t, err)
			assert.Equal(t, tt.want, err.Error())
			assert.Equal(t, []int{int(sigKILL)}, fake.sentSignals())
			assert.True(t, fake.forgotten())
			assert.ErrorContains(t, cmd.Wait(), "not started", "there is nothing left to wait for")
		})
	}
}

// TestCmdForgottenWhileFollowing pins that a command the plugin stops
// following ends with the plugin's reason, not an exit status it never had.
func TestCmdForgottenWhileFollowing(t *testing.T) {
	fake := newFakePlugin()
	fake.drop = "command forgotten"
	target := newTarget(t, fake)

	err := target.Command(t.Context(), "sleep", "300").Run()

	require.Error(t, err)
	assert.Equal(t, "command cmd-1: command forgotten", err.Error())
}

// TestCmdTakesAWholeFrame pins that the largest frame the plugin sends is read,
// past the WebSocket library's own default limit.
func TestCmdTakesAWholeFrame(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	want := strings.Repeat("x", 64<<10)
	fake.write(want, "")
	fake.exit(0)

	var stdout bytes.Buffer
	cmd := target.Command(t.Context(), "cat", "big")
	cmd.Stdout = &stdout

	require.NoError(t, cmd.Run())
	assert.Equal(t, want, stdout.String())
}

func TestAStoppedInstanceIsSaidInOneLine(t *testing.T) {
	target := newTargetServing(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<!doctype html>\n<html lang=\"en\">\n<head><title>Service not found</title></head>\n"+
			"<body>\nThere is no service on this URL.\n</body>\n</html>")
	}))

	err := target.Command(t.Context(), "true").Run()

	require.ErrorIs(t, err, ErrNotRunning)
	assert.Equal(t, `the instance is not running, or has no "sandbox" plugin`, err.Error())
}

// TestATransportFailureNamesNoURL pins that a node that cannot be reached is
// reported without its API URL.
func TestATransportFailureNamesNoURL(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	srv.Close()

	target := Target{
		Client:   plugin.NewClient(),
		Instance: platform.Instance{Uuid: "inst-1"},
		Plugin:   "sandbox",
		Opts: []plugin.Option{
			plugin.WithEndpoint(srv.URL),
			plugin.WithPluginName("sandbox"),
			plugin.WithToken("token"),
		},
	}

	err := target.Command(t.Context(), "true").Run()

	require.Error(t, err)
	require.ErrorContains(t, err, "failed to start command")
	assert.NotContains(t, err.Error(), "http")
	assert.NotContains(t, err.Error(), srv.URL)
}

func TestAPluginFailureKeepsItsOwnAccount(t *testing.T) {
	target := newTargetServing(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"status":"error","message":"no such command"}`)
	}))

	err := target.Command(t.Context(), "true").Run()

	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNotRunning)
	require.ErrorContains(t, err, "failed to start command")
	assert.ErrorContains(t, err, "no such command")
}

// TestCmdExitCode pins that a command which ran and failed is reported as an
// error carrying the status, so that a caller asking only whether it worked is
// answered without reaching for ExitCode.
func TestCmdExitCode(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	fake.exit(3)

	cmd := target.Command(t.Context(), "false")
	cmd.Stdout = io.Discard

	var exit *ExitError
	require.ErrorAs(t, cmd.Run(), &exit)
	assert.Equal(t, 3, exit.Code)
	assert.Equal(t, "cmd-1", exit.UUID)
	assert.Equal(t, 3, exit.ExitCode())
	assert.Equal(t, 3, cmd.ExitCode)
}

// TestCmdForget pins that a command the wait saw the end of drops the record
// the instance keeps of it without the caller asking, and that asking about a
// command that never started is not a request.
func TestCmdForget(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	target.CommandArgs(t.Context(), nil).Forget(t.Context())
	assert.False(t, fake.forgotten(), "nothing to forget")

	cmd := target.Command(t.Context(), "true")
	cmd.Stdout = io.Discard
	require.NoError(t, cmd.Start())
	fake.exit(0)
	require.NoError(t, cmd.Wait())

	assert.True(t, fake.forgotten(), "the wait dropped it, so no caller has to")
}

// TestCmdForgetOnFailure pins that a command that exits non-zero is forgotten
// all the same: its record and buffered output do not outlive the wait.
func TestCmdForgetOnFailure(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	cmd := target.Command(t.Context(), "false")
	cmd.Stdout = io.Discard
	require.NoError(t, cmd.Start())
	fake.exit(1)

	var exit *ExitError
	require.ErrorAs(t, cmd.Wait(), &exit)
	assert.Equal(t, 1, exit.Code)
	assert.True(t, fake.forgotten(), "a failed command is dropped too")
}

// TestCmdCancelClosesStdin pins that an interrupted command sees the end of its
// standard input, so that one waiting to be typed at stops rather than hanging
// on a reader nobody is feeding any more.
func TestCmdCancelClosesStdin(t *testing.T) {
	fake := newFakePlugin()
	target := newTarget(t, fake)

	blocked, writer := io.Pipe()
	t.Cleanup(func() { _ = writer.Close() })

	ctx, cancel := context.WithCancel(t.Context())
	cmd := target.Command(ctx, "cat")
	cmd.Stdout = io.Discard
	cmd.Stdin = blocked
	cmd.WaitDelay = 50 * time.Millisecond

	require.NoError(t, cmd.Start())
	cancel()
	require.Error(t, cmd.Wait())

	assert.Eventually(t, func() bool {
		_, eof := fake.stdinSeen()
		return eof
	}, 5*time.Second, 10*time.Millisecond)
}

type writerFunc func(p []byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }

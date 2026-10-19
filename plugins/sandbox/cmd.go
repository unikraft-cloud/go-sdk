// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"syscall"
	"time"

	"github.com/coder/websocket"

	plugin "unikraft.com/cloud/plugins/sandbox"
	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/x/log"
)

const (
	signalTimeout = 5 * time.Second
	forgetTimeout = 10 * time.Second
)

const PluginName = plugin.PluginName

type ExitError struct {
	UUID string
	Code int
}

func (e *ExitError) Error() string {
	return fmt.Sprintf("command %s exited with status %d", e.UUID, e.Code)
}

func (e *ExitError) ExitCode() int { return e.Code }

type Target struct {
	Client   *plugin.Client
	Instance platform.Instance
	Plugin   string
	Opts     []plugin.Option
}

func (t Target) Command(ctx context.Context, name string, args ...string) *Cmd {
	return t.CommandArgs(ctx, append([]string{name}, args...))
}

func (t Target) CommandArgs(ctx context.Context, args []string) *Cmd {
	c := &Cmd{Args: args, ctx: ctx, target: t}
	if len(args) == 0 {
		c.Err = errors.New("sandbox: no command given")
	}
	return c
}

func (t Target) CommandLine(ctx context.Context, cmdline string) *Cmd {
	c := &Cmd{Cmdline: cmdline, ctx: ctx, target: t}
	if cmdline == "" {
		c.Err = errors.New("sandbox: no command given")
	}
	return c
}

// Terminal is the size of the pty a command runs on.
type Terminal struct {
	Cols, Rows uint16
}

type Cmd struct {
	Args    []string
	Cmdline string

	Dir string
	Env map[string]string

	// Terminal runs the command on a pty; nil runs it on pipes.
	Terminal *Terminal

	Stdin          io.Reader
	Stdout, Stderr io.Writer
	Cancel         func() error
	WaitDelay      time.Duration
	Err            error
	UUID           string
	ExitCode       int

	ctx    context.Context
	target Target

	waitCtx  context.Context
	stopWait context.CancelFunc

	conn *websocket.Conn

	stdinErr chan error
	stdinEOF *sync.Once

	pipeOnce *sync.Once

	done   chan error
	closed bool

	waited    bool
	forgotten bool
}

// commandLine is the command as the plugin takes it: words run without a
// shell, or a line for its shell to read.
func (c *Cmd) commandLine() (plugin.CommandLineUnion, error) {
	if len(c.Args) == 0 {
		if c.Cmdline == "" {
			return nil, errors.New("sandbox: no command given")
		}
		return plugin.CommandLineShell(c.Cmdline), nil
	}
	if c.Cmdline != "" {
		return nil, errors.New("sandbox: only one of Args and Cmdline may be given")
	}
	return plugin.CommandLineArgs(c.Args), nil
}

func (c *Cmd) Start() error {
	if c.Err != nil {
		return c.Err
	}
	if c.UUID != "" {
		return errors.New("sandbox: command already started")
	}
	if errNoFollow != nil {
		return errNoFollow
	}

	log.G(c.ctx).Trace().Msg("executing command")

	cmdline, err := c.commandLine()
	if err != nil {
		return err
	}

	req := plugin.RunCommandRequest{Cmd: cmdline}
	if c.Dir != "" {
		req.Cwd = &c.Dir
	}
	if len(c.Env) > 0 {
		env := c.Env
		req.Env = &env
	}
	if c.Terminal != nil {
		pty := true
		req.Pty = &pty
		// Zero is refused; left out, the plugin's own default stands.
		if c.Terminal.Cols > 0 {
			req.Cols = &c.Terminal.Cols
		}
		if c.Terminal.Rows > 0 {
			req.Rows = &c.Terminal.Rows
		}
	}

	resp, err := c.target.Client.RunCommand(c.ctx, c.target.Instance, &req, c.target.Opts...)
	if err != nil {
		return c.target.apiError("failed to start command", err)
	}
	if resp.Data == nil || resp.Data.Uuid == "" {
		return fmt.Errorf("failed to start command: the %q plugin did not report a command UUID", c.target.Plugin)
	}
	c.UUID = resp.Data.Uuid

	c.waitCtx, c.stopWait = context.WithCancel(context.WithoutCancel(c.ctx))

	c.conn, err = c.attach()
	if err != nil {
		// Nobody can follow it, so it is not left running for nobody.
		c.kill()
		c.Forget(c.ctx)
		c.stopWait()
		c.UUID = ""
		return err
	}
	c.conn.SetReadLimit(wsReadLimit)

	c.stdinErr = make(chan error, 1)
	c.stdinEOF = new(sync.Once)
	c.pipeOnce = new(sync.Once)
	if c.Stdin != nil {
		go c.feed()
	}

	log.G(c.ctx).Trace().
		Str("cmd", c.UUID).
		Msg("waiting for command")

	c.done = make(chan error, 1)
	go c.follow(c.outputTo(c.Stdout), c.outputTo(c.Stderr))

	return nil
}

func (c *Cmd) Wait() error {
	if c.UUID == "" {
		return errors.New("sandbox: command not started")
	}
	if c.waited {
		return errors.New("sandbox: Wait was already called")
	}
	c.waited = true

	defer c.stop()

	cancelled := c.ctx.Done()
	var expired <-chan time.Time

	for {
		select {
		case <-cancelled:
			cancelled = nil
			log.G(c.ctx).Trace().Str("cmd", c.UUID).Msg("context cancelled, interrupting command")

			if err := c.cancel(); err != nil {
				return err
			}
			if c.WaitDelay > 0 {
				delay := time.NewTimer(c.WaitDelay)
				defer delay.Stop()
				expired = delay.C
			}

		case <-expired:
			c.kill()
			c.awaitEnd(signalTimeout)
			c.stop()
			c.Forget(c.ctx)
			return c.killed()

		case stdinErr := <-c.stdinErr:
			log.G(c.ctx).Warn().Err(stdinErr).Str("cmd", c.UUID).Msg("standard input failed")
			continue

		case err := <-c.done:
			c.closed = true
			// However it ended, the instance need not keep its record.
			c.Forget(c.ctx)
			if _, exited := errors.AsType[*ExitError](err); err == nil || exited {
				return err
			}
			if c.ctx.Err() != nil {
				return c.ctx.Err()
			}
			return err
		}
	}
}

// awaitEnd gives the stream up to d to report the command's end.
func (c *Cmd) awaitEnd(d time.Duration) {
	if c.closed {
		return
	}
	limit := time.NewTimer(d)
	defer limit.Stop()
	select {
	case <-c.done:
		c.closed = true
	case <-limit.C:
	}
}

func (c *Cmd) stop() {
	c.stopWait()
	_ = c.conn.CloseNow()
	if !c.closed {
		<-c.done
		c.closed = true
	}
}

func (c *Cmd) Run() error {
	if err := c.Start(); err != nil {
		return err
	}
	return c.Wait()
}

// Forget drops the instance's record of the command. The plugin keeps one for
// every command it has run, so a caller that has read all it wants of a
// finished command says so. It is not an error worth reporting: the record is
// the instance's, and it outliving the command harms nothing here.
func (c *Cmd) Forget(ctx context.Context) {
	if c.UUID == "" || c.forgotten {
		return
	}
	c.forgotten = true

	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), forgetTimeout)
	defer cancel()

	if _, err := c.target.Client.DeleteCommandByUuid(ctx, c.target.Instance, c.UUID, c.target.Opts...); err != nil {
		log.G(ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("could not drop the command record")
	}
}

func (c *Cmd) Signal(ctx context.Context, sig syscall.Signal) error {
	if c.UUID == "" {
		return errors.New("sandbox: command not started")
	}
	req := plugin.CommandSignalRequest{Signal: plugin.CommandSignalRequestSignalNumber(sig)}
	_, err := c.target.Client.SignalCommand(ctx, c.target.Instance, c.UUID, &req, c.target.Opts...)
	return c.target.apiError("failed to signal command", err)
}

func (c *Cmd) cancel() error {
	if c.Cancel != nil {
		return c.Cancel()
	}

	signalCtx, cancel := context.WithTimeout(c.waitCtx, signalTimeout)
	defer cancel()
	if err := c.Signal(signalCtx, sigINT); err != nil {
		log.G(c.ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("failed to signal remote command")
	}
	c.closeStdin(signalCtx)
	return nil
}

// kill sends SIGKILL to a command that did not stop on the interrupt.
func (c *Cmd) kill() {
	signalCtx, cancel := context.WithTimeout(c.waitCtx, signalTimeout)
	defer cancel()
	if err := c.Signal(signalCtx, sigKILL); err != nil {
		log.G(c.ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("failed to kill remote command")
	}
}

// closeStdin tells the command its input has ended, once.
func (c *Cmd) closeStdin(ctx context.Context) {
	if c.Stdin == nil {
		return
	}
	c.stdinEOF.Do(func() {
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), signalTimeout)
		defer cancel()
		if err := c.conn.Write(ctx, websocket.MessageText, []byte(`{"type":"eof"}`)); err != nil {
			log.G(c.ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("could not close standard input")
		}
	})
}

// Resize tells the command's terminal its new size.
func (c *Cmd) Resize(cols, rows uint16) error {
	if c.conn == nil {
		return errors.New("sandbox: command not started")
	}
	// The command's own context: a cancelled write would close the socket.
	return c.conn.Write(c.waitCtx, websocket.MessageText,
		fmt.Appendf(nil, `{"type":"resize","cols":%d,"rows":%d}`, cols, rows))
}

// killed reports a command that did not stop on the interrupt, and that this
// ended by force.
func (c *Cmd) killed() error {
	return fmt.Errorf("command %s was killed: %w", c.UUID, c.ctx.Err())
}

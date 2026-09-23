// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"bytes"
	"context"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/x/shell"
	"unikraft.com/x/stdio"
)

// deafWriter is a reader that has gone: every write fails the way a closed
// pipe does.
type deafWriter struct{}

func (deafWriter) Write([]byte) (int, error) { return 0, syscall.EPIPE }

func TestATimeoutLeftUnsetIsTheDefault(t *testing.T) {
	asked := Transport{Timeouts: Timeouts{InterruptGrace: time.Second, Reap: 2 * time.Second}}

	assert.Equal(t, time.Second, asked.waitDelay(false))
	assert.Equal(t, 2*time.Second, asked.waitDelay(true))

	none := Transport{}

	assert.Equal(t, interruptGrace, none.waitDelay(false),
		"zero would wait forever, which is not what a caller leaving it out means")
	assert.Equal(t, reapTimeout, none.waitDelay(true))
}

func TestAFailedCommandIsAStatusNotAnError(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{Target: newTargetServing(t, fake)}
	fake.write("said something", "")
	fake.exit(3)

	var out bytes.Buffer
	code, err := transport.Shell()(t.Context(), shell.Command{
		Args:    []string{"false"},
		Streams: stdio.Stdio{Stdout: &out},
	})

	require.NoError(t, err, "the command ran, so its status is the answer")
	assert.Equal(t, 3, code)
	assert.Equal(t, "said something", out.String())
}

func TestACommandThatRanReportsNoStatus(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{Target: newTargetServing(t, fake)}
	fake.exit(0)

	code, err := transport.Shell()(t.Context(), shell.Command{Args: []string{"true"}})

	require.NoError(t, err)
	assert.Zero(t, code)
}

func TestTheEnvironmentReachesTheInstance(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{Target: newTargetServing(t, fake)}
	fake.exit(0)

	_, err := transport.Shell()(t.Context(), shell.Command{
		Args: []string{"env"},
		Dir:  "/var/log",
		Env:  []string{"HOME=/", "MULTI=a=b"},
	})

	require.NoError(t, err)
	ran := fake.started()
	require.NotNil(t, ran.Env)
	assert.Equal(t, "/", (*ran.Env)["HOME"])
	assert.Equal(t, "a=b", (*ran.Env)["MULTI"], "only the first = separates a name from its value")
	require.NotNil(t, ran.Cwd)
	assert.Equal(t, "/var/log", *ran.Cwd)
}

func TestAnUninterruptibleCommandGivesThePromptBack(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{
		Target:   newTargetServing(t, fake),
		Timeouts: Timeouts{InterruptGrace: 200 * time.Millisecond},
	}

	ctx, cancel := context.WithCancel(t.Context())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	var said bytes.Buffer
	code, err := transport.Shell()(ctx, shell.Command{
		Args:    []string{"sleep", "forever"},
		Streams: stdio.Stdio{Stderr: &said},
	})

	require.NoError(t, err, "a command that would not end is still an answer to the prompt")
	assert.Equal(t, shell.StatusInterrupted, code)
	assert.Contains(t, fake.sentSignals(), 2, "Linux SIGINT")
	assert.NotEmpty(t, said.String(), "the user is told why the prompt came back")
}

func TestAClosedReaderReachesTheInstance(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{Target: newTargetServing(t, fake)}
	fake.write(strings.Repeat("x", 4096), "")
	fake.exit(0)

	code, err := transport.Shell()(t.Context(), shell.Command{
		Args:    []string{"cat", "big"},
		Streams: stdio.Stdio{Stdout: deafWriter{}},
	})

	require.NoError(t, err)
	assert.Zero(t, code)
	assert.Contains(t, fake.sentSignals(), 13,
		"the command is told its reader has gone, as a local pipeline would")
}

func TestADetachedCommandIsNotWaitedFor(t *testing.T) {
	fake := newFakePlugin()
	transport := Transport{Target: newTargetServing(t, fake)}

	ctx, cancel := context.WithCancel(shell.WithDetached(t.Context()))
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := transport.Shell()(ctx, shell.Command{Args: []string{"probe"}})

	require.ErrorIs(t, err, context.Canceled, "nobody typed it and nobody waits for it")
	fake.exit(0)
	assert.Eventually(t, fake.forgotten, 10*time.Second, 20*time.Millisecond,
		"the instance keeps no record of a command nobody collected")
}

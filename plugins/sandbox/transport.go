// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"time"

	"unikraft.com/x/log"
	"unikraft.com/x/shell"
)

const (
	// reapTimeout is how long a command nobody waits for is given to die.
	reapTimeout = 30 * time.Second

	// interruptGrace is how long an interrupted command gets to report its end.
	interruptGrace = 10 * time.Second
)

// Transport runs the shell's commands on the instance through the plugin.
type Transport struct {
	Target   Target
	Timeouts Timeouts
}

// Timeouts bound a command the shell has stopped waiting for; a zero field is
// the default.
type Timeouts struct {
	InterruptGrace time.Duration
	Reap           time.Duration
}

// waitDelay is how long the command has left once nobody is waiting on it.
func (t Transport) waitDelay(detached bool) time.Duration {
	if detached {
		return cmp.Or(t.Timeouts.Reap, reapTimeout)
	}
	return cmp.Or(t.Timeouts.InterruptGrace, interruptGrace)
}

// Shell answers the shell's file and environment questions through sh on
// the instance, given Exec to run each command.
func (t Transport) Shell() shell.ExecTransport { return t.Exec }

// Exec runs one command and reports its status, not an error, when it ran.
func (t Transport) Exec(ctx context.Context, sc shell.Command) (int, error) {
	detach := shell.IsDetached(ctx)
	streams := sc.Streams

	cmd := t.Target.CommandArgs(ctx, sc.Args)
	cmd.Dir = sc.Dir
	cmd.Env = EnvMap(sc.Env)
	cmd.Stdin = streams.Stdin
	cmd.Stdout = streams.Stdout
	cmd.Stderr = streams.Stderr

	cmd.WaitDelay = t.waitDelay(detach)
	if sc.Terminal != nil {
		cmd.Terminal = &Terminal{Cols: uint16(sc.Terminal.Size.Cols), Rows: uint16(sc.Terminal.Size.Rows)}
	}

	if err := cmd.Start(); err != nil {
		return 0, err
	}

	if sc.Terminal != nil {
		go func() {
			for {
				select {
				case size := <-sc.Terminal.Resize:
					if cmd.Resize(uint16(size.Cols), uint16(size.Rows)) != nil {
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	var detached <-chan struct{}
	if detach {
		detached = ctx.Done()
	}

	select {
	case <-detached:
		go t.reap(context.WithoutCancel(ctx), cmd, done)
		return 0, ctx.Err()

	case err := <-done:
		var exit *ExitError
		switch {
		case err == nil:
			return 0, nil
		case errors.As(err, &exit):
			return exit.Code, nil
		case ctx.Err() == nil:
			return 0, err
		default:
			// A terminal is still raw while the command has it.
			end := "\n"
			if sc.Terminal != nil {
				end = "\r\n"
			}
			fmt.Fprint(streams.Stderr, err, end)
			return shell.StatusInterrupted, nil
		}
	}
}

// reap waits out a command the shell no longer waits for, logs a bad end and
// drops the instance's record of it.
func (t Transport) reap(ctx context.Context, cmd *Cmd, done <-chan error) {
	if err := <-done; err != nil {
		log.G(ctx).Debug().Err(err).Str("cmd", cmd.UUID).Msg("the interrupted command did not finish")
	}
	cmd.Forget(ctx)
}

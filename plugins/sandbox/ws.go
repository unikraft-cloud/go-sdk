// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"runtime/debug"

	"github.com/coder/websocket"

	plugin "unikraft.com/cloud/plugins/sandbox"
	"unikraft.com/x/log"
)

const (
	stdinChunkSize = 32 << 10

	// A frame is 64 KiB of output and its stream byte; the default limit is 32 KiB.
	wsReadLimit = 1 << 20
)

// attach opens the socket that follows the command from its first byte.
func (c *Cmd) attach() (*websocket.Conn, error) {
	conn, err := c.target.Client.FollowCommand(c.ctx, c.target.Instance, c.UUID, nil, c.target.Opts...)
	if err == nil {
		return conn, nil
	}

	// It was started a moment ago, so the plugin is there but predates the socket.
	if apiErr, ok := plugin.GetAPIError(err); ok && apiErr.IsNotFound() {
		return nil, fmt.Errorf("failed to follow command: the %q plugin cannot stream commands; update its image", c.target.Plugin)
	}

	return nil, c.target.apiError("failed to follow command", err)
}

// follow delivers the command's output as the socket brings it, then its end.
func (c *Cmd) follow(stdout, stderr io.Writer) {
	// A panic in an output writer ends this command, not the process.
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		log.G(c.ctx).Warn().
			Str("cmd", c.UUID).
			Interface("panic", r).
			Bytes("stack", debug.Stack()).
			Msg("command output handler panicked")
		c.done <- fmt.Errorf("command %s: output handler panicked: %v", c.UUID, r)
	}()

	for {
		typ, msg, err := c.conn.Read(c.waitCtx)
		if err != nil {
			c.done <- c.cut(err)
			return
		}

		if typ == websocket.MessageBinary {
			var out io.Writer
			switch {
			case len(msg) == 0:
			case msg[0] == 1:
				out = stdout
			case msg[0] == 2:
				out = stderr
			}
			if out != nil {
				if _, err := out.Write(msg[1:]); err != nil {
					c.done <- fmt.Errorf("writing command output: %w", err)
					return
				}
			}
			continue
		}

		var event struct {
			Event    string
			Exitcode *int
			What     string
			Message  string
		}
		if err := json.Unmarshal(msg, &event); err != nil {
			log.G(c.ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("unreadable event")
			continue
		}
		switch event.Event {
		case "exit":
			c.done <- c.exited(event.Exitcode)
			return
		case "error":
			err := fmt.Errorf("%s: %s", event.What, event.Message)
			if event.What != "stdin" {
				// A resize that crossed the command's end, as often as not.
				log.G(c.ctx).Debug().Err(err).Str("cmd", c.UUID).Msg("the plugin refused a message")
				continue
			}
			select {
			case c.stdinErr <- err:
			default:
			}
		}
	}
}

func (c *Cmd) exited(code *int) error {
	if code == nil {
		return fmt.Errorf("command %s ended without an exit status", c.UUID)
	}

	c.ExitCode = *code
	log.G(c.ctx).Trace().
		Str("cmd", c.UUID).
		Int("exit_code", c.ExitCode).
		Msg("command ended")
	if c.ExitCode != 0 {
		return &ExitError{UUID: c.UUID, Code: c.ExitCode}
	}
	return nil
}

// cut is why the socket ended before the command's exit came through it.
func (c *Cmd) cut(err error) error {
	if c.waitCtx.Err() != nil {
		return c.waitCtx.Err()
	}
	if closed, ok := errors.AsType[websocket.CloseError](err); ok && closed.Code == websocket.StatusPolicyViolation {
		return fmt.Errorf("command %s: %s", c.UUID, closed.Reason)
	}
	return fmt.Errorf("failed to follow command %s: %w", c.UUID, err)
}

// feed sends Stdin to the command a chunk at a time, then closes its input.
func (c *Cmd) feed() {
	report := func(err error) {
		if c.waitCtx.Err() != nil {
			return
		}
		select {
		case c.stdinErr <- err:
		default:
		}
	}

	buf := make([]byte, stdinChunkSize)
	for {
		n, readErr := c.Stdin.Read(buf)
		if n > 0 {
			// The command's own context: a cancelled write would close the socket.
			if err := c.conn.Write(c.waitCtx, websocket.MessageBinary, buf[:n]); err != nil {
				report(fmt.Errorf("failed to send standard input to the command: %w", err))
				return
			}
		}
		if readErr == nil {
			continue
		}
		if !errors.Is(readErr, io.EOF) {
			report(fmt.Errorf("failed to read standard input: %w", readErr))
		}
		if c.waitCtx.Err() == nil {
			c.closeStdin(c.waitCtx)
		}
		return
	}
}

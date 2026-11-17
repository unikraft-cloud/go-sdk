// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	plugin "unikraft.com/cloud/plugins/sandbox"
)

// Ready returns nil when the sandbox plugin of the instance answers.
func (t Target) Ready(ctx context.Context) error {
	return t.apiError("sandbox plugin is not ready", t.Client.Ready(ctx, t.Instance, t.Opts...))
}

// WaitReady blocks until the sandbox plugin of the instance answers, or until
// timeout elapses.
func (t Target) WaitReady(ctx context.Context, timeout time.Duration) error {
	return t.apiError("waiting for the sandbox plugin", t.Client.WaitReady(ctx, t.Instance, timeout, t.Opts...))
}

// Commands returns the UUIDs of the commands that the plugin knows.
func (t Target) Commands(ctx context.Context) ([]string, error) {
	resp, err := t.Client.ListCommands(ctx, t.Instance, t.Opts...)
	if err != nil {
		return nil, t.apiError("failed to list commands", err)
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.Commands, nil
}

// CommandInfo returns the state of the command with uuid: its command line,
// environment, duration and exit code.
func (t Target) CommandInfo(ctx context.Context, uuid string) (*plugin.GetCommandData, error) {
	resp, err := t.Client.GetCommandByUuid(ctx, t.Instance, uuid, t.Opts...)
	if err != nil {
		return nil, t.apiError("failed to get command", err)
	}
	if resp.Data == nil {
		return nil, fmt.Errorf("failed to get command: the %q plugin returned no command", t.Plugin)
	}
	return resp.Data, nil
}

// CommandLogs is a part of the two output streams of a command.
type CommandLogs struct {
	Stdout, Stderr []byte
	// StdoutAvailable and StderrAvailable are the total sizes of the
	// streams, in bytes.
	StdoutAvailable, StderrAvailable uint64
}

// CommandLogsRange selects the part of each stream that
// [Target.CommandLogs] reads. A zero Limit reads to the end of the stream.
type CommandLogsRange struct {
	StdoutOffset, StderrOffset uint64
	Limit                      uint64
}

// CommandLogs reads the output that the command with uuid has written, from
// the offsets in r. It works while the command runs and after it ends.
func (t Target) CommandLogs(ctx context.Context, uuid string, r CommandLogsRange) (*CommandLogs, error) {
	var limit *uint64
	if r.Limit > 0 {
		limit = &r.Limit
	}
	req := plugin.CommandLogsRequest{
		Stdout: plugin.CommandLogsRange{Offset: &r.StdoutOffset, Limit: limit},
		Stderr: plugin.CommandLogsRange{Offset: &r.StderrOffset, Limit: limit},
	}
	resp, err := t.Client.GetCommandLogs(ctx, t.Instance, uuid, &req, t.Opts...)
	if err != nil {
		return nil, t.apiError("failed to read command logs", err)
	}
	if resp.Data == nil {
		return &CommandLogs{}, nil
	}
	stdout, err := base64.StdEncoding.DecodeString(resp.Data.Stdout)
	if err != nil {
		return nil, fmt.Errorf("decoding stdout: %w", err)
	}
	stderr, err := base64.StdEncoding.DecodeString(resp.Data.Stderr)
	if err != nil {
		return nil, fmt.Errorf("decoding stderr: %w", err)
	}
	return &CommandLogs{
		Stdout:          stdout,
		Stderr:          stderr,
		StdoutAvailable: resp.Data.StdoutAvailable,
		StderrAvailable: resp.Data.StderrAvailable,
	}, nil
}

// Stdout returns all the standard output of the command with uuid.
func (t Target) Stdout(ctx context.Context, uuid string) ([]byte, error) {
	return t.rawLog(ctx, uuid, plugin.CommandLogStreamStdout)
}

// Stderr returns all the standard error of the command with uuid.
func (t Target) Stderr(ctx context.Context, uuid string) ([]byte, error) {
	return t.rawLog(ctx, uuid, plugin.CommandLogStreamStderr)
}

func (t Target) rawLog(ctx context.Context, uuid string, stream plugin.CommandLogStream) ([]byte, error) {
	resp, err := t.Client.GetRawCommandLog(ctx, t.Instance, uuid, stream, nil, t.Opts...)
	if err != nil {
		return nil, t.apiError("failed to read "+string(stream), err)
	}
	return resp.RawBody(), nil
}

// DeleteCommand deletes the command with uuid and its output. The plugin
// refuses to delete a command that is still running.
func (t Target) DeleteCommand(ctx context.Context, uuid string) error {
	_, err := t.Client.DeleteCommandByUuid(ctx, t.Instance, uuid, t.Opts...)
	return t.apiError("failed to delete command", err)
}

// DeleteCommandLogs deletes the output of the command with uuid, and keeps
// the command.
func (t Target) DeleteCommandLogs(ctx context.Context, uuid string) error {
	_, err := t.Client.DeleteCommandLogsByUuid(ctx, t.Instance, uuid, t.Opts...)
	return t.apiError("failed to delete command logs", err)
}

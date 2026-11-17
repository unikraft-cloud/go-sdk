// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const pluginPrefix = "/v1/instances/inst-1/plugins/sandbox"

type seenRequest struct {
	method, path, body string
}

// commandsServer answers each path with its JSON envelope data, or with raw
// bytes for the raw log paths, and records the requests.
func commandsServer(t *testing.T, routes map[string]any) (Target, *[]seenRequest) {
	t.Helper()
	var seen []seenRequest
	target := newTargetServing(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		seen = append(seen, seenRequest{r.Method, r.URL.Path, string(body)})
		data, ok := routes[r.Method+" "+r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if raw, isRaw := data.([]byte); isRaw {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write(raw)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": data}))
	}))
	return target, &seen
}

func TestTargetCommands(t *testing.T) {
	t.Parallel()

	target, seen := commandsServer(t, map[string]any{
		"GET " + pluginPrefix + "/commands": map[string]any{"commands": []string{"c1", "c2"}},
	})

	uuids, err := target.Commands(t.Context())
	require.NoError(t, err)
	assert.Equal(t, []string{"c1", "c2"}, uuids)
	assert.Len(t, *seen, 1)
}

func TestTargetCommandInfo(t *testing.T) {
	t.Parallel()

	target, _ := commandsServer(t, map[string]any{
		"GET " + pluginPrefix + "/commands/c1": map[string]any{
			"uuid": "c1", "cmdline": []string{"ls"}, "cwd": "/", "env": map[string]string{},
			"pty": false, "duration_ms": 12, "exitcode": 3, "capped": false,
		},
	})

	info, err := target.CommandInfo(t.Context(), "c1")
	require.NoError(t, err)
	assert.Equal(t, "c1", info.Uuid)
	assert.EqualValues(t, 3, info.Exitcode)
}

func TestTargetCommandLogs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		r    CommandLogsRange
		want string
	}{
		{
			name: "to the end",
			r:    CommandLogsRange{StdoutOffset: 2},
			want: `{"stdout":{"offset":2},"stderr":{"offset":0}}`,
		},
		{
			name: "limited",
			r:    CommandLogsRange{StderrOffset: 4, Limit: 8},
			want: `{"stdout":{"offset":0,"limit":8},"stderr":{"offset":4,"limit":8}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			target, seen := commandsServer(t, map[string]any{
				"GET " + pluginPrefix + "/commands/c1/logs": map[string]any{
					"stdout":           base64.StdEncoding.EncodeToString([]byte("out")),
					"stderr":           base64.StdEncoding.EncodeToString([]byte("err")),
					"stdout_available": 10, "stderr_available": 20,
					"stdout_first": 0, "stderr_first": 0,
				},
			})

			logs, err := target.CommandLogs(t.Context(), "c1", tt.r)
			require.NoError(t, err)
			assert.Equal(t, &CommandLogs{
				Stdout: []byte("out"), Stderr: []byte("err"), StdoutAvailable: 10, StderrAvailable: 20,
			}, logs)
			require.Len(t, *seen, 1)
			assert.JSONEq(t, tt.want, (*seen)[0].body)
		})
	}
}

func TestTargetRawLogs(t *testing.T) {
	t.Parallel()

	target, _ := commandsServer(t, map[string]any{
		"GET " + pluginPrefix + "/commands/c1/logs/raw/stdout": []byte("all of stdout"),
		"GET " + pluginPrefix + "/commands/c1/logs/raw/stderr": []byte("all of stderr"),
	})

	stdout, err := target.Stdout(t.Context(), "c1")
	require.NoError(t, err)
	assert.Equal(t, "all of stdout", string(stdout))

	stderr, err := target.Stderr(t.Context(), "c1")
	require.NoError(t, err)
	assert.Equal(t, "all of stderr", string(stderr))
}

func TestTargetDeletes(t *testing.T) {
	t.Parallel()

	target, seen := commandsServer(t, map[string]any{
		"DELETE " + pluginPrefix + "/commands/c1":      map[string]any{},
		"DELETE " + pluginPrefix + "/commands/c1/logs": map[string]any{},
	})

	require.NoError(t, target.DeleteCommandLogs(t.Context(), "c1"))
	require.NoError(t, target.DeleteCommand(t.Context(), "c1"))
	assert.Equal(t, []string{pluginPrefix + "/commands/c1/logs", pluginPrefix + "/commands/c1"},
		[]string{(*seen)[0].path, (*seen)[1].path})
}

func TestTargetReady(t *testing.T) {
	t.Parallel()

	ready, _ := commandsServer(t, map[string]any{
		"GET " + pluginPrefix + "/commands": map[string]any{"commands": []string{}},
	})
	assert.NoError(t, ready.Ready(t.Context()))

	down, _ := commandsServer(t, map[string]any{})
	assert.Error(t, down.Ready(t.Context()))
}

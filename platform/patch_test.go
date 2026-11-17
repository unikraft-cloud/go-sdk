// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

type sentChange struct {
	Uuid  string `json:"uuid"`
	Prop  string `json:"prop"`
	Op    string `json:"op"`
	Value any    `json:"value"`
}

func okWithList(key string, item map[string]any) map[string]any {
	return map[string]any{
		"status": "success",
		"data":   map[string]any{key: []map[string]any{item}},
	}
}

func TestPatchSendsChangesThenRefreshes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		path  string
		key   string
		patch func(c gen.Client) error
		want  []sentChange
	}{
		{
			name: "instance",
			path: "/v1/instances",
			key:  "instances",
			patch: func(c gen.Client) error {
				return wrapInstance(c, gen.Instance{Uuid: "abc"}).Patch(t.Context(),
					Set(gen.MutableInstancePropertyHostname, "web"),
					Add(gen.MutableInstancePropertyTags, []string{"blue"}),
					Del(gen.MutableInstancePropertyTags, []string{"green"}),
				)
			},
			want: []sentChange{
				{Uuid: "abc", Prop: "hostname", Op: "set", Value: "web"},
				{Uuid: "abc", Prop: "tags", Op: "add", Value: []any{"blue"}},
				{Uuid: "abc", Prop: "tags", Op: "del", Value: []any{"green"}},
			},
		},
		{
			name: "volume",
			path: "/v1/volumes",
			key:  "volumes",
			patch: func(c gen.Client) error {
				return wrapVolume(c, gen.Volume{Uuid: "abc"}).Patch(t.Context(),
					Set(gen.MutableVolumePropertySizeMb, 2048),
					Add(gen.MutableVolumePropertyTags, []string{"blue"}),
				)
			},
			want: []sentChange{
				{Uuid: "abc", Prop: "size_mb", Op: "set", Value: float64(2048)},
				{Uuid: "abc", Prop: "tags", Op: "add", Value: []any{"blue"}},
			},
		},
		{
			name: "service group",
			path: "/v1/services",
			key:  "service_groups",
			patch: func(c gen.Client) error {
				return wrapServiceGroup(c, gen.ServiceGroup{Uuid: "abc"}).Patch(t.Context(),
					Set(gen.MutableServiceGroupPropertySoftLimit, 10),
					Del(gen.MutableServiceGroupPropertyDomains, []string{"a.example.com"}),
				)
			},
			want: []sentChange{
				{Uuid: "abc", Prop: "soft_limit", Op: "set", Value: float64(10)},
				{Uuid: "abc", Prop: "domains", Op: "del", Value: []any{"a.example.com"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, okWithList(tt.key, map[string]any{"uuid": "abc"}))
			})

			require.NoError(t, tt.patch(rawClientAgainst(srv)))
			require.Len(t, *reqs, 2)

			got := (*reqs)[0]
			assert.Equal(t, http.MethodPatch, got.Method)
			assert.Equal(t, tt.path, got.Path)
			var sent []sentChange
			require.NoError(t, json.Unmarshal(got.Body, &sent))
			assert.Equal(t, tt.want, sent)

			assert.Equal(t, http.MethodGet, (*reqs)[1].Method)
		})
	}
}

func TestPatchWithoutChangesSendsNothing(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request %s %s", r.Method, r.Path)
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	require.NoError(t, inst.Patch(t.Context()))
	assert.Empty(t, *reqs)
}

func TestInstanceWaitSendsStateAndTimeoutThenRefreshes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		opts []WaitOption
		want string
	}{
		{name: "defaults", want: `{}`},
		{
			name: "state and timeout",
			opts: []WaitOption{WithWaitState(gen.InstanceStateStopped), WithWaitTimeout(1500 * time.Millisecond)},
			want: `{"state":"stopped","timeout_s":2}`,
		},
		{
			name: "platform maximum",
			opts: []WaitOption{WithWaitTimeout(-1)},
			want: `{"timeout_s":-1}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "abc", "state": "stopped"}))
			})

			inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
			require.NoError(t, inst.Wait(t.Context(), tt.opts...))
			require.Len(t, *reqs, 2)

			got := (*reqs)[0]
			assert.Equal(t, http.MethodGet, got.Method)
			assert.Equal(t, "/v1/instances/abc/wait", got.Path)
			assert.JSONEq(t, tt.want, string(got.Body))
			assert.Equal(t, gen.InstanceStateStopped, inst.State)
		})
	}
}

func TestInstanceLogsPage(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{
			"uuid": "abc", "name": "foo", "output": "aGVsbG8=", "state": "running",
			"available": map[string]any{"start": 0, "end": 5},
			"range":     map[string]any{"start": 0, "end": 5},
		}))
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	page, err := inst.LogsPage(t.Context(), -5, 5)
	require.NoError(t, err)
	assert.Equal(t, "aGVsbG8=", page.Output)

	got := (*reqs)[0]
	assert.Equal(t, "/v1/instances/abc/log", got.Path)
	assert.JSONEq(t, `{"offset":-5,"limit":5}`, string(got.Body))
}

func notFoundResponse() map[string]any {
	return map[string]any{
		"status":  "error",
		"message": "not found",
		"data": map[string]any{"instances": []map[string]any{
			{"status": "error", "error": int(gen.APIHTTPErrorNotFound), "message": "not found"},
		}},
	}
}

func TestInstanceDeleteMissing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		opts    []DeleteOption
		wantErr bool
	}{
		{name: "error by default", wantErr: true},
		{name: "nil with missing OK", opts: []DeleteOption{WithDeleteMissingOK()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusNotFound, notFoundResponse())
			})

			err := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"}).Delete(t.Context(), tt.opts...)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestInstanceDeleteRetriesWhileBusy(t *testing.T) {
	t.Parallel()

	calls := 0
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls++
		if calls == 1 {
			writeJSON(t, w, http.StatusConflict, map[string]any{"status": "error", "message": "device EBUSY"})
			return
		}
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	require.NoError(t, inst.Delete(t.Context(), WithDeleteRetryBusy(5*time.Second)))
	assert.Len(t, *reqs, 2)
}

func TestIsBusy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "nil", err: nil, want: false},
		{name: "EBUSY", err: errors.New("detach failed: EBUSY"), want: true},
		{name: "errno -16", err: errors.New("Unknown error -16"), want: true},
		{name: "errno -160", err: errors.New("Unknown error -160"), want: false},
		{name: "other", err: errors.New("not found"), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, isBusy(tt.err))
		})
	}
}

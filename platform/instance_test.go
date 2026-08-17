// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

func rawClientAgainst(srv *httptest.Server) gen.Client {
	return newClientAgainst(srv).Raw()
}

func TestInstanceRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"uuid": "abc", "name": "foo", "state": "running", "memory_mb": 256},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	err := inst.Refresh(context.Background())
	require.NoError(t, err)
	assert.Equal(t, gen.InstanceStateRunning, inst.State)
	assert.EqualValues(t, 256, inst.MemoryMb)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "details=true", got.Query)

	var sent []gen.NameOrUUID
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.Len(t, sent, 1)
	require.NotNil(t, sent[0].Uuid)
	assert.Equal(t, "abc", *sent[0].Uuid)
}

func TestInstanceRefreshResetsUpdateSnapshot(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"uuid": "abc", "memory_mb": 512},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128})
	require.NoError(t, inst.Refresh(context.Background()))

	// Refresh must reset the diff snapshot to the freshly fetched value, so
	// an immediate Update (with no local edits since) has nothing to send.
	assert.Empty(t, inst.diff())
}

func TestInstanceRefreshNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"instances": []map[string]any{}},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	err := inst.Refresh(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "abc")
	assert.ErrorContains(t, err, "not found")
}

func TestInstanceStartRefreshesAfterwards(t *testing.T) {
	var calls []string
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls = append(calls, r.Method+" "+r.Path)
		switch {
		case r.Method == http.MethodPut && r.Path == "/v1/instances/abc/start":
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case r.Method == http.MethodGet && r.Path == "/v1/instances":
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"instances": []map[string]any{{"uuid": "abc", "state": "running"}}},
			})
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.Path)
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", State: gen.InstanceStateStopped})
	require.NoError(t, inst.Start(context.Background()))
	assert.Equal(t, gen.InstanceStateRunning, inst.State)
	assert.Equal(t, []string{"PUT /v1/instances/abc/start", "GET /v1/instances"}, calls)
}

func TestInstanceStopRefreshesAfterwards(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch {
		case r.Method == http.MethodPut && r.Path == "/v1/instances/abc/stop":
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case r.Method == http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"instances": []map[string]any{{"uuid": "abc", "state": "stopped"}}},
			})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", State: gen.InstanceStateRunning})
	require.NoError(t, inst.Stop(context.Background()))
	assert.Equal(t, gen.InstanceStateStopped, inst.State)
	assert.Len(t, *reqs, 2)
}

func TestInstanceSuspendRefreshesAfterwards(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch {
		case r.Method == http.MethodPut && r.Path == "/v1/instances/abc/suspend":
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case r.Method == http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"instances": []map[string]any{{"uuid": "abc", "state": "standby"}}},
			})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	require.NoError(t, inst.Suspend(context.Background()))
	assert.EqualValues(t, "standby", inst.State)
	assert.Len(t, *reqs, 2)
}

func TestInstanceDeleteDoesNotRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/instances/abc", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	require.NoError(t, inst.Delete(context.Background()))
	assert.Len(t, *reqs, 1, "Delete must not attempt to refresh a now-gone instance")
}

func TestInstanceStartPropagatesLifecycleError(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusInternalServerError, map[string]any{"status": "error", "message": "start failed"})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	err := inst.Start(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "start failed")
}

func TestInstanceUpdateWithNoChangesSkipsRequest(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request: %s %s", r.Method, r.Path)
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128})
	require.NoError(t, inst.Update(context.Background()))
	assert.Empty(t, *reqs)
}

func TestInstanceUpdateSendsOnlyChangedFieldsThenRefreshes(t *testing.T) {
	var patchBody []byte
	var calls []string
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls = append(calls, r.Method+" "+r.Path)
		switch r.Method {
		case http.MethodPatch:
			patchBody = r.Body
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"instances": []map[string]any{{"uuid": "abc", "memory_mb": 512}}},
			})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128, Vcpus: 2})
	inst.MemoryMb = 512 // only this field changes

	require.NoError(t, inst.Update(context.Background()))
	assert.Equal(t, []string{"PATCH /v1/instances", "GET /v1/instances"}, calls)

	var items []gen.UpdateInstancesRequestItem
	require.NoError(t, json.Unmarshal(patchBody, &items))
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableInstancePropertyMemoryMb, items[0].Prop)
	assert.Equal(t, gen.MutableInstanceOperationSet, items[0].Op)

	assert.EqualValues(t, 512, inst.MemoryMb, "Update refreshes the instance from the server response")
}

func TestInstanceUpdateJoinsRefreshErrorWhenPostUpdateRefreshFails(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case http.MethodGet:
			// The instance vanished between the update and the refresh.
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128})
	inst.MemoryMb = 256

	err := inst.Update(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "not found")
}

func TestInstanceUpdateReturnsAPIErrorWithoutRefreshWhenIfstateSet(t *testing.T) {
	refreshed := false
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "update rejected"})
		case http.MethodGet:
			refreshed = true
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"instances": []map[string]any{{"uuid": "abc", "memory_mb": 256}}},
			})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128})
	inst.MemoryMb = 256

	err := inst.Update(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "update rejected")
	assert.True(t, refreshed, "Update always refreshes even when the update call itself failed")
}

func TestInstanceDiffDetectsEachMutableField(t *testing.T) {
	base := gen.Instance{
		Uuid:     "abc",
		MemoryMb: 128,
		Vcpus:    1,
		Image:    "unikraft.io/base:latest",
		Args:     []string{"--foo"},
		Env:      map[string]string{"A": "1"},
		Tags:     []string{"prod"},
	}

	tests := []struct {
		name    string
		mutate  func(*Instance)
		wantLen int
		prop    gen.MutableInstanceProperty
	}{
		{"memory", func(i *Instance) { i.MemoryMb = 256 }, 1, gen.MutableInstancePropertyMemoryMb},
		{"vcpus", func(i *Instance) { i.Vcpus = 2 }, 1, gen.MutableInstancePropertyVcpus},
		{"image", func(i *Instance) { i.Image = "unikraft.io/base:v2" }, 1, gen.MutableInstancePropertyImage},
		{"args", func(i *Instance) { i.Args = []string{"--bar"} }, 1, gen.MutableInstancePropertyArgs},
		{"env", func(i *Instance) { i.Env = map[string]string{"A": "2"} }, 1, gen.MutableInstancePropertyEnv},
		{"tags", func(i *Instance) { i.Tags = []string{"staging"} }, 1, gen.MutableInstancePropertyTags},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := wrapInstance(nil, base)
			tt.mutate(inst)
			items := inst.diff()
			require.Len(t, items, tt.wantLen)
			assert.Equal(t, tt.prop, items[0].Prop)
			assert.Equal(t, gen.MutableInstanceOperationSet, items[0].Op)
		})
	}
}

func TestInstanceDiffNoChangesIsEmpty(t *testing.T) {
	base := gen.Instance{Uuid: "abc", MemoryMb: 128, Args: []string{"--foo"}, Env: map[string]string{"A": "1"}}
	inst := wrapInstance(nil, base)
	assert.Empty(t, inst.diff())
}

func TestInstanceDiffMultipleChangesProducesMultipleItems(t *testing.T) {
	base := gen.Instance{Uuid: "abc", MemoryMb: 128, Vcpus: 1}
	inst := wrapInstance(nil, base)
	inst.MemoryMb = 256
	inst.Vcpus = 4

	items := inst.diff()
	require.Len(t, items, 2)

	var props []gen.MutableInstanceProperty
	for _, it := range items {
		props = append(props, it.Prop)
	}
	assert.ElementsMatch(t, []gen.MutableInstanceProperty{gen.MutableInstancePropertyMemoryMb, gen.MutableInstancePropertyVcpus}, props)
}

func TestInstanceDiffPointerFields(t *testing.T) {
	trueVal := true
	base := gen.Instance{Uuid: "abc", DeleteLock: &trueVal}
	inst := wrapInstance(nil, base)

	falseVal := false
	inst.DeleteLock = &falseVal

	items := inst.diff()
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableInstancePropertyDeleteLock, items[0].Prop)
}

func TestInstanceDiffDoesNotTrackHostnameOrDependencies(t *testing.T) {
	// Hostname and Dependencies are write-only and have no field to diff
	// against, so mutating unrelated fields must not accidentally pick them up.
	base := gen.Instance{Uuid: "abc"}
	inst := wrapInstance(nil, base)
	assert.Empty(t, inst.diff())
}

func TestWrapInstanceClonesDataSoOrigIsNotAliased(t *testing.T) {
	data := gen.Instance{Uuid: "abc", Args: []string{"--foo"}}
	inst := wrapInstance(nil, data)

	// Mutating the caller's original slice must not affect the snapshot
	// wrapInstance took, since it deep-clones via JSON round-trip.
	data.Args[0] = "--mutated"
	assert.Equal(t, "--foo", inst.orig.Args[0])
	assert.Equal(t, "--foo", inst.Args[0])
}

func TestInstanceUpdateWrapsMarshalAndRefreshErrorsTogether(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "update rejected"})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
		}
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc", MemoryMb: 128})
	inst.MemoryMb = 256

	err := inst.Update(context.Background())
	require.Error(t, err)

	var joined interface{ Unwrap() []error }
	require.True(t, errors.As(err, &joined))
	assert.Len(t, joined.Unwrap(), 2)
}

func TestInstanceLogsReadsContentThroughRealClient(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, "/v1/instances/log", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{
						"uuid":      "abc",
						"output":    base64.StdEncoding.EncodeToString([]byte("hello world")),
						"available": map[string]any{"start": 0, "end": 11},
						"range":     map[string]any{"start": 0, "end": 11},
					},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	rc, err := inst.Logs(context.Background())
	require.NoError(t, err)
	defer rc.Close()

	got, err := io.ReadAll(rc)
	require.NoError(t, err)
	assert.Equal(t, "hello world", string(got))
}

func TestInstanceLogsCloseCancelsWithoutError(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{
						"uuid":      "abc",
						"output":    base64.StdEncoding.EncodeToString([]byte("hi")),
						"available": map[string]any{"start": 0, "end": 2},
						"range":     map[string]any{"start": 0, "end": 2},
					},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	rc, err := inst.Logs(context.Background())
	require.NoError(t, err)

	_, err = io.ReadAll(rc)
	require.NoError(t, err)
	assert.NoError(t, rc.Close())
}

func TestInstanceMetrics(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"rss_bytes": 1024, "cpu_time_ms": 42},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	metrics, err := inst.Metrics(context.Background())
	require.NoError(t, err)
	assert.EqualValues(t, 1024, metrics.RssBytes)
	assert.EqualValues(t, 42, metrics.CpuTimeMs)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/instances/abc/metrics", got.Path)
}

func TestInstanceMetricsNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"instances": []map[string]any{}},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	_, err := inst.Metrics(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "abc")
}

func TestInstanceHistory(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{
						"uuid": "abc",
						"history": []map[string]any{
							{"uuid": "chk1", "name": "checkpoint-1"},
						},
					},
				},
			},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	history, err := inst.History(context.Background())
	require.NoError(t, err)
	require.Len(t, history, 1)
	assert.Equal(t, "checkpoint-1", history[0].Name)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/instances/abc/history", got.Path)
}

func TestInstanceHistoryNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"instances": []map[string]any{}},
		})
	})

	inst := wrapInstance(rawClientAgainst(srv), gen.Instance{Uuid: "abc"})
	_, err := inst.History(context.Background())
	require.Error(t, err)
	assert.ErrorContains(t, err, "abc")
}

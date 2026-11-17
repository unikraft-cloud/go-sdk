// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/controlplane/gen"
)

func rawClientAgainst(srv *httptest.Server) gen.Client {
	return newClientAgainst(srv).Raw()
}

func TestNodeRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"nodes": []map[string]any{
					{"uuid": "abc", "name": "foo", "state": "ready", "vcpus": 4},
				},
			},
		})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	require.NoError(t, node.Refresh(context.Background()))
	assert.Equal(t, gen.NodeStateReady, node.State)
	assert.EqualValues(t, 4, node.Vcpus)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/nodes/abc", got.Path)
}

func TestNodeRefreshResetsUpdateSnapshot(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"nodes": []map[string]any{{"uuid": "abc", "delete_lock": true}}},
		})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: false})
	require.NoError(t, node.Refresh(context.Background()))

	items, err := node.diff()
	require.NoError(t, err)
	assert.Empty(t, items)
}

func TestNodeRefreshNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"nodes": []map[string]any{}},
		})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	err := node.Refresh(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "abc")
	require.ErrorContains(t, err, "not found")
}

func TestNodeDestroyDoesNotRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/nodes/abc", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	require.NoError(t, node.Destroy(context.Background()))
	assert.Len(t, *reqs, 1, "Destroy must not refresh a now-gone node")
}

func TestNodeDestroyForwardsOpts(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	require.NoError(t, node.Destroy(context.Background(), WithDestroyForce()))
	assert.Equal(t, "force=true", (*reqs)[0].Query)
}

func TestNodeDestroyPropagatesAPIError(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "delete-locked"})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	err := node.Destroy(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "delete-locked")
}

func TestNodeWaitRefreshesState(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"nodes": []map[string]any{{"uuid": "abc", "state": "ready"}}},
		})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", State: gen.NodeState("provisioning")})
	require.NoError(t, node.Wait(context.Background(), WithWaitStates(gen.NodeStateReady)))
	assert.Equal(t, gen.NodeStateReady, node.State)

	got := (*reqs)[0]
	assert.Equal(t, "/v1/nodes/abc/wait", got.Path)
	assert.Equal(t, "states=ready", got.Query)
}

func TestNodeWaitNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc"})
	err := node.Wait(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "not found")
}

func TestNodeUpdateWithNoChangesSkipsRequest(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request: %s %s", r.Method, r.Path)
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: true})
	require.NoError(t, node.Update(context.Background()))
	assert.Empty(t, *reqs)
}

func TestNodeUpdateSendsOnlyChangedFieldsThenRefreshes(t *testing.T) {
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
				"data":   map[string]any{"nodes": []map[string]any{{"uuid": "abc", "delete_lock": true}}},
			})
		}
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: false})
	node.DeleteLock = true

	require.NoError(t, node.Update(context.Background()))
	assert.Equal(t, []string{"PATCH /v1/nodes/abc", "GET /v1/nodes/abc"}, calls)

	var items []gen.UpdateNodePayload
	require.NoError(t, json.Unmarshal(patchBody, &items))
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableNodePropertyDeleteLock, items[0].Property)
	assert.Equal(t, gen.MutableNodeOperationSet, items[0].Operation)
	require.NotNil(t, items[0].Value)
	assert.Equal(t, "true", *items[0].Value)

	assert.True(t, node.DeleteLock)
}

func TestNodeUpdateDetectsTagsChange(t *testing.T) {
	var patchBody []byte
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			patchBody = r.Body
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"nodes": []map[string]any{{"uuid": "abc"}}},
			})
		}
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", Tags: map[string]string{"env": "prod"}})
	node.Tags = map[string]string{"env": "staging"}

	require.NoError(t, node.Update(context.Background()))

	var items []gen.UpdateNodePayload
	require.NoError(t, json.Unmarshal(patchBody, &items))
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableNodePropertyTags, items[0].Property)

	var sentTags map[string]string
	require.NoError(t, json.Unmarshal([]byte(*items[0].Value), &sentTags))
	assert.Equal(t, map[string]string{"env": "staging"}, sentTags)
}

func TestNodeUpdateJoinsRefreshErrorWhenPostUpdateRefreshFails(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
		}
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: false})
	node.DeleteLock = true

	err := node.Update(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "not found")
}

func TestNodeUpdateReturnsAPIErrorButStillRefreshes(t *testing.T) {
	refreshed := false
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "update rejected"})
		case http.MethodGet:
			refreshed = true
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{{"uuid": "abc"}}}})
		}
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: false})
	node.DeleteLock = true

	err := node.Update(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "update rejected")
	assert.True(t, refreshed)
}

func TestNodeDiffDoesNotTrackSshKeys(t *testing.T) {
	// SshKeys use add/del semantics rather than set, so it's intentionally
	// excluded from diffing even when it changes.
	node := wrapNode(nil, gen.Node{Uuid: "abc", SshKeys: []gen.SSHKey{{Name: "k1"}}})
	node.SshKeys = []gen.SSHKey{{Name: "k1"}, {Name: "k2"}}

	items, err := node.diff()
	require.NoError(t, err)
	assert.Empty(t, items)
}

func TestNodeDiffNoChangesIsEmpty(t *testing.T) {
	node := wrapNode(nil, gen.Node{Uuid: "abc", Tags: map[string]string{"a": "1"}})
	items, err := node.diff()
	require.NoError(t, err)
	assert.Empty(t, items)
}

func TestWrapNodeClonesDataSoOrigIsNotAliased(t *testing.T) {
	data := gen.Node{Uuid: "abc", Tags: map[string]string{"env": "prod"}}
	node := wrapNode(nil, data)

	data.Tags["env"] = "mutated"
	assert.Equal(t, "prod", node.orig.Tags["env"])
	assert.Equal(t, "prod", node.Tags["env"])
}

func TestNodeUpdateWrapsMarshalAndRefreshErrorsTogether(t *testing.T) {
	// Sanity check that errors.Join is actually used to combine the update
	// and refresh failures rather than one silently shadowing the other.
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch r.Method {
		case http.MethodPatch:
			writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "update rejected"})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
		}
	})

	node := wrapNode(rawClientAgainst(srv), gen.Node{Uuid: "abc", DeleteLock: false})
	node.DeleteLock = true

	err := node.Update(context.Background())
	require.Error(t, err)

	var joined interface{ Unwrap() []error }
	require.ErrorAs(t, err, &joined)
	assert.Len(t, joined.Unwrap(), 2)
}

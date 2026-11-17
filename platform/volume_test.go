// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

func TestVolumeRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"volumes": []map[string]any{
					{"uuid": "abc", "name": "vol", "size_mb": 1024},
				},
			},
		})
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc"})
	require.NoError(t, vol.Refresh(context.Background()))
	assert.EqualValues(t, 1024, vol.SizeMb)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/volumes", got.Path)
}

func TestVolumeDeleteDoesNotRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/volumes/abc", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc"})
	require.NoError(t, vol.Delete(context.Background()))
	assert.Len(t, *reqs, 1)
}

func TestVolumeAttachRefreshesAfterwards(t *testing.T) {
	var attachBody []byte
	var calls []string
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls = append(calls, r.Method+" "+r.Path)
		switch {
		case r.Method == http.MethodPut && r.Path == "/v1/volumes/abc/attach":
			attachBody = r.Body
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case r.Method == http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc"}}},
			})
		}
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc"})
	require.NoError(t, vol.Attach(context.Background(), Name("instance-1"), "/data", WithAttachReadOnly()))
	assert.Equal(t, []string{"PUT /v1/volumes/abc/attach", "GET /v1/volumes"}, calls)

	var sent gen.AttachVolumeByUUIDRequestBody
	require.NoError(t, json.Unmarshal(attachBody, &sent))
	assert.Equal(t, "/data", sent.At)
	require.NotNil(t, sent.AttachTo.Name)
	assert.Equal(t, "instance-1", *sent.AttachTo.Name)
	require.NotNil(t, sent.Readonly)
	assert.True(t, *sent.Readonly)
}

func TestVolumeDetachFromAllInstancesWhenRefNil(t *testing.T) {
	var detachBody []byte
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch {
		case r.Method == http.MethodPut && r.Path == "/v1/volumes/abc/detach":
			detachBody = r.Body
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case r.Method == http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc"}}},
			})
		}
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc"})
	require.NoError(t, vol.Detach(context.Background(), nil))

	var sent gen.DetachVolumeByUUIDRequestBody
	require.NoError(t, json.Unmarshal(detachBody, &sent))
	assert.Nil(t, sent.From)
}

func TestVolumeUpdateSendsOnlyChangedFieldsThenRefreshes(t *testing.T) {
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
				"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc", "size_mb": 2048}}},
			})
		}
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc", SizeMb: 1024})
	vol.SizeMb = 2048 // only this field changes

	require.NoError(t, vol.Update(context.Background()))
	assert.Equal(t, []string{"PATCH /v1/volumes", "GET /v1/volumes"}, calls)

	var items []gen.UpdateVolumesRequestItem
	require.NoError(t, json.Unmarshal(patchBody, &items))
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableVolumePropertySizeMb, items[0].Prop)
	assert.Equal(t, gen.MutableVolumeOperationSet, items[0].Op)
}

func TestVolumeUpdateWithNoChangesSkipsRequest(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request: %s %s", r.Method, r.Path)
	})

	vol := wrapVolume(rawClientAgainst(srv), gen.Volume{Uuid: "abc", SizeMb: 1024})
	require.NoError(t, vol.Update(context.Background()))
	assert.Empty(t, *reqs)
}

func TestClientCreateVolumeFetchesFullVolumeAfterAck(t *testing.T) {
	var calls []string
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls = append(calls, r.Method+" "+r.Path)
		switch r.Method {
		case http.MethodPost:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc", "name": "vol", "state": "available"}}},
			})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc", "name": "vol", "size_mb": 512}}},
			})
		}
	})

	c := newClientAgainst(srv)
	vol, err := c.CreateVolume(context.Background(), WithCreateVolumeName("vol"), WithCreateVolumeSizeMB(512))
	require.NoError(t, err)
	assert.Equal(t, "abc", vol.Uuid)
	assert.EqualValues(t, 512, vol.SizeMb)
	assert.Equal(t, []string{"POST /v1/volumes", "GET /v1/volumes"}, calls)
}

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

func TestServiceGroupRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"service_groups": []map[string]any{
					{"uuid": "abc", "name": "foo", "soft_limit": 5},
				},
			},
		})
	})

	sg := wrapServiceGroup(rawClientAgainst(srv), gen.ServiceGroup{Uuid: "abc"})
	require.NoError(t, sg.Refresh(context.Background()))
	assert.EqualValues(t, 5, sg.SoftLimit)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/services", got.Path)
}

func TestServiceGroupDeleteDoesNotRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/services/abc", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	sg := wrapServiceGroup(rawClientAgainst(srv), gen.ServiceGroup{Uuid: "abc"})
	require.NoError(t, sg.Delete(context.Background()))
	assert.Len(t, *reqs, 1)
}

func TestServiceGroupUpdateSendsOnlyChangedFieldsThenRefreshes(t *testing.T) {
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
				"data":   map[string]any{"service_groups": []map[string]any{{"uuid": "abc", "soft_limit": 10}}},
			})
		}
	})

	sg := wrapServiceGroup(rawClientAgainst(srv), gen.ServiceGroup{Uuid: "abc", SoftLimit: 5})
	sg.SoftLimit = 10 // only this field changes

	require.NoError(t, sg.Update(context.Background()))
	assert.Equal(t, []string{"PATCH /v1/services", "GET /v1/services"}, calls)

	var items []gen.UpdateServiceGroupsRequestItem
	require.NoError(t, json.Unmarshal(patchBody, &items))
	require.Len(t, items, 1)
	assert.Equal(t, gen.MutableServiceGroupPropertySoftLimit, items[0].Prop)
	assert.Equal(t, gen.MutableServiceGroupOperationSet, items[0].Op)
}

func TestServiceGroupUpdateWithNoChangesSkipsRequest(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request: %s %s", r.Method, r.Path)
	})

	sg := wrapServiceGroup(rawClientAgainst(srv), gen.ServiceGroup{Uuid: "abc", SoftLimit: 5})
	require.NoError(t, sg.Update(context.Background()))
	assert.Empty(t, *reqs)
}

func TestClientCreateServiceGroup(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"service_groups": []map[string]any{{"uuid": "abc", "name": "svc"}},
			},
		})
	})

	c := newClientAgainst(srv)
	sg, err := c.CreateServiceGroup(context.Background(), WithCreateServiceGroupName("svc"), WithCreateServiceGroupSoftLimit(3))
	require.NoError(t, err)
	assert.Equal(t, "abc", sg.Uuid)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodPost, got.Method)
	assert.Equal(t, "/v1/services", got.Path)

	var sent gen.CreateServiceGroupRequest
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.NotNil(t, sent.Name)
	assert.Equal(t, "svc", *sent.Name)
	require.NotNil(t, sent.SoftLimit)
	assert.EqualValues(t, 3, *sent.SoftLimit)
}

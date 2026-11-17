// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

func emptyList(key string) map[string]any {
	return map[string]any{"status": "success", "data": map[string]any{key: []any{}}}
}

func TestClientTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		response    map[string]any
		wantMissing bool
	}{
		{name: "found", response: okWithList("instances", map[string]any{"uuid": "t1", "name": "tpl"})},
		{name: "missing", response: emptyList("instances"), wantMissing: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, tt.response)
			})

			tmpl, err := newClientAgainst(srv).Template(t.Context(), Name("tpl"))
			assert.Equal(t, "/v1/instances/templates", (*reqs)[0].Path)
			if tt.wantMissing {
				assert.True(t, isTemplateMissing(err))
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "t1", tmpl.Uuid)
		})
	}
}

func TestClientTemplatesSendsTags(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "t1"}))
	})

	tmpls, err := newClientAgainst(srv).Templates(t.Context(), nil, WithTemplateTags("a", "b"), WithTemplateLimit(5))
	require.NoError(t, err)
	require.Len(t, tmpls, 1)
	assert.Equal(t, "count=5&tags=a&tags=b", (*reqs)[0].Query)
}

func TestTemplateDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		opts    []TemplateDeleteOption
		wantErr bool
	}{
		{name: "error by default", wantErr: true},
		{name: "nil with missing OK", opts: []TemplateDeleteOption{WithTemplateDeleteMissingOK()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusNotFound, notFoundResponse())
			})

			err := wrapTemplate(rawClientAgainst(srv), gen.Instance{Uuid: "t1"}).Delete(t.Context(), tt.opts...)
			assert.Equal(t, "/v1/instances/templates/t1", (*reqs)[0].Path)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestTemplateCloneSendsTemplateUUID(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "i1"}))
	})

	tmpl := wrapTemplate(rawClientAgainst(srv), gen.Instance{Uuid: "t1"})
	instances, err := tmpl.Clone(t.Context(), WithCreateName("clone"))
	require.NoError(t, err)
	require.Len(t, instances, 1)
	assert.Equal(t, "i1", instances[0].Uuid)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodPost, got.Method)
	assert.Equal(t, "/v1/instances", got.Path)
	assert.JSONEq(t, `{"name":"clone","template":{"uuid":"t1"}}`, string(got.Body))
}

func TestTemplateCloneRejectsImage(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request %s", r.Path)
	})

	tmpl := wrapTemplate(rawClientAgainst(srv), gen.Instance{Uuid: "t1"})
	_, err := tmpl.Clone(t.Context(), WithCreateImage(gen.ImageReference("nginx")))
	assert.Error(t, err)
	assert.Empty(t, *reqs)
}

func TestPrepareTemplateReturnsExisting(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "t1", "name": "tpl"}))
	})

	tmpl, err := newClientAgainst(srv).PrepareTemplate(t.Context(), "tpl", gen.Instance{})
	require.NoError(t, err)
	assert.Equal(t, "t1", tmpl.Uuid)
	assert.Len(t, *reqs, 1)
}

func TestPrepareTemplateSeedsWaitsAndCleansUp(t *testing.T) {
	t.Parallel()

	lookups := 0
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		switch {
		case r.Method == http.MethodGet && r.Path == "/v1/instances/templates":
			lookups++
			if lookups < 3 {
				writeJSON(t, w, http.StatusOK, emptyList("instances"))
				return
			}
			writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "t1", "name": "tpl"}))
		case r.Method == http.MethodPost && r.Path == "/v1/instances":
			writeJSON(t, w, http.StatusOK, okWithList("instances", map[string]any{"uuid": "seed"}))
		case r.Method == http.MethodDelete && r.Path == "/v1/instances/seed":
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.Path)
		}
	})

	source := gen.Instance{Image: "org/worker:v3", MemoryMb: 1024, Tags: []string{"team"}}
	tmpl, err := newClientAgainst(srv).PrepareTemplate(t.Context(), "tpl", source,
		WithPrepareAutokill(time.Minute), WithPrepareTimeout(30*time.Second))
	require.NoError(t, err)
	assert.Equal(t, "t1", tmpl.Uuid)

	var create map[string]any
	require.NoError(t, json.Unmarshal((*reqs)[1].Body, &create))
	assert.Equal(t, false, create["autostart"])
	assert.Equal(t, []any{"team"}, create["tags"])
	template := create["template"].(map[string]any)
	assert.Equal(t, "tpl", template["name"])
	assert.Equal(t, true, template["prepare"])
	assert.EqualValues(t, 30, template["prepare_timeout_s"])
	assert.Equal(t, map[string]any{"time_ms": float64(60000)}, template["autokill"])

	last := (*reqs)[len(*reqs)-1]
	assert.Equal(t, http.MethodDelete, last.Method)
}

func TestClientImagesSendsFilters(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, okWithList("images", map[string]any{"digest": "sha256:aa"}))
	})

	images, err := newClientAgainst(srv).Images(t.Context(), WithImageTag("org/app:v1"), WithImageDigest("sha256:aa"))
	require.NoError(t, err)
	require.Len(t, images, 1)
	assert.Equal(t, "/v1/images", (*reqs)[0].Path)
	assert.Equal(t, "digest=sha256%3Aaa&tag=org%2Fapp%3Av1", (*reqs)[0].Query)
}

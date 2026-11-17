// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBulkSendsOneRequest(t *testing.T) {
	t.Parallel()

	refs := []Ref{Name("web"), UUID("u2")}

	tests := []struct {
		name   string
		call   func(c Client) error
		method string
		path   string
		body   string
	}{
		{
			name:   "start",
			call:   func(c Client) error { return c.StartInstances(t.Context(), refs, WithStartTimeout(5*time.Second)) },
			method: http.MethodPut, path: "/v1/instances/start",
			body: `[{"name":"web","timeout_s":5},{"uuid":"u2","timeout_s":5}]`,
		},
		{
			name:   "stop",
			call:   func(c Client) error { return c.StopInstances(t.Context(), refs, WithStopForce()) },
			method: http.MethodPut, path: "/v1/instances/stop",
			body: `[{"name":"web","force":true},{"uuid":"u2","force":true}]`,
		},
		{
			name:   "suspend",
			call:   func(c Client) error { return c.SuspendInstances(t.Context(), refs) },
			method: http.MethodPut, path: "/v1/instances/suspend",
			body: `[{"name":"web"},{"uuid":"u2"}]`,
		},
		{
			name:   "delete instances",
			call:   func(c Client) error { return c.DeleteInstances(t.Context(), refs, WithoutRetention()) },
			method: http.MethodDelete, path: "/v1/instances",
			body: `[{"name":"web","dont_retain":true},{"uuid":"u2","dont_retain":true}]`,
		},
		{
			name:   "delete volumes",
			call:   func(c Client) error { return c.DeleteVolumes(t.Context(), refs) },
			method: http.MethodDelete, path: "/v1/volumes",
			body: `[{"name":"web"},{"uuid":"u2"}]`,
		},
		{
			name:   "delete service groups",
			call:   func(c Client) error { return c.DeleteServiceGroups(t.Context(), refs) },
			method: http.MethodDelete, path: "/v1/services",
			body: `[{"name":"web"},{"uuid":"u2"}]`,
		},
		{
			name:   "delete certificates",
			call:   func(c Client) error { return c.DeleteCertificates(t.Context(), refs) },
			method: http.MethodDelete, path: "/v1/certificates",
			body: `[{"name":"web"},{"uuid":"u2"}]`,
		},
		{
			name:   "delete templates",
			call:   func(c Client) error { return c.DeleteTemplates(t.Context(), refs) },
			method: http.MethodDelete, path: "/v1/instances/templates",
			body: `[{"name":"web"},{"uuid":"u2"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
			})

			require.NoError(t, tt.call(newClientAgainst(srv)))
			require.Len(t, *reqs, 1)
			got := (*reqs)[0]
			assert.Equal(t, tt.method, got.Method)
			assert.Equal(t, tt.path, got.Path)
			assert.JSONEq(t, tt.body, string(got.Body))
		})
	}
}

func TestDeleteInstancesMissingOK(t *testing.T) {
	t.Parallel()

	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusNotFound, notFoundResponse())
	})

	c := newClientAgainst(srv)
	assert.Error(t, c.DeleteInstances(t.Context(), []Ref{Name("web")}))
	assert.NoError(t, c.DeleteInstances(t.Context(), []Ref{Name("web")}, WithDeleteMissingOK()))
}

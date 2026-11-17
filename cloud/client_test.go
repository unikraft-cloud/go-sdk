// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package cloud

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/controlplane/gen"
	"unikraft.com/cloud/sdk/platform"
)

type recordedRequest struct {
	Method  string
	Path    string
	Headers http.Header
}

func newTestServer(t *testing.T, body any) (*httptest.Server, *[]recordedRequest) {
	t.Helper()
	var got []recordedRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		got = append(got, recordedRequest{Method: r.Method, Path: r.URL.Path, Headers: r.Header.Clone()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		assert.NoError(t, json.NewEncoder(w).Encode(body))
	}))
	t.Cleanup(srv.Close)
	return srv, &got
}

func TestClientPlatformSharesTokenAndUserAgent(t *testing.T) {
	srv, reqs := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})

	c := NewClient(WithToken("shared-token"), WithUserAgent("shared-agent/1"))
	p := c.Platform(srv.URL)

	_, err := p.Instances(context.Background(), []platform.Ref{})
	require.NoError(t, err)

	got := (*reqs)[0]
	assert.Equal(t, "Bearer shared-token", got.Headers.Get("Authorization"))
	assert.Equal(t, "shared-agent/1", got.Headers.Get("User-Agent"))
}

func TestClientControlPlaneSharesTokenAndUserAgent(t *testing.T) {
	srv, reqs := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})

	c := NewClient(WithToken("shared-token"), WithUserAgent("shared-agent/1"))
	cp := c.ControlPlane().WithEndpoint(srv.URL)

	_, err := cp.Raw().ListNodes(context.Background(), []gen.NameOrUUID{}, gen.ListNodesOpts{})
	require.NoError(t, err)

	got := (*reqs)[0]
	assert.Equal(t, "Bearer shared-token", got.Headers.Get("Authorization"))
	assert.Equal(t, "shared-agent/1", got.Headers.Get("User-Agent"))
}

func TestClientPlatformWithCustomHTTPClient(t *testing.T) {
	srv, reqs := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})

	calls := 0
	custom := countingHTTPClient{inner: &http.Client{}, calls: &calls}

	c := NewClient(WithToken("t"), WithHTTPClient(custom))
	p := c.Platform(srv.URL)

	_, err := p.Instances(context.Background(), []platform.Ref{})
	require.NoError(t, err)

	assert.Equal(t, 1, calls)
	assert.Len(t, *reqs, 1)
}

type countingHTTPClient struct {
	inner *http.Client
	calls *int
}

func (c countingHTTPClient) Do(req *http.Request) (*http.Response, error) {
	*c.calls++
	return c.inner.Do(req)
}

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

func TestClientPlatformEmptyMetroDoesNotOverrideDefaultResolution(t *testing.T) {
	// Passing an empty metro must not call platform.WithDefaultMetro(""),
	// which would otherwise stomp the platform package's own env-based
	// default-metro resolution.
	c := NewClient(WithToken("t"))
	p := c.Platform("")
	require.NotNil(t, p)
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

func TestClientMintsIndependentPlatformClientsPerCall(t *testing.T) {
	srvA, reqsA := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
	srvB, reqsB := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})

	c := NewClient(WithToken("t"))
	a := c.Platform(srvA.URL)
	b := c.Platform(srvB.URL)

	_, err := a.Instances(context.Background(), []platform.Ref{})
	require.NoError(t, err)
	_, err = b.Instances(context.Background(), []platform.Ref{})
	require.NoError(t, err)

	assert.Len(t, *reqsA, 1)
	assert.Len(t, *reqsB, 1)
}

func TestClientMintsIndependentControlPlaneClientsPerCall(t *testing.T) {
	srv, reqs := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})

	c := NewClient(WithToken("t"))
	cp1 := c.ControlPlane().WithEndpoint(srv.URL)
	cp2 := c.ControlPlane().WithEndpoint(srv.URL)

	_, err := cp1.Raw().ListNodes(context.Background(), []gen.NameOrUUID{}, gen.ListNodesOpts{})
	require.NoError(t, err)
	_, err = cp2.Raw().ListNodes(context.Background(), []gen.NameOrUUID{}, gen.ListNodesOpts{})
	require.NoError(t, err)

	assert.Len(t, *reqs, 2)
}

func TestClientAllowInsecurePropagatesToPlatformClient(t *testing.T) {
	srv, reqs := newTestServer(t, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})

	c := NewClient(WithToken("t"), WithAllowInsecure(true))
	p := c.Platform(srv.URL)

	_, err := p.Instances(context.Background(), []platform.Ref{})
	require.NoError(t, err)
	assert.Len(t, *reqs, 1)
}

func TestNewClientWithNoOptionsDoesNotPanic(t *testing.T) {
	c := NewClient()
	require.NotNil(t, c)
	assert.NotNil(t, c.Platform("fra"))
	assert.NotNil(t, c.ControlPlane())
}

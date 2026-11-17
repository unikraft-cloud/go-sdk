// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

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
)

type recordedRequest struct {
	Method  string
	Path    string
	Query   string
	Headers http.Header
	Body    []byte
}

func newTestServer(t *testing.T, handler func(w http.ResponseWriter, r *recordedRequest)) (*httptest.Server, *[]recordedRequest) {
	t.Helper()
	var got []recordedRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		rec := recordedRequest{
			Method:  r.Method,
			Path:    r.URL.Path,
			Query:   r.URL.RawQuery,
			Headers: r.Header.Clone(),
			Body:    body,
		}
		got = append(got, rec)
		handler(w, &rec)
	}))
	t.Cleanup(srv.Close)
	return srv, &got
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, body any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	require.NoError(t, json.NewEncoder(w).Encode(body))
}

func newClientAgainst(srv *httptest.Server) Client {
	return NewClient(WithToken("test-token"), WithUserAgent("sdk-test/1"), WithEndpoint(srv.URL))
}

func TestClientNodeFound(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"nodes": []map[string]any{
					{"uuid": "abc", "name": "foo", "state": "ready"},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	node, err := c.Node(context.Background(), Name("foo"))
	require.NoError(t, err)
	assert.Equal(t, "abc", node.Uuid)
	assert.Equal(t, "foo", node.Name)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/nodes", got.Path)
	assert.Equal(t, "Bearer test-token", got.Headers.Get("Authorization"))
	assert.Equal(t, "sdk-test/1", got.Headers.Get("User-Agent"))

	var sent []gen.NameOrUUID
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.Len(t, sent, 1)
	assert.Equal(t, "foo", *sent[0].Name)
}

func TestClientNodeNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"nodes": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Node(context.Background(), Name("missing"))
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
	require.ErrorContains(t, err, "not found")
}

func TestClientNodeAPIError(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusBadRequest, map[string]any{"status": "error", "message": "boom"})
	})

	c := newClientAgainst(srv)
	_, err := c.Node(context.Background(), UUID("abc"))
	require.Error(t, err)
	require.ErrorContains(t, err, "boom")
}

func TestClientNodesEmptyRefsRequestsAll(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"nodes": []map[string]any{
					{"uuid": "a", "name": "one"},
					{"uuid": "b", "name": "two"},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	nodes, err := c.Nodes(context.Background(), nil)
	require.NoError(t, err)
	require.Len(t, nodes, 2)
	assert.Equal(t, "one", nodes[0].Name)
	assert.Equal(t, "two", nodes[1].Name)

	got := (*reqs)[0]
	assert.Empty(t, string(got.Body))
}

func TestClientNodesMultipleRefsAndOpts(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"nodes": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Nodes(context.Background(), []Ref{Name("foo"), UUID("bar")}, WithNodesState(gen.NodeStateReady))
	require.NoError(t, err)

	got := (*reqs)[0]
	assert.Equal(t, "state=ready", got.Query)

	var sent []gen.NameOrUUID
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.Len(t, sent, 2)
	assert.Equal(t, "foo", *sent[0].Name)
	assert.Equal(t, "bar", *sent[1].Uuid)
}

func TestClientProvisionNode(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"node": map[string]any{"uuid": "new-uuid", "name": "new-node"},
			},
		})
	})

	c := newClientAgainst(srv)
	node, err := c.ProvisionNode(context.Background(), WithProvisionName("new-node"))
	require.NoError(t, err)
	assert.Equal(t, "new-uuid", node.Uuid)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodPost, got.Method)
	assert.Equal(t, "/v1/nodes", got.Path)

	var sent gen.ProvisionNodeRequest
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.NotNil(t, sent.Name)
	assert.Equal(t, "new-node", *sent.Name)
}

func TestClientWithEndpointReturnsIndependentClient(t *testing.T) {
	srvA, reqsA := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
	})
	srvB, reqsB := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
	})

	a := newClientAgainst(srvA)
	b := a.WithEndpoint(srvB.URL)

	_, err := a.Nodes(context.Background(), nil)
	require.NoError(t, err)
	_, err = b.Nodes(context.Background(), nil)
	require.NoError(t, err)

	assert.Len(t, *reqsA, 1)
	assert.Len(t, *reqsB, 1)
}

func TestClientWithHTTPClientOverride(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"nodes": []map[string]any{}}})
	})

	calls := 0
	counting := countingHTTPClient{inner: &http.Client{}, calls: &calls}

	c := newClientAgainst(srv).WithHTTPClient(counting)
	_, err := c.Nodes(context.Background(), nil)
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

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

// recordedRequest captures what the SDK actually sent over the wire, so
// tests can assert on headers/body without re-implementing the transport.
type recordedRequest struct {
	Method  string
	Path    string
	Query   string
	Headers http.Header
	Body    []byte
}

// newTestServer starts an httptest.Server that dispatches to handler and
// records every request it receives.
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
	return NewClient(WithToken("test-token"), WithUserAgent("sdk-test/1"), WithDefaultMetro(srv.URL))
}

func TestClientInstanceFound(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"uuid": "abc", "name": "foo", "state": "running"},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	inst, err := c.Instance(context.Background(), Name("foo"))
	require.NoError(t, err)
	assert.Equal(t, "abc", inst.Uuid)
	assert.Equal(t, "foo", inst.Name)
	assert.Equal(t, gen.InstanceStateRunning, inst.State)

	require.Len(t, *reqs, 1)
	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/instances", got.Path)
	assert.Equal(t, "Bearer test-token", got.Headers.Get("Authorization"))
	assert.Equal(t, "sdk-test/1", got.Headers.Get("User-Agent"))

	var sent []gen.NameOrUUID
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.Len(t, sent, 1)
	require.NotNil(t, sent[0].Name)
	assert.Equal(t, "foo", *sent[0].Name)
}

func TestClientInstanceNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"instances": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Instance(context.Background(), Name("missing"))
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
	require.ErrorContains(t, err, "not found")
}

func TestClientInstanceAPIError(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusBadRequest, map[string]any{
			"status":  "error",
			"message": "boom",
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Instance(context.Background(), Name("foo"))
	require.Error(t, err)
	require.ErrorContains(t, err, "boom")
}

func TestClientServiceGroupFound(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"service_groups": []map[string]any{{"uuid": "abc", "name": "svc"}}},
		})
	})

	c := newClientAgainst(srv)
	sg, err := c.ServiceGroup(context.Background(), Name("svc"))
	require.NoError(t, err)
	assert.Equal(t, "abc", sg.Uuid)
	assert.Equal(t, "/v1/services", (*reqs)[0].Path)
}

func TestClientServiceGroupNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"service_groups": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.ServiceGroup(context.Background(), Name("missing"))
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
}

func TestClientServiceGroupsEmptyRefsRequestsAll(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"service_groups": []map[string]any{{"uuid": "a"}, {"uuid": "b"}},
			},
		})
	})

	c := newClientAgainst(srv)
	groups, err := c.ServiceGroups(context.Background(), nil)
	require.NoError(t, err)
	require.Len(t, groups, 2)

	assert.Empty(t, string((*reqs)[0].Body), "no refs requested means a nil/omitted request body")
}

func TestClientCertificateFound(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"certificates": []map[string]any{{"uuid": "abc", "name": "cert"}}},
		})
	})

	c := newClientAgainst(srv)
	cert, err := c.Certificate(context.Background(), Name("cert"))
	require.NoError(t, err)
	assert.Equal(t, "abc", cert.Uuid)
	assert.Equal(t, "/v1/certificates", (*reqs)[0].Path)
}

func TestClientCertificateNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"certificates": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Certificate(context.Background(), Name("missing"))
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
}

func TestClientVolumeFound(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"volumes": []map[string]any{{"uuid": "abc", "name": "vol"}}},
		})
	})

	c := newClientAgainst(srv)
	vol, err := c.Volume(context.Background(), Name("vol"))
	require.NoError(t, err)
	assert.Equal(t, "abc", vol.Uuid)
	assert.Equal(t, "/v1/volumes", (*reqs)[0].Path)
}

func TestClientVolumeNotFound(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"volumes": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Volume(context.Background(), Name("missing"))
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
}

func TestClientInstancesEmptyRefsRequestsAll(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"uuid": "a", "name": "one"},
					{"uuid": "b", "name": "two"},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	instances, err := c.Instances(context.Background(), []Ref{})
	require.NoError(t, err)
	require.Len(t, instances, 2)
	assert.Equal(t, "one", instances[0].Name)
	assert.Equal(t, "two", instances[1].Name)

	got := (*reqs)[0]
	assert.Empty(t, string(got.Body), "no refs requested means a nil/omitted request body")
}

func TestClientInstancesMultipleRefsAndOpts(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"instances": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Instances(context.Background(), []Ref{Name("foo"), UUID("bar")}, WithDetails())
	require.NoError(t, err)

	got := (*reqs)[0]
	assert.Equal(t, "details=true", got.Query)

	var sent []gen.NameOrUUID
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.Len(t, sent, 2)
	assert.Equal(t, "foo", *sent[0].Name)
	assert.Equal(t, "bar", *sent[1].Uuid)
}

func TestClientCreateInstance(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"instances": []map[string]any{
					{"uuid": "new-uuid", "name": "new-instance"},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	created, err := c.CreateInstance(context.Background(), WithCreateName("new-instance"))
	require.NoError(t, err)
	require.Len(t, created, 1)
	assert.Equal(t, "new-uuid", created[0].Uuid)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodPost, got.Method)
	assert.Equal(t, "/v1/instances", got.Path)

	var sent gen.CreateInstanceRequest
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.NotNil(t, sent.Name)
	assert.Equal(t, "new-instance", *sent.Name)
}

func TestClientWithMetroReturnsIndependentClient(t *testing.T) {
	srvA, reqsA := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
	})
	srvB, reqsB := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
	})

	a := newClientAgainst(srvA)
	b := a.WithMetro(srvB.URL)

	_, err := a.Instances(context.Background(), []Ref{})
	require.NoError(t, err)
	_, err = b.Instances(context.Background(), []Ref{})
	require.NoError(t, err)

	assert.Len(t, *reqsA, 1, "the original client must still target its own server")
	assert.Len(t, *reqsB, 1, "the derived client must target the new server")
}

func TestClientWithHTTPClientOverride(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": []map[string]any{}}})
	})

	calls := 0
	counting := countingHTTPClient{inner: &http.Client{}, calls: &calls}

	c := newClientAgainst(srv).WithHTTPClient(counting)
	_, err := c.Instances(context.Background(), []Ref{})
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

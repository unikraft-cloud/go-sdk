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

func TestCertificateRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"certificates": []map[string]any{
					{"uuid": "abc", "name": "cert", "common_name": "example.com"},
				},
			},
		})
	})

	cert := wrapCertificate(rawClientAgainst(srv), gen.Certificate{Uuid: "abc"})
	require.NoError(t, cert.Refresh(context.Background()))
	assert.Equal(t, "example.com", cert.CommonName)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/certificates", got.Path)
}

func TestCertificateDeleteDoesNotRefresh(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/v1/certificates/abc", r.Path)
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
	})

	cert := wrapCertificate(rawClientAgainst(srv), gen.Certificate{Uuid: "abc"})
	require.NoError(t, cert.Delete(context.Background()))
	assert.Len(t, *reqs, 1)
}

func TestCertificateUpdateReplacesChainAndKeyThenRefreshes(t *testing.T) {
	var putBody []byte
	var calls []string
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		calls = append(calls, r.Method+" "+r.Path)
		switch r.Method {
		case http.MethodPut:
			putBody = r.Body
			writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{}})
		case http.MethodGet:
			writeJSON(t, w, http.StatusOK, map[string]any{
				"status": "success",
				"data":   map[string]any{"certificates": []map[string]any{{"uuid": "abc"}}},
			})
		}
	})

	cert := wrapCertificate(rawClientAgainst(srv), gen.Certificate{Uuid: "abc"})
	require.NoError(t, cert.Update(context.Background(), "new-chain", "new-pkey"))
	assert.Equal(t, []string{"PUT /v1/certificates/abc", "GET /v1/certificates"}, calls)

	var sent gen.UpdateCertificateByUUIDRequestBody
	require.NoError(t, json.Unmarshal(putBody, &sent))
	assert.Equal(t, "new-chain", sent.Chain)
	assert.Equal(t, "new-pkey", sent.Pkey)
}

func TestClientCreateCertificate(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"certificates": []map[string]any{{"uuid": "abc", "name": "cert"}},
			},
		})
	})

	c := newClientAgainst(srv)
	cert, err := c.CreateCertificate(context.Background(),
		WithCreateCertificateName("cert"),
		WithCreateCertificateCommonName("example.com"),
		WithCreateCertificateChain("chain"),
		WithCreateCertificatePKey("pkey"),
	)
	require.NoError(t, err)
	assert.Equal(t, "abc", cert.Uuid)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodPost, got.Method)
	assert.Equal(t, "/v1/certificates", got.Path)

	var sent gen.CreateCertificateRequest
	require.NoError(t, json.Unmarshal(got.Body, &sent))
	require.NotNil(t, sent.CommonName)
	assert.Equal(t, "example.com", *sent.CommonName)
	assert.Equal(t, "chain", sent.Chain)
	assert.Equal(t, "pkey", sent.Pkey)
}

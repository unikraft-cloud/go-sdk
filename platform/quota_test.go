// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientQuotas(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data": map[string]any{
				"quotas": []map[string]any{
					{"uuid": "abc", "hard": map[string]any{"instances": 10}},
				},
			},
		})
	})

	c := newClientAgainst(srv)
	q, err := c.Quotas(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "abc", q.Uuid)
	assert.EqualValues(t, 10, q.Hard.Instances)

	got := (*reqs)[0]
	assert.Equal(t, http.MethodGet, got.Method)
	assert.Equal(t, "/v1/users/quotas", got.Path)
}

func TestClientQuotasForUser(t *testing.T) {
	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"quotas": []map[string]any{{"uuid": "def"}}},
		})
	})

	c := newClientAgainst(srv)
	q, err := c.QuotasForUser(context.Background(), "def")
	require.NoError(t, err)
	assert.Equal(t, "def", q.Uuid)

	got := (*reqs)[0]
	assert.Equal(t, "/v1/users/def/quotas", got.Path)
}

func TestClientQuotasErrorsWhenEmpty(t *testing.T) {
	srv, _ := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		writeJSON(t, w, http.StatusOK, map[string]any{
			"status": "success",
			"data":   map[string]any{"quotas": []map[string]any{}},
		})
	})

	c := newClientAgainst(srv)
	_, err := c.Quotas(context.Background())
	require.Error(t, err)
}

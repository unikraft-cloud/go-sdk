// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func metrosResponse(metros ...map[string]any) map[string]any {
	return map[string]any{"status": "success", "data": map[string]any{"metros": metros}}
}

func TestAvailableMetros(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, metrosResponse(
		map[string]any{"name": "Frankfurt", "iata_code": "FRA", "endpoint": "https://api.fra.example.com/v1/"},
		map[string]any{"name": "Sin", "endpoint": "sin"},
		map[string]any{"name": ""},
	))

	c := NewClient(WithToken("t"), WithControlPlaneEndpoint(srv.URL))
	want := []Metro{
		{Code: "fra", Endpoint: "https://api.fra.example.com"},
		{Code: "sin", Endpoint: "https://api.sin.unikraft.cloud"},
	}

	for range 2 {
		metros, err := c.AvailableMetros(t.Context())
		require.NoError(t, err)
		assert.Equal(t, want, metros)
	}
	require.Len(t, *reqs, 1, "a successful answer is kept")
	assert.Equal(t, "/v1/metros", (*reqs)[0].Path)
}

func TestAvailableMetrosNone(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, metrosResponse())
	c := NewClient(WithToken("t"), WithControlPlaneEndpoint(srv.URL))

	for range 2 {
		_, err := c.AvailableMetros(t.Context())
		require.Error(t, err)
	}
	assert.Len(t, *reqs, 2, "a failed answer is not kept")
}

func TestMetros(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		codes     []string
		wantNames []string
		wantCalls int
	}{
		{name: "named", codes: []string{"was", "fra"}, wantNames: []string{"fra", "was"}},
		{name: "all available", wantNames: []string{"dal", "fra"}, wantCalls: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, metrosResponse(
				map[string]any{"iata_code": "FRA"},
				map[string]any{"iata_code": "DAL"},
			))
			c := NewClient(WithToken("t"), WithControlPlaneEndpoint(srv.URL))

			g, err := c.Metros(t.Context(), tt.codes...)
			require.NoError(t, err)
			assert.ElementsMatch(t, tt.wantNames, g.Names())
			assert.Len(t, *reqs, tt.wantCalls)
		})
	}
}

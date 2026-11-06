// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/secrets"
	"github.com/moby/buildkit/session/secrets/secretsprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/build"
)

func TestAttach(t *testing.T) {
	t.Parallel()

	t.Run("serves the session to the build", func(t *testing.T) {
		t.Parallel()

		ctx := testContext(t)

		sm, err := session.NewManager()
		require.NoError(t, err)
		srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet || r.URL.Path != "/v1/builds/b1/session" || r.Header.Get("Authorization") != "Bearer tok" {
				http.Error(w, "unexpected request", http.StatusBadRequest)
				return
			}
			ws, err := websocket.Accept(w, r, nil)
			if err != nil {
				return
			}
			_ = sm.HandleConn(r.Context(), websocket.NetConn(r.Context(), ws, websocket.MessageBinary), r.Header)
		}))
		// The server offers HTTP/2, as a load balancer does.
		srv.EnableHTTP2 = true
		srv.StartTLS()
		t.Cleanup(srv.Close)

		hc := srv.Client()
		hc.Timeout = time.Minute
		c, err := build.New(srv.URL, build.WithToken("tok"), build.WithHTTPClient(hc))
		require.NoError(t, err)

		s, err := session.NewSession(ctx, "app")
		require.NoError(t, err)
		s.Allow(secretsprovider.FromMap(map[string][]byte{"npmrc": []byte("hunter2")}))

		require.NoError(t, c.Attach(ctx, "b1", s))
		t.Cleanup(func() { _ = s.Close() })

		caller, err := sm.Get(ctx, s.ID(), false)
		require.NoError(t, err)
		assert.Equal(t, "app", caller.SharedKey())

		secret, err := secrets.GetSecret(ctx, caller, "npmrc")
		require.NoError(t, err)
		assert.Equal(t, "hunter2", string(secret))
	})

	t.Run("a refused session leaves the build alone", func(t *testing.T) {
		t.Parallel()

		ctx := testContext(t)

		requests := make(chan request, 2)
		c := serve(t, requests, answer{status: http.StatusConflict, body: `{"error": {"code": "conflict", "message": "the build already has a session"}}`})

		s, err := session.NewSession(ctx, "")
		require.NoError(t, err)

		err = c.Attach(ctx, "b1", s)
		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, build.ErrorCodeConflict, apiErr.Code)

		<-requests // the attach
		assert.Empty(t, requests)
	})
}

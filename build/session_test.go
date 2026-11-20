// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
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

	t.Run("a session whose connection drops attaches again with the same ID", func(t *testing.T) {
		t.Parallel()
		ctx := testContext(t)

		sm, err := session.NewManager()
		require.NoError(t, err)
		var mu sync.Mutex
		var ids []string // of the session, by dial
		var gets atomic.Int32
		drop, dropped := make(chan struct{}), make(chan struct{})
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/builds/b1" {
				w.Header().Set("Content-Type", "application/json")
				if gets.Add(1) == 1 { // a replica that stops
					w.WriteHeader(http.StatusServiceUnavailable)
					_, _ = w.Write([]byte(`{"error": {"code": "internal", "message": "the service is shutting down"}}`))
					return
				}
				_, _ = w.Write([]byte(`{"id": "b1", "status": "running", "created_at": "2026-10-09T12:00:00Z"}`))
				return
			}
			mu.Lock()
			ids = append(ids, r.Header.Get("X-Docker-Expose-Session-Uuid"))
			dial := len(ids)
			mu.Unlock()
			if dial == 2 || dial == 3 { // the service has not seen the drop yet
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte(`{"error": {"code": "conflict", "message": "the build does not wait for a session: it has one already"}}`))
				return
			}
			ws, err := websocket.Accept(w, r, nil)
			if err != nil {
				return
			}
			hctx := r.Context()
			if dial == 1 {
				var cancel context.CancelFunc
				hctx, cancel = context.WithCancel(hctx)
				go func() { <-drop; cancel() }()
				defer close(dropped) // once HandleConn has removed the session
			}
			_ = sm.HandleConn(hctx, websocket.NetConn(hctx, ws, websocket.MessageBinary), r.Header)
		}))
		t.Cleanup(srv.Close)
		c, err := build.New(srv.URL, build.WithToken("tok"), build.WithRetry(time.Millisecond, 10*time.Millisecond))
		require.NoError(t, err)

		s, err := session.NewSession(ctx, "app")
		require.NoError(t, err)
		s.Allow(secretsprovider.FromMap(map[string][]byte{"npmrc": []byte("hunter2")}))
		t.Cleanup(func() { _ = s.Close() })
		require.NoError(t, c.Attach(ctx, "b1", s))
		_, err = sm.Get(ctx, s.ID(), false)
		require.NoError(t, err)

		close(drop)
		<-dropped
		caller, err := sm.Get(ctx, s.ID(), false)
		require.NoError(t, err)
		secret, err := secrets.GetSecret(ctx, caller, "npmrc")
		require.NoError(t, err)
		assert.Equal(t, "hunter2", string(secret))
		mu.Lock()
		defer mu.Unlock()
		assert.Equal(t, []string{s.ID(), s.ID(), s.ID(), s.ID()}, ids)
	})

	t.Run("a session is not attached again once its build has ended", func(t *testing.T) {
		t.Parallel()
		ctx := testContext(t)

		sm, err := session.NewManager()
		require.NoError(t, err)
		var dials atomic.Int32
		asked := make(chan struct{}, 1)
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.URL.Path == "/v1/builds/b1":
				_, _ = w.Write([]byte(`{"id": "b1", "status": "failed", "created_at": "2026-10-09T12:00:00Z"}`))
				select {
				case asked <- struct{}{}:
				default:
				}
			case dials.Add(1) == 1:
				ws, err := websocket.Accept(w, r, nil)
				if err != nil {
					return
				}
				// The build ends, thus the service ends the session.
				hctx, cancel := context.WithCancel(r.Context())
				cancel()
				_ = sm.HandleConn(hctx, websocket.NetConn(r.Context(), ws, websocket.MessageBinary), r.Header)
			default:
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte(`{"error": {"code": "conflict", "message": "the build does not wait for a session: it has finished"}}`))
			}
		}))
		t.Cleanup(srv.Close)
		c, err := build.New(srv.URL, build.WithToken("tok"), build.WithRetry(time.Millisecond, 10*time.Millisecond))
		require.NoError(t, err)

		s, err := session.NewSession(ctx, "")
		require.NoError(t, err)
		require.NoError(t, c.Attach(ctx, "b1", s))

		select {
		case <-asked:
		case <-ctx.Done():
			t.Fatal("the client did not ask for the build")
		}
		// A client that dials again does so within its longest wait, 10 ms.
		<-time.After(100 * time.Millisecond)
		assert.EqualValues(t, 2, dials.Load(), "the client dialed again after the build ended")
		require.NoError(t, s.Close())
	})
}

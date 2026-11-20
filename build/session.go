// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

//go:build !js

package build

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"math/rand/v2"
	"net"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/moby/buildkit/session"
)

const (
	// reattachBudget is how many times in a row the session dials again after
	// its connection ended, before it gives up.
	reattachBudget = 30

	// dialTimeout bounds the handshake of a dial. Session.Run holds the lock of
	// Close while it dials, thus Close waits at most this long.
	dialTimeout = 10 * time.Second
)

// Attach attaches s to the waiting build id; s then serves it until the build, ctx or s ends.
// When its connection ends while the build runs, s dials again, thus close s once the build ends.
func (c *Client) Attach(ctx context.Context, id string, s *session.Session) error {
	idSegment, err := pathSegment("id", id)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/builds/%s/session", idSegment)

	attached := make(chan error, 1)
	go c.serveSession(ctx, id, path, s, attached)

	if err := <-attached; err != nil {
		return fmt.Errorf("attaching the session: %w", err)
	}

	return nil
}

// serveSession runs s over a connection to path, and dials again each time the connection
// ends, until ctx or s ends or the build no longer runs. first gets the result of the first dial.
func (c *Client) serveSession(ctx context.Context, id, path string, s *session.Session, first chan<- error) {
	var failed int
	for n := 0; ; n++ {
		var dialed bool
		var dialErr error
		start := time.Now()
		_ = s.Run(ctx, func(ctx context.Context, _ string, meta map[string][]string) (net.Conn, error) {
			dialed = true
			conn, err := c.dial(ctx, path, meta)
			dialErr = err
			if n == 0 {
				first <- err
			}
			return conn, err
		})

		switch {
		case n == 0 && !dialed:
			// Run returns without dialing when s is already closed.
			first <- errors.New("the session is closed")
			return
		case n == 0 && dialErr != nil, !dialed, ctx.Err() != nil:
			return
		case dialErr == nil && time.Since(start) >= c.retryMax:
			// A connection that served for a while resets the budget, one that ends at once does not.
			failed = 0
		case dialErr != nil && !c.redialable(ctx, id, dialErr):
			return
		}
		if failed++; failed > reattachBudget {
			return
		}

		wait := c.retryMax
		if failed < 16 {
			wait = min(c.retryMin<<failed, c.retryMax)
		}
		// The sessions of a replica that stops drop at once, thus they dial again at other times.
		wait = wait/2 + rand.N(wait/2+1)
		select {
		case <-time.After(wait):
		case <-ctx.Done():
			return
		}
	}
}

// redialable tells whether the session dials again after the dial failed with err. The service
// answers 409 also once the build has ended, thus the state of the build decides.
func (c *Client) redialable(ctx context.Context, id string, err error) bool {
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusConflict {
		return transient(err)
	}
	// A lookup that stalls fails at the timeout, thus the dials go on.
	lctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()
	b, err := c.GetBuild(lctx, id)
	if err != nil {
		return transient(err)
	}
	return b.Status == BuildStatusWaiting || b.Status == BuildStatusRunning
}

// transient tells whether err can pass: the service could not be reached, or
// it answered 429 or 5xx.
func transient(err error) bool {
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		return true
	}
	return apiErr.StatusCode == http.StatusTooManyRequests || apiErr.StatusCode >= http.StatusInternalServerError
}

// dial opens a WebSocket to path, whose binary messages carry the session. Go
// dials a WebSocket over HTTP/1.1, also when the server offers HTTP/2.
func (c *Client) dial(ctx context.Context, path string, header map[string][]string) (net.Conn, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	maps.Copy(req.Header, header)

	// The timeouts apply only to the handshake, thus the session can go on.
	hctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()
	ws, resp, err := websocket.Dial(hctx, req.URL.String(), &websocket.DialOptions{HTTPClient: c.hc, HTTPHeader: req.Header})
	switch {
	case err == nil:
		return websocket.NetConn(ctx, ws, websocket.MessageBinary), nil
	case resp == nil:
		return nil, err
	case resp.StatusCode >= http.StatusBadRequest:
		return nil, c.readError(resp)
	case resp.StatusCode >= http.StatusMultipleChoices:
		return nil, redirected(resp)
	default:
		return nil, err
	}
}

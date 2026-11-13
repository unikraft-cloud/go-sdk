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
	"net"
	"net/http"

	"github.com/coder/websocket"
	"github.com/moby/buildkit/session"
)

// Attach attaches s to the waiting build id; s then serves it until the build, ctx or s ends.
func (c *Client) Attach(ctx context.Context, id string, s *session.Session) error {
	idSegment, err := pathSegment("id", id)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("/v1/builds/%s/session", idSegment)

	attached := make(chan error, 1)
	go func() {
		_ = s.Run(ctx, func(ctx context.Context, _ string, meta map[string][]string) (net.Conn, error) {
			conn, err := c.dial(ctx, path, meta)
			attached <- err
			return conn, err
		})
		// Run returns without dialing when s is already closed.
		select {
		case attached <- errors.New("the session is closed"):
		default:
		}
	}()

	if err := <-attached; err != nil {
		return fmt.Errorf("attaching the session: %w", err)
	}

	return nil
}

// dial opens a WebSocket to path, whose binary messages carry the session. Go
// dials a WebSocket over HTTP/1.1, also when the server offers HTTP/2.
func (c *Client) dial(ctx context.Context, path string, header map[string][]string) (net.Conn, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	maps.Copy(req.Header, header)

	// The timeout of the client applies only to the handshake, thus the session can go on.
	ws, resp, err := websocket.Dial(ctx, req.URL.String(), &websocket.DialOptions{HTTPClient: c.hc, HTTPHeader: req.Header})
	switch {
	case err == nil:
		return websocket.NetConn(ctx, ws, websocket.MessageBinary), nil
	case resp == nil:
		return nil, err
	case resp.StatusCode >= http.StatusBadRequest:
		return nil, readError(resp)
	case resp.StatusCode >= http.StatusMultipleChoices:
		return nil, redirected(resp)
	default:
		return nil, err
	}
}

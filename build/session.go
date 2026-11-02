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
	"io"
	"maps"
	"net"
	"net/http"
	"net/http/httptrace"

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
		_ = s.Run(ctx, func(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error) {
			conn, err := c.upgrade(ctx, path, proto, meta)
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

// upgrade switches a POST to path over to proto.
// shortcut: HTTP/2 cannot switch, UKC's edge speaks HTTP/1.1; force HTTP/1.1 if a server negotiates h2.
func (c *Client) upgrade(ctx context.Context, path, proto string, header map[string][]string) (net.Conn, error) {
	var conn net.Conn
	ctx = httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { conn = info.Conn },
	})

	req, err := c.newRequest(ctx, http.MethodPost, path, nil, nil)
	if err != nil {
		return nil, err
	}
	maps.Copy(req.Header, header)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", proto)

	// A timeout would end the session, and net/http's timeout body cannot be written.
	hc := *c.hc
	hc.Timeout = 0
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	switch {
	case resp.StatusCode == http.StatusSwitchingProtocols:
	case resp.StatusCode >= http.StatusBadRequest:
		defer resp.Body.Close()
		return nil, readError(resp)
	case resp.StatusCode >= http.StatusMultipleChoices:
		defer resp.Body.Close()
		return nil, redirected(resp)
	default:
		defer resp.Body.Close()
		return nil, fmt.Errorf("the service answered %s instead of switching to %s", resp.Status, proto)
	}

	// net/http hands out the switched connection as the body.
	rwc, ok := resp.Body.(io.ReadWriteCloser)
	if !ok || conn == nil {
		resp.Body.Close()
		return nil, fmt.Errorf("the connection that switched to %s is not writable", proto)
	}

	return upgradedConn{Conn: conn, rwc: rwc}, nil
}

// upgradedConn reads through the 101's body, which holds what net/http read ahead.
type upgradedConn struct {
	net.Conn
	rwc io.ReadWriteCloser
}

func (c upgradedConn) Read(p []byte) (int, error)  { return c.rwc.Read(p) }
func (c upgradedConn) Write(p []byte) (int, error) { return c.rwc.Write(p) }
func (c upgradedConn) Close() error                { return c.rwc.Close() }

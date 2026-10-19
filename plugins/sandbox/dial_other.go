// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

//go:build !js

package sandbox

import (
	"context"
	"net/http"

	"github.com/coder/websocket"
)

// errNoFollow is why no command can be followed from here; nil where one can.
var errNoFollow error

func dial(ctx context.Context, req *http.Request, client *http.Client) (*websocket.Conn, *http.Response, error) {
	return websocket.Dial(ctx, req.URL.String(), &websocket.DialOptions{HTTPClient: client, HTTPHeader: req.Header})
}

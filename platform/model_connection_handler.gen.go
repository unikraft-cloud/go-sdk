// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Connection handlers to use for the service.  Handlers define how the service
// will handle incoming connections and forward traffic from the Internet to
// your application.  For example, a service can be configured to terminate TLS
// connections, redirect HTTP traffic, or enable HTTP mode for load balancing.
// You configure the handlers for every published service port individually.
//
// There are currently 3 supported handlers:
//
// | Handler    | Description |
// |------------|-------------|
// | `tls`      | Terminate the TLS connection at the Unikraft Cloud gateway using our wildcard certificate issued for the kraft.cloud domain. The gateway forwards the unencrypted traffic to your application. |
// | `http`     | Enable HTTP mode on the load balancer to load balance on the level of individual HTTP requests. In this mode, only HTTP connections are accepted. If this option is not set the load balancer works in TCP mode and distributes TCP connections. |
// | `redirect` | Redirect traffic from the source port to the destination port. |
//
// Note that there is a set of constraints when publishing ports:
//   - Port 80: MUST have "http" and MUST not have "tls" set;
//   - Port 443: MUST have http and tls set;
//   - The `redirect` handler can only be set on port 80 (HTTP) to redirect to
//     port 443 (HTTPS);
//   - All other ports MUST have tls and MUST not have http set.
type ConnectionHandler string

const (
	ConnectionHandlerTLS      ConnectionHandler = "tls"
	ConnectionHandlerHTTP     ConnectionHandler = "http"
	ConnectionHandlerREDIRECT ConnectionHandler = "redirect"
)

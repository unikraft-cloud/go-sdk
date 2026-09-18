// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package httpclient provides an HTTP client interface and default
// implementations.
package httpclient

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// HTTPClient interface abstracts a generic HTTP request issuing client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Option configures the HTTP client returned by NewHTTPClient.
type Option func(*options)

type options struct {
	insecure  bool
	userAgent string
	timeout   time.Duration
	transport *http.Transport
}

// WithInsecure configures the client to skip TLS certificate verification.
func WithInsecure() Option {
	return func(o *options) {
		o.insecure = true
	}
}

// WithUserAgent configures the client to set the given User-Agent header on
// outgoing requests that do not already set one.
func WithUserAgent(ua string) Option {
	return func(o *options) {
		o.userAgent = ua
	}
}

// WithTimeout sets a deadline on the full exchange, bodies included.
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

// WithTransport configures the client to use transport instead of the defaults
// from http.DefaultTransport, which is how a caller sets its own connection
// pool limits.  A clone is taken, so the returned client owns its pool and
// WithInsecure does not modify the given transport.
func WithTransport(transport *http.Transport) Option {
	return func(o *options) {
		o.transport = transport
	}
}

// NewHTTPClient creates a default Go HTTP client.
func NewHTTPClient(opts ...Option) *http.Client {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	base := o.transport
	if base == nil {
		// TODO: move these hardcoded values to configurable env variables.
		base = http.DefaultTransport.(*http.Transport).Clone()
		// Stdlib default is 30s; fail faster on an unreachable peer.
		base.DialContext = (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext
		// Stdlib default is 100; this client talks to many hosts at once.
		base.MaxIdleConns = 500
		// Stdlib default is 2; keep more idle connections per host open.
		base.MaxIdleConnsPerHost = 100
		// Stops once headers arrive, so a body of any size can stream.
		base.ResponseHeaderTimeout = 30 * time.Second
	}

	transport := base.Clone()
	if o.insecure {
		if transport.TLSClientConfig == nil {
			transport.TLSClientConfig = &tls.Config{}
		}
		transport.TLSClientConfig.InsecureSkipVerify = true //nolint:gosec // explicit opt-in via WithInsecure
	}

	var rt http.RoundTripper = transport
	if o.userAgent != "" {
		rt = &userAgentTransport{
			userAgent: o.userAgent,
			next:      transport,
		}
	}

	return &http.Client{
		Transport: rt,
		Timeout:   o.timeout,
	}
}

// NewInsecureHTTPClient creates a default Go HTTP client with insecure checks
// skipped.
//
// Deprecated: use NewHTTPClient(WithInsecure()) instead.
func NewInsecureHTTPClient(opts ...Option) *http.Client {
	return NewHTTPClient(append(opts, WithInsecure())...)
}

// userAgentTransport sets a default User-Agent header on outgoing requests
// that do not already have one set.
type userAgentTransport struct {
	userAgent string
	next      http.RoundTripper
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("User-Agent") == "" {
		req = req.Clone(req.Context())
		req.Header.Set("User-Agent", t.userAgent)
	}
	return t.next.RoundTrip(req)
}

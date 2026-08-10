// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package httpclient provides an HTTP client interface and default
// implementations.
package httpclient

import (
	"crypto/tls"
	"net/http"
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

// NewHTTPClient creates a default Go HTTP client.
func NewHTTPClient(opts ...Option) *http.Client {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
	if o.insecure {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true, // Allow insecure connections
		}
	}

	var rt http.RoundTripper = transport
	if o.userAgent != "" {
		rt = &userAgentTransport{
			userAgent: o.userAgent,
			next:      transport,
		}
	}

	return &http.Client{Transport: rt}
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

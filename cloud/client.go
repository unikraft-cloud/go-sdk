// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package cloud holds shared auth/config and mints scoped platform.Client
// and controlplane.Client instances.
package cloud

import (
	"sync"

	"unikraft.com/cloud/sdk/controlplane"
	"unikraft.com/cloud/sdk/pkg/httpclient"
	"unikraft.com/cloud/sdk/platform"
)

// ClientOption is an option function used during initialization of a Client.
type ClientOption func(*clientConfig)

type clientConfig struct {
	token         string
	httpClient    httpclient.HTTPClient
	userAgent     string
	allowInsecure bool
	controlPlane  string
}

// WithToken sets the access token shared by clients minted from this Client.
func WithToken(token string) ClientOption {
	return func(c *clientConfig) { c.token = token }
}

// WithHTTPClient sets the HTTP client shared by clients minted from this Client.
func WithHTTPClient(hc httpclient.HTTPClient) ClientOption {
	return func(c *clientConfig) { c.httpClient = hc }
}

// WithUserAgent sets the user agent string shared by clients minted from this Client.
func WithUserAgent(ua string) ClientOption {
	return func(c *clientConfig) { c.userAgent = ua }
}

// WithAllowInsecure configures clients minted from this Client to skip
// checks of the API certificates.
func WithAllowInsecure(allow bool) ClientOption {
	return func(c *clientConfig) { c.allowInsecure = allow }
}

// WithControlPlaneEndpoint sets the URL of the control plane API.
func WithControlPlaneEndpoint(endpoint string) ClientOption {
	return func(c *clientConfig) { c.controlPlane = endpoint }
}

// Client holds auth/config shared across the platform and control plane
// APIs, and mints scoped clients for each.
type Client struct {
	cfg clientConfig

	metrosMu sync.Mutex
	metros   []Metro
}

// NewClient creates a new Client.
func NewClient(opts ...ClientOption) *Client {
	var cfg clientConfig
	for _, o := range opts {
		o(&cfg)
	}
	return &Client{cfg: cfg}
}

// Platform mints a platform.Client scoped to metro, sharing this Client's
// token/http-client/user-agent/allow-insecure settings. Pass an empty metro
// to use the platform package's own default-metro resolution.
func (c *Client) Platform(metro string) platform.Client {
	var opts []platform.ClientOption
	if c.cfg.token != "" {
		opts = append(opts, platform.WithToken(c.cfg.token))
	}
	if c.cfg.userAgent != "" {
		opts = append(opts, platform.WithUserAgent(c.cfg.userAgent))
	}
	if c.cfg.allowInsecure {
		opts = append(opts, platform.WithAllowInsecure(c.cfg.allowInsecure))
	}
	if c.cfg.httpClient != nil {
		opts = append(opts, platform.WithHTTPClient(c.cfg.httpClient))
	}
	if metro != "" {
		opts = append(opts, platform.WithDefaultMetro(metro))
	}
	return platform.NewClient(opts...)
}

// ControlPlane mints a controlplane.Client sharing this Client's
// token/http-client/user-agent/allow-insecure settings.
func (c *Client) ControlPlane() controlplane.Client {
	var opts []controlplane.ClientOption
	if c.cfg.token != "" {
		opts = append(opts, controlplane.WithToken(c.cfg.token))
	}
	if c.cfg.userAgent != "" {
		opts = append(opts, controlplane.WithUserAgent(c.cfg.userAgent))
	}
	if c.cfg.allowInsecure {
		opts = append(opts, controlplane.WithAllowInsecure(c.cfg.allowInsecure))
	}
	if c.cfg.httpClient != nil {
		opts = append(opts, controlplane.WithHTTPClient(c.cfg.httpClient))
	}
	if c.cfg.controlPlane != "" {
		opts = append(opts, controlplane.WithEndpoint(c.cfg.controlPlane))
	}
	return controlplane.NewClient(opts...)
}

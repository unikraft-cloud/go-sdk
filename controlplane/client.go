// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"context"
	"fmt"
	"time"

	"unikraft.com/cloud/sdk/controlplane/gen"
	"unikraft.com/cloud/sdk/pkg/httpclient"
)

// Client is a façade over the control plane API, exposing resource
// accessors that return rich objects, plus Raw for everything else.
type Client interface {
	// Raw returns the underlying generated client, for API surface not
	// (yet) covered by Client's rich accessors.
	Raw() gen.Client

	Node(ctx context.Context, ref Ref) (*Node, error)
	// Nodes returns all nodes if refs is empty.
	Nodes(ctx context.Context, refs []Ref, opts ...NodesOption) ([]*Node, error)
	// ProvisionNode provisions a new node, returning the created node.
	ProvisionNode(ctx context.Context, opts ...ProvisionNodeOption) (*Node, error)

	// FindImage returns the registry record of an image, or nil if the
	// registry does not have it.
	FindImage(ctx context.Context, ref string) (*RegistryImage, error)
	// ImageExists reports whether the registry has an image.
	ImageExists(ctx context.Context, ref string) (bool, error)

	WithEndpoint(endpoint string) Client
	WithHTTPClient(hc httpclient.HTTPClient) Client
	WithTimeout(to time.Duration) Client
}

// ClientOption is an option function used during initialization of a Client.
type ClientOption func(*clientConfig)

type clientConfig struct {
	token         string
	endpoint      string
	httpClient    httpclient.HTTPClient
	userAgent     string
	allowInsecure bool
}

// WithToken sets the access token used for authentication with the API.
func WithToken(token string) ClientOption {
	return func(c *clientConfig) { c.token = token }
}

// WithEndpoint sets the endpoint used when connecting to the API.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *clientConfig) { c.endpoint = endpoint }
}

// WithHTTPClient sets the HTTP client used when connecting to the API.
func WithHTTPClient(hc httpclient.HTTPClient) ClientOption {
	return func(c *clientConfig) { c.httpClient = hc }
}

// WithUserAgent sets the user agent string used for API requests.
func WithUserAgent(ua string) ClientOption {
	return func(c *clientConfig) { c.userAgent = ua }
}

// WithAllowInsecure configures the client to skip checks of the API certificates.
func WithAllowInsecure(allow bool) ClientOption {
	return func(c *clientConfig) { c.allowInsecure = allow }
}

// NewClient creates a new Client for the control plane API.
func NewClient(opts ...ClientOption) Client {
	var cfg clientConfig
	for _, o := range opts {
		o(&cfg)
	}

	var genOpts []gen.ClientOption
	if cfg.token != "" {
		genOpts = append(genOpts, gen.WithToken(cfg.token))
	}
	if cfg.endpoint != "" {
		genOpts = append(genOpts, gen.WithDefaultEndpoint(cfg.endpoint))
	}
	if cfg.httpClient != nil {
		genOpts = append(genOpts, gen.WithHTTPClient(cfg.httpClient))
	}
	if cfg.userAgent != "" {
		genOpts = append(genOpts, gen.WithUserAgent(cfg.userAgent))
	}
	if cfg.allowInsecure {
		genOpts = append(genOpts, gen.WithAllowInsecure(cfg.allowInsecure))
	}

	return &controlplaneClient{raw: gen.NewClientFromEnv(genOpts...)}
}

type controlplaneClient struct {
	raw gen.Client
}

func (c *controlplaneClient) Raw() gen.Client { return c.raw }

func (c *controlplaneClient) Node(ctx context.Context, ref Ref) (*Node, error) {
	resp, err := c.raw.ListNodes(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.ListNodesOpts{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Nodes) == 0 {
		return nil, fmt.Errorf("node %s not found", ref)
	}
	return wrapNode(c.raw, resp.Data.Nodes[0]), nil
}

func (c *controlplaneClient) Nodes(ctx context.Context, refs []Ref, opts ...NodesOption) ([]*Node, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.ListNodesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.ListNodes(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	nodes := make([]*Node, len(resp.Data.Nodes))
	for i, data := range resp.Data.Nodes {
		nodes[i] = wrapNode(c.raw, data)
	}
	return nodes, nil
}

func (c *controlplaneClient) ProvisionNode(ctx context.Context, opts ...ProvisionNodeOption) (*Node, error) {
	var request gen.ProvisionNodeRequest
	for _, opt := range opts {
		opt(&request)
	}
	resp, err := c.raw.ProvisionNode(ctx, request)
	if err != nil {
		return nil, err
	}
	return wrapNode(c.raw, resp.Data.Node), nil
}

func (c *controlplaneClient) WithEndpoint(endpoint string) Client {
	return &controlplaneClient{raw: c.raw.WithEndpoint(endpoint)}
}

func (c *controlplaneClient) WithHTTPClient(hc httpclient.HTTPClient) Client {
	return &controlplaneClient{raw: c.raw.WithHTTPClient(hc)}
}

func (c *controlplaneClient) WithTimeout(to time.Duration) Client {
	return &controlplaneClient{raw: c.raw.WithTimeout(to)}
}

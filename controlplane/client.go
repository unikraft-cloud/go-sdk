// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"unikraft.com/cloud/sdk/pkg/httpclient"
)

type Client interface {
	// Check the authorization status of a request.  The responses are streamed
	// back to the client, indicating whether the signin request is authorized or
	// not.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/auth/check
	//
	// See: https://unikraft.com/docs/api/platform/v1/auth#check-authorization
	CheckAuthorization(ctx context.Context, request CheckAuthorizationRequest, ropts ...RequestOption) (<-chan *Response[CheckAuthorizationResponseData], error)
	// Initiate the sign-in process and return an authorization URL.  The user
	// should be redirected to this URL to complete the sign-in.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/auth/signin
	//
	// See: https://unikraft.com/docs/api/platform/v1/auth#request-signin
	RequestSignin(ctx context.Context, request RequestSigninRequest, ropts ...RequestOption) (*Response[RequestSigninResponseData], error)
	// Activates a new node.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/node/activate
	//
	// See: https://unikraft.com/docs/api/platform/v1/node#node-activate
	NodeActivate(ctx context.Context, request NodeActivateRequest, ropts ...RequestOption) (*Response[NodeActivateResponseData], error)
	// Renews a node's license.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/node/renew
	//
	// See: https://unikraft.com/docs/api/platform/v1/node#node-renew
	NodeRenew(ctx context.Context, request NodeRenewRequest, ropts ...RequestOption) (*Response[NodeRenewResponseData], error)
	// WithEndpoint sets the endpoint to use when connecting to the API.
	WithEndpoint(string) Client
	// WithTimeout sets the timeout when making the request.
	WithTimeout(time.Duration) Client
	// WithHTTPClient overwrites the base HTTP client.
	WithHTTPClient(httpclient.HTTPClient) Client
}

// NewClient creates a new client for the API.
func NewClient(copts ...ClientOption) Client {
	options := ClientOptions{}

	for _, opt := range copts {
		opt(&options)
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UKC_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UNIKRAFT_CLOUD_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("KRAFTCLOUD_TOKEN"))
	}

	if options.DefaultEndpoint() == "" {
		options.SetDefaultEndpoint(DefaultEndpoint)
	}

	if options.AllowInsecure() && options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewInsecureHTTPClient())
	}

	if options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewHTTPClient())
	}

	return &client{
		request: &Request{
			copts: &options,
		},
	}
}

type client struct {
	request *Request
}

// WithEndpoint sets the endpoint to use when connecting to the API.
func (c *client) WithEndpoint(m string) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithEndpoint(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *client) WithHTTPClient(hc httpclient.HTTPClient) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *client) WithTimeout(to time.Duration) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *client) clone() *client {
	ccpy := *c
	return &ccpy
}

func (c *client) CheckAuthorization(ctx context.Context, request CheckAuthorizationRequest, ropts ...RequestOption) (<-chan *Response[CheckAuthorizationResponseData], error) {
	requestPath := "/v1/auth/check"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CheckAuthorizationResponseData]{}
	if err := doRequest[CheckAuthorizationResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}
	return resp.Events()
}

func (c *client) RequestSignin(ctx context.Context, request RequestSigninRequest, ropts ...RequestOption) (*Response[RequestSigninResponseData], error) {
	requestPath := "/v1/auth/signin"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[RequestSigninResponseData]{}
	if err := doRequest[RequestSigninResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) NodeActivate(ctx context.Context, request NodeActivateRequest, ropts ...RequestOption) (*Response[NodeActivateResponseData], error) {
	requestPath := "/v1/node/activate"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[NodeActivateResponseData]{}
	if err := doRequest[NodeActivateResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) NodeRenew(ctx context.Context, request NodeRenewRequest, ropts ...RequestOption) (*Response[NodeRenewResponseData], error) {
	requestPath := "/v1/node/renew"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[NodeRenewResponseData]{}
	if err := doRequest[NodeRenewResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

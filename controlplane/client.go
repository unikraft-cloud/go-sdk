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

	"github.com/unikraft-cloud/go-sdk/pkg/httpclient"
)

type Client interface {
	// CheckAuthorization is used to check the authorization status of a request.
	// It streams responses back to the client, indicating whether the signin
	// request is authorized or not.
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
	// RequestSignin initiates the sign-in process and returns an authorization
	// URL.  The user should be redirected to this URL to complete the sign-in.
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

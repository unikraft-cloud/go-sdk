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
	"net/url"
	"os"
	"strings"
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
	// ListMetros lists all available metros.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/metros
	//
	// See: https://unikraft.com/docs/api/platform/v1/metros#list-metros
	ListMetros(ctx context.Context, ropts ...RequestOption) (*Response[ListMetroResponseData], error)
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
	// See: https://unikraft.com/docs/api/platform/v1/node-activation-service#node-activate
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
	// See: https://unikraft.com/docs/api/platform/v1/node-activation-service#node-renew
	NodeRenew(ctx context.Context, request NodeRenewRequest, ropts ...RequestOption) (*Response[NodeRenewResponseData], error)
	// Delete one or more nodes.
	//
	// Batch deletion of nodes by their identifiers.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: DELETE /v1/node
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#destroy-node
	DestroyNode(ctx context.Context, request []NameOrUUID, opts DestroyNodeOpts, ropts ...RequestOption) (*Response[DestroyNodeResponseData], error)
	// Destroy (delete) a node by its UUID.
	//
	// Deprovisions and deletes the node. The node will transition to
	// DEPROVISIONING state and then be removed. If the node has a delete lock,
	// this operation will fail unless the lock is first removed.
	//
	// @param `uuid`
	// 	The UUID of the node to delete.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: DELETE /v1/node/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#destroy-node-by-uuid
	DestroyNodeByUUID(ctx context.Context, uuid string, opts DestroyNodeByUUIDOpts, ropts ...RequestOption) (*Response[DestroyNodeResponseData], error)
	// Get a single node by its UUID.
	//
	// @param `uuid`
	// 	The UUID of the node to retrieve.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#get-node-by-uuid
	GetNodeByUUID(ctx context.Context, uuid string, ropts ...RequestOption) (*Response[ListNodesResponseData], error)
	// List available machine types for a provider.
	//
	// Returns the machine types (instance types) available for provisioning
	// on the specified cloud provider.
	//
	// @param `provider`
	// 	The provider to list machine types for.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node/provider/{provider}/types
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#list-machine-types
	ListMachineTypes(ctx context.Context, provider string, opts ListMachineTypesOpts, ropts ...RequestOption) (*Response[ListMachineTypesResponseData], error)
	// Get one or more nodes.
	//
	// Returns nodes matching the specified filters. If no filters are
	// provided, all nodes are returned.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#list-nodes
	ListNodes(ctx context.Context, request []NameOrUUID, opts ListNodesOpts, ropts ...RequestOption) (*Response[ListNodesResponseData], error)
	// List available regions for a provider.
	//
	// Returns the regions available for provisioning on the specified cloud
	// provider, including availability zone information.
	//
	// @param `provider`
	// 	The provider to list regions for.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node/provider/{provider}/regions
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#list-regions
	ListRegions(ctx context.Context, provider string, ropts ...RequestOption) (*Response[ListRegionsResponseData], error)
	// Create a new node.
	//
	// Creates a new compute node on the specified cloud provider. The node
	// will go through provisioning (PROVISIONING state) and configuration
	// (CONFIGURING state) before becoming ready (READY state) for use.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/node
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#provision-node
	ProvisionNode(ctx context.Context, request ProvisionNodeRequest, ropts ...RequestOption) (*Response[ProvisionNodeResponseData], error)
	// Update a node by its UUID.
	//
	// Updates mutable properties of a node such as tags, delete lock,
	// or SSH keys.
	//
	// @param `uuid`
	// 	The UUID of the node to update.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: PATCH /v1/node/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#update-node-by-uuid
	UpdateNodeByUUID(ctx context.Context, uuid string, request []UpdateNodePayload, ropts ...RequestOption) (*Response[UpdateNodesResponseData], error)
	// Update one or more nodes.
	//
	// Batch update of mutable properties for multiple nodes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: PATCH /v1/node
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#update-nodes
	UpdateNodes(ctx context.Context, request []NameOrUUID, ropts ...RequestOption) (*Response[UpdateNodesResponseData], error)
	// Wait for a node to reach a specific state.
	//
	// Blocks until the specified node reaches one of the desired states
	// or the timeout is reached.
	//
	// @param `uuid`
	// 	The UUID of the node to wait for.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node/{uuid}/wait
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#wait-node-by-uuid
	WaitNodeByUUID(ctx context.Context, uuid string, opts WaitNodeByUUIDOpts, ropts ...RequestOption) (*Response[WaitNodesResponseData], error)
	// Wait for multiple nodes to reach specific states.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/node/wait
	//
	// See: https://unikraft.com/docs/api/platform/v1/node-service#wait-nodes
	WaitNodes(ctx context.Context, opts WaitNodesOpts, ropts ...RequestOption) (*Response[WaitNodesResponseData], error)
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
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ListMetros(ctx context.Context, ropts ...RequestOption) (*Response[ListMetroResponseData], error) {
	requestPath := "/v1/metros"

	resp := &Response[ListMetroResponseData]{}
	if err := doRequest[ListMetroResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
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
		return resp, fmt.Errorf("performing the request: %w", err)
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
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) DestroyNode(ctx context.Context, request []NameOrUUID, opts DestroyNodeOpts, ropts ...RequestOption) (*Response[DestroyNodeResponseData], error) {
	requestPath := "/v1/node"

	query := make(url.Values)
	if opts.Force != nil {
		query.Add("force", fmt.Sprintf("%t", *opts.Force))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DestroyNodeResponseData]{}
	if err := doRequest[DestroyNodeResponseData](ctx, c.request, http.MethodDelete, requestPath, query, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) DestroyNodeByUUID(ctx context.Context, uuid string, opts DestroyNodeByUUIDOpts, ropts ...RequestOption) (*Response[DestroyNodeResponseData], error) {
	requestPath := "/v1/node/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(uuid))

	query := make(url.Values)
	if opts.Force != nil {
		query.Add("force", fmt.Sprintf("%t", *opts.Force))
	}

	resp := &Response[DestroyNodeResponseData]{}
	if err := doRequest[DestroyNodeResponseData](ctx, c.request, http.MethodDelete, requestPath, query, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) GetNodeByUUID(ctx context.Context, uuid string, ropts ...RequestOption) (*Response[ListNodesResponseData], error) {
	requestPath := "/v1/node/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(uuid))

	resp := &Response[ListNodesResponseData]{}
	if err := doRequest[ListNodesResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ListMachineTypes(ctx context.Context, provider string, opts ListMachineTypesOpts, ropts ...RequestOption) (*Response[ListMachineTypesResponseData], error) {
	requestPath := "/v1/node/provider/{provider}/types"
	requestPath = strings.ReplaceAll(requestPath, "{provider}", url.PathEscape(provider))

	query := make(url.Values)
	if opts.Region != nil {
		query.Add("region", *opts.Region)
	}

	resp := &Response[ListMachineTypesResponseData]{}
	if err := doRequest[ListMachineTypesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ListNodes(ctx context.Context, request []NameOrUUID, opts ListNodesOpts, ropts ...RequestOption) (*Response[ListNodesResponseData], error) {
	requestPath := "/v1/node"

	query := make(url.Values)
	if opts.Provider != nil {
		query.Add("provider", *opts.Provider)
	}
	if opts.State != nil {
		query.Add("state", *opts.State)
	}
	if opts.Metro != nil {
		query.Add("metro", *opts.Metro)
	}
	if opts.Region != nil {
		query.Add("region", *opts.Region)
	}
	if opts.Limit != nil {
		query.Add("limit", fmt.Sprintf("%d", *opts.Limit))
	}
	if opts.Offset != nil {
		query.Add("offset", fmt.Sprintf("%d", *opts.Offset))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[ListNodesResponseData]{}
	if err := doRequest[ListNodesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ListRegions(ctx context.Context, provider string, ropts ...RequestOption) (*Response[ListRegionsResponseData], error) {
	requestPath := "/v1/node/provider/{provider}/regions"
	requestPath = strings.ReplaceAll(requestPath, "{provider}", url.PathEscape(provider))

	resp := &Response[ListRegionsResponseData]{}
	if err := doRequest[ListRegionsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ProvisionNode(ctx context.Context, request ProvisionNodeRequest, ropts ...RequestOption) (*Response[ProvisionNodeResponseData], error) {
	requestPath := "/v1/node"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[ProvisionNodeResponseData]{}
	if err := doRequest[ProvisionNodeResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) UpdateNodeByUUID(ctx context.Context, uuid string, request []UpdateNodePayload, ropts ...RequestOption) (*Response[UpdateNodesResponseData], error) {
	requestPath := "/v1/node/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(uuid))

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateNodesResponseData]{}
	if err := doRequest[UpdateNodesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) UpdateNodes(ctx context.Context, request []NameOrUUID, ropts ...RequestOption) (*Response[UpdateNodesResponseData], error) {
	requestPath := "/v1/node"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateNodesResponseData]{}
	if err := doRequest[UpdateNodesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) WaitNodeByUUID(ctx context.Context, uuid string, opts WaitNodeByUUIDOpts, ropts ...RequestOption) (*Response[WaitNodesResponseData], error) {
	requestPath := "/v1/node/{uuid}/wait"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(uuid))

	query := make(url.Values)
	for _, v := range opts.States {
		query.Add("states", v)
	}
	if opts.TimeoutMs != nil {
		query.Add("timeout_ms", fmt.Sprintf("%d", *opts.TimeoutMs))
	}

	resp := &Response[WaitNodesResponseData]{}
	if err := doRequest[WaitNodesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) WaitNodes(ctx context.Context, opts WaitNodesOpts, ropts ...RequestOption) (*Response[WaitNodesResponseData], error) {
	requestPath := "/v1/node/wait"

	query := make(url.Values)
	for _, v := range opts.States {
		query.Add("states", v)
	}
	if opts.TimeoutMs != nil {
		query.Add("timeout_ms", fmt.Sprintf("%d", *opts.TimeoutMs))
	}

	resp := &Response[WaitNodesResponseData]{}
	if err := doRequest[WaitNodesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

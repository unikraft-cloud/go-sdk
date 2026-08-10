// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

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
	// Forget a finished command.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: DELETE /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}
	DeleteInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Execute a new command in the specified plugin.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/commands
	ExecInstanceCommand(ctx context.Context, instance string, pluginName string, request ExecInstanceCommandRequestBody, ropts ...RequestOption) (*Response[CommandResponseData], error)
	// Feed data into a command's standard input.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/stdin
	FeedInstanceCommandStdin(ctx context.Context, instance string, pluginName string, uuid string, request FeedInstanceCommandStdinRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Fetch collected logs from a command.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/logs
	GetInstanceCommandLogs(ctx context.Context, instance string, pluginName string, uuid string, request GetInstanceCommandLogsRequestBody, ropts ...RequestOption) (*Response[CommandLogsResponseData], error)
	// Fetch raw stdout or stderr from a command.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `stream`
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/logs/raw/{stream}
	GetInstanceCommandRawLogs(ctx context.Context, instance string, pluginName string, uuid string, stream SandboxLogStream, ropts ...RequestOption) (*Response[GetInstanceCommandRawLogsResponseData], error)
	// Inspect an existing command.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}
	InspectInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[CommandInspectResponseData], error)
	// List commands running in the specified plugin.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: GET /v1/instances/{instance}/plugins/{plugin_name}/commands
	ListInstanceCommands(ctx context.Context, instance string, pluginName string, ropts ...RequestOption) (*Response[ListInstanceCommandsResponseData], error)
	// Create a directory via the filesystem API.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/fs/mkdir
	MakeInstanceDirectory(ctx context.Context, instance string, pluginName string, request MakeInstanceDirectoryRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Read a file via the filesystem API.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/fs/read
	ReadInstanceFile(ctx context.Context, instance string, pluginName string, request ReadInstanceFileRequestBody, ropts ...RequestOption) (*Response[ReadInstanceFileResponseData], error)
	// Download a file's raw contents via the filesystem API.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/fs/read_raw
	ReadInstanceFileRaw(ctx context.Context, instance string, pluginName string, request ReadInstanceFileRawRequestBody, ropts ...RequestOption) (*Response[ReadInstanceFileRawResponseData], error)
	// Send a signal to a running command.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/signal
	SignalInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, request SignalInstanceCommandRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Wait for a command to finish (blocking).
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/wait
	WaitInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Wait for a command to finish with a timeout.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/wait_timeout
	WaitInstanceCommandWithTimeout(ctx context.Context, instance string, pluginName string, uuid string, request WaitInstanceCommandWithTimeoutRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// Write a file via the filesystem API.
	//
	// @param `instance`
	//
	// @param `pluginName`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `ropts`
	// 	Optional request modifiers.
	//
	// Performs: POST /v1/instances/{instance}/plugins/{plugin_name}/fs/write
	WriteInstanceFile(ctx context.Context, instance string, pluginName string, request WriteInstanceFileRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error)
	// WithMetro sets the metro to use when connecting to the API.
	WithMetro(string) Client
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
		options.SetDefaultMetro(os.Getenv("UKC_METRO"))
	}
	if options.DefaultEndpoint() == "" {
		options.SetDefaultMetro(DefaultMetro)
	}

	if options.AllowInsecure() && options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewHTTPClient(httpclient.WithInsecure()))
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

// WithMetro sets the metro to use when connecting to the API.
func (c *client) WithMetro(m string) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
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

func (c *client) DeleteInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ExecInstanceCommand(ctx context.Context, instance string, pluginName string, request ExecInstanceCommandRequestBody, ropts ...RequestOption) (*Response[CommandResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CommandResponseData]{}
	if err := doRequest[CommandResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) FeedInstanceCommandStdin(ctx context.Context, instance string, pluginName string, uuid string, request FeedInstanceCommandStdinRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/stdin"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) GetInstanceCommandLogs(ctx context.Context, instance string, pluginName string, uuid string, request GetInstanceCommandLogsRequestBody, ropts ...RequestOption) (*Response[CommandLogsResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/logs"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CommandLogsResponseData]{}
	if err := doRequest[CommandLogsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) GetInstanceCommandRawLogs(ctx context.Context, instance string, pluginName string, uuid string, stream SandboxLogStream, ropts ...RequestOption) (*Response[GetInstanceCommandRawLogsResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/logs/raw/{stream}"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))
	requestPath = strings.ReplaceAll(requestPath, "{stream}", url.PathEscape(string(stream)))

	resp := &Response[GetInstanceCommandRawLogsResponseData]{}
	if err := doRequest[GetInstanceCommandRawLogsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) InspectInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[CommandInspectResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[CommandInspectResponseData]{}
	if err := doRequest[CommandInspectResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ListInstanceCommands(ctx context.Context, instance string, pluginName string, ropts ...RequestOption) (*Response[ListInstanceCommandsResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	resp := &Response[ListInstanceCommandsResponseData]{}
	if err := doRequest[ListInstanceCommandsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) MakeInstanceDirectory(ctx context.Context, instance string, pluginName string, request MakeInstanceDirectoryRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/fs/mkdir"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ReadInstanceFile(ctx context.Context, instance string, pluginName string, request ReadInstanceFileRequestBody, ropts ...RequestOption) (*Response[ReadInstanceFileResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/fs/read"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[ReadInstanceFileResponseData]{}
	if err := doRequest[ReadInstanceFileResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) ReadInstanceFileRaw(ctx context.Context, instance string, pluginName string, request ReadInstanceFileRawRequestBody, ropts ...RequestOption) (*Response[ReadInstanceFileRawResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/fs/read_raw"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[ReadInstanceFileRawResponseData]{}
	if err := doRequest[ReadInstanceFileRawResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) SignalInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, request SignalInstanceCommandRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/signal"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) WaitInstanceCommand(ctx context.Context, instance string, pluginName string, uuid string, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/wait"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, nil, resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) WaitInstanceCommandWithTimeout(ctx context.Context, instance string, pluginName string, uuid string, request WaitInstanceCommandWithTimeoutRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/commands/{uuid}/wait_timeout"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

func (c *client) WriteInstanceFile(ctx context.Context, instance string, pluginName string, request WriteInstanceFileRequestBody, ropts ...RequestOption) (*Response[PluginEmptyResponseData], error) {
	requestPath := "/v1/instances/{instance}/plugins/{plugin_name}/fs/write"
	requestPath = strings.ReplaceAll(requestPath, "{instance}", url.PathEscape(string(instance)))
	requestPath = strings.ReplaceAll(requestPath, "{plugin_name}", url.PathEscape(string(pluginName)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[PluginEmptyResponseData]{}
	if err := doRequest[PluginEmptyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp, ropts...); err != nil {
		return resp, fmt.Errorf("performing the request: %w", err)
	}
	return resp, nil
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"fmt"
	"time"

	"unikraft.com/cloud/sdk/pkg/httpclient"
	"unikraft.com/cloud/sdk/platform/gen"
)

// Client is a façade over the platform API, exposing resource accessors
// that return rich objects, plus Raw for everything else.
type Client interface {
	// Raw returns the underlying generated client, for API surface not
	// (yet) covered by Client's rich accessors.
	Raw() gen.Client

	Instance(ctx context.Context, ref Ref) (*Instance, error)
	// Instances returns all instances if refs is empty.
	Instances(ctx context.Context, refs []Ref, opts ...InstancesOption) ([]*Instance, error)
	// CreateInstance may create more than one instance if WithCreateReplicas is set.
	CreateInstance(ctx context.Context, opts ...CreateInstanceOption) ([]*Instance, error)

	ServiceGroup(ctx context.Context, ref Ref) (*ServiceGroup, error)
	// ServiceGroups returns all service groups if refs is empty.
	ServiceGroups(ctx context.Context, refs []Ref, opts ...ServiceGroupsOption) ([]*ServiceGroup, error)
	CreateServiceGroup(ctx context.Context, opts ...CreateServiceGroupOption) (*ServiceGroup, error)

	Certificate(ctx context.Context, ref Ref) (*Certificate, error)
	// Certificates returns all certificates if refs is empty.
	Certificates(ctx context.Context, refs []Ref, opts ...CertificatesOption) ([]*Certificate, error)
	CreateCertificate(ctx context.Context, opts ...CreateCertificateOption) (*Certificate, error)

	Volume(ctx context.Context, ref Ref) (*Volume, error)
	// Volumes returns all volumes if refs is empty.
	Volumes(ctx context.Context, refs []Ref, opts ...VolumesOption) ([]*Volume, error)
	CreateVolume(ctx context.Context, opts ...CreateVolumeOption) (*Volume, error)

	Template(ctx context.Context, ref Ref) (*Template, error)
	// Templates returns all templates if refs is empty.
	Templates(ctx context.Context, refs []Ref, opts ...TemplatesOption) ([]*Template, error)
	// PrepareTemplate returns the template with name, and makes it from
	// source if it does not exist.
	PrepareTemplate(ctx context.Context, name string, source gen.Instance, opts ...PrepareTemplateOption) (*Template, error)

	// Images returns the images that the nodes of the metro have cached.
	Images(ctx context.Context, opts ...ImagesOption) ([]gen.Image, error)

	// Quotas reports the current user's resource usage and limits.
	Quotas(ctx context.Context) (*gen.Quotas, error)
	// QuotasForUser reports resource usage and limits for the user
	// identified by uuid.
	QuotasForUser(ctx context.Context, uuid string) (*gen.Quotas, error)

	WithMetro(metro string) Client
	WithHTTPClient(hc httpclient.HTTPClient) Client
	WithTimeout(to time.Duration) Client
}

// ClientOption is an option function used during initialization of a Client.
type ClientOption func(*clientConfig)

type clientConfig struct {
	token         string
	defaultMetro  string
	httpClient    httpclient.HTTPClient
	userAgent     string
	allowInsecure bool
}

// WithToken sets the access token used for authentication with the API.
func WithToken(token string) ClientOption {
	return func(c *clientConfig) { c.token = token }
}

// WithDefaultMetro sets the default metro used when connecting to the API.
func WithDefaultMetro(metro string) ClientOption {
	return func(c *clientConfig) { c.defaultMetro = metro }
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

// NewClient creates a new Client for the platform API. The token and
// default metro are read from the UKC_TOKEN/UNIKRAFT_CLOUD_TOKEN/
// KRAFTCLOUD_TOKEN and UKC_METRO environment variables unless overridden by
// options.
func NewClient(opts ...ClientOption) Client {
	var cfg clientConfig
	for _, o := range opts {
		o(&cfg)
	}

	var genOpts []gen.ClientOption
	if cfg.token != "" {
		genOpts = append(genOpts, gen.WithToken(cfg.token))
	}
	if cfg.defaultMetro != "" {
		genOpts = append(genOpts, gen.WithDefaultMetro(cfg.defaultMetro))
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

	return &platformClient{raw: gen.NewClientFromEnv(genOpts...)}
}

type platformClient struct {
	raw gen.Client
}

func (c *platformClient) Raw() gen.Client { return c.raw }

func (c *platformClient) Instance(ctx context.Context, ref Ref) (*Instance, error) {
	resp, err := c.raw.GetInstances(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.GetInstancesOpts{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Instances) == 0 {
		return nil, fmt.Errorf("instance %s not found", ref)
	}
	return wrapInstance(c.raw, resp.Data.Instances[0]), nil
}

func (c *platformClient) Instances(ctx context.Context, refs []Ref, opts ...InstancesOption) ([]*Instance, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.GetInstancesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetInstances(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	instances := make([]*Instance, len(resp.Data.Instances))
	for i, data := range resp.Data.Instances {
		instances[i] = wrapInstance(c.raw, data)
	}
	return instances, nil
}

func (c *platformClient) CreateInstance(ctx context.Context, opts ...CreateInstanceOption) ([]*Instance, error) {
	var request gen.CreateInstanceRequest
	for _, opt := range opts {
		opt(&request)
	}
	resp, err := c.raw.CreateInstance(ctx, request)
	if err != nil {
		return nil, err
	}
	instances := make([]*Instance, len(resp.Data.Instances))
	for i, data := range resp.Data.Instances {
		instances[i] = wrapInstance(c.raw, data)
	}
	return instances, nil
}

func (c *platformClient) ServiceGroup(ctx context.Context, ref Ref) (*ServiceGroup, error) {
	resp, err := c.raw.GetServiceGroups(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.GetServiceGroupsOpts{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.ServiceGroups) == 0 {
		return nil, fmt.Errorf("service group %s not found", ref)
	}
	return wrapServiceGroup(c.raw, resp.Data.ServiceGroups[0]), nil
}

func (c *platformClient) ServiceGroups(ctx context.Context, refs []Ref, opts ...ServiceGroupsOption) ([]*ServiceGroup, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.GetServiceGroupsOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetServiceGroups(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	groups := make([]*ServiceGroup, len(resp.Data.ServiceGroups))
	for i, data := range resp.Data.ServiceGroups {
		groups[i] = wrapServiceGroup(c.raw, data)
	}
	return groups, nil
}

func (c *platformClient) CreateServiceGroup(ctx context.Context, opts ...CreateServiceGroupOption) (*ServiceGroup, error) {
	var request gen.CreateServiceGroupRequest
	for _, opt := range opts {
		opt(&request)
	}
	resp, err := c.raw.CreateServiceGroup(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.ServiceGroups) == 0 {
		return nil, fmt.Errorf("service group create response contained no service group")
	}
	return wrapServiceGroup(c.raw, resp.Data.ServiceGroups[0]), nil
}

func (c *platformClient) Certificate(ctx context.Context, ref Ref) (*Certificate, error) {
	resp, err := c.raw.GetCertificates(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.GetCertificatesOpts{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Certificates) == 0 {
		return nil, fmt.Errorf("certificate %s not found", ref)
	}
	return wrapCertificate(c.raw, resp.Data.Certificates[0]), nil
}

func (c *platformClient) Certificates(ctx context.Context, refs []Ref, opts ...CertificatesOption) ([]*Certificate, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.GetCertificatesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetCertificates(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	certs := make([]*Certificate, len(resp.Data.Certificates))
	for i, data := range resp.Data.Certificates {
		certs[i] = wrapCertificate(c.raw, data)
	}
	return certs, nil
}

func (c *platformClient) CreateCertificate(ctx context.Context, opts ...CreateCertificateOption) (*Certificate, error) {
	var request gen.CreateCertificateRequest
	for _, opt := range opts {
		opt(&request)
	}
	resp, err := c.raw.CreateCertificate(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Certificates) == 0 {
		return nil, fmt.Errorf("certificate create response contained no certificate")
	}
	return wrapCertificate(c.raw, resp.Data.Certificates[0]), nil
}

func (c *platformClient) Volume(ctx context.Context, ref Ref) (*Volume, error) {
	resp, err := c.raw.GetVolumes(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.GetVolumesOpts{})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Volumes) == 0 {
		return nil, fmt.Errorf("volume %s not found", ref)
	}
	return wrapVolume(c.raw, resp.Data.Volumes[0]), nil
}

func (c *platformClient) Volumes(ctx context.Context, refs []Ref, opts ...VolumesOption) ([]*Volume, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.GetVolumesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetVolumes(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	volumes := make([]*Volume, len(resp.Data.Volumes))
	for i, data := range resp.Data.Volumes {
		volumes[i] = wrapVolume(c.raw, data)
	}
	return volumes, nil
}

func (c *platformClient) CreateVolume(ctx context.Context, opts ...CreateVolumeOption) (*Volume, error) {
	var request gen.CreateVolumeRequest
	for _, opt := range opts {
		opt(&request)
	}
	resp, err := c.raw.CreateVolume(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Volumes) == 0 {
		return nil, fmt.Errorf("volume create response contained no volume")
	}
	// CreateVolume's response only acknowledges creation (uuid/name/state);
	// it doesn't carry the full volume, so fetch it to populate the wrapper.
	return c.Volume(ctx, UUID(resp.Data.Volumes[0].Uuid))
}

func (c *platformClient) WithMetro(metro string) Client {
	return &platformClient{raw: c.raw.WithMetro(metro)}
}

func (c *platformClient) WithHTTPClient(hc httpclient.HTTPClient) Client {
	return &platformClient{raw: c.raw.WithHTTPClient(hc)}
}

func (c *platformClient) WithTimeout(to time.Duration) Client {
	return &platformClient{raw: c.raw.WithTimeout(to)}
}

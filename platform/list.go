// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"iter"

	"unikraft.com/cloud/sdk/platform/gen"
)

// ListOption configures the List methods of [Client].
type ListOption func(*listConfig)

type listConfig struct {
	pageSize int
	details  bool
	tags     []string
}

// WithListPageSize sets how many items each request reads. The default is
// 100.
func WithListPageSize(n int) ListOption {
	return func(c *listConfig) { c.pageSize = n }
}

// WithListDetails reads full details for each item.
func WithListDetails() ListOption {
	return func(c *listConfig) { c.details = true }
}

// WithListTags lists only items that have all of tags. Service groups and
// certificates have no tags, and ignore this option.
func WithListTags(tags ...string) ListOption {
	return func(c *listConfig) { c.tags = tags }
}

func newListConfig(opts []ListOption) listConfig {
	var cfg listConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func detailsOpt(details bool) *bool {
	if !details {
		return nil
	}
	return new(true)
}

func (c *platformClient) ListInstances(ctx context.Context, opts ...ListOption) iter.Seq2[*Instance, error] {
	cfg := newListConfig(opts)
	fetch := func(ctx context.Context, count uint32, from *string) ([]*Instance, error) {
		resp, err := c.raw.GetInstances(ctx, nil, gen.GetInstancesOpts{
			Count: &count, From: from, Details: detailsOpt(cfg.details), Tags: cfg.tags,
		})
		if err != nil {
			return nil, err
		}
		items := make([]*Instance, len(resp.Data.Instances))
		for i, data := range resp.Data.Instances {
			items[i] = wrapInstance(c.raw, data)
		}
		return items, nil
	}
	return paginate(ctx, cfg.pageSize, fetch, func(i *Instance) string { return i.Uuid })
}

func (c *platformClient) ListVolumes(ctx context.Context, opts ...ListOption) iter.Seq2[*Volume, error] {
	cfg := newListConfig(opts)
	fetch := func(ctx context.Context, count uint32, from *string) ([]*Volume, error) {
		resp, err := c.raw.GetVolumes(ctx, nil, gen.GetVolumesOpts{
			Count: &count, From: from, Details: detailsOpt(cfg.details), Tags: cfg.tags,
		})
		if err != nil {
			return nil, err
		}
		items := make([]*Volume, len(resp.Data.Volumes))
		for i, data := range resp.Data.Volumes {
			items[i] = wrapVolume(c.raw, data)
		}
		return items, nil
	}
	return paginate(ctx, cfg.pageSize, fetch, func(v *Volume) string { return v.Uuid })
}

func (c *platformClient) ListServiceGroups(ctx context.Context, opts ...ListOption) iter.Seq2[*ServiceGroup, error] {
	cfg := newListConfig(opts)
	fetch := func(ctx context.Context, count uint32, from *string) ([]*ServiceGroup, error) {
		resp, err := c.raw.GetServiceGroups(ctx, nil, gen.GetServiceGroupsOpts{
			Count: &count, From: from, Details: detailsOpt(cfg.details),
		})
		if err != nil {
			return nil, err
		}
		items := make([]*ServiceGroup, len(resp.Data.ServiceGroups))
		for i, data := range resp.Data.ServiceGroups {
			items[i] = wrapServiceGroup(c.raw, data)
		}
		return items, nil
	}
	return paginate(ctx, cfg.pageSize, fetch, func(g *ServiceGroup) string { return g.Uuid })
}

func (c *platformClient) ListCertificates(ctx context.Context, opts ...ListOption) iter.Seq2[*Certificate, error] {
	cfg := newListConfig(opts)
	fetch := func(ctx context.Context, count uint32, from *string) ([]*Certificate, error) {
		resp, err := c.raw.GetCertificates(ctx, nil, gen.GetCertificatesOpts{
			Count: &count, From: from, Details: detailsOpt(cfg.details),
		})
		if err != nil {
			return nil, err
		}
		items := make([]*Certificate, len(resp.Data.Certificates))
		for i, data := range resp.Data.Certificates {
			items[i] = wrapCertificate(c.raw, data)
		}
		return items, nil
	}
	return paginate(ctx, cfg.pageSize, fetch, func(cert *Certificate) string { return cert.Uuid })
}

func (c *platformClient) ListTemplates(ctx context.Context, opts ...ListOption) iter.Seq2[*Template, error] {
	cfg := newListConfig(opts)
	fetch := func(ctx context.Context, count uint32, from *string) ([]*Template, error) {
		resp, err := c.raw.GetTemplateInstances(ctx, nil, gen.GetTemplateInstancesOpts{
			Count: &count, From: from, Details: detailsOpt(cfg.details), Tags: cfg.tags,
		})
		if err != nil {
			return nil, err
		}
		items := make([]*Template, len(resp.Data.Instances))
		for i, data := range resp.Data.Instances {
			items[i] = wrapTemplate(c.raw, data)
		}
		return items, nil
	}
	return paginate(ctx, cfg.pageSize, fetch, func(t *Template) string { return t.Uuid })
}

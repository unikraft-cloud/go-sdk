// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"

	"unikraft.com/cloud/sdk/platform/gen"
)

func namesOrUUIDs(refs []Ref) []gen.NameOrUUID {
	out := make([]gen.NameOrUUID, len(refs))
	for i, ref := range refs {
		out[i] = ref.toNameOrUUID()
	}
	return out
}

func (c *platformClient) StartInstances(ctx context.Context, refs []Ref, opts ...StartOption) error {
	var body gen.StartInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	items := make([]gen.StartInstancesRequestItem, len(refs))
	for i, id := range namesOrUUIDs(refs) {
		items[i] = gen.StartInstancesRequestItem{
			Uuid: id.Uuid, Name: id.Name, TimeoutS: body.TimeoutS,
		}
	}
	_, err := c.raw.StartInstances(ctx, items)
	return err
}

func (c *platformClient) StopInstances(ctx context.Context, refs []Ref, opts ...StopOption) error {
	var body gen.StopInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	items := make([]gen.StopInstancesRequestItem, len(refs))
	for i, id := range namesOrUUIDs(refs) {
		items[i] = gen.StopInstancesRequestItem{
			Uuid: id.Uuid, Name: id.Name, Force: body.Force, DrainTimeoutMs: body.DrainTimeoutMs,
			Quick: body.Quick, Ifstate: body.Ifstate, Dump: body.Dump,
		}
	}
	_, err := c.raw.StopInstances(ctx, items)
	return err
}

func (c *platformClient) SuspendInstances(ctx context.Context, refs []Ref, opts ...SuspendOption) error {
	var body gen.SuspendInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	items := make([]gen.SuspendInstancesRequestItem, len(refs))
	for i, id := range namesOrUUIDs(refs) {
		items[i] = gen.SuspendInstancesRequestItem{
			Uuid: id.Uuid, Name: id.Name, DrainTimeoutMs: body.DrainTimeoutMs,
		}
	}
	_, err := c.raw.SuspendInstances(ctx, items)
	return err
}

func (c *platformClient) DeleteInstances(ctx context.Context, refs []Ref, opts ...DeleteOption) error {
	var cfg deleteConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	items := make([]gen.DeleteInstanceRequestItem, len(refs))
	for i, id := range namesOrUUIDs(refs) {
		items[i] = gen.DeleteInstanceRequestItem{
			Uuid: id.Uuid, Name: id.Name, TimeoutS: cfg.body.TimeoutS, DontRetain: cfg.body.DontRetain,
		}
	}
	err := whileBusy(ctx, cfg.retryBusy, func() error {
		_, err := c.raw.DeleteInstances(ctx, items)
		return err
	})
	if cfg.missingOK && gen.ErrorContainsOnly(err, gen.APIHTTPErrorNotFound) {
		return nil
	}
	return err
}

func (c *platformClient) DeleteVolumes(ctx context.Context, refs []Ref) error {
	_, err := c.raw.DeleteVolumes(ctx, namesOrUUIDs(refs))
	return err
}

func (c *platformClient) DeleteServiceGroups(ctx context.Context, refs []Ref) error {
	_, err := c.raw.DeleteServiceGroups(ctx, namesOrUUIDs(refs))
	return err
}

func (c *platformClient) DeleteCertificates(ctx context.Context, refs []Ref) error {
	_, err := c.raw.DeleteCertificates(ctx, namesOrUUIDs(refs))
	return err
}

func (c *platformClient) DeleteTemplates(ctx context.Context, refs []Ref) error {
	_, err := c.raw.DeleteTemplateInstances(ctx, namesOrUUIDs(refs))
	return err
}

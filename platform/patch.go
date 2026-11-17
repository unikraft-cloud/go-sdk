// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"errors"

	"unikraft.com/cloud/sdk/platform/gen"
)

// Change is one property change that Patch sends in a single update request.
// Use Set, Add or Del to make one.
type Change[P ~string] struct {
	Prop  P
	Op    string
	Value any
}

// Set replaces the value of prop with value.
func Set[P ~string](prop P, value any) Change[P] {
	return Change[P]{Prop: prop, Op: "set", Value: value}
}

// Add adds value to prop, for list and map properties such as tags or env.
func Add[P ~string](prop P, value any) Change[P] {
	return Change[P]{Prop: prop, Op: "add", Value: value}
}

// Del removes value from prop, for list and map properties such as tags or
// env.
func Del[P ~string](prop P, value any) Change[P] {
	return Change[P]{Prop: prop, Op: "del", Value: value}
}

// InstanceChange is a Change to an instance property.
type InstanceChange = Change[gen.MutableInstanceProperty]

// VolumeChange is a Change to a volume property.
type VolumeChange = Change[gen.MutableVolumeProperty]

// ServiceGroupChange is a Change to a service group property.
type ServiceGroupChange = Change[gen.MutableServiceGroupProperty]

// Patch sends changes in one update request, then refreshes the instance.
// Unlike Update, it can also change write-only properties such as hostname
// and dependencies.
func (i *Instance) Patch(ctx context.Context, changes ...InstanceChange) error {
	if len(changes) == 0 {
		return nil
	}
	items := make([]gen.UpdateInstancesRequestItem, len(changes))
	for n, c := range changes {
		items[n] = gen.UpdateInstancesRequestItem{
			Uuid: &i.Uuid, Prop: c.Prop, Op: gen.MutableInstanceOperation(c.Op), Value: &c.Value,
		}
	}
	_, err := i.client.UpdateInstances(ctx, items)
	return errors.Join(err, i.Refresh(ctx))
}

// Patch sends changes in one update request, then refreshes the volume.
func (v *Volume) Patch(ctx context.Context, changes ...VolumeChange) error {
	if len(changes) == 0 {
		return nil
	}
	items := make([]gen.UpdateVolumesRequestItem, len(changes))
	for n, c := range changes {
		items[n] = gen.UpdateVolumesRequestItem{
			Uuid: &v.Uuid, Prop: c.Prop, Op: gen.MutableVolumeOperation(c.Op), Value: &c.Value,
		}
	}
	_, err := v.client.UpdateVolumes(ctx, items)
	return errors.Join(err, v.Refresh(ctx))
}

// Patch sends changes in one update request, then refreshes the service
// group.
func (g *ServiceGroup) Patch(ctx context.Context, changes ...ServiceGroupChange) error {
	if len(changes) == 0 {
		return nil
	}
	items := make([]gen.UpdateServiceGroupsRequestItem, len(changes))
	for n, c := range changes {
		items[n] = gen.UpdateServiceGroupsRequestItem{
			Uuid: &g.Uuid, Prop: c.Prop, Op: gen.MutableServiceGroupOperation(c.Op), Value: &c.Value,
		}
	}
	_, err := g.client.UpdateServiceGroups(ctx, items)
	return errors.Join(err, g.Refresh(ctx))
}

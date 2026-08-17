// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"unikraft.com/cloud/sdk/platform/gen"
)

// ServiceGroup is a live, resource-oriented view of a service group. It
// embeds gen.ServiceGroup so its fields are accessed directly
// (sg.SoftLimit, sg.Domains, ...), and adds behavior (Refresh, Update,
// Delete) on top of the flat generated client.
type ServiceGroup struct {
	gen.ServiceGroup

	client gen.Client
	orig   gen.ServiceGroup // snapshot as of the last fetch/refresh, for diffing in Update
}

// cloneServiceGroup deep-clones a gen.ServiceGroup via a JSON round-trip, so
// that slice/pointer fields aren't aliased between a ServiceGroup and its
// snapshot. in was itself just unmarshaled from the API response, so a
// marshal/unmarshal round-trip of it failing indicates a broken invariant,
// not a normal error condition.
func cloneServiceGroup(in gen.ServiceGroup) gen.ServiceGroup {
	b, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("platform: cloning service group: %v", err))
	}
	var out gen.ServiceGroup
	if err := json.Unmarshal(b, &out); err != nil {
		panic(fmt.Sprintf("platform: cloning service group: %v", err))
	}
	return out
}

func wrapServiceGroup(client gen.Client, data gen.ServiceGroup) *ServiceGroup {
	return &ServiceGroup{
		ServiceGroup: cloneServiceGroup(data),
		client:       client,
		orig:         cloneServiceGroup(data),
	}
}

// Refresh re-fetches the service group from the API and resets the diffing
// snapshot used by Update.
func (g *ServiceGroup) Refresh(ctx context.Context) error {
	resp, err := g.client.GetServiceGroups(ctx, []gen.NameOrUUID{{Uuid: &g.Uuid}}, gen.GetServiceGroupsOpts{Details: new(true)})
	if err != nil {
		return err
	}
	if len(resp.Data.ServiceGroups) == 0 {
		return fmt.Errorf("service group %s not found", g.Uuid)
	}
	fresh := resp.Data.ServiceGroups[0]
	g.ServiceGroup, g.orig = cloneServiceGroup(fresh), cloneServiceGroup(fresh)
	return nil
}

// Delete deletes the service group. Unlike Update, it does not refresh
// afterwards: the service group is gone, and re-fetching it would just 404.
func (g *ServiceGroup) Delete(ctx context.Context) error {
	_, err := g.client.DeleteServiceGroupByUUID(ctx, g.Uuid)
	return err
}

// Update diffs the service group's current field values against the
// snapshot taken at the last fetch/refresh, sends the minimal patch for
// whatever changed, and always refreshes afterwards.
func (g *ServiceGroup) Update(ctx context.Context) error {
	items := g.diff()
	if len(items) == 0 {
		return nil
	}
	_, err := g.client.UpdateServiceGroups(ctx, items)
	if rerr := g.Refresh(ctx); rerr != nil {
		err = errors.Join(err, rerr)
	}
	return err
}

func (g *ServiceGroup) diff() []gen.UpdateServiceGroupsRequestItem {
	var items []gen.UpdateServiceGroupsRequestItem
	set := func(prop gen.MutableServiceGroupProperty, value any) {
		items = append(items, gen.UpdateServiceGroupsRequestItem{
			Uuid: &g.Uuid, Prop: prop, Op: gen.MutableServiceGroupOperationSet, Value: &value,
		})
	}

	if !reflect.DeepEqual(g.Services, g.orig.Services) {
		set(gen.MutableServiceGroupPropertyServices, g.Services)
	}
	if !reflect.DeepEqual(g.Domains, g.orig.Domains) {
		set(gen.MutableServiceGroupPropertyDomains, g.Domains)
	}
	if g.SoftLimit != g.orig.SoftLimit {
		set(gen.MutableServiceGroupPropertySoftLimit, g.SoftLimit)
	}
	if g.HardLimit != g.orig.HardLimit {
		set(gen.MutableServiceGroupPropertyHardLimit, g.HardLimit)
	}
	if !reflect.DeepEqual(g.Autokill, g.orig.Autokill) {
		set(gen.MutableServiceGroupPropertyAutokill, g.Autokill)
	}

	return items
}

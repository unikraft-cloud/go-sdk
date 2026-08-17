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
	"slices"

	"unikraft.com/cloud/sdk/platform/gen"
)

// Volume is a live, resource-oriented view of a volume. It embeds gen.Volume
// so its fields are accessed directly (vol.SizeMb, vol.State, ...), and adds
// behavior (Refresh, Attach, Detach, Update, Delete) on top of the flat
// generated client.
type Volume struct {
	gen.Volume

	client gen.Client
	orig   gen.Volume // snapshot as of the last fetch/refresh, for diffing in Update
}

// cloneVolume deep-clones a gen.Volume via a JSON round-trip, so that
// slice/map/pointer fields aren't aliased between a Volume and its
// snapshot. in was itself just unmarshaled from the API response, so a
// marshal/unmarshal round-trip of it failing indicates a broken invariant,
// not a normal error condition.
func cloneVolume(in gen.Volume) gen.Volume {
	b, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("platform: cloning volume: %v", err))
	}
	var out gen.Volume
	if err := json.Unmarshal(b, &out); err != nil {
		panic(fmt.Sprintf("platform: cloning volume: %v", err))
	}
	return out
}

func wrapVolume(client gen.Client, data gen.Volume) *Volume {
	return &Volume{
		Volume: cloneVolume(data),
		client: client,
		orig:   cloneVolume(data),
	}
}

// Refresh re-fetches the volume from the API and resets the diffing
// snapshot used by Update.
func (v *Volume) Refresh(ctx context.Context) error {
	resp, err := v.client.GetVolumes(ctx, []gen.NameOrUUID{{Uuid: &v.Uuid}}, gen.GetVolumesOpts{Details: new(true)})
	if err != nil {
		return err
	}
	if len(resp.Data.Volumes) == 0 {
		return fmt.Errorf("volume %s not found", v.Uuid)
	}
	fresh := resp.Data.Volumes[0]
	v.Volume, v.orig = cloneVolume(fresh), cloneVolume(fresh)
	return nil
}

// Delete deletes the volume. Unlike Update, it does not refresh afterwards:
// the volume is gone, and re-fetching it would just 404.
func (v *Volume) Delete(ctx context.Context) error {
	_, err := v.client.DeleteVolumeByUUID(ctx, v.Uuid)
	return err
}

// AttachOption configures a call to Volume.Attach.
type AttachOption func(*gen.AttachVolumeByUUIDRequestBody)

// WithAttachReadOnly mounts the volume read-only.
func WithAttachReadOnly() AttachOption {
	return func(b *gen.AttachVolumeByUUIDRequestBody) { b.Readonly = new(true) }
}

// Attach attaches the volume to the instance identified by to, mounted at
// path, then refreshes the volume to pick up the resulting state.
func (v *Volume) Attach(ctx context.Context, to Ref, path string, opts ...AttachOption) error {
	body := gen.AttachVolumeByUUIDRequestBody{AttachTo: to.toNameOrUUID(), At: path}
	for _, opt := range opts {
		opt(&body)
	}
	if _, err := v.client.AttachVolumeByUUID(ctx, v.Uuid, body); err != nil {
		return err
	}
	return v.Refresh(ctx)
}

// Detach detaches the volume from the instance identified by from, or from
// all instances if from is nil, then refreshes the volume to pick up the
// resulting state.
func (v *Volume) Detach(ctx context.Context, from Ref) error {
	var body gen.DetachVolumeByUUIDRequestBody
	if from != nil {
		nou := from.toNameOrUUID()
		body.From = &nou
	}
	if _, err := v.client.DetachVolumeByUUID(ctx, v.Uuid, body); err != nil {
		return err
	}
	return v.Refresh(ctx)
}

// Update diffs the volume's current field values against the snapshot taken
// at the last fetch/refresh, sends the minimal patch for whatever changed,
// and always refreshes afterwards.
func (v *Volume) Update(ctx context.Context) error {
	items := v.diff()
	if len(items) == 0 {
		return nil
	}
	_, err := v.client.UpdateVolumes(ctx, items)
	if rerr := v.Refresh(ctx); rerr != nil {
		err = errors.Join(err, rerr)
	}
	return err
}

func (v *Volume) diff() []gen.UpdateVolumesRequestItem {
	var items []gen.UpdateVolumesRequestItem
	set := func(prop gen.MutableVolumeProperty, value any) {
		items = append(items, gen.UpdateVolumesRequestItem{
			Uuid: &v.Uuid, Prop: prop, Op: gen.MutableVolumeOperationSet, Value: &value,
		})
	}

	if v.SizeMb != v.orig.SizeMb {
		set(gen.MutableVolumePropertySizeMb, v.SizeMb)
	}
	if !slices.Equal(v.Tags, v.orig.Tags) {
		set(gen.MutableVolumePropertyTags, v.Tags)
	}
	if v.QuotaPolicy != v.orig.QuotaPolicy {
		set(gen.MutableVolumePropertyQuotaPolicy, v.QuotaPolicy)
	}
	if !reflect.DeepEqual(v.DeleteLock, v.orig.DeleteLock) {
		set(gen.MutableVolumePropertyDeleteLock, v.DeleteLock)
	}

	return items
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"unikraft.com/cloud/sdk/platform/gen"
)

// CreateVolumeOption configures a call to [Client.CreateVolume].
type CreateVolumeOption func(*gen.CreateVolumeRequest)

// WithCreateVolumeName sets the volume's name. If not set, a random name is generated.
func WithCreateVolumeName(name string) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.Name = new(name) }
}

// WithCreateVolumeSizeMB sets the size of the volume, in megabytes.
func WithCreateVolumeSizeMB(mb uint64) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.SizeMb = new(mb) }
}

// WithCreateVolumeQuotaPolicy sets the quota policy for the volume.
func WithCreateVolumeQuotaPolicy(policy gen.VolumeQuotaPolicy) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.QuotaPolicy = &policy }
}

// WithCreateVolumeFilesystem sets the filesystem type to format or configure.
func WithCreateVolumeFilesystem(fs string) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.Filesystem = new(fs) }
}

// WithCreateVolumeTags associates tags with the volume.
func WithCreateVolumeTags(tags ...string) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.Tags = tags }
}

// WithCreateVolumeAccessMode sets the access mode, controlling volume
// sharing behavior. Defaults to rwo if not set.
func WithCreateVolumeAccessMode(mode gen.VolumeAccessMode) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.AccessMode = &mode }
}

// WithCreateVolumeTemplate creates the volume by cloning a template volume.
func WithCreateVolumeTemplate(tmpl gen.NameOrUUID) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.Template = &tmpl }
}

// WithCreateVolumeHostPath creates a managed volume backed by the given host
// path.
func WithCreateVolumeHostPath(path string) CreateVolumeOption {
	return func(r *gen.CreateVolumeRequest) { r.HostPath = new(path) }
}

// VolumesOption configures a call to [Client.Volumes].
type VolumesOption func(*gen.GetVolumesOpts)

// WithVolumeDetails includes full details for each returned volume.
func WithVolumeDetails() VolumesOption {
	return func(o *gen.GetVolumesOpts) { o.Details = new(true) }
}

// WithVolumeLimit caps the number of volumes returned.
func WithVolumeLimit(n uint32) VolumesOption {
	return func(o *gen.GetVolumesOpts) { o.Count = new(n) }
}

// WithVolumePage continues listing from the pagination cursor returned by a
// previous call.
func WithVolumePage(from string) VolumesOption {
	return func(o *gen.GetVolumesOpts) { o.From = new(from) }
}

// WithVolumeTags filters volumes to those matching all of the given tags.
func WithVolumeTags(tags ...string) VolumesOption {
	return func(o *gen.GetVolumesOpts) { o.Tags = tags }
}

// WithVolumeOrder sets the sort order for the returned volumes.
func WithVolumeOrder(order PaginationOrder) VolumesOption {
	return func(o *gen.GetVolumesOpts) { o.Order = new(gen.PaginationOrder(order)) }
}

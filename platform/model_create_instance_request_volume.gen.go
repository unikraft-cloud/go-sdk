// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A volume defines a storage volume that can be attached to the instance.
type CreateInstanceRequestVolume struct {
	// The UUID of an existing volume.
	//
	// If this is the only specified field, then it will look up an existing
	// volume by this UUID.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the volume.
	//
	// If this is the only specified field, then it will look up an existing
	// volume by this name.  If the volume does not exist, the request will
	// fail.  If a new volume is intended to be created, then this field must be
	// specified along with the mount point in the instance and a provisioning
	// source (size_mb or host_path).
	Name *string `json:"name,omitzero"`
	// The mount point for the volume in the instance.
	At string `json:"at"`
	// Whether the volume is read-only.
	//
	// If this field is set to true, the volume will be mounted as read-only in
	// the instance.  This field is optional and defaults to false and is only
	// applicable when using an existing volume.
	Readonly *bool `json:"readonly,omitzero"`
	// Quota policy for the volume.
	QuotaPolicy *string `json:"quota_policy,omitzero"`
	// Filesystem type to format or configure.
	// Without custom configuration, this is either `ext4` or `virtiofs`.
	Filesystem *string `json:"filesystem,omitzero"`
	// Tags to assign to the new volume.
	Tags []string `json:"tags,omitzero"`
	// Guest UID for managed volumes (host_path mode only).
	Uid *uint32 `json:"uid,omitzero"`
	// Guest GID for managed volumes (host_path mode only).
	Gid *uint32 `json:"gid,omitzero"`
	// Script arguments passed to volume initialization scripts.
	Args map[string]string `json:"args,omitzero"`
	// Access mode of the volume, controlling sharing behavior.
	// Defaults to read-write by a single instance (RWO).
	AccessMode *VolumeAccessMode `json:"access_mode,omitzero"`
	// The size of the volume when creating a new volume.
	//
	// When creating a new volume as part of the instance create request,
	// specify the size of the volume in MiB.
	SizeMb *uint64 `json:"size_mb,omitzero"`
	// A host path to create a managed volume from.
	HostPath *string `json:"host_path,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestVolume) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestVolume
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestVolume) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestVolume
	return json.Marshal((Alias)(m))
}

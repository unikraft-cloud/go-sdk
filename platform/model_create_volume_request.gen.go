// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// The request message for creating a volume.
type CreateVolumeRequest struct {
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name of the form `vol-X` is generated for you, where
	// `X` is a 5 character long random alphanumeric suffix..  The name can also
	// be used to identify the volume in API calls.
	Name *string `json:"name,omitzero"`
	// Quota policy for the volume.
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitzero"`
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
	// The access mode of the volume, controlling volume sharing behavior.
	// Defaults to `rwo` if not specified.
	AccessMode *VolumeAccessMode `json:"access_mode,omitzero"`
	// The size of the volume in megabytes.
	SizeMb *uint64 `json:"size_mb,omitzero"`
	// A host path to create a managed volume from.
	HostPath *string `json:"host_path,omitzero"`
	// Source template volume to clone from.
	Template *NameOrUUID `json:"template,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateVolumeRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateVolumeRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateVolumeRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateVolumeRequest
	return json.Marshal((Alias)(m))
}

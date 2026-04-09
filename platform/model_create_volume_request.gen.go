// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for creating a volume.
// Quota policy for the volume.
type CreateVolumeRequestQuotaPolicy string

const (
	CreateVolumeRequestQuotaPolicyStatic  CreateVolumeRequestQuotaPolicy = "static"
	CreateVolumeRequestQuotaPolicyDynamic CreateVolumeRequestQuotaPolicy = "dynamic"
)

type CreateVolumeRequest struct {
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name of the form `vol-X` is generated for you, where
	// `X` is a 5 character long random alphanumeric suffix..  The name can also
	// be used to identify the volume in API calls.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Quota policy for the volume.
	QuotaPolicy *CreateVolumeRequestQuotaPolicy `json:"quota_policy,omitempty"`
	// Filesystem type to format or configure.
	// Without custom configuration, this is either `ext4` or `virtiofs`.
	Filesystem *string `json:"filesystem,omitempty"`
	// Tags to assign to the new volume.
	Tags []string `json:"tags,omitempty"`
	// Guest UID for managed volumes (host_path mode only).
	Uid *uint32 `json:"uid,omitempty"`
	// Guest GID for managed volumes (host_path mode only).
	Gid *uint32 `json:"gid,omitempty"`
	// Script arguments passed to volume initialization scripts.
	Args map[string]string `json:"args,omitempty"`
	// The size of the volume in megabytes.
	SizeMb *uint64 `json:"size_mb,omitempty"`
	// A host path to create a managed volume from.
	HostPath *string `json:"host_path,omitempty"`
	// Source template volume to clone from.
	Template *NameOrUUID `json:"template,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateVolumeRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateVolumeRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":         {},
		"metro":        {},
		"quota_policy": {},
		"filesystem":   {},
		"tags":         {},
		"uid":          {},
		"gid":          {},
		"args":         {},
		"size_mb":      {},
		"host_path":    {},
		"template":     {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m CreateVolumeRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateVolumeRequest
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

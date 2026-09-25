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
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name        *string           `json:"name,omitzero"`
	SizeMb      *uint64           `json:"size_mb,omitzero"`
	HostPath    *string           `json:"host_path,omitzero"`
	At          string            `json:"at"`
	Readonly    *bool             `json:"readonly,omitzero"`
	QuotaPolicy *string           `json:"quota_policy,omitzero"`
	Filesystem  *string           `json:"filesystem,omitzero"`
	Tags        []string          `json:"tags,omitzero"`
	Uid         *uint32           `json:"uid,omitzero"`
	Gid         *uint32           `json:"gid,omitzero"`
	AccessMode  *VolumeAccessMode `json:"access_mode,omitzero"`
	Args        map[string]string `json:"args,omitzero"`

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

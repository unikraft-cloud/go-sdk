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

// The request message for creating a volume.
type CreateVolumeRequest struct {
	Name        *string            `json:"name,omitzero"`
	SizeMb      *uint64            `json:"size_mb,omitzero"`
	HostPath    *string            `json:"host_path,omitzero"`
	Template    *NameOrUUID        `json:"template,omitzero"`
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitzero"`
	Filesystem  *string            `json:"filesystem,omitzero"`
	Tags        []string           `json:"tags,omitzero"`
	Uid         *uint32            `json:"uid,omitzero"`
	Gid         *uint32            `json:"gid,omitzero"`
	AccessMode  *VolumeAccessMode  `json:"access_mode,omitzero"`
	Args        map[string]string  `json:"args,omitzero"`

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

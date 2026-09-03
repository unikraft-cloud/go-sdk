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

// A single request item describing the volume to clone.
type CloneVolumesRequestItem struct {
	// The name of the new cloned volume.  If not provided, a random name
	// of the form `vol-X` is generated for you, where `X` is a 5 character
	// long random alphanumeric suffix.
	VolName *string `json:"vol_name,omitzero"`
	// The quota policy for the new cloned volume.  If not provided, the quota
	// policy of the source volume is used.
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitzero"`
	// A list of tags to assign to the new cloned volume.
	Tags []string `json:"tags,omitzero"`
	// The UUID of the volume to clone.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the volume to clone.  Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CloneVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CloneVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumesRequestItem
	return json.Marshal((Alias)(m))
}

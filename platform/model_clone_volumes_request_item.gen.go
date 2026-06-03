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

// A single request item describing the volume to clone.

type CloneVolumesRequestItem struct {
	// The name of the new cloned volume.  If not provided, a random name
	// of the form `vol-X` is generated for you, where `X` is a 5 character
	// long random alphanumeric suffix.
	VolName *string `json:"vol_name,omitempty"`
	// The quota policy for the new cloned volume.  If not provided, the quota
	// policy of the source volume is used.
	QuotaPolicy *VolumeQuotaPolicy `json:"quota_policy,omitempty"`
	// A list of tags to assign to the new cloned volume.
	Tags []string `json:"tags,omitempty"`
	// The UUID of the volume to clone.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to clone.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CloneVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CloneVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumesRequestItem
	return json.Marshal((Alias)(m))
}

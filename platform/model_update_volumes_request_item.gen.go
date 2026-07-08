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

// A single request item for updating a volume.
type UpdateVolumesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop MutableVolumeProperty `json:"prop"`
	// The operation to perform.
	Op MutableVolumeOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "size_mb": unsigned integer
	// - For "quota_policy": "static" or "dynamic"
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *interface{} `json:"value,omitempty"`
	// The UUID of the volume to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UpdateVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateVolumesRequestItem
	return json.Marshal((Alias)(m))
}

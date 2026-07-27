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

// A single update operation to be applied to a template volume.
type UpdateTemplateVolumesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop MutableTemplateVolumeProperty `json:"prop"`
	// The operation to perform.
	Op MutableTemplateVolumeOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *interface{} `json:"value,omitempty"`
	// The UUID of the template volume to update.  Mutually exclusive with
	// name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the template volume to update.  Mutually exclusive with
	// UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateTemplateVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateTemplateVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateTemplateVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateTemplateVolumesRequestItem
	return json.Marshal((Alias)(m))
}

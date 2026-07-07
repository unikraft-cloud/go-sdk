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

// A single update operation to be applied to a checkpoint instance.
type UpdateCheckpointInstancesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the checkpoint instance.
	Metro *string `json:"metro,omitempty"`
	// The property to modify.
	Prop MutableCheckpointInstanceProperty `json:"prop"`
	// The operation to perform on the property.
	Op MutableCheckpointInstanceOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitempty"`
	// The UUID of the checkpoint instance to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the checkpoint instance to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UpdateCheckpointInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateCheckpointInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateCheckpointInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateCheckpointInstancesRequestItem
	return json.Marshal((Alias)(m))
}

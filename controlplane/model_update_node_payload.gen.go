// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// UpdateNodePayload contains the changes to apply to a node.
type UpdateNodePayload struct {
	// The property to update.
	Property MutableNodeProperty `json:"property"`
	// The operation to perform on the property.
	Operation MutableNodeOperation `json:"operation"`
	// The value for the update operation. The type depends on the property:
	// - TAGS: google.protobuf.Struct with key-value pairs
	// - DELETE_LOCK: google.protobuf.Value with boolean
	// - SSH_KEYS: google.protobuf.ListValue with SSHKey objects
	Value *string `json:"value,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateNodePayload) UnmarshalJSON(data []byte) error {
	type Alias UpdateNodePayload
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateNodePayload) MarshalJSON() ([]byte, error) {
	type Alias UpdateNodePayload
	return json.Marshal((Alias)(m))
}

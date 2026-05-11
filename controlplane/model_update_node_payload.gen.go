// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

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
	Value *string `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateNodePayload) UnmarshalJSON(data []byte) error {
	type Alias UpdateNodePayload
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"property":  {},
		"operation": {},
		"value":     {},
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

func (m UpdateNodePayload) MarshalJSON() ([]byte, error) {
	type Alias UpdateNodePayload
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

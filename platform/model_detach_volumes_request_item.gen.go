// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request of detaching a volume.

type DetachVolumesRequestItem struct {
	// (Optional).  UUID or name of the instance to detach the volume from.
	// If not specified, the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitempty"`
	// The UUID of the volume to detach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to detach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *DetachVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias DetachVolumesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"from": {},
		"uuid": {},
		"name": {},
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

func (m DetachVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias DetachVolumesRequestItem
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

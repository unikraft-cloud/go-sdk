// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for detaching one or more volume(s) from instances by
// their UUID(s) or name(s).

type DetachVolumesRequest struct {
	// The UUID of the volume to detach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to detach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string                   `json:"name,omitempty"`
	From *DetachVolumesRequestFrom `json:"from,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *DetachVolumesRequest) UnmarshalJSON(data []byte) error {
	type Alias DetachVolumesRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid": {},
		"name": {},
		"from": {},
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

func (m DetachVolumesRequest) MarshalJSON() ([]byte, error) {
	type Alias DetachVolumesRequest
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

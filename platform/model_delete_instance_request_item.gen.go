// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request item for deleting an instance by its UUID or name.

type DeleteInstanceRequestItem struct {
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// Timeout in seconds to wait for the instance to be deleted.  No wait
	// performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *DeleteInstanceRequestItem) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstanceRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"metro":     {},
		"timeout_s": {},
		"uuid":      {},
		"name":      {},
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

func (m DeleteInstanceRequestItem) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstanceRequestItem
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

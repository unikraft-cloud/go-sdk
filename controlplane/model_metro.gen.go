// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

type Metro struct {
	// The UUID of the metro.
	Uuid *string `json:"uuid,omitempty"`
	// The API endpoint for the metro.
	Endpoint *string `json:"endpoint,omitempty"`
	// The name of the metro.
	Name *string `json:"name,omitempty"`
	// The IATA code of the metro.
	IataCode *string `json:"iata_code,omitempty"`
	// The country where the metro is located.
	Country *string `json:"country,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *Metro) UnmarshalJSON(data []byte) error {
	type Alias Metro
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":      {},
		"endpoint":  {},
		"name":      {},
		"iata_code": {},
		"country":   {},
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

func (m Metro) MarshalJSON() ([]byte, error) {
	type Alias Metro
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

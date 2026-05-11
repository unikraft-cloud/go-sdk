// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// Information about an available region.

type Region struct {
	// The region identifier (e.g., "us-east-1", "us-central1").
	Name string `json:"name"`
	// Human-readable display name.
	DisplayName *string `json:"display_name,omitempty"`
	// The country code where this region is located.
	Country *string `json:"country,omitempty"`
	// Geographic coordinates of the region.
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	// Availability zones within this region.
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// Whether this region is currently available for provisioning.
	Available bool `json:"available"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *Region) UnmarshalJSON(data []byte) error {
	type Alias Region
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":               {},
		"display_name":       {},
		"country":            {},
		"latitude":           {},
		"longitude":          {},
		"availability_zones": {},
		"available":          {},
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

func (m Region) MarshalJSON() ([]byte, error) {
	type Alias Region
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

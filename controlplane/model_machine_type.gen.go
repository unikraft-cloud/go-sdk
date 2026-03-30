// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// Information about an available machine type.

type MachineType struct {
	// The machine type identifier (e.g., "m5.xlarge", "n2-standard-4").
	Name *string `json:"name,omitempty"`
	// Human-readable description.
	Description *string `json:"description,omitempty"`
	// Number of vCPUs.
	Vcpus *uint32 `json:"vcpus,omitempty"`
	// Memory in MiB.
	MemoryMib *uint64 `json:"memory_mib,omitempty"`
	// Machine category (e.g., "general-purpose", "compute-optimized",
	// "memory-optimized").
	Category *string `json:"category,omitempty"`
	// Regions where this machine type is available.
	AvailableRegions []string `json:"available_regions,omitempty"`
	// Whether this machine type supports nested virtualization.
	NestedVirtualization *bool `json:"nested_virtualization,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *MachineType) UnmarshalJSON(data []byte) error {
	type Alias MachineType
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":                  {},
		"description":           {},
		"vcpus":                 {},
		"memory_mib":            {},
		"category":              {},
		"available_regions":     {},
		"nested_virtualization": {},
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

func (m MachineType) MarshalJSON() ([]byte, error) {
	type Alias MachineType
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

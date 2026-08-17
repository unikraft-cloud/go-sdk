// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

// Information about an available machine type.
type MachineType struct {
	// The machine type identifier (e.g., "m5.xlarge", "n2-standard-4").
	Name string `json:"name"`
	// Human-readable description.
	Description *string `json:"description,omitzero"`
	// Number of vCPUs.
	Vcpus uint32 `json:"vcpus"`
	// Memory in MiB.
	MemoryMib uint64 `json:"memory_mib"`
	// Machine category (e.g., "general-purpose", "compute-optimized",
	// "memory-optimized").
	Category *string `json:"category,omitzero"`
	// Regions where this machine type is available.
	AvailableRegions []string `json:"available_regions,omitzero"`
	// Whether this machine type supports nested virtualization.
	NestedVirtualization bool `json:"nested_virtualization"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *MachineType) UnmarshalJSON(data []byte) error {
	type Alias MachineType
	return json.Unmarshal(data, (*Alias)(m))
}

func (m MachineType) MarshalJSON() ([]byte, error) {
	type Alias MachineType
	return json.Marshal((Alias)(m))
}

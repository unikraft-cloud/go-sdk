// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Additional limits

type QuotasLimits struct {
	// Minimum amount of memory assigned to live instances in megabytes
	MinMemoryMb *int64 `json:"min_memory_mb,omitempty"`
	// Maximum amount of memory assigned to live instances in megabytes
	MaxMemoryMb *int64 `json:"max_memory_mb,omitempty"`
	// Minimum size of a volume in megabytes
	MinVolumeMb *int64 `json:"min_volume_mb,omitempty"`
	// Maximum size of a volume in megabytes
	MaxVolumeMb *int64 `json:"max_volume_mb,omitempty"`
	// Minimum size of an autoscale group
	MinAutoscaleSize *int64 `json:"min_autoscale_size,omitempty"`
	// Maximum size of an autoscale group
	MaxAutoscaleSize *int64 `json:"max_autoscale_size,omitempty"`
	// Minimum number of vCPUs
	MinVcpus *int64 `json:"min_vcpus,omitempty"`
	// Maximum number of vCPUs
	MaxVcpus *int64 `json:"max_vcpus,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *QuotasLimits) UnmarshalJSON(data []byte) error {
	type Alias QuotasLimits
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"min_memory_mb":      {},
		"max_memory_mb":      {},
		"min_volume_mb":      {},
		"max_volume_mb":      {},
		"min_autoscale_size": {},
		"max_autoscale_size": {},
		"min_vcpus":          {},
		"max_vcpus":          {},
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

func (m QuotasLimits) MarshalJSON() ([]byte, error) {
	type Alias QuotasLimits
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Configured quota limits

type QuotasHard struct {
	// Number of instances
	Instances *int64 `json:"instances,omitempty"`
	// Number of instances that are not in the `stopped` state
	LiveInstances *int64 `json:"live_instances,omitempty"`
	// Number of vCPUs
	LiveVcpus *int64 `json:"live_vcpus,omitempty"`
	// Amount of memory assigned to instances that are not in the `stopped`
	// state in megabytes
	LiveMemoryMb *int64 `json:"live_memory_mb,omitempty"`
	// Number of services
	ServiceGroups *int64 `json:"service_groups,omitempty"`
	// Number of published network ports over all existing services
	Services *int64 `json:"services,omitempty"`
	// Number of volumes
	Volumes *int64 `json:"volumes,omitempty"`
	// Total size of all volumes in megabytes
	TotalVolumeMb *int64 `json:"total_volume_mb,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *QuotasHard) UnmarshalJSON(data []byte) error {
	type Alias QuotasHard
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"instances":       {},
		"live_instances":  {},
		"live_vcpus":      {},
		"live_memory_mb":  {},
		"service_groups":  {},
		"services":        {},
		"volumes":         {},
		"total_volume_mb": {},
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

func (m QuotasHard) MarshalJSON() ([]byte, error) {
	type Alias QuotasHard
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

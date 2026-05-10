// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type UserVmdb struct {
	// Maximum number of VM instances the user can have at one moment.
	MaxInstances *int32 `json:"max_instances,omitempty"`
	// Minimum amount of memory assigned to a VM in MB.
	MinMemoryMb *int32 `json:"min_memory_mb,omitempty"`
	// Default amount of memory assigned to a VM in MB.
	DefMemoryMb *int32 `json:"def_memory_mb,omitempty"`
	// Maximum amount of memory assigned to a VM in MB.
	MaxMemoryMb *int32 `json:"max_memory_mb,omitempty"`
	// Minimum number of vCPUs assigned to a VM.
	MinVcpus *int32 `json:"min_vcpus,omitempty"`
	// Maximum number of vCPUs assigned to a VM.
	MaxVcpus *int32 `json:"max_vcpus,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UserVmdb) UnmarshalJSON(data []byte) error {
	type Alias UserVmdb
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"max_instances": {},
		"min_memory_mb": {},
		"def_memory_mb": {},
		"max_memory_mb": {},
		"min_vcpus":     {},
		"max_vcpus":     {},
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

func (m UserVmdb) MarshalJSON() ([]byte, error) {
	type Alias UserVmdb
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

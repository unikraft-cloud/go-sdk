// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type UserVmm struct {
	// Maximum number of vCPUs the user can have assigned to live instances.
	MaxVcpus *int32 `json:"max_vcpus,omitempty"`
	// Maximum amount of memory in MB the user can have assigned to live
	// instances.
	MaxMemoryMb *int32 `json:"max_memory_mb,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UserVmm) UnmarshalJSON(data []byte) error {
	type Alias UserVmm
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UserVmm) MarshalJSON() ([]byte, error) {
	type Alias UserVmm
	return json.Marshal((Alias)(m))
}

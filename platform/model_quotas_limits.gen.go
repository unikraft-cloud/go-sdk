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

type QuotasLimits struct {
	// Minimum amount of memory assigned to live instances in megabytes
	MinMemoryMb int64 `json:"min_memory_mb"`
	// Maximum amount of memory assigned to live instances in megabytes
	MaxMemoryMb int64 `json:"max_memory_mb"`
	// Minimum size of a volume in megabytes
	MinVolumeMb int64 `json:"min_volume_mb"`
	// Maximum size of a volume in megabytes
	MaxVolumeMb int64 `json:"max_volume_mb"`
	// Minimum size of an autoscale group
	MinAutoscaleSize int64 `json:"min_autoscale_size"`
	// Maximum size of an autoscale group
	MaxAutoscaleSize int64 `json:"max_autoscale_size"`
	// Minimum number of vCPUs
	MinVcpus int64 `json:"min_vcpus"`
	// Maximum number of vCPUs
	MaxVcpus int64 `json:"max_vcpus"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *QuotasLimits) UnmarshalJSON(data []byte) error {
	type Alias QuotasLimits
	return json.Unmarshal(data, (*Alias)(m))
}

func (m QuotasLimits) MarshalJSON() ([]byte, error) {
	type Alias QuotasLimits
	return json.Marshal((Alias)(m))
}

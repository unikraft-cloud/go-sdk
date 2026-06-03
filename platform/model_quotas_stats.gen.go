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

type QuotasStats struct {
	// Number of instances
	Instances int64 `json:"instances"`
	// Number of instances that are not in the `stopped` state
	LiveInstances int64 `json:"live_instances"`
	// Number of vCPUs
	LiveVcpus int64 `json:"live_vcpus"`
	// Amount of memory assigned to instances that are not in the `stopped`
	// state in megabytes
	LiveMemoryMb int64 `json:"live_memory_mb"`
	// Number of services
	ServiceGroups int64 `json:"service_groups"`
	// Number of published network ports over all existing services
	Services int64 `json:"services"`
	// Number of volumes
	Volumes int64 `json:"volumes"`
	// Total size of all volumes in megabytes
	TotalVolumeMb int64 `json:"total_volume_mb"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *QuotasStats) UnmarshalJSON(data []byte) error {
	type Alias QuotasStats
	return json.Unmarshal(data, (*Alias)(m))
}

func (m QuotasStats) MarshalJSON() ([]byte, error) {
	type Alias QuotasStats
	return json.Marshal((Alias)(m))
}

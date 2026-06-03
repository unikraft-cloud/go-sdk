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

// A single request item to suspend an instance.

type SuspendInstancesRequestItem struct {
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Timeout for draining connections in milliseconds.  No draining
	// will occur if set to 0.  Use -1 for the largest possible value.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitempty"`
	// The UUID of the instance to suspend.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to suspend.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *SuspendInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m SuspendInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstancesRequestItem
	return json.Marshal((Alias)(m))
}

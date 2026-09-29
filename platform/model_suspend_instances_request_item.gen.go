// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A single request item to suspend an instance.
type SuspendInstancesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`
	// Timeout for draining connections in milliseconds. No draining
	// will occur if set to 0. Use -1 for the largest possible value.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *SuspendInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m SuspendInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstancesRequestItem
	return json.Marshal((Alias)(m))
}

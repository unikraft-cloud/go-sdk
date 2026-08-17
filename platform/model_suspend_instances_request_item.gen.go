// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

// A single request item to suspend an instance.
type SuspendInstancesRequestItem struct {
	// Timeout for draining connections in milliseconds.  No draining
	// will occur if set to 0.  Use -1 for the largest possible value.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitzero"`
	// The UUID of the instance to suspend.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the instance to suspend.  Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

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

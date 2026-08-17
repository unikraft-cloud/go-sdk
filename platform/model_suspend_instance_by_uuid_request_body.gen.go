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

// Parameters for suspending the instance.
type SuspendInstanceByUUIDRequestBody struct {
	// Timeout for draining connections in milliseconds.  No draining
	// will occur if set to 0.  Use -1 for the largest possible value.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *SuspendInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m SuspendInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

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

// Parameters for starting the instance.
type StartInstanceByUUIDRequestBody struct {
	// Deprecated: Use `timeout_s` instead.  Timeout in milliseconds to
	// wait for the instance to reach running state.  If `timeout_s` is
	// not set, this value is converted by rounding up to the next full
	// second.  No wait performed for a value of 0.
	WaitTimeoutMs *int64 `json:"wait_timeout_ms,omitzero"`
	// Timeout in seconds to wait for the instance to reach running
	// state.  If you start your instance, you can wait for it to
	// finish starting with a blocking API call if you specify a wait
	// timeout greater than zero.  No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *StartInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias StartInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m StartInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias StartInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

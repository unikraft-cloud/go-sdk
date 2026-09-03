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

// A single request item to start an instance.
type StartInstancesRequestItem struct {
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
	// The UUID of the instance to start.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the instance to start.  Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *StartInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias StartInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m StartInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias StartInstancesRequestItem
	return json.Marshal((Alias)(m))
}

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

// Automatic delete-on-idle/request-limit configuration for non-template instances.
// Not used for template instances.

type InstanceAutokill struct {
	// Time in milliseconds after the instance was stopped before it is deleted.
	// A value of 0 disables time-based autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
	// Maximum number of requests/connections the instance serves before it is
	// deleted. A value of 0 disables request-based autokill.
	NumRequests *uint32 `json:"num_requests,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *InstanceAutokill) UnmarshalJSON(data []byte) error {
	type Alias InstanceAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceAutokill) MarshalJSON() ([]byte, error) {
	type Alias InstanceAutokill
	return json.Marshal((Alias)(m))
}

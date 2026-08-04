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

// Automatic delete-on-idle/request-limit configuration for non-template
// instances.
type CreateInstanceRequestAutokill struct {
	// Time in milliseconds after the instance was stopped before it is deleted.
	// A value of 0 disables time-based autokill.
	TimeMs *uint64 `json:"time_ms,omitzero"`
	// Maximum number of requests/connections the instance serves before it is
	// deleted. A value of 0 disables request-based autokill.
	NumRequests *uint32 `json:"num_requests,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestAutokill) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestAutokill) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestAutokill
	return json.Marshal((Alias)(m))
}

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

// Automatic delete-on-idle configuration for service groups.
type ServiceGroupAutokill struct {
	// Time in milliseconds after the service group becomes empty before it is
	// deleted. A value of 0 disables autokill.
	TimeMs *uint64 `json:"time_ms,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ServiceGroupAutokill) UnmarshalJSON(data []byte) error {
	type Alias ServiceGroupAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ServiceGroupAutokill) MarshalJSON() ([]byte, error) {
	type Alias ServiceGroupAutokill
	return json.Marshal((Alias)(m))
}

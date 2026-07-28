// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Information about an available region.
type Region struct {
	// The region identifier (e.g., "us-east-1", "us-central1").
	Name string `json:"name,omitzero"`
	// The IATA code for the region.
	Iata string `json:"iata,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Region) UnmarshalJSON(data []byte) error {
	type Alias Region
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Region) MarshalJSON() ([]byte, error) {
	type Alias Region
	return json.Marshal((Alias)(m))
}

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

type Metro struct {
	// The UUID of the metro.
	Uuid string `json:"uuid,omitzero"`
	// The API endpoint for the metro.
	Endpoint string `json:"endpoint,omitzero"`
	// The name of the metro.
	Name string `json:"name,omitzero"`
	// The IATA code of the metro.
	IataCode string `json:"iata_code,omitzero"`
	// The country where the metro is located.
	Country string `json:"country,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Metro) UnmarshalJSON(data []byte) error {
	type Alias Metro
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Metro) MarshalJSON() ([]byte, error) {
	type Alias Metro
	return json.Marshal((Alias)(m))
}

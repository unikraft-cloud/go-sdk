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

// License information (admin only).
type DataLicense struct {
	// The serial number of the license certificate, hex-encoded.
	Serial string `json:"serial"`
	// Whether the license is currently valid.
	Valid bool `json:"valid"`
	// List of enabled features.
	Features []string `json:"features,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DataLicense) UnmarshalJSON(data []byte) error {
	type Alias DataLicense
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DataLicense) MarshalJSON() ([]byte, error) {
	type Alias DataLicense
	return json.Marshal((Alias)(m))
}

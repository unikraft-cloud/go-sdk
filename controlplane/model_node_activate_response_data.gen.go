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

type NodeActivateResponseData struct {
	// The issued license certificate in base64 URL encoded PEM format.
	License string `json:"license,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NodeActivateResponseData) UnmarshalJSON(data []byte) error {
	type Alias NodeActivateResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeActivateResponseData) MarshalJSON() ([]byte, error) {
	type Alias NodeActivateResponseData
	return json.Marshal((Alias)(m))
}

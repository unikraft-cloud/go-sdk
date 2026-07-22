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

type NodeRenewResponseData struct {
	// The renewed license certificate in base64 URL encoded PEM format.
	License *string `json:"license,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *NodeRenewResponseData) UnmarshalJSON(data []byte) error {
	type Alias NodeRenewResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeRenewResponseData) MarshalJSON() ([]byte, error) {
	type Alias NodeRenewResponseData
	return json.Marshal((Alias)(m))
}

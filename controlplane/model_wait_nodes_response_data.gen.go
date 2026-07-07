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

type WaitNodesResponseData struct {
	// The nodes after reaching the desired state.
	Nodes []Node `json:"nodes,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *WaitNodesResponseData) UnmarshalJSON(data []byte) error {
	type Alias WaitNodesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m WaitNodesResponseData) MarshalJSON() ([]byte, error) {
	type Alias WaitNodesResponseData
	return json.Marshal((Alias)(m))
}

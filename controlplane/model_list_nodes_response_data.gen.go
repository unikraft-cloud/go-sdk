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

type ListNodesResponseData struct {
	// The list of nodes matching the request.
	Nodes []Node `json:"nodes,omitzero"`
	// Total count of nodes matching the filters (for pagination).
	TotalCount uint32 `json:"total_count"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ListNodesResponseData) UnmarshalJSON(data []byte) error {
	type Alias ListNodesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ListNodesResponseData) MarshalJSON() ([]byte, error) {
	type Alias ListNodesResponseData
	return json.Marshal((Alias)(m))
}

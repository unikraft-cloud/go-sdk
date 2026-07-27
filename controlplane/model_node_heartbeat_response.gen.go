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

// The response message for a node heartbeat.
type NodeHeartbeatResponse struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the response.
	Message *string `json:"message,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NodeHeartbeatResponse) UnmarshalJSON(data []byte) error {
	type Alias NodeHeartbeatResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeHeartbeatResponse) MarshalJSON() ([]byte, error) {
	type Alias NodeHeartbeatResponse
	return json.Marshal((Alias)(m))
}

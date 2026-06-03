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

// The request message for a node heartbeat.

type NodeHeartbeatRequest struct {
	// The unique machine identifier of the node sending the heartbeat.
	MachineId string `json:"machine_id"`
	// The current platform status of the node.
	PlatformStatus PlatformStatus `json:"platform_status"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *NodeHeartbeatRequest) UnmarshalJSON(data []byte) error {
	type Alias NodeHeartbeatRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeHeartbeatRequest) MarshalJSON() ([]byte, error) {
	type Alias NodeHeartbeatRequest
	return json.Marshal((Alias)(m))
}

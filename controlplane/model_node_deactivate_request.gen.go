// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

// The request message for license deactivation. The caller is identified by
// an HTTP message signature (RFC 9421) made with the node's stable private
// key, the same mechanism used by the private node-facing APIs (e.g.
// NodeHeartbeat) -- not by any field in this message.
type NodeDeactivateRequest struct {
	// The serial number of the last issued certificate.
	Serial string `json:"serial"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NodeDeactivateRequest) UnmarshalJSON(data []byte) error {
	type Alias NodeDeactivateRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeDeactivateRequest) MarshalJSON() ([]byte, error) {
	type Alias NodeDeactivateRequest
	return json.Marshal((Alias)(m))
}

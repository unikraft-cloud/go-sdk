// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// Request message for registering a node that an operator or a provisioner
// runs.
type RegisterNodeRequest struct {
	// The name or UUID of the metro that the node joins.
	Metro string `json:"metro"`
	// The URL of the platform API of the node.  It must be an http(s) URL
	// with only a host and an optional port.
	ApiEndpoint string `json:"api_endpoint"`
	// Optional name for the node.  It is also the hostname of the node.  If
	// you do not provide it, the control plane generates a name.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *RegisterNodeRequest) UnmarshalJSON(data []byte) error {
	type Alias RegisterNodeRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m RegisterNodeRequest) MarshalJSON() ([]byte, error) {
	type Alias RegisterNodeRequest
	return json.Marshal((Alias)(m))
}

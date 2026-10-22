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

type RegisterNodeResponseData struct {
	// The registered node.
	Node Node `json:"node"`
	// The single-use activation token of the node.  The node activates with
	// it and takes over this node record.  The control plane returns the
	// token only once.
	ActivationToken string `json:"activation_token"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *RegisterNodeResponseData) UnmarshalJSON(data []byte) error {
	type Alias RegisterNodeResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m RegisterNodeResponseData) MarshalJSON() ([]byte, error) {
	type Alias RegisterNodeResponseData
	return json.Marshal((Alias)(m))
}

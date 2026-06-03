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

// The request message for certificate activation.
type NodeActivateRequest struct {
	// The certificate signing request (CSR) for the license which is base64
	// encoded.
	Csr string `json:"csr"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *NodeActivateRequest) UnmarshalJSON(data []byte) error {
	type Alias NodeActivateRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeActivateRequest) MarshalJSON() ([]byte, error) {
	type Alias NodeActivateRequest
	return json.Marshal((Alias)(m))
}

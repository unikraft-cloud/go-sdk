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
	// An opaque secret for first-time activation. On renewal (where the node's
	// public key is already known), this field is omitted and the CSR
	// self-signature is used as proof of key possession.
	Secret *string `json:"secret,omitzero"`
	// The serial number of the last issued certificate.
	Serial *string `json:"serial,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NodeActivateRequest) UnmarshalJSON(data []byte) error {
	type Alias NodeActivateRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NodeActivateRequest) MarshalJSON() ([]byte, error) {
	type Alias NodeActivateRequest
	return json.Marshal((Alias)(m))
}

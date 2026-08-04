// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type UpdateCertificateByUUIDRequestBody struct {
	// The new certificate chain.
	//
	// This is the public chain of the certificate in PEM format. The chain
	// should include the certificate and any intermediate certificates.
	Chain string `json:"chain,omitzero"`
	// The new private key.
	//
	// This is the private key of the certificate in PEM format. The private
	// key must match the public key in the certificate chain.
	Pkey string `json:"pkey,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateCertificateByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateCertificateByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateCertificateByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateCertificateByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

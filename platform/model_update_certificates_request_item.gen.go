// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A single update operation to be applied to a certificate.
type UpdateCertificatesRequestItem struct {
	// The new certificate chain.
	//
	// This is the public chain of the certificate in PEM format. The chain
	// should include the certificate and any intermediate certificates.
	Chain string `json:"chain"`
	// The new private key.
	//
	// This is the private key of the certificate in PEM format. The private
	// key must match the public key in the certificate chain.
	Pkey string `json:"pkey"`
	// The UUID of the certificate to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the certificate to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateCertificatesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateCertificatesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateCertificatesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateCertificatesRequestItem
	return json.Marshal((Alias)(m))
}

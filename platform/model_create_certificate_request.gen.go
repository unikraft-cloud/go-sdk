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

// The request message for creating/uploading a new certificate.
type CreateCertificateRequest struct {
	// The name of the certificate.
	//
	// This is a human-readable name that can be used to identify the certificate.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name is generated for you.  The name can also be used
	// to identify the certificate in API calls.
	Name *string `json:"name,omitzero"`
	// The common name (CN) of the certificate.
	//
	// Deprecated: Use `common_name` instead.
	Cn *string `json:"cn,omitzero"`
	// The common name (CN) of the certificate.
	//
	// This must be a fully-qualified domain name (FQDN). Exactly one of `cn`
	// or `common_name` must be specified.
	CommonName *string `json:"common_name,omitzero"`
	// The chain of the certificate.
	Chain string `json:"chain"`
	// The private key of the certificate.
	Pkey string `json:"pkey"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateCertificateRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateCertificateRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateCertificateRequest
	return json.Marshal((Alias)(m))
}

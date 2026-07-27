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

type CreateCertificateResponseData struct {
	// The certificate which was created by this request.
	//
	// Note: only one certificate can be specified in the request, so this
	// will always contain a single entry.
	Certificates []Certificate `json:"certificates,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateCertificateResponseData) UnmarshalJSON(data []byte) error {
	type Alias CreateCertificateResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateCertificateResponseData) MarshalJSON() ([]byte, error) {
	type Alias CreateCertificateResponseData
	return json.Marshal((Alias)(m))
}

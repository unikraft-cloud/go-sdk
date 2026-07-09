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

// The request message for updating one or more certificate(s) by their
// UUID(s) or name(s).
type UpdateCertificatesRequest struct {
	// A list of update operations to apply to certificates.
	Body []UpdateCertificatesRequestItem `json:"body"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UpdateCertificatesRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateCertificatesRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateCertificatesRequest) MarshalJSON() ([]byte, error) {
	type Alias UpdateCertificatesRequest
	return json.Marshal((Alias)(m))
}

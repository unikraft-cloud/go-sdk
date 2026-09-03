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

// A domain name
type CreateServiceGroupRequestDomain struct {
	// Publicly accessible domain name.  If this name ends in a period `.` it must
	// be a valid Full Qualified Domain Name (FQDN), otherwise it will become a
	// subdomain of the target metro.
	Name string `json:"name"`
	// Use an existing certificate for the domain.  If this field is
	// specified, the domain must be associated with a valid certificate.
	Certificate *NameOrUUID `json:"certificate,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateServiceGroupRequestDomain) UnmarshalJSON(data []byte) error {
	type Alias CreateServiceGroupRequestDomain
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateServiceGroupRequestDomain) MarshalJSON() ([]byte, error) {
	type Alias CreateServiceGroupRequestDomain
	return json.Marshal((Alias)(m))
}

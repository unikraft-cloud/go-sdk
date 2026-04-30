// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A domain name

type CreateServiceGroupRequestDomain struct {
	// Publicly accessible domain name.  If this name ends in a period `.` it must
	// be a valid Full Qualified Domain Name (FQDN), otherwise it will become a
	// subdomain of the target metro.
	Name string `json:"name"`
	// Use an existing certificate for the domain.  If this field is
	// specified, the domain must be associated with a valid certificate.
	Certificate *NameOrUUID `json:"certificate,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateServiceGroupRequestDomain) UnmarshalJSON(data []byte) error {
	type Alias CreateServiceGroupRequestDomain
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":        {},
		"certificate": {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m CreateServiceGroupRequestDomain) MarshalJSON() ([]byte, error) {
	type Alias CreateServiceGroupRequestDomain
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

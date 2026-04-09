// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Use an existing certificate for the domain.  If this field is
// specified, the domain must be associated with a valid certificate.

type CreateServiceGroupRequestDomainCertificate struct {
	// (Only applies when using global control plane).
	// The metro of the resource.
	Metro *string `json:"metro,omitempty"`
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateServiceGroupRequestDomainCertificate) UnmarshalJSON(data []byte) error {
	type Alias CreateServiceGroupRequestDomainCertificate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"metro": {},
		"uuid":  {},
		"name":  {},
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

func (m CreateServiceGroupRequestDomainCertificate) MarshalJSON() ([]byte, error) {
	type Alias CreateServiceGroupRequestDomainCertificate
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

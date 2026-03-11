// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The domain configuration for the service group.
//
// Domain names are completely specified with all labels in the hierarchy of
// the DNS, having no parts omitted.  The domain can be associated with an
// existing certificate by specifying the certificate's name or UUID.  If no
// certificate is specified and a FQDN is provided, Unikraft Cloud will
// automatically generate a new certificate for the domain based on Let's
// Encrypt and seek to accomplish a DNS-01 challenge.

type ServiceGroupInstanceDomain struct {
	// Publicly accessible domain name.
	//
	// If this name ends in a period `.` it must be a valid Full Qualified
	// Domain Name (FQDN), otherwise it will become a subdomain of the target
	// metro.
	Fqdn *string `json:"fqdn,omitempty"`
	// The certificate associated with the domain.
	//
	// The certificate is used to secure the domain with TLS/SSL.  If no
	// certificate is specified, Unikraft Cloud will automatically generate a
	// new certificate for the domain based on Let's Encrypt and seek to
	// accomplish a DNS-01 challenge.
	Certificate *NameOrUUID `json:"certificate,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *ServiceGroupInstanceDomain) UnmarshalJSON(data []byte) error {
	type Alias ServiceGroupInstanceDomain
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"fqdn":        {},
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

func (m ServiceGroupInstanceDomain) MarshalJSON() ([]byte, error) {
	type Alias ServiceGroupInstanceDomain
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

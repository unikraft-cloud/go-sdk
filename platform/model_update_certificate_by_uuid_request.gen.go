// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for updating a certificate by its UUID.

type UpdateCertificateByUUIDRequest struct {
	// The UUID of the certificate to update.
	Uuid string `json:"uuid"`
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

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateCertificateByUUIDRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateCertificateByUUIDRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":  {},
		"chain": {},
		"pkey":  {},
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

func (m UpdateCertificateByUUIDRequest) MarshalJSON() ([]byte, error) {
	type Alias UpdateCertificateByUUIDRequest
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

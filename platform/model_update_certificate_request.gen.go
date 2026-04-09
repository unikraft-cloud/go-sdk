// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for updating one or more certificate(s) by their
// UUID(s) or name(s).

type UpdateCertificateRequest struct {
	Id UpdateCertificateRequestId `json:"id"`
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

func (m *UpdateCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateCertificateRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
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

func (m UpdateCertificateRequest) MarshalJSON() ([]byte, error) {
	type Alias UpdateCertificateRequest
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

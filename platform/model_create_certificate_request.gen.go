// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for creating/uploading a new certificate.

type CreateCertificateRequest struct {
	// The common name (CN) of the certificate.
	Cn string `json:"cn"`
	// The chain of the certificate.
	Chain string `json:"chain"`
	// The private key of the certificate.
	Pkey string `json:"pkey"`
	// The name of the certificate.
	//
	// This is a human-readable name that can be used to identify the certificate.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name is generated for you.  The name can also be used
	// to identify the certificate in API calls.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateCertificateRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"cn":    {},
		"chain": {},
		"pkey":  {},
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

func (m CreateCertificateRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateCertificateRequest
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

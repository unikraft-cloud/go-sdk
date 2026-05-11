// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Details of the certificate which was deleted by this request.

type DeleteCertificatesResponseDeletedCertificate struct {
	// Indicates whether the delete operation was successful or not for this
	// certificate.
	Status ResponseStatus `json:"status"`
	// The UUID of the certificate which was deleted.
	Uuid string `json:"uuid"`
	// The name of the certificate which was deleted.
	Name string `json:"name"`
	// (Only applies when using global control plane).
	// The metro of the certificate.
	Metro *string `json:"metro,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *DeleteCertificatesResponseDeletedCertificate) UnmarshalJSON(data []byte) error {
	type Alias DeleteCertificatesResponseDeletedCertificate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"status":  {},
		"uuid":    {},
		"name":    {},
		"metro":   {},
		"message": {},
		"error":   {},
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

func (m DeleteCertificatesResponseDeletedCertificate) MarshalJSON() ([]byte, error) {
	type Alias DeleteCertificatesResponseDeletedCertificate
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

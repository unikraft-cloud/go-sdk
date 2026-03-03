// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// The response data for this request.

type RequestSigninResponseData struct {
	// The full URL which can be used to remotely confirm the signin.
	AuthorizationUrl *string `json:"authorization_url,omitempty"`
	// A unique identifier for the request.  This can be used to track the
	// request in the system.
	RequestId *string `json:"request_id,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *RequestSigninResponseData) UnmarshalJSON(data []byte) error {
	type Alias RequestSigninResponseData
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"authorization_url": {},
		"request_id":        {},
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

func (m RequestSigninResponseData) MarshalJSON() ([]byte, error) {
	type Alias RequestSigninResponseData
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

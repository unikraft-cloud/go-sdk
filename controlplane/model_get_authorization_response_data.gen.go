// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// The response data for this request.

type GetAuthorizationResponseData struct {
	// The organization name the token is associated with.
	OrganizationName *string `json:"organization_name,omitempty"`
	// The display name of the organization the token is associated with.
	OrganizationDisplayName *string `json:"organization_display_name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *GetAuthorizationResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetAuthorizationResponseData
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"organization_name":         {},
		"organization_display_name": {},
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

func (m GetAuthorizationResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetAuthorizationResponseData
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

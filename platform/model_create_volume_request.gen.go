// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for creating a volume.

type CreateVolumeRequest struct {
	// The size of the volume in megabytes.
	SizeMb uint64 `json:"size_mb"`
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name of the form `vol-X` is generated for you, where
	// `X` is a 5 character long random alphanumeric suffix..  The name can also
	// be used to identify the volume in API calls.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateVolumeRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateVolumeRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"size_mb": {},
		"name":    {},
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

func (m CreateVolumeRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateVolumeRequest
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

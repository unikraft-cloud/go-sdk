// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type CloneVolumeByUUIDRequestBody struct {
	// The name of the new cloned volume.
	VolName *string `json:"vol_name,omitempty"`
	// The tags associated with the volume.
	// Maximum 16 tags are allowed, and each tag may not be longer than 256 characters.
	Tags []string `json:"tags,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CloneVolumeByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumeByUUIDRequestBody
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"vol_name": {},
		"tags":     {},
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

func (m CloneVolumeByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumeByUUIDRequestBody
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

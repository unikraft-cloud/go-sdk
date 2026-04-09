// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The quota policy for the new cloned volume.  If not provided, the quota
// policy of the source volume is used.
type CloneVolumeByUUIDRequestBodyQuotaPolicy string

const (
	CloneVolumeByUUIDRequestBodyQuotaPolicyStatic  CloneVolumeByUUIDRequestBodyQuotaPolicy = "static"
	CloneVolumeByUUIDRequestBodyQuotaPolicyDynamic CloneVolumeByUUIDRequestBodyQuotaPolicy = "dynamic"
)

type CloneVolumeByUUIDRequestBody struct {
	// The name of the new cloned volume.  If not provided, a random name
	// of the form `vol-X` is generated for you, where `X` is a 5 character
	// long random alphanumeric suffix.
	VolName *string `json:"vol_name,omitempty"`
	// The quota policy for the new cloned volume.  If not provided, the quota
	// policy of the source volume is used.
	QuotaPolicy *CloneVolumeByUUIDRequestBodyQuotaPolicy `json:"quota_policy,omitempty"`
	// A list of tags to assign to the new cloned volume.
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
		"vol_name":     {},
		"quota_policy": {},
		"tags":         {},
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

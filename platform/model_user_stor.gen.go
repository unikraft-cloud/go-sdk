// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type UserStor struct {
	// Maximum number of volumes the user can have at one moment.
	MaxVolumes *int32 `json:"max_volumes,omitempty"`
	// Minimum size of a volume in MB.
	MinVolumeMb *int32 `json:"min_volume_mb,omitempty"`
	// Maximum size of a volume in MB.
	MaxVolumeMb *int32 `json:"max_volume_mb,omitempty"`
	// Maximum total size of all volumes in MB.
	MaxTotalVolumeMb *int32 `json:"max_total_volume_mb,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UserStor) UnmarshalJSON(data []byte) error {
	type Alias UserStor
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"max_volumes":         {},
		"min_volume_mb":       {},
		"max_volume_mb":       {},
		"max_total_volume_mb": {},
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

func (m UserStor) MarshalJSON() ([]byte, error) {
	type Alias UserStor
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

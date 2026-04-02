// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request item for attaching a volume to an instance.

type AttachVolumesRequestItem struct {
	// The UUID of the volume to attach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid string `json:"uuid"`
	// The name of the volume to attach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name     string                           `json:"name"`
	AttachTo AttachVolumesRequestItemAttachTo `json:"attach_to"`
	// Path of the mountpoint.
	//
	// The path must be absolute, not contain `.` and `..` components, and not
	// contain colons (`:`). The path must point to an empty directory. If the
	// directory does not exist, it is created.
	At string `json:"at"`
	// Whether the volume should be mounted read-only.
	Readonly *bool `json:"readonly,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *AttachVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias AttachVolumesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":      {},
		"name":      {},
		"attach_to": {},
		"at":        {},
		"readonly":  {},
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

func (m AttachVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias AttachVolumesRequestItem
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

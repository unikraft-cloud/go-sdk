// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Current state of the volume.
type CloneVolumeResponseVolumeState string

const (
	CloneVolumeResponseVolumeStateUninitialized CloneVolumeResponseVolumeState = "uninitialized"
	CloneVolumeResponseVolumeStateInitializing  CloneVolumeResponseVolumeState = "initializing"
	CloneVolumeResponseVolumeStateAvailable     CloneVolumeResponseVolumeState = "available"
	CloneVolumeResponseVolumeStateIdle          CloneVolumeResponseVolumeState = "idle"
	CloneVolumeResponseVolumeStateMounted       CloneVolumeResponseVolumeState = "mounted"
	CloneVolumeResponseVolumeStateBusy          CloneVolumeResponseVolumeState = "busy"
	CloneVolumeResponseVolumeStateError         CloneVolumeResponseVolumeState = "error"
)

type CloneVolumeResponseVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// Current state of the volume.
	State *CloneVolumeResponseVolumeState `json:"state,omitempty"`
	// UUID of the newly created volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly created volume.
	Name *string `json:"name,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CloneVolumeResponseVolume) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumeResponseVolume
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"status":  {},
		"state":   {},
		"uuid":    {},
		"name":    {},
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

func (m CloneVolumeResponseVolume) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumeResponseVolume
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

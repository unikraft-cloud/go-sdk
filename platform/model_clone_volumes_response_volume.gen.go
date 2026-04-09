// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The state of the volume.
type CloneVolumesResponseVolumeState string

const (
	CloneVolumesResponseVolumeStateUninitialized CloneVolumesResponseVolumeState = "uninitialized"
	CloneVolumesResponseVolumeStateInitializing  CloneVolumesResponseVolumeState = "initializing"
	CloneVolumesResponseVolumeStateAvailable     CloneVolumesResponseVolumeState = "available"
	CloneVolumesResponseVolumeStateIdle          CloneVolumesResponseVolumeState = "idle"
	CloneVolumesResponseVolumeStateMounted       CloneVolumesResponseVolumeState = "mounted"
	CloneVolumesResponseVolumeStateBusy          CloneVolumesResponseVolumeState = "busy"
	CloneVolumesResponseVolumeStateError         CloneVolumesResponseVolumeState = "error"
	CloneVolumesResponseVolumeStateTemplate      CloneVolumesResponseVolumeState = "template"
)

type CloneVolumesResponseVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// The UUID of the newly cloned volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly cloned volume.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the volume.
	Metro *string `json:"metro,omitempty"`
	// The state of the volume.
	State *CloneVolumesResponseVolumeState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CloneVolumesResponseVolume) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumesResponseVolume
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
		"state":   {},
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

func (m CloneVolumesResponseVolume) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumesResponseVolume
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

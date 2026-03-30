// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The state of the volume.
type CreateTemplateVolumesResponseTemplateVolumeState string

const (
	CreateTemplateVolumesResponseTemplateVolumeStateUninitialized CreateTemplateVolumesResponseTemplateVolumeState = "uninitialized"
	CreateTemplateVolumesResponseTemplateVolumeStateInitializing  CreateTemplateVolumesResponseTemplateVolumeState = "initializing"
	CreateTemplateVolumesResponseTemplateVolumeStateAvailable     CreateTemplateVolumesResponseTemplateVolumeState = "available"
	CreateTemplateVolumesResponseTemplateVolumeStateIdle          CreateTemplateVolumesResponseTemplateVolumeState = "idle"
	CreateTemplateVolumesResponseTemplateVolumeStateMounted       CreateTemplateVolumesResponseTemplateVolumeState = "mounted"
	CreateTemplateVolumesResponseTemplateVolumeStateBusy          CreateTemplateVolumesResponseTemplateVolumeState = "busy"
	CreateTemplateVolumesResponseTemplateVolumeStateError         CreateTemplateVolumesResponseTemplateVolumeState = "error"
	CreateTemplateVolumesResponseTemplateVolumeStateTemplate      CreateTemplateVolumesResponseTemplateVolumeState = "template"
)

type CreateTemplateVolumesResponseTemplateVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// The UUID of the volume converted into a template.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume converted into a template.
	Name *string `json:"name,omitempty"`
	// The state of the volume.
	State *CreateTemplateVolumesResponseTemplateVolumeState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateTemplateVolumesResponseTemplateVolume) UnmarshalJSON(data []byte) error {
	type Alias CreateTemplateVolumesResponseTemplateVolume
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

func (m CreateTemplateVolumesResponseTemplateVolume) MarshalJSON() ([]byte, error) {
	type Alias CreateTemplateVolumesResponseTemplateVolume
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

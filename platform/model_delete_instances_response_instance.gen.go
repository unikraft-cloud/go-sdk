// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Details of the instance which was deleted by this request.

type DeleteInstancesResponseInstance struct {
	// Indicates whether the start operation was successful or not for this
	// instance.
	Status ResponseStatus `json:"status"`
	// The UUID of the instance which was deleted.
	Uuid string `json:"uuid"`
	// The name of the instance which was deleted.
	Name string `json:"name"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// The previous state of the instance before it was deleted.
	PreviousState string `json:"previous_state"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *DeleteInstancesResponseInstance) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstancesResponseInstance
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"status":         {},
		"uuid":           {},
		"name":           {},
		"metro":          {},
		"previous_state": {},
		"message":        {},
		"error":          {},
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

func (m DeleteInstancesResponseInstance) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstancesResponseInstance
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

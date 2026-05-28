// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type UpdateInstancesResponseUpdatedInstance struct {
	// The UUID of the instance that was updated.
	Uuid string `json:"uuid"`
	// The name of the instance that was updated.
	Name string `json:"name"`
	// The status of this particular instance update operation.
	Status ResponseStatus `json:"status"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// (Optional).  The client-provided ID from the request.
	Id *string `json:"id,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateInstancesResponseUpdatedInstance) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstancesResponseUpdatedInstance
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":    {},
		"name":    {},
		"status":  {},
		"metro":   {},
		"id":      {},
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

func (m UpdateInstancesResponseUpdatedInstance) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstancesResponseUpdatedInstance
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

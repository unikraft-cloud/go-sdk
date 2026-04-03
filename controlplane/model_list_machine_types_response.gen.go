// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// Response message for listing machine types.

type ListMachineTypesResponse struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the response.
	Message *string                       `json:"message,omitempty"`
	Data    *ListMachineTypesResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.
	OpTimeUs *uint64 `json:"op_time_us,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *ListMachineTypesResponse) UnmarshalJSON(data []byte) error {
	type Alias ListMachineTypesResponse
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"status":     {},
		"message":    {},
		"data":       {},
		"errors":     {},
		"op_time_us": {},
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

func (m ListMachineTypesResponse) MarshalJSON() ([]byte, error) {
	type Alias ListMachineTypesResponse
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request item to suspend an instance.

type SuspendInstancesRequestItem struct {
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Timeout for draining connections in milliseconds.  No draining
	// will occur if set to 0.  Use -1 for the largest possible value.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitempty"`
	// The UUID of the instance to suspend.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to suspend.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *SuspendInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"metro":            {},
		"drain_timeout_ms": {},
		"uuid":             {},
		"name":             {},
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

func (m SuspendInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstancesRequestItem
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

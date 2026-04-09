// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request item to start an instance.

type StartInstancesRequestItem struct {
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// Deprecated: Use `timeout_s` instead.  Timeout in milliseconds to
	// wait for the instance to reach running state.  If `timeout_s` is
	// not set, this value is converted by rounding up to the next full
	// second.  No wait performed for a value of 0.
	WaitTimeoutMs *int64 `json:"wait_timeout_ms,omitempty"`
	// Timeout in seconds to wait for the instance to reach running
	// state.  If you start your instance, you can wait for it to
	// finish starting with a blocking API call if you specify a wait
	// timeout greater than zero.  No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// The UUID of the instance to start.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to start.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *StartInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias StartInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"metro":           {},
		"wait_timeout_ms": {},
		"timeout_s":       {},
		"uuid":            {},
		"name":            {},
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

func (m StartInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias StartInstancesRequestItem
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Automatic delete-on-idle/request-limit configuration.
// Not used for template instances.

type InstanceAutokill struct {
	// Time in milliseconds after the instance was stopped before it is deleted.
	// A value of 0 disables time-based autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`
	// Maximum number of requests/connections the instance serves before it is
	// deleted. A value of 0 disables request-based autokill.
	NumRequests *uint32 `json:"num_requests,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstanceAutokill) UnmarshalJSON(data []byte) error {
	type Alias InstanceAutokill
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"time_ms":      {},
		"num_requests": {},
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

func (m InstanceAutokill) MarshalJSON() ([]byte, error) {
	type Alias InstanceAutokill
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

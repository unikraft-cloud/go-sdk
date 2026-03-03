// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json"
	"time"
)

// The current restart attempt for the instance.
// Not used for template instances.

type InstanceRestart struct {
	// Current restart attempt number. This is incremented each time the instance
	// is restarted automatically by the platform.
	Attempt *uint32 `json:"attempt,omitempty"`
	// Timestamp of the next scheduled restart attempt.
	NextAt *time.Time `json:"next_at,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstanceRestart) UnmarshalJSON(data []byte) error {
	type Alias InstanceRestart
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"attempt": {},
		"next_at": {},
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

func (m InstanceRestart) MarshalJSON() ([]byte, error) {
	type Alias InstanceRestart
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

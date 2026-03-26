// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request item to stop an instance.

type StopInstancesRequestItem struct {
	// The UUID of the instance to stop.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to stop.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// Whether to immediately force stop the instance.
	Force *bool `json:"force,omitempty"`
	// Timeout for draining connections in milliseconds.  The instance does not
	// receive new connections in the draining phase.  The instance is stopped
	// when the last connection has been closed or the timeout expired.  The
	// maximum timeout may vary.  Use -1 for the largest possible value.
	//
	// Note: This endpoint does not block.  Use the wait endpoint for the
	// instance to reach the stopped state.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitempty"`
	// Whether to perform a quick shutdown.  This flag is
	// overridden by force.
	Quick *bool `json:"quick,omitempty"`
	// Only stop the instance if it is in this state.
	Ifstate *string `json:"ifstate,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *StopInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias StopInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":             {},
		"name":             {},
		"force":            {},
		"drain_timeout_ms": {},
		"quick":            {},
		"ifstate":          {},
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

func (m StopInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias StopInstancesRequestItem
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

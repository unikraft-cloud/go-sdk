// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single wait operation to be applied to an instance.
// The desired state to wait for.  Default is `running`.
type WaitInstancesRequestItemState string

const (
	WaitInstancesRequestItemStateStopped  WaitInstancesRequestItemState = "stopped"
	WaitInstancesRequestItemStateStarting WaitInstancesRequestItemState = "starting"
	WaitInstancesRequestItemStateRunning  WaitInstancesRequestItemState = "running"
	WaitInstancesRequestItemStateDraining WaitInstancesRequestItemState = "draining"
	WaitInstancesRequestItemStateStopping WaitInstancesRequestItemState = "stopping"
	WaitInstancesRequestItemStateTemplate WaitInstancesRequestItemState = "template"
	WaitInstancesRequestItemStateStandby  WaitInstancesRequestItemState = "standby"
)

type WaitInstancesRequestItem struct {
	// The UUID of the instance to wait for.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to wait for.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// The desired state to wait for.  Default is `running`.
	State *WaitInstancesRequestItemState `json:"state,omitempty"`
	// Timeout in milliseconds to wait for the instance to reach the desired
	// state.  If the timeout is reached, the request will fail with an error.
	// A value of -1 means to wait indefinitely until the instance reaches the
	// desired state.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *WaitInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias WaitInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":       {},
		"name":       {},
		"state":      {},
		"timeout_ms": {},
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

func (m WaitInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias WaitInstancesRequestItem
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

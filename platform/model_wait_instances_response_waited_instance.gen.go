// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The current state of the instance.
type WaitInstancesResponseWaitedInstanceState string

const (
	WaitInstancesResponseWaitedInstanceStateStopped  WaitInstancesResponseWaitedInstanceState = "stopped"
	WaitInstancesResponseWaitedInstanceStateStarting WaitInstancesResponseWaitedInstanceState = "starting"
	WaitInstancesResponseWaitedInstanceStateRunning  WaitInstancesResponseWaitedInstanceState = "running"
	WaitInstancesResponseWaitedInstanceStateDraining WaitInstancesResponseWaitedInstanceState = "draining"
	WaitInstancesResponseWaitedInstanceStateStopping WaitInstancesResponseWaitedInstanceState = "stopping"
	WaitInstancesResponseWaitedInstanceStateTemplate WaitInstancesResponseWaitedInstanceState = "template"
	WaitInstancesResponseWaitedInstanceStateStandby  WaitInstancesResponseWaitedInstanceState = "standby"
)

type WaitInstancesResponseWaitedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// The current state of the instance.
	State *WaitInstancesResponseWaitedInstanceState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *WaitInstancesResponseWaitedInstance) UnmarshalJSON(data []byte) error {
	type Alias WaitInstancesResponseWaitedInstance
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
		"state":   {},
		"message": {},
		"error":   {},
		"status":  {},
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

func (m WaitInstancesResponseWaitedInstance) MarshalJSON() ([]byte, error) {
	type Alias WaitInstancesResponseWaitedInstance
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

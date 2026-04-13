// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The current state of the instance.
type SuspendInstancesResponseSuspendedInstanceState string

const (
	SuspendInstancesResponseSuspendedInstanceStateStopped  SuspendInstancesResponseSuspendedInstanceState = "stopped"
	SuspendInstancesResponseSuspendedInstanceStateStarting SuspendInstancesResponseSuspendedInstanceState = "starting"
	SuspendInstancesResponseSuspendedInstanceStateRunning  SuspendInstancesResponseSuspendedInstanceState = "running"
	SuspendInstancesResponseSuspendedInstanceStateDraining SuspendInstancesResponseSuspendedInstanceState = "draining"
	SuspendInstancesResponseSuspendedInstanceStateStopping SuspendInstancesResponseSuspendedInstanceState = "stopping"
	SuspendInstancesResponseSuspendedInstanceStateTemplate SuspendInstancesResponseSuspendedInstanceState = "template"
	SuspendInstancesResponseSuspendedInstanceStateStandby  SuspendInstancesResponseSuspendedInstanceState = "standby"
)

// The previous state of the instance before the suspend operation was invoked.
type SuspendInstancesResponseSuspendedInstancePreviousState string

const (
	SuspendInstancesResponseSuspendedInstancePreviousStateStopped  SuspendInstancesResponseSuspendedInstancePreviousState = "stopped"
	SuspendInstancesResponseSuspendedInstancePreviousStateStarting SuspendInstancesResponseSuspendedInstancePreviousState = "starting"
	SuspendInstancesResponseSuspendedInstancePreviousStateRunning  SuspendInstancesResponseSuspendedInstancePreviousState = "running"
	SuspendInstancesResponseSuspendedInstancePreviousStateDraining SuspendInstancesResponseSuspendedInstancePreviousState = "draining"
	SuspendInstancesResponseSuspendedInstancePreviousStateStopping SuspendInstancesResponseSuspendedInstancePreviousState = "stopping"
	SuspendInstancesResponseSuspendedInstancePreviousStateTemplate SuspendInstancesResponseSuspendedInstancePreviousState = "template"
	SuspendInstancesResponseSuspendedInstancePreviousStateStandby  SuspendInstancesResponseSuspendedInstancePreviousState = "standby"
)

type SuspendInstancesResponseSuspendedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// The current state of the instance.
	State *SuspendInstancesResponseSuspendedInstanceState `json:"state,omitempty"`
	// The previous state of the instance before the suspend operation was invoked.
	PreviousState *SuspendInstancesResponseSuspendedInstancePreviousState `json:"previous_state,omitempty"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *SuspendInstancesResponseSuspendedInstance) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstancesResponseSuspendedInstance
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":           {},
		"name":           {},
		"metro":          {},
		"state":          {},
		"previous_state": {},
		"status":         {},
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

func (m SuspendInstancesResponseSuspendedInstance) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstancesResponseSuspendedInstance
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

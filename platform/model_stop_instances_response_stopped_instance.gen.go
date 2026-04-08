// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The current state of the instance.
type StopInstancesResponseStoppedInstanceState string

const (
	StopInstancesResponseStoppedInstanceStateStopped  StopInstancesResponseStoppedInstanceState = "stopped"
	StopInstancesResponseStoppedInstanceStateStarting StopInstancesResponseStoppedInstanceState = "starting"
	StopInstancesResponseStoppedInstanceStateRunning  StopInstancesResponseStoppedInstanceState = "running"
	StopInstancesResponseStoppedInstanceStateDraining StopInstancesResponseStoppedInstanceState = "draining"
	StopInstancesResponseStoppedInstanceStateStopping StopInstancesResponseStoppedInstanceState = "stopping"
	StopInstancesResponseStoppedInstanceStateTemplate StopInstancesResponseStoppedInstanceState = "template"
	StopInstancesResponseStoppedInstanceStateStandby  StopInstancesResponseStoppedInstanceState = "standby"
)

// The previous state of the instance before the stop operation was invoked.
type StopInstancesResponseStoppedInstancePreviousState string

const (
	StopInstancesResponseStoppedInstancePreviousStateStopped  StopInstancesResponseStoppedInstancePreviousState = "stopped"
	StopInstancesResponseStoppedInstancePreviousStateStarting StopInstancesResponseStoppedInstancePreviousState = "starting"
	StopInstancesResponseStoppedInstancePreviousStateRunning  StopInstancesResponseStoppedInstancePreviousState = "running"
	StopInstancesResponseStoppedInstancePreviousStateDraining StopInstancesResponseStoppedInstancePreviousState = "draining"
	StopInstancesResponseStoppedInstancePreviousStateStopping StopInstancesResponseStoppedInstancePreviousState = "stopping"
	StopInstancesResponseStoppedInstancePreviousStateTemplate StopInstancesResponseStoppedInstancePreviousState = "template"
	StopInstancesResponseStoppedInstancePreviousStateStandby  StopInstancesResponseStoppedInstancePreviousState = "standby"
)

type StopInstancesResponseStoppedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// The current state of the instance.
	State *StopInstancesResponseStoppedInstanceState `json:"state,omitempty"`
	// The previous state of the instance before the stop operation was invoked.
	PreviousState *StopInstancesResponseStoppedInstancePreviousState `json:"previous_state,omitempty"`
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

func (m *StopInstancesResponseStoppedInstance) UnmarshalJSON(data []byte) error {
	type Alias StopInstancesResponseStoppedInstance
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

func (m StopInstancesResponseStoppedInstance) MarshalJSON() ([]byte, error) {
	type Alias StopInstancesResponseStoppedInstance
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// State of the instance when the logs were retrieved.
type GetInstancesLogsResponseLoggedInstanceState string

const (
	GetInstancesLogsResponseLoggedInstanceStateStopped  GetInstancesLogsResponseLoggedInstanceState = "stopped"
	GetInstancesLogsResponseLoggedInstanceStateStarting GetInstancesLogsResponseLoggedInstanceState = "starting"
	GetInstancesLogsResponseLoggedInstanceStateRunning  GetInstancesLogsResponseLoggedInstanceState = "running"
	GetInstancesLogsResponseLoggedInstanceStateDraining GetInstancesLogsResponseLoggedInstanceState = "draining"
	GetInstancesLogsResponseLoggedInstanceStateStopping GetInstancesLogsResponseLoggedInstanceState = "stopping"
	GetInstancesLogsResponseLoggedInstanceStateTemplate GetInstancesLogsResponseLoggedInstanceState = "template"
	GetInstancesLogsResponseLoggedInstanceStateStandby  GetInstancesLogsResponseLoggedInstanceState = "standby"
)

type GetInstancesLogsResponseLoggedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// Base64 encoded log output of the instance.
	Output *string `json:"output,omitempty"`
	// Description of the log availability.
	Available *GetInstancesLogsResponseLoggedInstanceAvailable `json:"available,omitempty"`
	// Description of the range that was returned.  Useful for requests with
	// offset relative to end.
	Range *GetInstancesLogsResponseLoggedInstanceRange `json:"range,omitempty"`
	// State of the instance when the logs were retrieved.
	State *GetInstancesLogsResponseLoggedInstanceState `json:"state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *GetInstancesLogsResponseLoggedInstance) UnmarshalJSON(data []byte) error {
	type Alias GetInstancesLogsResponseLoggedInstance
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":      {},
		"name":      {},
		"output":    {},
		"available": {},
		"range":     {},
		"state":     {},
		"message":   {},
		"error":     {},
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

func (m GetInstancesLogsResponseLoggedInstance) MarshalJSON() ([]byte, error) {
	type Alias GetInstancesLogsResponseLoggedInstance
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

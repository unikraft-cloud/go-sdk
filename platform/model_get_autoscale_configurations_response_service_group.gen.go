// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type GetAutoscaleConfigurationsResponseServiceGroup struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// The UUID of the service where the configuration was created.
	Uuid string `json:"uuid"`
	// The name of the service where the configuration was created.
	Name string `json:"name"`
	// (Only applies when using global control plane).
	// The metro of the service group.
	Metro *string `json:"metro,omitempty"`
	// If the autoscale configuration is enabled.
	Enabled bool `json:"enabled"`
	// The minimum number of instances to keep running.
	// Only if enabled is true.
	MinSize *int64 `json:"min_size,omitempty"`
	// The maximum number of instances to keep running.
	// Only if enabled is true.
	MaxSize *int64 `json:"max_size,omitempty"`
	// The warmup time in seconds for new instances.
	// Only if enabled is true.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitempty"`
	// The cooldown time in seconds for the autoscale configuration.
	// Only if enabled is true.
	CooldownTimeMs *int64 `json:"cooldown_time_ms,omitempty"`
	// The instance template used for the autoscale configuration.
	// Only if enabled is true.
	Template *ServiceGroupTemplate `json:"template,omitempty"`
	// The policies applied to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *GetAutoscaleConfigurationsResponseServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationsResponseServiceGroup
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"status":           {},
		"uuid":             {},
		"name":             {},
		"metro":            {},
		"enabled":          {},
		"min_size":         {},
		"max_size":         {},
		"warmup_time_ms":   {},
		"cooldown_time_ms": {},
		"template":         {},
		"policies":         {},
		"message":          {},
		"error":            {},
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

func (m GetAutoscaleConfigurationsResponseServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationsResponseServiceGroup
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

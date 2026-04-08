// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type CreateAutoscaleConfigurationsRequestConfiguration struct {
	// The UUID of the service to create a configuration for.
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service to create a configuration for.
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// (Optional, only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// The minimum number of instances to keep running.
	MinSize *int64 `json:"min_size,omitempty"`
	// The maximum number of instances to keep running.
	MaxSize *int64 `json:"max_size,omitempty"`
	// The warmup time in milliseconds for new instances.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitempty"`
	// The cooldown time in milliseconds for the autoscale configuration.
	CooldownTimeMs *int64                                                       `json:"cooldown_time_ms,omitempty"`
	CreateArgs     *CreateAutoscaleConfigurationsRequestConfigurationCreateArgs `json:"create_args,omitempty"`
	// The policies to apply to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateAutoscaleConfigurationsRequestConfiguration) UnmarshalJSON(data []byte) error {
	type Alias CreateAutoscaleConfigurationsRequestConfiguration
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
		"metro":            {},
		"min_size":         {},
		"max_size":         {},
		"warmup_time_ms":   {},
		"cooldown_time_ms": {},
		"create_args":      {},
		"policies":         {},
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

func (m CreateAutoscaleConfigurationsRequestConfiguration) MarshalJSON() ([]byte, error) {
	type Alias CreateAutoscaleConfigurationsRequestConfiguration
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

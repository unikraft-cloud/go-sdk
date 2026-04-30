// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message to create an autoscale configuration policy for a
// service.

type CreateAutoscaleConfigurationPolicyRequest struct {
	// The Name of the service to add a policy to.
	Name string `json:"name"`
	// The policy type to add to the autoscale configuration.
	Type AutoscalePolicy `json:"type"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateAutoscaleConfigurationPolicyRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateAutoscaleConfigurationPolicyRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name": {},
		"type": {},
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

func (m CreateAutoscaleConfigurationPolicyRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateAutoscaleConfigurationPolicyRequest
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

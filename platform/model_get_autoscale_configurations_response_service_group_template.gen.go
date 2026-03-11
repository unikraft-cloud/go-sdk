// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The instance template used for the autoscale configuration.
// Only if enabled is true.

type GetAutoscaleConfigurationsResponseServiceGroupTemplate struct {
	// The name of the template used for the autoscale configuration.
	Name *string `json:"name,omitempty"`
	// The UUID of the template used for the autoscale configuration.
	Uuid *string `json:"uuid,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *GetAutoscaleConfigurationsResponseServiceGroupTemplate) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationsResponseServiceGroupTemplate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name": {},
		"uuid": {},
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

func (m GetAutoscaleConfigurationsResponseServiceGroupTemplate) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationsResponseServiceGroupTemplate
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

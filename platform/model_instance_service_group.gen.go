// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The service group configuration for the instance.

type InstanceServiceGroup struct {
	// The UUID of the service group.
	//
	// This is a unique identifier for the service group that is generated when
	// the service is created.  The UUID is used to reference the service group
	// in API calls and can be used to identify the service in all API calls
	// that require an service identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service group.
	//
	// This is a human-readable name that can be used to identify the service
	// group.  The name is unique within the context of your account.  The name
	// can also be used to identify the service group in API calls.
	Name *string `json:"name,omitempty"`
	// The domain configuration for the service group.
	Domains []ServiceGroupInstanceDomain `json:"domains,omitempty"`
	// (Only applies when using global control plane).
	// Where the service group is located.
	Metro *string `json:"metro,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstanceServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias InstanceServiceGroup
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
		"domains": {},
		"metro":   {},
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

func (m InstanceServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias InstanceServiceGroup
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

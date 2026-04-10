// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// Template instances.
// An existing instance can be saved as a template. This template is then
// used to create new instances that inherit the exact configuration and
// state the original instance had when the template was created.

type CreateInstanceRequestTemplate struct {
	// (Optional).  Whether the instance needs to run in order to reach template state
	Prepare *bool `json:"prepare,omitempty"`
	// (Optional).  The UUID of a template instance to create the instance from.
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// (Optional).  The name of a template instance to create the instance from.
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// Where the volume is located.
	Metro      *string                                  `json:"metro,omitempty"`
	CreateArgs *CreateInstanceRequestTemplateCreateArgs `json:"create_args,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateInstanceRequestTemplate) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestTemplate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"prepare":     {},
		"uuid":        {},
		"name":        {},
		"metro":       {},
		"create_args": {},
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

func (m CreateInstanceRequestTemplate) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestTemplate
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

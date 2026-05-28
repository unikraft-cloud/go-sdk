// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type UpdateInstanceByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop MutableInstanceProperty `json:"prop"`
	// The operation to perform on the property.
	Op MutableInstanceOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "image": string
	// - For "args": string or array of strings
	// - For "env": object (for SET/ADD) or string/array of strings (for DEL)
	// - For "memory_mb": integer
	// - For "vcpus": integer
	// - For "scale_to_zero": object with cooldown_time_ms, policy, and stateful fields
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "schedules": array of schedule objects (with name, when, action, and optional args fields).
	//   Use action "exec" together with args to execute a command at the scheduled time.
	// - For "autokill": object with time_ms and num_requests fields
	// - For "hostname": string (valid DNS label)
	// - For "roms": array of ROM objects (with name and image fields) for SET/ADD, or array of ROM names for DEL
	// - For "dependencies": array of instance identifiers (name or UUID)
	// - For "sched_priority": SchedPriority enum value ("normal", "medium", "high", "admin")
	Value *interface{} `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstanceByUUIDRequestBody
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
		"prop":  {},
		"op":    {},
		"value": {},
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

func (m UpdateInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstanceByUUIDRequestBody
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single update operation to be applied to an instance.
// The property to modify.
type UpdateInstancesRequestItemProp string

const (
	UpdateInstancesRequestItemPropImage         UpdateInstancesRequestItemProp = "image"
	UpdateInstancesRequestItemPropArgs          UpdateInstancesRequestItemProp = "args"
	UpdateInstancesRequestItemPropEnv           UpdateInstancesRequestItemProp = "env"
	UpdateInstancesRequestItemPropMemory_mb     UpdateInstancesRequestItemProp = "memory_mb"
	UpdateInstancesRequestItemPropVcpus         UpdateInstancesRequestItemProp = "vcpus"
	UpdateInstancesRequestItemPropScale_to_zero UpdateInstancesRequestItemProp = "scale_to_zero"
	UpdateInstancesRequestItemPropTags          UpdateInstancesRequestItemProp = "tags"
	UpdateInstancesRequestItemPropDelete_lock   UpdateInstancesRequestItemProp = "delete_lock"
	UpdateInstancesRequestItemPropSchedules     UpdateInstancesRequestItemProp = "schedules"
	UpdateInstancesRequestItemPropAutokill      UpdateInstancesRequestItemProp = "autokill"
)

// The operation to perform on the property.
type UpdateInstancesRequestItemOp string

const (
	UpdateInstancesRequestItemOpSet UpdateInstancesRequestItemOp = "set"
	UpdateInstancesRequestItemOpAdd UpdateInstancesRequestItemOp = "add"
	UpdateInstancesRequestItemOpDel UpdateInstancesRequestItemOp = "del"
)

type UpdateInstancesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The UUID of the instance to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// The property to modify.
	Prop UpdateInstancesRequestItemProp `json:"prop"`
	// The operation to perform on the property.
	Op UpdateInstancesRequestItemOp `json:"op"`
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
	Value *interface{} `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
		"uuid":  {},
		"name":  {},
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

func (m UpdateInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstancesRequestItem
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

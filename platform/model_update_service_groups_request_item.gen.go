// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A single update operation to be applied to a service group.
type UpdateServiceGroupsRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`
	// A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitzero"`
	// The property to modify.
	Prop MutableServiceGroupProperty `json:"prop"`
	// The operation to perform.
	Op MutableServiceGroupOperation `json:"op"`
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
	// Use action "exec" together with args to execute a command at the scheduled time.
	// - For "autokill": object with time_ms and num_requests fields
	// - For "hostname": string (valid DNS label)
	// - For "roms": array of ROM objects (with name and image fields) for SET/ADD, or array of ROM names for DEL
	// - For "dependencies": array of instance identifiers (name or UUID)
	// - For "sched_priority": SchedPriority enum value ("normal", "medium", "high", "admin")
	Value *interface{} `json:"value,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateServiceGroupsRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateServiceGroupsRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateServiceGroupsRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateServiceGroupsRequestItem
	return json.Marshal((Alias)(m))
}

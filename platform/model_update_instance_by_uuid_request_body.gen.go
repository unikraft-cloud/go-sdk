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

type UpdateInstanceByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitzero"`
	// The property to modify.
	Prop MutableInstanceProperty `json:"prop"`
	// The operation to perform on the property.
	Op MutableInstanceOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "image": image reference string, or object with image url,
	//   credentials, headers and pull policy
	// - For "args": string or array of strings
	// - For "env": object (for SET/ADD) or string/array of strings (for DEL)
	// - For "memory_mb": integer
	// - For "vcpus": integer
	// - For "scale_to_zero": object with cooldown_time_ms, policy, and stateful fields
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "schedules": array of schedule objects (with name, when, action, and optional args fields) for SET/ADD, or array of schedule names for DEL.
	//   Use action "exec" together with args to execute a command at the scheduled time.
	// - For "autokill": object with time_ms and num_requests fields
	// - For "hostname": string (valid DNS label)
	// - For "roms": array of ROM objects (with name and image fields) for SET/ADD, or array of ROM names for DEL
	// - For "plugins": array of plugin objects (with name, rom, and optional config fields) for SET/ADD
	// - For "dependencies": array of instance identifiers (name or UUID)
	// - For "sched_priority": SchedPriority enum value ("normal", "medium", "high", "admin")
	// - For "annotations": object (for SET/ADD) or string/array of strings (for DEL)
	Value *interface{} `json:"value,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

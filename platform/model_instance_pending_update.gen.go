// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// A queued property change awaiting application (typically on next restart).
type InstancePendingUpdate struct {
	// The property being updated.
	Prop MutableInstanceProperty `json:"prop"`
	// The patch operation type.
	Op MutableInstanceOperation `json:"op"`
	// The new value for the property.  Type depends on the property being
	// updated.
	Value interface{} `json:"value"`
	// The status of this update.
	Status InstancePendingUpdateStatus `json:"status"`
	// Error message.  Only present when status is "failed".
	Error *string `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstancePendingUpdate) UnmarshalJSON(data []byte) error {
	type Alias InstancePendingUpdate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstancePendingUpdate) MarshalJSON() ([]byte, error) {
	type Alias InstancePendingUpdate
	return json.Marshal((Alias)(m))
}

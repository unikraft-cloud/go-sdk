// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A queued property change awaiting application (typically on next restart).
// The status of this update.
type InstancePendingUpdateStatus string

const (
	InstancePendingUpdateStatusPending InstancePendingUpdateStatus = "pending"
	InstancePendingUpdateStatusFailed  InstancePendingUpdateStatus = "failed"
)

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
	Error *string `json:"error,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstancePendingUpdate) UnmarshalJSON(data []byte) error {
	type Alias InstancePendingUpdate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"prop":   {},
		"op":     {},
		"value":  {},
		"status": {},
		"error":  {},
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

func (m InstancePendingUpdate) MarshalJSON() ([]byte, error) {
	type Alias InstancePendingUpdate
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

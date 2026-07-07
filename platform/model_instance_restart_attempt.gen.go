// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"time"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Records the current restart attempt of an instance.
type InstanceRestartAttempt struct {
	// Current restart attempt number. This is incremented each time the instance
	// is restarted automatically by the platform.
	Attempt uint32 `json:"attempt"`
	// Timestamp of the next scheduled restart attempt.
	NextAt *time.Time `json:"next_at,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *InstanceRestartAttempt) UnmarshalJSON(data []byte) error {
	type Alias InstanceRestartAttempt
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceRestartAttempt) MarshalJSON() ([]byte, error) {
	type Alias InstanceRestartAttempt
	return json.Marshal((Alias)(m))
}

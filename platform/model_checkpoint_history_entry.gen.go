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

// A checkpoint history entry, representing a single checkpoint in the
// history of an instance.
type CheckpointHistoryEntry struct {
	// The UUID of the checkpoint.
	Uuid string `json:"uuid"`
	// The name of the checkpoint.
	Name string `json:"name"`
	// The time the checkpoint was created.
	CreatedAt time.Time `json:"created_at"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CheckpointHistoryEntry) UnmarshalJSON(data []byte) error {
	type Alias CheckpointHistoryEntry
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CheckpointHistoryEntry) MarshalJSON() ([]byte, error) {
	type Alias CheckpointHistoryEntry
	return json.Marshal((Alias)(m))
}

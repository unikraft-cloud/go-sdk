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

// Automatic delete-on-idle configuration for the checkpoint instance.
type ItemCheckpointAutokill struct {
	// Time in milliseconds after the checkpoint was last used for restoring
	// before it is deleted. A value of 0 disables checkpoint autokill.
	TimeMs *uint64 `json:"time_ms,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ItemCheckpointAutokill) UnmarshalJSON(data []byte) error {
	type Alias ItemCheckpointAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ItemCheckpointAutokill) MarshalJSON() ([]byte, error) {
	type Alias ItemCheckpointAutokill
	return json.Marshal((Alias)(m))
}

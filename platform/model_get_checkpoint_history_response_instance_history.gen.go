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

// Per-item result for a get checkpoint/instance history operation.
type GetCheckpointHistoryResponseInstanceHistory struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The checkpoint history entries.
	History []CheckpointHistoryEntry `json:"history,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetCheckpointHistoryResponseInstanceHistory) UnmarshalJSON(data []byte) error {
	type Alias GetCheckpointHistoryResponseInstanceHistory
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetCheckpointHistoryResponseInstanceHistory) MarshalJSON() ([]byte, error) {
	type Alias GetCheckpointHistoryResponseInstanceHistory
	return json.Marshal((Alias)(m))
}

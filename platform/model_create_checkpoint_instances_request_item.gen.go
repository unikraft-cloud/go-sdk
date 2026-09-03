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

// A single checkpoint creation request.
type CreateCheckpointInstancesRequestItem struct {
	// The source instance to create a checkpoint from (by name or UUID).
	From NameOrUUID `json:"from"`
	// (Optional).  The name of the checkpoint.
	// If not provided, a name will be generated.
	Name *string `json:"name,omitzero"`
	// Timeout in seconds to wait for the checkpoint to be created.
	// No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitzero"`
	// (Optional).  Tags to associate with the checkpoint.
	Tags []string `json:"tags,omitzero"`
	// (Optional). Automatic delete-on-idle configuration for the new
	// checkpoint.
	Autokill *ItemCheckpointAutokill `json:"autokill,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateCheckpointInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias CreateCheckpointInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateCheckpointInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias CreateCheckpointInstancesRequestItem
	return json.Marshal((Alias)(m))
}

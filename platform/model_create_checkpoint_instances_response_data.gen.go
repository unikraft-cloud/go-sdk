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

type CreateCheckpointInstancesResponseData struct {
	// List of checkpoint instances that were created during the operation.
	Instances []CreateCheckpointInstancesResponseCheckpointInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateCheckpointInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias CreateCheckpointInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateCheckpointInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias CreateCheckpointInstancesResponseData
	return json.Marshal((Alias)(m))
}

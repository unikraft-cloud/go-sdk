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

type DeleteCheckpointInstancesResponseData struct {
	// List of checkpoint instances that were processed during the delete operation.
	Instances []DeleteCheckpointInstancesResponseCheckpointInstance `json:"instances,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteCheckpointInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias DeleteCheckpointInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteCheckpointInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias DeleteCheckpointInstancesResponseData
	return json.Marshal((Alias)(m))
}

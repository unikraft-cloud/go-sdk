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

type GetCheckpointInstancesResponseData struct {
	// List of checkpoint instances that were retrieved during the operation.
	Instances []Instance `json:"instances,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *GetCheckpointInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetCheckpointInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetCheckpointInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetCheckpointInstancesResponseData
	return json.Marshal((Alias)(m))
}

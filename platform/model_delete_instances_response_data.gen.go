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

type DeleteInstancesResponseData struct {
	// The instance(s) which were deleted by the request.
	Instances []DeleteInstancesResponseInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstancesResponseData
	return json.Marshal((Alias)(m))
}

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

type UpdateInstancesResponseData struct {
	// List of instances that were processed during the update operation.
	Instances []UpdateInstancesResponseUpdatedInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstancesResponseData
	return json.Marshal((Alias)(m))
}

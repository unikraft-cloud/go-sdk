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

type DetachVolumeByUUIDRequestBody struct {
	// (Optional).  UUID or name of the instance to detach the volume from.
	// If not specified, the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DetachVolumeByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias DetachVolumeByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DetachVolumeByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias DetachVolumeByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

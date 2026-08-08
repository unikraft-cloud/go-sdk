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

type CloneVolumesResponseVolume struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// The UUID of the newly cloned volume.
	Uuid string `json:"uuid"`
	// The name of the newly cloned volume.
	Name string `json:"name"`
	// The state of the volume.
	State VolumeState `json:"state"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CloneVolumesResponseVolume) UnmarshalJSON(data []byte) error {
	type Alias CloneVolumesResponseVolume
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CloneVolumesResponseVolume) MarshalJSON() ([]byte, error) {
	type Alias CloneVolumesResponseVolume
	return json.Marshal((Alias)(m))
}

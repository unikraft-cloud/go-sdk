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

type CreateVolumeResponseVolume struct {
	// The status of the response.
	Status ResponseStatus `json:"status,omitzero"`
	// UUID of the newly created volume.
	Uuid string `json:"uuid,omitzero"`
	// The name of the newly created volume.
	Name string `json:"name,omitzero"`
	// The state of the volume.
	State VolumeState `json:"state,omitzero"`
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

func (m *CreateVolumeResponseVolume) UnmarshalJSON(data []byte) error {
	type Alias CreateVolumeResponseVolume
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateVolumeResponseVolume) MarshalJSON() ([]byte, error) {
	type Alias CreateVolumeResponseVolume
	return json.Marshal((Alias)(m))
}

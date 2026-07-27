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

// Reference to the instance to attach the volume to.
type VolumeInstanceID struct {
	// The UUID of the instance that the volume is attached to.
	Uuid string `json:"uuid"`
	// The name of the instance that the volume is attached to.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *VolumeInstanceID) UnmarshalJSON(data []byte) error {
	type Alias VolumeInstanceID
	return json.Unmarshal(data, (*Alias)(m))
}

func (m VolumeInstanceID) MarshalJSON() ([]byte, error) {
	type Alias VolumeInstanceID
	return json.Marshal((Alias)(m))
}

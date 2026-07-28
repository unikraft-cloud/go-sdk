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

type VolumeInstanceMount struct {
	// The UUID of the instance that the volume is mounted in.
	Uuid string `json:"uuid,omitzero"`
	// The name of the instance that the volume is mounted in.
	Name string `json:"name,omitzero"`
	// Whether the volume is mounted read-only or read-write.
	Readonly bool `json:"readonly,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *VolumeInstanceMount) UnmarshalJSON(data []byte) error {
	type Alias VolumeInstanceMount
	return json.Unmarshal(data, (*Alias)(m))
}

func (m VolumeInstanceMount) MarshalJSON() ([]byte, error) {
	type Alias VolumeInstanceMount
	return json.Marshal((Alias)(m))
}

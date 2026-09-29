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

// A volume defines a storage which can be attached to the instance. Volumes
// can be used to store persistent data which should remain available even if
// the instance is stopped or restarted.
type InstanceVolume struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The mount point of the volume in the instance. This is the directory in
	// the instance where the volume will be mounted.
	At string `json:"at"`
	// Whether the volume is read-only or not.
	Readonly bool `json:"readonly"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceVolume) UnmarshalJSON(data []byte) error {
	type Alias InstanceVolume
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceVolume) MarshalJSON() ([]byte, error) {
	type Alias InstanceVolume
	return json.Marshal((Alias)(m))
}

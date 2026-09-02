// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type AttachVolumeByUUIDRequestBody struct {
	// UUID or name of the instance to attach the volume to.
	AttachTo NameOrUUID `json:"attach_to"`
	// Path of the mountpoint.
	//
	// The path must be absolute, not contain `.` and `..` components, and not
	// contain colons (`:`). The path must point to an empty directory. If the
	// directory does not exist, it is created.
	At string `json:"at"`
	// Whether the volume should be mounted read-only.
	Readonly *bool `json:"readonly,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AttachVolumeByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias AttachVolumeByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AttachVolumeByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias AttachVolumeByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

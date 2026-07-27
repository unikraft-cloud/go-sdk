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

// A single request item for attaching a volume to an instance.
type AttachVolumesRequestItem struct {
	// UUID or name of the instance to attach the volume to.
	AttachTo NameOrUUID `json:"attach_to"`
	// Path of the mountpoint.
	//
	// The path must be absolute, not contain `.` and `..` components, and not
	// contain colons (`:`). The path must point to an empty directory. If the
	// directory does not exist, it is created.
	At string `json:"at"`
	// Whether the volume should be mounted read-only.
	Readonly *bool `json:"readonly,omitempty"`
	// The UUID of the volume to attach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to attach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AttachVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias AttachVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AttachVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias AttachVolumesRequestItem
	return json.Marshal((Alias)(m))
}

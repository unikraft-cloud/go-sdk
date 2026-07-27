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

// A single request of detaching a volume.
type DetachVolumesRequestItem struct {
	// (Optional).  UUID or name of the instance to detach the volume from.
	// If not specified, the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitempty"`
	// The UUID of the volume to detach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to detach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DetachVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias DetachVolumesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DetachVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias DetachVolumesRequestItem
	return json.Marshal((Alias)(m))
}

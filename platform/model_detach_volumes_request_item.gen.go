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

// A single request item to detach a volume from an instance.
type DetachVolumesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`
	// UUID or name of the instance to detach the volume from. If not specified,
	// the volume is detached from all instances.
	From *NameOrUUID `json:"from,omitzero"`

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

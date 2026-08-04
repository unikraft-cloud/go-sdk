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

// The request item for deleting an instance by its UUID or name.
type DeleteInstanceRequestItem struct {
	// Timeout in seconds to wait for the instance to be deleted.  No wait
	// performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitzero"`
	// Delete immediately without retention.  If the instance is already
	// being retained, this will force its deletion.  Ignored if retention
	// for instances is not configured.
	DontRetain *bool `json:"dont_retain,omitzero"`
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteInstanceRequestItem) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstanceRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteInstanceRequestItem) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstanceRequestItem
	return json.Marshal((Alias)(m))
}

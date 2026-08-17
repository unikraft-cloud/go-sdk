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

// An identifier for a resource.  Either a name or a UUID.
type NameOrUUID struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NameOrUUID) UnmarshalJSON(data []byte) error {
	type Alias NameOrUUID
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NameOrUUID) MarshalJSON() ([]byte, error) {
	type Alias NameOrUUID
	return json.Marshal((Alias)(m))
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// An identifier for a resource.  Either a name or a UUID.
type NameOrUUID struct {
	// (Only applies when using global control plane).
	// The metro of the resource.
	Metro *string `json:"metro,omitempty"`
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *NameOrUUID) UnmarshalJSON(data []byte) error {
	type Alias NameOrUUID
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NameOrUUID) MarshalJSON() ([]byte, error) {
	type Alias NameOrUUID
	return json.Marshal((Alias)(m))
}

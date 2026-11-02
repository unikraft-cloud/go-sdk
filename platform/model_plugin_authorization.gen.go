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

// Authorization for requests to a plugin endpoint.
type PluginAuthorization struct {
	// The authorization method.
	Type *PluginAuthorizationType `json:"type,omitzero"`
	// The token that requests must present as `Authorization: Bearer <token>`.
	// Required for, and only allowed with, the `bearer` type.
	Token *string `json:"token,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *PluginAuthorization) UnmarshalJSON(data []byte) error {
	type Alias PluginAuthorization
	return json.Unmarshal(data, (*Alias)(m))
}

func (m PluginAuthorization) MarshalJSON() ([]byte, error) {
	type Alias PluginAuthorization
	return json.Marshal((Alias)(m))
}

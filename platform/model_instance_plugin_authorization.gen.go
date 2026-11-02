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

// The authorization method of a plugin, as reported in the instance status.
// The token is never returned.
type InstancePluginAuthorization struct {
	// The authorization method.
	Type PluginAuthorizationType `json:"type"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstancePluginAuthorization) UnmarshalJSON(data []byte) error {
	type Alias InstancePluginAuthorization
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstancePluginAuthorization) MarshalJSON() ([]byte, error) {
	type Alias InstancePluginAuthorization
	return json.Marshal((Alias)(m))
}

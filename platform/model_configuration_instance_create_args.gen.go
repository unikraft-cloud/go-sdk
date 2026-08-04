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

type ConfigurationInstanceCreateArgs struct {
	// The ROM to use for the autoscale configuration.
	Roms *InstanceCreateArgsInstanceCreateRequestRoms `json:"roms,omitzero"`
	// The template to use for the autoscale configuration.
	Template *NameOrUUID `json:"template,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ConfigurationInstanceCreateArgs) UnmarshalJSON(data []byte) error {
	type Alias ConfigurationInstanceCreateArgs
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ConfigurationInstanceCreateArgs) MarshalJSON() ([]byte, error) {
	type Alias ConfigurationInstanceCreateArgs
	return json.Marshal((Alias)(m))
}

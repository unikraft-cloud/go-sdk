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

// Defines the source template used to build a new instance.
type CreateInstanceRequestTemplate struct {
	// (Optional).  Whether the instance needs to run in order to reach template state
	Prepare *bool `json:"prepare,omitempty"`
	// (Optional).  The UUID of a template instance to create the instance from.
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// (Optional).  The name of a template instance to create the instance from.
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// Where the volume is located.
	Metro *string `json:"metro,omitempty"`
	// (Optional). Configuration parameters to apply when building the new instance from the source template.
	CreateArgs *Instance `json:"create_args,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CreateInstanceRequestTemplate) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestTemplate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestTemplate) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestTemplate
	return json.Marshal((Alias)(m))
}

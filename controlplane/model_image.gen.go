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

type Image struct {
	// The canonical name of the image.
	Name string `json:"name"`
	// The list of tags for the image.  Only populated when details are requested.
	Tags []ImageTag `json:"tags,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *Image) UnmarshalJSON(data []byte) error {
	type Alias Image
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Image) MarshalJSON() ([]byte, error) {
	type Alias Image
	return json.Marshal((Alias)(m))
}

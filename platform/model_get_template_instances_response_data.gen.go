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

type GetTemplateInstancesResponseData struct {
	// List of template instances that were retrieved during the operation.
	Instances []Instance `json:"instances,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *GetTemplateInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetTemplateInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetTemplateInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetTemplateInstancesResponseData
	return json.Marshal((Alias)(m))
}

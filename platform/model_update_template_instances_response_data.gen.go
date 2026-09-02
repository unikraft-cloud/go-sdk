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

type UpdateTemplateInstancesResponseData struct {
	// List of template instances that were processed during the update operation.
	Instances []UpdateTemplateInstancesResponseTemplateInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateTemplateInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias UpdateTemplateInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateTemplateInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias UpdateTemplateInstancesResponseData
	return json.Marshal((Alias)(m))
}

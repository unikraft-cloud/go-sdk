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

type CreateTemplateInstancesResponseData struct {
	// List of template instances that were created during the operation.
	Instances []CreateTemplateInstancesResponseTemplateInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateTemplateInstancesResponseData) UnmarshalJSON(data []byte) error {
	type Alias CreateTemplateInstancesResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateTemplateInstancesResponseData) MarshalJSON() ([]byte, error) {
	type Alias CreateTemplateInstancesResponseData
	return json.Marshal((Alias)(m))
}

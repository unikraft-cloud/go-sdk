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

type ServiceGroupTemplate struct {
	// The name of the template used for the autoscale configuration.
	Name string `json:"name"`
	// The UUID of the template used for the autoscale configuration.
	Uuid string `json:"uuid"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ServiceGroupTemplate) UnmarshalJSON(data []byte) error {
	type Alias ServiceGroupTemplate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ServiceGroupTemplate) MarshalJSON() ([]byte, error) {
	type Alias ServiceGroupTemplate
	return json.Marshal((Alias)(m))
}

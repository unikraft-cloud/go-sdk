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

// Additional data returned by the health check.

type HealthzResponseData struct {
	Checks   map[string]string `json:"checks,omitempty"`
	Versions map[string]string `json:"versions,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *HealthzResponseData) UnmarshalJSON(data []byte) error {
	type Alias HealthzResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m HealthzResponseData) MarshalJSON() ([]byte, error) {
	type Alias HealthzResponseData
	return json.Marshal((Alias)(m))
}

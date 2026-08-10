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
	// The health state of each registered checker, keyed by checker name.
	// Valid keys are "images", "systemd", and "user-defined"; a checker's
	// key is only present if it is enabled. Checkers report only their
	// aggregate state; per-check detail (e.g. which default image is
	// missing, or which user-defined script failed) is not exposed here.
	Checks   map[string]HealthState `json:"checks,omitzero"`
	Versions map[string]string      `json:"versions,omitzero"`
	License  *DataLicense           `json:"license,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *HealthzResponseData) UnmarshalJSON(data []byte) error {
	type Alias HealthzResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m HealthzResponseData) MarshalJSON() ([]byte, error) {
	type Alias HealthzResponseData
	return json.Marshal((Alias)(m))
}

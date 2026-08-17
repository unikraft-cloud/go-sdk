// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type ListRegionsResponseData struct {
	// The list of available regions.
	Regions []Region `json:"regions,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ListRegionsResponseData) UnmarshalJSON(data []byte) error {
	type Alias ListRegionsResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ListRegionsResponseData) MarshalJSON() ([]byte, error) {
	type Alias ListRegionsResponseData
	return json.Marshal((Alias)(m))
}

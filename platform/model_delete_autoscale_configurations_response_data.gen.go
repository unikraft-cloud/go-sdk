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

type DeleteAutoscaleConfigurationsResponseData struct {
	// The configuration(s) which were deleted by the request.
	ServiceGroups []DeleteAutoscaleConfigurationsResponseServiceGroup `json:"service_groups,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteAutoscaleConfigurationsResponseData) UnmarshalJSON(data []byte) error {
	type Alias DeleteAutoscaleConfigurationsResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteAutoscaleConfigurationsResponseData) MarshalJSON() ([]byte, error) {
	type Alias DeleteAutoscaleConfigurationsResponseData
	return json.Marshal((Alias)(m))
}

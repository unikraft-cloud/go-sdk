// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

type GetAutoscaleConfigurationsResponseData struct {
	// The configuration(s) which were retrieved by the request.
	ServiceGroups []GetAutoscaleConfigurationsResponseServiceGroup `json:"service_groups,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetAutoscaleConfigurationsResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationsResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetAutoscaleConfigurationsResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationsResponseData
	return json.Marshal((Alias)(m))
}

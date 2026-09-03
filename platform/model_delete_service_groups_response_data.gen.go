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

type DeleteServiceGroupsResponseData struct {
	// The service group(s) which were deleted by the request.
	ServiceGroups []DeleteServiceGroupsResponseDeletedServiceGroup `json:"service_groups,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteServiceGroupsResponseData) UnmarshalJSON(data []byte) error {
	type Alias DeleteServiceGroupsResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteServiceGroupsResponseData) MarshalJSON() ([]byte, error) {
	type Alias DeleteServiceGroupsResponseData
	return json.Marshal((Alias)(m))
}

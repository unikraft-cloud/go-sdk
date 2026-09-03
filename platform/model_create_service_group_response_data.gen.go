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

type CreateServiceGroupResponseData struct {
	// The service group which was created by this request.
	//
	// Note: only one service group can be specified in the request, so this
	// will always contain a single entry.
	ServiceGroups []ServiceGroup `json:"service_groups,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateServiceGroupResponseData) UnmarshalJSON(data []byte) error {
	type Alias CreateServiceGroupResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateServiceGroupResponseData) MarshalJSON() ([]byte, error) {
	type Alias CreateServiceGroupResponseData
	return json.Marshal((Alias)(m))
}

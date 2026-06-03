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

// Parameters for deleting the instance by UUID.

type DeleteInstanceByUUIDRequestBody struct {
	// Timeout in seconds to wait for the instance to be deleted.  No wait
	// performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *DeleteInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

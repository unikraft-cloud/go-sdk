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

// Per-item result for a delete service groups operation.
type DeleteServiceGroupsResponseDeletedServiceGroup struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the status.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteServiceGroupsResponseDeletedServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias DeleteServiceGroupsResponseDeletedServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteServiceGroupsResponseDeletedServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias DeleteServiceGroupsResponseDeletedServiceGroup
	return json.Marshal((Alias)(m))
}

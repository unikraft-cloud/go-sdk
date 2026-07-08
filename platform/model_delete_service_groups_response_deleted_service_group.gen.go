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

// Details of the service group which was deleted by this request.
type DeleteServiceGroupsResponseDeletedServiceGroup struct {
	// Indicates whether the delete operation was successful or not for this
	// service group.
	Status ResponseStatus `json:"status"`
	// The UUID of the service group which was deleted.
	Uuid string `json:"uuid"`
	// The name of the service group which was deleted.
	Name string `json:"name"`
	// An optional message providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *DeleteServiceGroupsResponseDeletedServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias DeleteServiceGroupsResponseDeletedServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteServiceGroupsResponseDeletedServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias DeleteServiceGroupsResponseDeletedServiceGroup
	return json.Marshal((Alias)(m))
}

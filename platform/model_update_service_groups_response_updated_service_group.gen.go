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

type UpdateServiceGroupsResponseUpdatedServiceGroup struct {
	// The UUID of the service group that was updated.
	Uuid string `json:"uuid"`
	// The name of the service group that was updated.
	Name string `json:"name"`
	// The status of this particular service group update operation.
	Status ResponseStatus `json:"status"`
	// (Optional).  The client-provided ID from the request.
	Id *string `json:"id,omitempty"`
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

func (m *UpdateServiceGroupsResponseUpdatedServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias UpdateServiceGroupsResponseUpdatedServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateServiceGroupsResponseUpdatedServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias UpdateServiceGroupsResponseUpdatedServiceGroup
	return json.Marshal((Alias)(m))
}

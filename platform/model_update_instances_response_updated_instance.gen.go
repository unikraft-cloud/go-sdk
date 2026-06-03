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

type UpdateInstancesResponseUpdatedInstance struct {
	// The UUID of the instance that was updated.
	Uuid string `json:"uuid"`
	// The name of the instance that was updated.
	Name string `json:"name"`
	// The status of this particular instance update operation.
	Status ResponseStatus `json:"status"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// (Optional).  The client-provided ID from the request.
	Id *string `json:"id,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UpdateInstancesResponseUpdatedInstance) UnmarshalJSON(data []byte) error {
	type Alias UpdateInstancesResponseUpdatedInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateInstancesResponseUpdatedInstance) MarshalJSON() ([]byte, error) {
	type Alias UpdateInstancesResponseUpdatedInstance
	return json.Marshal((Alias)(m))
}

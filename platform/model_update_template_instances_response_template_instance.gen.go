// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type UpdateTemplateInstancesResponseTemplateInstance struct {
	// The UUID of the template instance that was updated.
	Uuid string `json:"uuid"`
	// The name of the template instance that was updated.
	Name string `json:"name"`
	// The status of this particular template instance update operation.
	Status ResponseStatus `json:"status"`
	// (Optional).  The client-provided ID from the request.
	Id *string `json:"id,omitzero"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateTemplateInstancesResponseTemplateInstance) UnmarshalJSON(data []byte) error {
	type Alias UpdateTemplateInstancesResponseTemplateInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateTemplateInstancesResponseTemplateInstance) MarshalJSON() ([]byte, error) {
	type Alias UpdateTemplateInstancesResponseTemplateInstance
	return json.Marshal((Alias)(m))
}

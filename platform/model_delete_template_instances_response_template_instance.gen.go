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

type DeleteTemplateInstancesResponseTemplateInstance struct {
	// The UUID of the template instance that was deleted.
	Uuid string `json:"uuid"`
	// The name of the template instance that was deleted.
	Name string `json:"name"`
	// The status of this particular template instance deletion operation.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteTemplateInstancesResponseTemplateInstance) UnmarshalJSON(data []byte) error {
	type Alias DeleteTemplateInstancesResponseTemplateInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteTemplateInstancesResponseTemplateInstance) MarshalJSON() ([]byte, error) {
	type Alias DeleteTemplateInstancesResponseTemplateInstance
	return json.Marshal((Alias)(m))
}

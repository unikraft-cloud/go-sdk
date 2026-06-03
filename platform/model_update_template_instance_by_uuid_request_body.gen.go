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

type UpdateTemplateInstanceByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop MutableTemplateInstanceProperty `json:"prop"`
	// The operation to perform on the property.
	Op MutableTemplateInstanceOperation `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *UpdateTemplateInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateTemplateInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateTemplateInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateTemplateInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

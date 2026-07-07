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

// A single template instance to be created.
type CreateTemplateInstancesRequestItem struct {
	// Timeout in seconds to wait for the template instances to be created.
	// A value of -1 means to wait indefinitely until the instance reaches the
	// desired state. No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// (Optional). Automatic delete-on-idle configuration for the new template.
	Autokill *ItemAutokill `json:"autokill,omitempty"`
	// The UUID of the instance to convert into template. Mutually exclusive
	// with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to convert into template. Mutually exclusive
	// with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CreateTemplateInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias CreateTemplateInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateTemplateInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias CreateTemplateInstancesRequestItem
	return json.Marshal((Alias)(m))
}

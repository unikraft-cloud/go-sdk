// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single update operation to be applied to a template instance.
// The property to modify.
type UpdateTemplateInstancesRequestItemProp string

const (
	UpdateTemplateInstancesRequestItemPropTags        UpdateTemplateInstancesRequestItemProp = "tags"
	UpdateTemplateInstancesRequestItemPropDelete_lock UpdateTemplateInstancesRequestItemProp = "delete_lock"
	UpdateTemplateInstancesRequestItemPropAutokill    UpdateTemplateInstancesRequestItemProp = "autokill"
)

// The operation to perform on the property.
type UpdateTemplateInstancesRequestItemOp string

const (
	UpdateTemplateInstancesRequestItemOpSet UpdateTemplateInstancesRequestItemOp = "set"
	UpdateTemplateInstancesRequestItemOpAdd UpdateTemplateInstancesRequestItemOp = "add"
	UpdateTemplateInstancesRequestItemOpDel UpdateTemplateInstancesRequestItemOp = "del"
)

type UpdateTemplateInstancesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in
	// the response.
	Id *string `json:"id,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the template instance.
	Metro *string `json:"metro,omitempty"`
	// The property to modify.
	Prop UpdateTemplateInstancesRequestItemProp `json:"prop"`
	// The operation to perform on the property.
	Op UpdateTemplateInstancesRequestItemOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitempty"`
	// The UUID of the template instance to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the template instance to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateTemplateInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateTemplateInstancesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
		"metro": {},
		"prop":  {},
		"op":    {},
		"value": {},
		"uuid":  {},
		"name":  {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m UpdateTemplateInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateTemplateInstancesRequestItem
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

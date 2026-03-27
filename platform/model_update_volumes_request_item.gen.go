// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single request item for updating a volume.
// The property to modify.
type UpdateVolumesRequestItemProp string

const (
	UpdateVolumesRequestItemPropSize_mb      UpdateVolumesRequestItemProp = "size_mb"
	UpdateVolumesRequestItemPropTags         UpdateVolumesRequestItemProp = "tags"
	UpdateVolumesRequestItemPropQuota_policy UpdateVolumesRequestItemProp = "quota_policy"
	UpdateVolumesRequestItemPropDelete_lock  UpdateVolumesRequestItemProp = "delete_lock"
)

// The operation to perform.
type UpdateVolumesRequestItemOp string

const (
	UpdateVolumesRequestItemOpSet UpdateVolumesRequestItemOp = "set"
	UpdateVolumesRequestItemOpAdd UpdateVolumesRequestItemOp = "add"
	UpdateVolumesRequestItemOpDel UpdateVolumesRequestItemOp = "del"
)

type UpdateVolumesRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The UUID of the volume to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// The property to modify.
	Prop UpdateVolumesRequestItemProp `json:"prop"`
	// The operation to perform.
	Op UpdateVolumesRequestItemOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "size_mb": unsigned integer
	// - For "quota_policy": 1 - static reservation, 2 - dynamic reservation
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *interface{} `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateVolumesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateVolumesRequestItem
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
		"uuid":  {},
		"name":  {},
		"prop":  {},
		"op":    {},
		"value": {},
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

func (m UpdateVolumesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateVolumesRequestItem
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

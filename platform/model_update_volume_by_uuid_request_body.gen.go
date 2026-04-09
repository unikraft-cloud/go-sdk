// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The property to modify.
type UpdateVolumeByUUIDRequestBodyProp string

const (
	UpdateVolumeByUUIDRequestBodyPropSize_mb      UpdateVolumeByUUIDRequestBodyProp = "size_mb"
	UpdateVolumeByUUIDRequestBodyPropTags         UpdateVolumeByUUIDRequestBodyProp = "tags"
	UpdateVolumeByUUIDRequestBodyPropQuota_policy UpdateVolumeByUUIDRequestBodyProp = "quota_policy"
	UpdateVolumeByUUIDRequestBodyPropDelete_lock  UpdateVolumeByUUIDRequestBodyProp = "delete_lock"
)

// The operation to perform.
type UpdateVolumeByUUIDRequestBodyOp string

const (
	UpdateVolumeByUUIDRequestBodyOpSet UpdateVolumeByUUIDRequestBodyOp = "set"
	UpdateVolumeByUUIDRequestBodyOpAdd UpdateVolumeByUUIDRequestBodyOp = "add"
	UpdateVolumeByUUIDRequestBodyOpDel UpdateVolumeByUUIDRequestBodyOp = "del"
)

type UpdateVolumeByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateVolumeByUUIDRequestBodyProp `json:"prop"`
	// The operation to perform.
	Op UpdateVolumeByUUIDRequestBodyOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "size_mb": unsigned integer
	// - For "quota_policy": "static" or "dynamic"
	// - For "tags": array of Strings
	// - For "delete_lock": boolean
	Value *interface{} `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateVolumeByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateVolumeByUUIDRequestBody
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"id":    {},
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

func (m UpdateVolumeByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateVolumeByUUIDRequestBody
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The property to modify.
type UpdateServiceGroupByUUIDRequestBodyProp string

const (
	UpdateServiceGroupByUUIDRequestBodyPropServices   UpdateServiceGroupByUUIDRequestBodyProp = "services"
	UpdateServiceGroupByUUIDRequestBodyPropDomains    UpdateServiceGroupByUUIDRequestBodyProp = "domains"
	UpdateServiceGroupByUUIDRequestBodyPropSoft_limit UpdateServiceGroupByUUIDRequestBodyProp = "soft_limit"
	UpdateServiceGroupByUUIDRequestBodyPropHard_limit UpdateServiceGroupByUUIDRequestBodyProp = "hard_limit"
	UpdateServiceGroupByUUIDRequestBodyPropAutokill   UpdateServiceGroupByUUIDRequestBodyProp = "autokill"
)

// The operation to perform.
type UpdateServiceGroupByUUIDRequestBodyOp string

const (
	UpdateServiceGroupByUUIDRequestBodyOpSet UpdateServiceGroupByUUIDRequestBodyOp = "set"
	UpdateServiceGroupByUUIDRequestBodyOpAdd UpdateServiceGroupByUUIDRequestBodyOp = "add"
	UpdateServiceGroupByUUIDRequestBodyOpDel UpdateServiceGroupByUUIDRequestBodyOp = "del"
)

type UpdateServiceGroupByUUIDRequestBody struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateServiceGroupByUUIDRequestBodyProp `json:"prop"`
	// The operation to perform.
	Op UpdateServiceGroupByUUIDRequestBodyOp `json:"op"`
	// The value for the update operation:
	// - For "services": array of Service objects (same as for creation)
	// - For "domains": array of Domain objects (same as for creation)
	// - For "soft_limit": integer (1–65535), must be <= "hard_limit"
	// - For "hard_limit": integer (1–65535), must be >= "soft_limit"
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateServiceGroupByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias UpdateServiceGroupByUUIDRequestBody
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

func (m UpdateServiceGroupByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias UpdateServiceGroupByUUIDRequestBody
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A single update operation to be applied to a service group
// The property to modify.
type UpdateServiceGroupsRequestItemProp string

const (
	UpdateServiceGroupsRequestItemPropServices   UpdateServiceGroupsRequestItemProp = "services"
	UpdateServiceGroupsRequestItemPropDomains    UpdateServiceGroupsRequestItemProp = "domains"
	UpdateServiceGroupsRequestItemPropSoft_limit UpdateServiceGroupsRequestItemProp = "soft_limit"
	UpdateServiceGroupsRequestItemPropHard_limit UpdateServiceGroupsRequestItemProp = "hard_limit"
	UpdateServiceGroupsRequestItemPropAutokill   UpdateServiceGroupsRequestItemProp = "autokill"
)

// The operation to perform.
type UpdateServiceGroupsRequestItemOp string

const (
	UpdateServiceGroupsRequestItemOpSet UpdateServiceGroupsRequestItemOp = "set"
	UpdateServiceGroupsRequestItemOpAdd UpdateServiceGroupsRequestItemOp = "add"
	UpdateServiceGroupsRequestItemOpDel UpdateServiceGroupsRequestItemOp = "del"
)

type UpdateServiceGroupsRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The property to modify.
	Prop UpdateServiceGroupsRequestItemProp `json:"prop"`
	// The operation to perform.
	Op UpdateServiceGroupsRequestItemOp `json:"op"`
	// The value for the update operation:
	// - For "services": array of Service objects (same as for creation)
	// - For "domains": array of Domain objects (same as for creation)
	// - For "soft_limit": integer (1–65535), must be <= "hard_limit"
	// - For "hard_limit": integer (1–65535), must be >= "soft_limit"
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitempty"`
	// The UUID of the service group to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service group to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *UpdateServiceGroupsRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateServiceGroupsRequestItem
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

func (m UpdateServiceGroupsRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateServiceGroupsRequestItem
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

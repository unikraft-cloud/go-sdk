// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A single update operation to be applied to a service group
type UpdateServiceGroupsRequestItem struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitzero"`
	// The property to modify.
	Prop MutableServiceGroupProperty `json:"prop"`
	// The operation to perform.
	Op MutableServiceGroupOperation `json:"op"`
	// The value for the update operation:
	// - For "services": array of Service objects (same as for creation)
	// - For "domains": array of Domain objects (same as for creation)
	// - For "soft_limit": integer (1–65535), must be <= "hard_limit"
	// - For "hard_limit": integer (1–65535), must be >= "soft_limit"
	// - For "autokill": object with time_ms field
	Value *interface{} `json:"value,omitzero"`
	// The UUID of the service group to update.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the service group to update.  Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UpdateServiceGroupsRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UpdateServiceGroupsRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UpdateServiceGroupsRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UpdateServiceGroupsRequestItem
	return json.Marshal((Alias)(m))
}

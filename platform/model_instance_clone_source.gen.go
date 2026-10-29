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

// The instance, template, or checkpoint that an instance was cloned from.
type InstanceCloneSource struct {
	// The UUID of the source.
	Uuid string `json:"uuid"`
	// The name of the source.
	Name *string `json:"name,omitzero"`
	// The name of the node the source is placed on.
	Node *string `json:"node,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceCloneSource) UnmarshalJSON(data []byte) error {
	type Alias InstanceCloneSource
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceCloneSource) MarshalJSON() ([]byte, error) {
	type Alias InstanceCloneSource
	return json.Marshal((Alias)(m))
}

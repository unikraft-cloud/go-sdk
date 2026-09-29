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

// The object an audit event is about.
type AuditObject struct {
	// The kind of object.
	Type AuditObjectType `json:"type"`
	// The object's UUID.
	Uuid string `json:"uuid"`
	// The UUID of the user the object belongs to.
	Owner *string `json:"owner,omitzero"`
	// The tags set on the object.
	Tags []string `json:"tags,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AuditObject) UnmarshalJSON(data []byte) error {
	type Alias AuditObject
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AuditObject) MarshalJSON() ([]byte, error) {
	type Alias AuditObject
	return json.Marshal((Alias)(m))
}

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

// The migration of an instance.
type InstanceMigration struct {
	// Whether the instance is the copy of an incoming migration. Only present
	// on the destination node.
	Incoming *bool `json:"incoming,omitzero"`
	// The migration URL of the destination. Only present on the source node.
	Destination *string `json:"destination,omitzero"`
	// The phase of the migration.
	State InstanceMigrationState `json:"state"`
	// The size of the files the migration copies in bytes, including the holes
	// of sparse files. Only present once the migration knows the files.
	TotalBytes *uint64 `json:"total_bytes,omitzero"`
	// The bytes of these files that have been copied. Only present together
	// with total_bytes.
	CopiedBytes *uint64 `json:"copied_bytes,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceMigration) UnmarshalJSON(data []byte) error {
	type Alias InstanceMigration
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceMigration) MarshalJSON() ([]byte, error) {
	type Alias InstanceMigration
	return json.Marshal((Alias)(m))
}

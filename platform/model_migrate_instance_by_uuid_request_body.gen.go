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

// Parameters for migrating the instance to another node.
type MigrateInstanceByUUIDRequestBody struct {
	// The destination, as returned by its accept request.
	Destination MigrateInstanceDestination `json:"destination"`
	// Timeout in seconds to wait for the migration to end. No wait if 0
	// (default). Use -1 for the largest possible value.
	//
	// A migration that waits in the queue or for the instance to start or
	// stop has not ended. If the timeout expires first, the request fails
	// with a timeout and the instance reports the state of its migration.
	// The migration goes on.
	TimeoutS *int64 `json:"timeout_s,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *MigrateInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias MigrateInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m MigrateInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias MigrateInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

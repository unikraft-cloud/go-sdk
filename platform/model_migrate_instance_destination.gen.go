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

// The destination of a migration, as issued by its accept request.
type MigrateInstanceDestination struct {
	// The migration URL that the destination returned.
	Url string `json:"url"`
	// The token that the destination returned.
	Token string `json:"token"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *MigrateInstanceDestination) UnmarshalJSON(data []byte) error {
	type Alias MigrateInstanceDestination
	return json.Unmarshal(data, (*Alias)(m))
}

func (m MigrateInstanceDestination) MarshalJSON() ([]byte, error) {
	type Alias MigrateInstanceDestination
	return json.Marshal((Alias)(m))
}

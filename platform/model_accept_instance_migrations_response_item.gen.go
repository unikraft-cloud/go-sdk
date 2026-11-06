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

// Per-item result for an accept migrations operation.
type AcceptInstanceMigrationsResponseItem struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the instance the token is valid for, if one was given.
	Uuid *string `json:"uuid,omitzero"`
	// The migration URL of this node, to pass to the migrate request.
	Url string `json:"url"`
	// The token for the migration, to pass to the migrate request.
	Token string `json:"token"`
	// The lifetime of the token until its first use in seconds.
	TtlS int32 `json:"ttl_s"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AcceptInstanceMigrationsResponseItem) UnmarshalJSON(data []byte) error {
	type Alias AcceptInstanceMigrationsResponseItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AcceptInstanceMigrationsResponseItem) MarshalJSON() ([]byte, error) {
	type Alias AcceptInstanceMigrationsResponseItem
	return json.Marshal((Alias)(m))
}

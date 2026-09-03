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

// A single item in the request.
type GetInstancesLogsRequestItem struct {
	// The byte offset of the log output to receive.  A negative sign makes the
	// offset relative to the end of the log.
	Offset *int64 `json:"offset,omitzero"`
	// The amount of bytes to return at most.
	Limit *int64 `json:"limit,omitzero"`
	// The UUID of the instance to retrieve logs for.  Mutually exclusive with
	// name.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the instance to retrieve logs for.  Mutually exclusive with
	// UUID.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetInstancesLogsRequestItem) UnmarshalJSON(data []byte) error {
	type Alias GetInstancesLogsRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetInstancesLogsRequestItem) MarshalJSON() ([]byte, error) {
	type Alias GetInstancesLogsRequestItem
	return json.Marshal((Alias)(m))
}

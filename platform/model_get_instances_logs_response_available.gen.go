// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type GetInstancesLogsResponseAvailable struct {
	// The first byte offset that can be retrieved.
	Start int64 `json:"start"`
	// The last byte offset that can be retrieved.
	End int64 `json:"end"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetInstancesLogsResponseAvailable) UnmarshalJSON(data []byte) error {
	type Alias GetInstancesLogsResponseAvailable
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetInstancesLogsResponseAvailable) MarshalJSON() ([]byte, error) {
	type Alias GetInstancesLogsResponseAvailable
	return json.Marshal((Alias)(m))
}

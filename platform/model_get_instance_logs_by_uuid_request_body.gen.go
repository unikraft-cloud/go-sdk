// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type GetInstanceLogsByUUIDRequestBody struct {
	// The byte offset of the log output to receive.  A negative sign makes the
	// offset relative to the end of the log.
	Offset *int64 `json:"offset,omitempty"`
	// The amount of bytes to return at most.
	Limit *int64 `json:"limit,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *GetInstanceLogsByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias GetInstanceLogsByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetInstanceLogsByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias GetInstanceLogsByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

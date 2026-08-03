// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type GetInstanceCommandLogsRequestBody struct {
	// Range request for stdout.
	Stdout *BodyStreamRange `json:"stdout,omitzero"`
	// Range request for stderr.
	Stderr *BodyStreamRange `json:"stderr,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetInstanceCommandLogsRequestBody) UnmarshalJSON(data []byte) error {
	type Alias GetInstanceCommandLogsRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetInstanceCommandLogsRequestBody) MarshalJSON() ([]byte, error) {
	type Alias GetInstanceCommandLogsRequestBody
	return json.Marshal((Alias)(m))
}

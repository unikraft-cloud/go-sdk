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

type ExecInstanceCommandRequestBody struct {
	Cmd string `json:"cmd"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ExecInstanceCommandRequestBody) UnmarshalJSON(data []byte) error {
	type Alias ExecInstanceCommandRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ExecInstanceCommandRequestBody) MarshalJSON() ([]byte, error) {
	type Alias ExecInstanceCommandRequestBody
	return json.Marshal((Alias)(m))
}

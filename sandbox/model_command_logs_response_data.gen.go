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

type CommandLogsResponseData struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	// Total available length of stdout, in bytes.
	StdoutAvailable *uint64 `json:"stdout_available,omitzero"`
	// Total available length of stderr, in bytes.
	StderrAvailable *uint64 `json:"stderr_available,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CommandLogsResponseData) UnmarshalJSON(data []byte) error {
	type Alias CommandLogsResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CommandLogsResponseData) MarshalJSON() ([]byte, error) {
	type Alias CommandLogsResponseData
	return json.Marshal((Alias)(m))
}

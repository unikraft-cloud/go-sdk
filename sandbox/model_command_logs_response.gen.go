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

// The response message containing the standard output and error logs.
type CommandLogsResponse struct {
	Status   ResponseStatus           `json:"status"`
	Message  *string                  `json:"message,omitzero"`
	Data     *CommandLogsResponseData `json:"data,omitzero"`
	Errors   []ResponseError          `json:"errors,omitzero"`
	OpTimeUs uint64                   `json:"op_time_us"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CommandLogsResponse) UnmarshalJSON(data []byte) error {
	type Alias CommandLogsResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CommandLogsResponse) MarshalJSON() ([]byte, error) {
	type Alias CommandLogsResponse
	return json.Marshal((Alias)(m))
}

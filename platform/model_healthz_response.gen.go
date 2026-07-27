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

// The response message for a health check of the platform.
type HealthzResponse struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the response.
	Message *string `json:"message,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The response data for this request.
	Data *HealthzResponseData `json:"data,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs uint64 `json:"op_time_us"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *HealthzResponse) UnmarshalJSON(data []byte) error {
	type Alias HealthzResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m HealthzResponse) MarshalJSON() ([]byte, error) {
	type Alias HealthzResponse
	return json.Marshal((Alias)(m))
}

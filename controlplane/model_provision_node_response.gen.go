// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Response message for creating a node.
type ProvisionNodeResponse struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the response.
	Message *string `json:"message,omitempty"`
	// The response data for this request.
	Data *ProvisionNodeResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.
	OpTimeUs *uint64 `json:"op_time_us,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *ProvisionNodeResponse) UnmarshalJSON(data []byte) error {
	type Alias ProvisionNodeResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ProvisionNodeResponse) MarshalJSON() ([]byte, error) {
	type Alias ProvisionNodeResponse
	return json.Marshal((Alias)(m))
}

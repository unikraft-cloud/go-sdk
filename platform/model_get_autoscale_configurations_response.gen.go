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

// The response message for a GetAutoscaleConfigurationsRequest.
type GetAutoscaleConfigurationsResponse struct {
	// The status of the response.
	Status GetAutoscaleConfigurationsResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// The response data for this request.
	Data GetAutoscaleConfigurationsResponseData `json:"data,omitzero"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitzero"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs uint64 `json:"op_time_us,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetAutoscaleConfigurationsResponse) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationsResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetAutoscaleConfigurationsResponse) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationsResponse
	return json.Marshal((Alias)(m))
}

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

// The response message for creating a volume.
type CreateVolumeResponse struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the status.
	Message *string `json:"message,omitzero"`
	// The response data for this request.
	Data *CreateVolumeResponseData `json:"data,omitzero"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitzero"`
	// The operation time in microseconds.
	OpTimeUs uint64 `json:"op_time_us"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateVolumeResponse) UnmarshalJSON(data []byte) error {
	type Alias CreateVolumeResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateVolumeResponse) MarshalJSON() ([]byte, error) {
	type Alias CreateVolumeResponse
	return json.Marshal((Alias)(m))
}

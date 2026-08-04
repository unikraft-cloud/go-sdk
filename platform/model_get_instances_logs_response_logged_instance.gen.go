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

type GetInstancesLogsResponseLoggedInstance struct {
	// The UUID of the instance.
	Uuid string `json:"uuid,omitzero"`
	// The name of the instance.
	Name string `json:"name,omitzero"`
	// Base64 encoded log output of the instance.
	Output string `json:"output,omitzero"`
	// Description of the log availability.
	Available GetInstancesLogsResponseAvailable `json:"available,omitzero"`
	// Description of the range that was returned.  Useful for requests with
	// offset relative to end.
	Range GetInstancesLogsResponseRange `json:"range,omitzero"`
	// State of the instance when the logs were retrieved.
	State InstanceState `json:"state,omitzero"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetInstancesLogsResponseLoggedInstance) UnmarshalJSON(data []byte) error {
	type Alias GetInstancesLogsResponseLoggedInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetInstancesLogsResponseLoggedInstance) MarshalJSON() ([]byte, error) {
	type Alias GetInstancesLogsResponseLoggedInstance
	return json.Marshal((Alias)(m))
}

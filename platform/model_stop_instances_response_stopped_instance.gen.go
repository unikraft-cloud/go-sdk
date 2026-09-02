// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type StopInstancesResponseStoppedInstance struct {
	// The UUID of the instance.
	Uuid string `json:"uuid"`
	// The name of the instance.
	Name string `json:"name"`
	// The current state of the instance.
	State InstanceState `json:"state"`
	// The previous state of the instance before the stop operation was invoked.
	PreviousState InstanceState `json:"previous_state"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *StopInstancesResponseStoppedInstance) UnmarshalJSON(data []byte) error {
	type Alias StopInstancesResponseStoppedInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m StopInstancesResponseStoppedInstance) MarshalJSON() ([]byte, error) {
	type Alias StopInstancesResponseStoppedInstance
	return json.Marshal((Alias)(m))
}

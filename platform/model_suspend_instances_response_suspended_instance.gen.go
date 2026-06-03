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

type SuspendInstancesResponseSuspendedInstance struct {
	// The UUID of the instance.
	Uuid string `json:"uuid"`
	// The name of the instance.
	Name string `json:"name"`
	// (Only applies when using global control plane).
	// The metro of the instance.
	Metro *string `json:"metro,omitempty"`
	// The current state of the instance.
	State InstanceState `json:"state"`
	// The previous state of the instance before the suspend operation was invoked.
	PreviousState InstanceState `json:"previous_state"`
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *SuspendInstancesResponseSuspendedInstance) UnmarshalJSON(data []byte) error {
	type Alias SuspendInstancesResponseSuspendedInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m SuspendInstancesResponseSuspendedInstance) MarshalJSON() ([]byte, error) {
	type Alias SuspendInstancesResponseSuspendedInstance
	return json.Marshal((Alias)(m))
}

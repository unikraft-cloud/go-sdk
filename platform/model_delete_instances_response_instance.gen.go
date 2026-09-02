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

// Details of the instance which was deleted by this request.
type DeleteInstancesResponseInstance struct {
	// Indicates whether the start operation was successful or not for this
	// instance.
	Status ResponseStatus `json:"status"`
	// The UUID of the instance which was deleted.
	Uuid string `json:"uuid"`
	// The name of the instance which was deleted.
	Name string `json:"name"`
	// The previous state of the instance before it was deleted.
	PreviousState string `json:"previous_state"`
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

func (m *DeleteInstancesResponseInstance) UnmarshalJSON(data []byte) error {
	type Alias DeleteInstancesResponseInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteInstancesResponseInstance) MarshalJSON() ([]byte, error) {
	type Alias DeleteInstancesResponseInstance
	return json.Marshal((Alias)(m))
}

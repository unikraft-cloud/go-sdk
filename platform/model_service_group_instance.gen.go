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

// An instance belonging to a service group.
type ServiceGroupInstance struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ServiceGroupInstance) UnmarshalJSON(data []byte) error {
	type Alias ServiceGroupInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ServiceGroupInstance) MarshalJSON() ([]byte, error) {
	type Alias ServiceGroupInstance
	return json.Marshal((Alias)(m))
}

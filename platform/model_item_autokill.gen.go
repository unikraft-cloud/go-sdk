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

// Automatic delete-on-idle configuration for the template instance.
type ItemAutokill struct {
	// Time in milliseconds after the template was last used for cloning
	// before it is deleted. A value of 0 disables template autokill.
	TimeMs *uint64 `json:"time_ms,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *ItemAutokill) UnmarshalJSON(data []byte) error {
	type Alias ItemAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ItemAutokill) MarshalJSON() ([]byte, error) {
	type Alias ItemAutokill
	return json.Marshal((Alias)(m))
}

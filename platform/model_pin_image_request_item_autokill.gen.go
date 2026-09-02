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

type PinImageRequestItemAutokill struct {
	// Automatically unpin the image after this many milliseconds of
	// inactivity.  `0` (the default) disables this.
	TimeMs uint64 `json:"time_ms"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *PinImageRequestItemAutokill) UnmarshalJSON(data []byte) error {
	type Alias PinImageRequestItemAutokill
	return json.Unmarshal(data, (*Alias)(m))
}

func (m PinImageRequestItemAutokill) MarshalJSON() ([]byte, error) {
	type Alias PinImageRequestItemAutokill
	return json.Marshal((Alias)(m))
}

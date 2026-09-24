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

// Why an instance stopped.
type AuditStop struct {
	// The origins that contributed to the stop.
	Reason []string `json:"reason,omitzero"`
	// The kernel stop code.
	Code *uint32 `json:"code,omitzero"`
	// The stop cause.
	Cause *string `json:"cause,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AuditStop) UnmarshalJSON(data []byte) error {
	type Alias AuditStop
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AuditStop) MarshalJSON() ([]byte, error) {
	type Alias AuditStop
	return json.Marshal((Alias)(m))
}

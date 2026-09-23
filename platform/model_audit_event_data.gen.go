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

// The payload of an audit event.
//
// Which fields are present depends on the event type:
//
// - `vm.state_change`
// - `vm.start_failed`
type AuditEventData struct {
	// The state before the transition.
	Prev *string `json:"prev,omitzero"`
	// The state after the transition.
	New *string `json:"new,omitzero"`
	// The state the instance was left in after a failed start.
	State *string `json:"state,omitzero"`
	// The error that failed the start, as an errno name such as `EDQUOT`.
	Error *string `json:"error,omitzero"`
	// Why the instance stopped.
	Stop *AuditStop `json:"stop,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AuditEventData) UnmarshalJSON(data []byte) error {
	type Alias AuditEventData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AuditEventData) MarshalJSON() ([]byte, error) {
	type Alias AuditEventData
	return json.Marshal((Alias)(m))
}

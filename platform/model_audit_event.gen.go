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

// One audit event.
type AuditEvent struct {
	// The type of event.
	Type AuditEventType `json:"type"`
	// When the event was raised.
	Timestamp time.Time `json:"timestamp"`
	// The object the event is about.
	Object AuditObject `json:"object"`
	// The operation the event belongs to.
	Attribution *AuditAttribution `json:"attribution,omitzero"`
	// The event payload. Its fields depend on `type`.
	Data *AuditEventData `json:"data,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AuditEvent) UnmarshalJSON(data []byte) error {
	type Alias AuditEvent
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AuditEvent) MarshalJSON() ([]byte, error) {
	type Alias AuditEvent
	return json.Marshal((Alias)(m))
}

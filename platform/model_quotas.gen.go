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

// Quotas with per-item response envelope fields merged in.
type Quotas struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the quota.
	Uuid string `json:"uuid"`
	// Used quota.
	Used QuotasStats `json:"used"`
	// Configured quota limits.
	Hard QuotasStats `json:"hard"`
	// Additional limits.
	Limits QuotasLimits `json:"limits"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Quotas) UnmarshalJSON(data []byte) error {
	type Alias Quotas
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Quotas) MarshalJSON() ([]byte, error) {
	type Alias Quotas
	return json.Marshal((Alias)(m))
}

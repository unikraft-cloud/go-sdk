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

type Quotas struct {
	// The UUID of the quota.
	Uuid string `json:"uuid"`
	// Used quota
	Used QuotasStats `json:"used"`
	// Configured quota limits
	Hard QuotasStats `json:"hard"`
	// Additional limits
	Limits QuotasLimits `json:"limits"`
	// An optional field representing the status of the request.  This field is
	// only set when this message object is used as a response message.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the user.
	Metro *string `json:"metro,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *Quotas) UnmarshalJSON(data []byte) error {
	type Alias Quotas
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Quotas) MarshalJSON() ([]byte, error) {
	type Alias Quotas
	return json.Marshal((Alias)(m))
}

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

type CreateAutoscaleConfigurationPolicyResponsePolicy struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// The UUID of the service of the added policy.
	Uuid string `json:"uuid"`
	// The name of the service of the added policy.
	Name string `json:"name"`
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

func (m *CreateAutoscaleConfigurationPolicyResponsePolicy) UnmarshalJSON(data []byte) error {
	type Alias CreateAutoscaleConfigurationPolicyResponsePolicy
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateAutoscaleConfigurationPolicyResponsePolicy) MarshalJSON() ([]byte, error) {
	type Alias CreateAutoscaleConfigurationPolicyResponsePolicy
	return json.Marshal((Alias)(m))
}

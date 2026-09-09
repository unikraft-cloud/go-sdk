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

// Per-item result for a delete autoscale configuration policy operation.
type DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse struct {
	// The name of the deleted policy.
	Name string `json:"name"`
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information about the status.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse) UnmarshalJSON(data []byte) error {
	type Alias DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse) MarshalJSON() ([]byte, error) {
	type Alias DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse
	return json.Marshal((Alias)(m))
}

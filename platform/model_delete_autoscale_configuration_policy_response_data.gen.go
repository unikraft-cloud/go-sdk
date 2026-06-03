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

type DeleteAutoscaleConfigurationPolicyResponseData struct {
	// The policies which were deleted by the request.
	Policies []DeleteAutoscaleConfigurationPolicyResponsePoliciesResponse `json:"policies,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *DeleteAutoscaleConfigurationPolicyResponseData) UnmarshalJSON(data []byte) error {
	type Alias DeleteAutoscaleConfigurationPolicyResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteAutoscaleConfigurationPolicyResponseData) MarshalJSON() ([]byte, error) {
	type Alias DeleteAutoscaleConfigurationPolicyResponseData
	return json.Marshal((Alias)(m))
}

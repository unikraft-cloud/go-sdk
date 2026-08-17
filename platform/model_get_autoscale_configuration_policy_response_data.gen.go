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

type GetAutoscaleConfigurationPolicyResponseData struct {
	// The policy which was retrieved by the request.
	Policies []GetAutoscaleConfigurationPolicyResponsePolicyResponse `json:"policies,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetAutoscaleConfigurationPolicyResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationPolicyResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetAutoscaleConfigurationPolicyResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationPolicyResponseData
	return json.Marshal((Alias)(m))
}

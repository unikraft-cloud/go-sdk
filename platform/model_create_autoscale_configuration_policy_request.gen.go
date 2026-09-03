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

// The request message to create an autoscale configuration policy for a
// service.
type CreateAutoscaleConfigurationPolicyRequest struct {
	// The Name of the service to add a policy to.
	Name string `json:"name"`
	// The policy type to add to the autoscale configuration.
	Type AutoscalePolicy `json:"type"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateAutoscaleConfigurationPolicyRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateAutoscaleConfigurationPolicyRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateAutoscaleConfigurationPolicyRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateAutoscaleConfigurationPolicyRequest
	return json.Marshal((Alias)(m))
}

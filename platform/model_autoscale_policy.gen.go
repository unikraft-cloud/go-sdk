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

// AutoscalePolicy defines the autoscale policy for a service.
// Right now it contains fields from both the `ondemand` and `step` policies.
// They are marked both as optional, so only one of them should be set at a
// time. This is a current limitation of the API design.
type AutoscalePolicy struct {
	// The name of the policy.
	Name string `json:"name"`
	// If the policy is enabled.
	Enabled *bool `json:"enabled,omitzero"`
	// Metric to use for the step policy.
	Metric *StepPolicyMetric `json:"metric,omitzero"`
	// The type of adjustment to be made in the step policy.
	AdjustmentType *AdjustmentType `json:"adjustment_type,omitzero"`
	// The steps for the step policy.
	// Each step defines an adjustment value and optional bounds.
	Steps []AutoscalePolicyStep `json:"steps,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AutoscalePolicy) UnmarshalJSON(data []byte) error {
	type Alias AutoscalePolicy
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AutoscalePolicy) MarshalJSON() ([]byte, error) {
	type Alias AutoscalePolicy
	return json.Marshal((Alias)(m))
}

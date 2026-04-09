// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The policy which was retrieved by the request.
// Metric to use for the step policy.
type GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyMetric string

const (
	GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyMetricCpu GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyMetric = "cpu"
)

// The type of adjustment to be made in the step policy.
type GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentType string

const (
	GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentTypeChange     GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentType = "change"
	GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentTypeExact      GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentType = "exact"
	GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentTypePercentage GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentType = "percentage"
)

type GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy struct {
	// The name of the policy.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the service group the policy applies to.
	Metro *string `json:"metro,omitempty"`
	// If the policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Metric to use for the step policy.
	Metric *GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyMetric `json:"metric,omitempty"`
	// The type of adjustment to be made in the step policy.
	AdjustmentType *GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicyAdjustmentType `json:"adjustment_type,omitempty"`
	// The steps for the step policy.
	// Each step defines an adjustment value and optional bounds.
	Steps []AutoscalePolicyStep `json:"steps,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":            {},
		"metro":           {},
		"enabled":         {},
		"metric":          {},
		"adjustment_type": {},
		"steps":           {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

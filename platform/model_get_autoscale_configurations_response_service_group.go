// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type GetAutoscaleConfigurationsResponseServiceGroupStatus string

const (
	GetAutoscaleConfigurationsResponseServiceGroupStatusSuccess GetAutoscaleConfigurationsResponseServiceGroupStatus = "success"
	GetAutoscaleConfigurationsResponseServiceGroupStatusError   GetAutoscaleConfigurationsResponseServiceGroupStatus = "error"
)

type GetAutoscaleConfigurationsResponseServiceGroup struct {
	// The status of the response.
	Status *GetAutoscaleConfigurationsResponseServiceGroupStatus `json:"status,omitempty"`
	// The UUID of the service where the configuration was created.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service where the configuration was created.
	Name *string `json:"name,omitempty"`
	// If the autoscale configuration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The minimum number of instances to keep running.
	// Only if enabled is true.
	MinSize *int64 `json:"min_size,omitempty"`
	// The maximum number of instances to keep running.
	// Only if enabled is true.
	MaxSize *int64 `json:"max_size,omitempty"`
	// The warmup time in seconds for new instances.
	// Only if enabled is true.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitempty"`
	// The cooldown time in seconds for the autoscale configuration.
	// Only if enabled is true.
	CooldownTimeMs *int64                                                  `json:"cooldown_time_ms,omitempty"`
	Template       *GetAutoscaleConfigurationsResponseServiceGroupTemplate `json:"template,omitempty"`
	// The policies applied to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitempty"`
}

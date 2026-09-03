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

type GetAutoscaleConfigurationsResponseServiceGroup struct {
	// The status of the response.
	Status ResponseStatus `json:"status"`
	// The UUID of the service where the configuration was created.
	Uuid string `json:"uuid"`
	// The name of the service where the configuration was created.
	Name string `json:"name"`
	// If the autoscale configuration is enabled.
	Enabled bool `json:"enabled"`
	// The minimum number of instances to keep running.
	// Only if enabled is true.
	MinSize *int64 `json:"min_size,omitzero"`
	// The maximum number of instances to keep running.
	// Only if enabled is true.
	MaxSize *int64 `json:"max_size,omitzero"`
	// The warmup time in seconds for new instances.
	// Only if enabled is true.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitzero"`
	// The cooldown time in seconds for the autoscale configuration.
	// Only if enabled is true.
	CooldownTimeMs *int64 `json:"cooldown_time_ms,omitzero"`
	// The instance template used for the autoscale configuration.
	// Only if enabled is true.
	Template *ServiceGroupTemplate `json:"template,omitzero"`
	// The policies applied to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitzero"`
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

func (m *GetAutoscaleConfigurationsResponseServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias GetAutoscaleConfigurationsResponseServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetAutoscaleConfigurationsResponseServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias GetAutoscaleConfigurationsResponseServiceGroup
	return json.Marshal((Alias)(m))
}

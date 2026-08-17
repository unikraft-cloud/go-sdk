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

// The request message to create an autoscale configuration for a service group
// based on its UUID.
type CreateAutoscaleConfigurationByServiceGroupUUIDRequest struct {
	// The UUID of the service to create a configuration for.
	// Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// The minimum number of instances to keep running.
	MinSize *int64 `json:"min_size,omitzero"`
	// The maximum number of instances to keep running.
	MaxSize *int64 `json:"max_size,omitzero"`
	// The warmup time in milliseconds for new instances.
	WarmupTimeMs *int64 `json:"warmup_time_ms,omitzero"`
	// The cooldown time in milliseconds for the autoscale configuration.
	CooldownTimeMs *int64 `json:"cooldown_time_ms,omitzero"`
	// The arguments to use when creating the autoscale configuration.
	CreateArgs CreateAutoscaleConfigurationByServiceGroupUUIDRequestInstanceCreateArgs `json:"create_args"`
	// The policies to apply to the autoscale configuration.
	Policies []AutoscalePolicy `json:"policies,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateAutoscaleConfigurationByServiceGroupUUIDRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateAutoscaleConfigurationByServiceGroupUUIDRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateAutoscaleConfigurationByServiceGroupUUIDRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateAutoscaleConfigurationByServiceGroupUUIDRequest
	return json.Marshal((Alias)(m))
}

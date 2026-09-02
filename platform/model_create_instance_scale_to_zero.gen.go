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

type CreateInstanceScaleToZero struct {
	// The specific policy to use for scaling the instance to zero.
	Policy *InstanceScaleToZeroPolicy `json:"policy,omitzero"`
	// Whether the instance should be stateful when scaled to zero. If set to
	// true, the instance will retain its state (e.g., RAM contents) when scaled
	// to zero.  This is useful for instances that need to maintain their state
	// across scale-to-zero operations.  If set to false, the instance will lose
	// its state when scaled to zero, and it will be restarted from scratch when
	// scaled back up.
	Stateful *bool `json:"stateful,omitzero"`
	// The cooldown time in milliseconds before the instance can be scaled to
	// zero again.  This is useful to prevent rapid scaling to zero and back up,
	// which can lead to performance issues or resource exhaustion.
	CooldownTimeMs *int32 `json:"cooldown_time_ms,omitzero"`
	// The notification time in milliseconds before the instance is scaled to
	// zero. This allows the instance to perform any necessary cleanup or state
	// saving before being scaled down.
	NotifyTimeMs *int32 `json:"notify_time_ms,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceScaleToZero) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceScaleToZero
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceScaleToZero) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceScaleToZero
	return json.Marshal((Alias)(m))
}

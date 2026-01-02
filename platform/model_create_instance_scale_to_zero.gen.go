// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The specific policy to use for scaling the instance to zero.
type CreateInstanceScaleToZeroPolicy string

const (
	CreateInstanceScaleToZeroPolicyOn   CreateInstanceScaleToZeroPolicy = "on"
	CreateInstanceScaleToZeroPolicyOff  CreateInstanceScaleToZeroPolicy = "off"
	CreateInstanceScaleToZeroPolicyIdle CreateInstanceScaleToZeroPolicy = "idle"
)

type CreateInstanceScaleToZero struct {
	// The specific policy to use for scaling the instance to zero.
	Policy *CreateInstanceScaleToZeroPolicy `json:"policy,omitempty"`
	// Whether the instance should be stateful when scaled to zero. If set to
	// true, the instance will retain its state (e.g., RAM contents) when scaled
	// to zero.  This is useful for instances that need to maintain their state
	// across scale-to-zero operations.  If set to false, the instance will lose
	// its state when scaled to zero, and it will be restarted from scratch when
	// scaled back up.
	Stateful *bool `json:"stateful,omitempty"`
	// The cooldown time in milliseconds before the instance can be scaled to
	// zero again.  This is useful to prevent rapid scaling to zero and back up,
	// which can lead to performance issues or resource exhaustion.
	CooldownTimeMs *int32 `json:"cooldown_time_ms,omitempty"`
	// The notification time in milliseconds before the instance is scaled to
	// zero. This allows the instance to perform any necessary cleanup or state
	// saving before being scaled down.
	NotifyTimeMs *int32 `json:"notify_time_ms,omitempty"`
}

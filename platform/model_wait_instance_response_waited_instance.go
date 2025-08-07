// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The current state of the instance.
type WaitInstanceResponseWaitedInstanceState string

const (
	WaitInstanceResponseWaitedInstanceStateStopped  WaitInstanceResponseWaitedInstanceState = "stopped"
	WaitInstanceResponseWaitedInstanceStateStarting WaitInstanceResponseWaitedInstanceState = "starting"
	WaitInstanceResponseWaitedInstanceStateRunning  WaitInstanceResponseWaitedInstanceState = "running"
	WaitInstanceResponseWaitedInstanceStateDraining WaitInstanceResponseWaitedInstanceState = "draining"
	WaitInstanceResponseWaitedInstanceStateStopping WaitInstanceResponseWaitedInstanceState = "stopping"
	WaitInstanceResponseWaitedInstanceStateStandby  WaitInstanceResponseWaitedInstanceState = "standby"
)

type WaitInstanceResponseWaitedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// The current state of the instance.
	State *WaitInstanceResponseWaitedInstanceState `json:"state,omitempty"`
}

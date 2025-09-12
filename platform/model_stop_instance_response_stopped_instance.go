// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The current state of the instance.
type StopInstanceResponseStoppedInstanceState string

const (
	StopInstanceResponseStoppedInstanceStateStopped  StopInstanceResponseStoppedInstanceState = "stopped"
	StopInstanceResponseStoppedInstanceStateStarting StopInstanceResponseStoppedInstanceState = "starting"
	StopInstanceResponseStoppedInstanceStateRunning  StopInstanceResponseStoppedInstanceState = "running"
	StopInstanceResponseStoppedInstanceStateDraining StopInstanceResponseStoppedInstanceState = "draining"
	StopInstanceResponseStoppedInstanceStateStopping StopInstanceResponseStoppedInstanceState = "stopping"
	StopInstanceResponseStoppedInstanceStateStandby  StopInstanceResponseStoppedInstanceState = "standby"
)

// The previous state of the instance before the stop operation was invoked.
type StopInstanceResponseStoppedInstancePreviousState string

const (
	StopInstanceResponseStoppedInstancePreviousStateStopped  StopInstanceResponseStoppedInstancePreviousState = "stopped"
	StopInstanceResponseStoppedInstancePreviousStateStarting StopInstanceResponseStoppedInstancePreviousState = "starting"
	StopInstanceResponseStoppedInstancePreviousStateRunning  StopInstanceResponseStoppedInstancePreviousState = "running"
	StopInstanceResponseStoppedInstancePreviousStateDraining StopInstanceResponseStoppedInstancePreviousState = "draining"
	StopInstanceResponseStoppedInstancePreviousStateStopping StopInstanceResponseStoppedInstancePreviousState = "stopping"
	StopInstanceResponseStoppedInstancePreviousStateStandby  StopInstanceResponseStoppedInstancePreviousState = "standby"
)

type StopInstanceResponseStoppedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// The current state of the instance.
	State *StopInstanceResponseStoppedInstanceState `json:"state,omitempty"`
	// The previous state of the instance before the stop operation was invoked.
	PreviousState *StopInstanceResponseStoppedInstancePreviousState `json:"previous_state,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}

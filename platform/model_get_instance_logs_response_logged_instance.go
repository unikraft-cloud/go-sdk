// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// State of the instance when the logs were retrieved.
type GetInstanceLogsResponseLoggedInstanceState string

const (
	GetInstanceLogsResponseLoggedInstanceStateStopped  GetInstanceLogsResponseLoggedInstanceState = "stopped"
	GetInstanceLogsResponseLoggedInstanceStateStarting GetInstanceLogsResponseLoggedInstanceState = "starting"
	GetInstanceLogsResponseLoggedInstanceStateRunning  GetInstanceLogsResponseLoggedInstanceState = "running"
	GetInstanceLogsResponseLoggedInstanceStateDraining GetInstanceLogsResponseLoggedInstanceState = "draining"
	GetInstanceLogsResponseLoggedInstanceStateStopping GetInstanceLogsResponseLoggedInstanceState = "stopping"
	GetInstanceLogsResponseLoggedInstanceStateStandby  GetInstanceLogsResponseLoggedInstanceState = "standby"
)

type GetInstanceLogsResponseLoggedInstance struct {
	// The UUID of the instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance.
	Name *string `json:"name,omitempty"`
	// Base64 encoded log output of the instance.
	Output    *string                                         `json:"output,omitempty"`
	Available *GetInstanceLogsResponseLoggedInstanceAvailable `json:"available,omitempty"`
	Range     *GetInstanceLogsResponseLoggedInstanceRange     `json:"range,omitempty"`
	// State of the instance when the logs were retrieved.
	State *GetInstanceLogsResponseLoggedInstanceState `json:"state,omitempty"`
}

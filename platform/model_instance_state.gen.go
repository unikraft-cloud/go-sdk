// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The current state of an instance.
type InstanceState string

const (
	InstanceStateStopped  InstanceState = "stopped"
	InstanceStateStarting InstanceState = "starting"
	InstanceStateRunning  InstanceState = "running"
	InstanceStateDraining InstanceState = "draining"
	InstanceStateStopping InstanceState = "stopping"
	InstanceStateTemplate InstanceState = "template"
	InstanceStateStandby  InstanceState = "standby"
)

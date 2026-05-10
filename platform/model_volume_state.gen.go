// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// VolumeState defines the state of a volume at a given moment.
type VolumeState string

const (
	VolumeStateUNINITIALIZED VolumeState = "uninitialized"
	VolumeStateINITIALIZING  VolumeState = "initializing"
	VolumeStateAVAILABLE     VolumeState = "available"
	VolumeStateIDLE          VolumeState = "idle"
	VolumeStateMOUNTED       VolumeState = "mounted"
	VolumeStateBUSY          VolumeState = "busy"
	VolumeStateERROR         VolumeState = "error"
	VolumeStateTEMPLATE      VolumeState = "template"
)

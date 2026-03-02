// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Current state of the volume.
type CreateVolumeResponseVolumeState string

const (
	CreateVolumeResponseVolumeStateUninitialized CreateVolumeResponseVolumeState = "uninitialized"
	CreateVolumeResponseVolumeStateInitializing  CreateVolumeResponseVolumeState = "initializing"
	CreateVolumeResponseVolumeStateAvailable     CreateVolumeResponseVolumeState = "available"
	CreateVolumeResponseVolumeStateIdle          CreateVolumeResponseVolumeState = "idle"
	CreateVolumeResponseVolumeStateMounted       CreateVolumeResponseVolumeState = "mounted"
	CreateVolumeResponseVolumeStateBusy          CreateVolumeResponseVolumeState = "busy"
	CreateVolumeResponseVolumeStateError         CreateVolumeResponseVolumeState = "error"
)

type CreateVolumeResponseVolume struct {
	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`
	// Current state of the volume.
	State *CreateVolumeResponseVolumeState `json:"state,omitempty"`
	// UUID of the newly created volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly created volume.
	Name *string `json:"name,omitempty"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
}

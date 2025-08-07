// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type CreateVolumeResponseVolumeStatus string

const (
	CreateVolumeResponseVolumeStatusSuccess CreateVolumeResponseVolumeStatus = "success"
	CreateVolumeResponseVolumeStatusError   CreateVolumeResponseVolumeStatus = "error"
)

type CreateVolumeResponseVolume struct {
	// The status of the response.
	Status *CreateVolumeResponseVolumeStatus `json:"status,omitempty"`
	// UUID of the newly created volume.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the newly created volume.
	Name *string `json:"name,omitempty"`
}

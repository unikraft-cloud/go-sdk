// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type AttachVolumesResponseAttachedVolumeStatus string

const (
	AttachVolumesResponseAttachedVolumeStatusSuccess AttachVolumesResponseAttachedVolumeStatus = "success"
	AttachVolumesResponseAttachedVolumeStatusError   AttachVolumesResponseAttachedVolumeStatus = "error"
)

type AttachVolumesResponseAttachedVolume struct {
	// The status of the response.
	Status *AttachVolumesResponseAttachedVolumeStatus `json:"status,omitempty"`
	// The UUID of the volume that was attached.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume that was attached.
	Name *string `json:"name,omitempty"`
}

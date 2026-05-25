// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// VolumeAccessMode defines the access mode of a volume, which controls sharing
// behavior and caching strategy.
type VolumeAccessMode string

const (
	VolumeAccessModeRwo VolumeAccessMode = "rwo"
	VolumeAccessModeRox VolumeAccessMode = "rox"
	VolumeAccessModeRwx VolumeAccessMode = "rwx"
)

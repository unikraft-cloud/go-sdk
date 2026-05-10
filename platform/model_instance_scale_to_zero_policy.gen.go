// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Scale-to-zero policy types to use when creating an instance.
//
// The policy can be one of the following values:
//
// | Policy | Description |
// |--------|-------------|
// | `on`   | The instance should be scaled-to-zero (when possible). |
// | `off`  | The instance should never be scaled-to-zero. |
// | `idle` | The instance should be scaled-to-zero even with established but idle TCP connections. |
type InstanceScaleToZeroPolicy string

const (
	InstanceScaleToZeroPolicyOFF  InstanceScaleToZeroPolicy = "off"
	InstanceScaleToZeroPolicyON   InstanceScaleToZeroPolicy = "on"
	InstanceScaleToZeroPolicyIDLE InstanceScaleToZeroPolicy = "idle"
)

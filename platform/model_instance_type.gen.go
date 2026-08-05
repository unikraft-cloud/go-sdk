// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The type of virtual machine used to run an instance.
//
// | Type    | Description |
// |---------|-------------|
// | `micro` | A lightweight microVM (default).  Boots in milliseconds and is suitable for most workloads. |
// | `full`  | A full virtual machine with broader hardware support, such as GPU passthrough.  Requires a plan with full VM support. |
type InstanceType string

const (
	InstanceTypeMicro InstanceType = "micro"
	InstanceTypeFull  InstanceType = "full"
)

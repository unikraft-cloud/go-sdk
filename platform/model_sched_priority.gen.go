// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// SchedPriority defines the scheduling priority for an instance.
// User requires the `override_vm_priority` permission to change it.
//
// The list of available scheduling priorities:
//
// | Priority | Description |
// |----------|-------------|
// | `normal` | Default scheduling priority. |
// | `medium` | Medium scheduling priority. |
// | `high`   | High scheduling priority. |
// | `admin`  | Admin scheduling priority. |
type SchedPriority string

const (
	SchedPriorityNormal SchedPriority = "normal"
	SchedPriorityMedium SchedPriority = "medium"
	SchedPriorityHigh   SchedPriority = "high"
	SchedPriorityAdmin  SchedPriority = "admin"
)

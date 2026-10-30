// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// What resuming a user does with an instance that the suspension of the user
// brought to rest.
//
// | Action    | Description |
// |-----------|-------------|
// | `standby` | Autostart is enabled again. The instance is in standby and starts on the next incoming request. |
// | `start`   | The instance is started. |
type UserResumeAction string

const (
	UserResumeActionStandby UserResumeAction = "standby"
	UserResumeActionStart   UserResumeAction = "start"
)

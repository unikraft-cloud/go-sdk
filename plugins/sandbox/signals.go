// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import "syscall"

// The instance runs Linux, so the signals sent to a command carry the Linux
// numbers, not the numbers of the host.
const (
	sigINT  syscall.Signal = 2
	sigKILL syscall.Signal = 9
	sigPIPE syscall.Signal = 13
)

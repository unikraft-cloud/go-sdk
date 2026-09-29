// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

//go:build !windows

package sandbox

import (
	"errors"
	"syscall"
)

// isBrokenPipe reports whether a write failed because the reader has gone.
func isBrokenPipe(err error) bool {
	return errors.Is(err, syscall.EPIPE)
}

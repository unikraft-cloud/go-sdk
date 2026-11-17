// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrnoStringZeroIsEmpty(t *testing.T) {
	assert.Empty(t, Errno(0).String())
}

func TestErrnoStringKnownValue(t *testing.T) {
	assert.Equal(t, "ENOMEM", Errno(syscall.ENOMEM).String())
	assert.Equal(t, "EACCES", Errno(syscall.EACCES).String())
}

func TestErrnoStringUnknownValueFallsBackToNumeric(t *testing.T) {
	assert.Equal(t, "errno(999999)", Errno(999999).String())
}

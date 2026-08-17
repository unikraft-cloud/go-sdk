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

func TestFlatStopCodeString(t *testing.T) {
	assert.Equal(t, "unknown", FlatStopCodeUnknown.String())
	assert.Equal(t, "image pull failed", FlatStopCodeImagePullFailed.String())
	assert.Equal(t, "code(42)", FlatStopCode(42).String())
}

func TestKernelStopCodeFieldAccessors(t *testing.T) {
	// errno=ENOMEM(0xc), shutdown=1, initlevel=5, reason=PGFAULT(4)
	sc := KernelStopCode(0xc)<<16 | KernelStopCode(1)<<15 | KernelStopCode(5)<<8 | KernelStopCode(StopCodeReasonPGFAULT)

	assert.Equal(t, syscall.ENOMEM, sc.Errno())
	assert.Equal(t, "ENOMEM", sc.ErrnoString())
	assert.Equal(t, uint8(1), sc.ShutdownTable())
	assert.Equal(t, uint8(5), sc.InitLevel())
	assert.Equal(t, StopCodeReasonPGFAULT, sc.Reason())
}

func TestKernelStopCodeErrnoStringEmptyWhenZero(t *testing.T) {
	sc := KernelStopCode(StopCodeReasonOK)
	assert.Equal(t, syscall.Errno(0), sc.Errno())
	assert.Equal(t, "", sc.ErrnoString())
}

func TestKernelStopCodeDescription(t *testing.T) {
	tests := []struct {
		name   string
		reason KernelStopCodeReason
		errno  syscall.Errno
		want   string
	}{
		{"ok", StopCodeReasonOK, 0, ""},
		{"assertion", StopCodeReasonEXP, 0, "assertion error"},
		{"pgfault enomem", StopCodeReasonPGFAULT, syscall.ENOMEM, "out of memory"},
		{"pgfault efault", StopCodeReasonPGFAULT, syscall.EFAULT, "illegal memory access"},
		{"pgfault eperm", StopCodeReasonPGFAULT, syscall.EPERM, "illegal memory access"},
		{"pgfault other", StopCodeReasonPGFAULT, syscall.EIO, "page fault"},
		{"segfault", StopCodeReasonSEGFAULT, 0, "segmentation fault"},
		{"math", StopCodeReasonMATH, 0, "arithmetic error"},
		{"invlop", StopCodeReasonINVLOP, 0, "instruction error"},
		{"hwerr", StopCodeReasonHWERR, 0, "hardware error"},
		{"secerr", StopCodeReasonSECERR, 0, "security violation"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := KernelStopCode(tt.errno)<<16 | KernelStopCode(tt.reason)
			assert.Equal(t, tt.want, sc.Description())
		})
	}
}

func TestKernelStopCodeDescriptionUnexpectedReason(t *testing.T) {
	sc := KernelStopCode(0xFF) // reason byte outside the known enum values
	assert.Equal(t, "unexpected error", sc.Description())
}

func TestKernelStopCodeStringCombinesDescriptionAndErrno(t *testing.T) {
	sc := KernelStopCode(syscall.ENOMEM)<<16 | KernelStopCode(StopCodeReasonPGFAULT)
	assert.Equal(t, "out of memory (ENOMEM)", sc.String())
}

func TestKernelStopCodeStringWithoutErrno(t *testing.T) {
	sc := KernelStopCode(StopCodeReasonSEGFAULT)
	assert.Equal(t, "segmentation fault", sc.String())
}

func TestKernelStopCodeReasonString(t *testing.T) {
	assert.Equal(t, "OK", StopCodeReasonOK.String())
	assert.Equal(t, "SEGFAULT", StopCodeReasonSEGFAULT.String())
	assert.Equal(t, "reason(200)", KernelStopCodeReason(200).String())
}

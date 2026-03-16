// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import (
	"fmt"
	"syscall"
)

// FlatStopCode is a simple stop code that is created by the platform.
type FlatStopCode uint32

const (
	// FlatStopCodeUnknown is the default stop code when the platform does not
	// have a more specific code to provide.
	FlatStopCodeUnknown FlatStopCode = 0

	// FlatStopCodeImagePullFailed indicates that the platform failed to pull the
	// image for the instance.
	FlatStopCodeImagePullFailed FlatStopCode = 1
)

func (f FlatStopCode) String() string {
	switch f {
	case FlatStopCodeUnknown:
		return "unknown"
	case FlatStopCodeImagePullFailed:
		return "image pull failed"
	default:
		return fmt.Sprintf("code(%d)", uint32(f))
	}
}

// KernelStopCode is the kernel stop code.
//
// It is a packed bitfield described by the KernelStopCodeMask* constants.
type KernelStopCode uint32

// Stop code of the kernel. This value encodes multiple details about the stop
// irrespective of the application.
//
// MSB                                                     LSB
// ┌──────────────┬──────────┬──────────┬───────────┬────────┐
// │ 31 ────── 24 │ 23 ── 16 │    15    │ 14 ──── 8 │ 7 ── 0 │
// ├──────────────┼──────────┼──────────┼───────────┼────────┤
// │ reserved[^1] │ errno    │ shutdown │ initlevel │ reason │
// └──────────────┴──────────┴──────────┴───────────┴────────┘
//
// [^1]: Reserved for future use.
// errno:     The application errno (Linux errno.h value; optional, may be 0).
// shutdown:  Whether the stop originated from the inittable (0) or termtable (1).
// initlevel: The initlevel at the time of the stop.
// reason:    The reason for the stop (see KernelStopCodeReason).
const (
	KernelStopCodeMaskErrno     KernelStopCode = 0xFF0000
	KernelStopCodeMaskShutdown  KernelStopCode = 0x008000
	KernelStopCodeMaskInitLevel KernelStopCode = 0x007F00
	KernelStopCodeMaskReason    KernelStopCode = 0x0000FF
)

func (sc KernelStopCode) String() string {
	result := sc.Description()
	if errnoStr := sc.ErrnoString(); errnoStr != "" {
		result += fmt.Sprintf(" (%s)", errnoStr)
	}
	return result
}

func (sc KernelStopCode) Errno() syscall.Errno {
	return syscall.Errno((sc & KernelStopCodeMaskErrno) >> 16)
}

func (sc KernelStopCode) ErrnoString() string {
	errno := sc.Errno()
	if errno == 0 {
		return ""
	}
	return Errno(errno).String()
}

// ShutdownTable returns whether the stop originated from the inittable (0) or
// from the termtable (1).
func (sc KernelStopCode) ShutdownTable() uint8 {
	return uint8((sc & KernelStopCodeMaskShutdown) >> 15)
}

// InitLevel returns the initlevel at the time of the stop.
func (sc KernelStopCode) InitLevel() uint8 {
	return uint8((sc & KernelStopCodeMaskInitLevel) >> 8)
}

// Reason returns the stop code reason.
func (sc KernelStopCode) Reason() KernelStopCodeReason {
	return KernelStopCodeReason(sc & KernelStopCodeMaskReason)
}

func (sc KernelStopCode) Description() string {
	switch sc.Reason() {
	case StopCodeReasonOK:
		return ""
	case StopCodeReasonEXP:
		return "assertion error"
	case StopCodeReasonPGFAULT:
		switch sc.Errno() {
		case syscall.ENOMEM:
			return "out of memory"
		case syscall.EFAULT, syscall.EPERM:
			return "illegal memory access"
		}
		return "page fault"
	case StopCodeReasonSEGFAULT:
		return "segmentation fault"
	case StopCodeReasonMATH:
		return "arithmetic error"
	case StopCodeReasonINVLOP:
		return "instruction error"
	case StopCodeReasonHWERR:
		return "hardware error"
	case StopCodeReasonSECERR:
		return "security violation"
	default:
		return "unexpected error"
	}
}

// KernelStopCodeReason is the reason component of KernelStopCode.
type KernelStopCodeReason uint8

const (
	// 0 - Success
	StopCodeReasonOK KernelStopCodeReason = iota

	// 1 - Explicit crash (bugon/assert/crash/unhandled breakpoint)
	StopCodeReasonEXP

	// 2 - Arithmetic error
	StopCodeReasonMATH

	// 3 - Invalid CPU instruction or instruction error (e.g., operand alignment)
	StopCodeReasonINVLOP

	// 4 - Page fault - see errno (out of mem, EFAULT)
	StopCodeReasonPGFAULT

	// 5 - Segmentation fault
	StopCodeReasonSEGFAULT

	// 6 - Hardware error, NMI, alignment checks
	StopCodeReasonHWERR

	// 7 - Security violation, control protection (MTE, shadow stacks, PKU?)
	StopCodeReasonSECERR
)

// kernelStopCodeReasons is the list of stop code reason string representations.
var kernelStopCodeReasons = []string{
	"OK",
	"EXP",
	"MATH",
	"INVLOP",
	"PGFAULT",
	"SEGFAULT",
	"HWERR",
	"SECERR",
}

func (r KernelStopCodeReason) String() string {
	if int(r) < len(kernelStopCodeReasons) {
		return kernelStopCodeReasons[r]
	}
	return fmt.Sprintf("reason(%d)", uint8(r))
}

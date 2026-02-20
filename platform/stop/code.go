// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import (
	"fmt"
	"syscall"
)

// StopCode is the kernel stop code.
//
// It is a packed bitfield described by the StopCodeMask* constants.
type StopCode uint32

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
// reason:    The reason for the stop (see StopCodeReason).
const (
	StopCodeMaskErrno     StopCode = 0xFF0000
	StopCodeMaskShutdown  StopCode = 0x008000
	StopCodeMaskInitLevel StopCode = 0x007F00
	StopCodeMaskReason    StopCode = 0x0000FF
)

func (sc StopCode) Errno() syscall.Errno {
	return syscall.Errno((sc & StopCodeMaskErrno) >> 16)
}

func (sc StopCode) ErrnoString() string {
	errno := sc.Errno()
	if errno == 0 {
		return ""
	}
	if name, ok := errnoNames[errno]; ok {
		return name
	}
	return fmt.Sprintf("errno(%d)", uint32(errno))
}

// ShutdownTable returns whether the stop originated from the inittable (0) or
// from the termtable (1).
func (sc StopCode) ShutdownTable() uint8 {
	return uint8((sc & StopCodeMaskShutdown) >> 15)
}

// InitLevel returns the initlevel at the time of the stop.
func (sc StopCode) InitLevel() uint8 {
	return uint8((sc & StopCodeMaskInitLevel) >> 8)
}

// Reason returns the stop code reason.
func (sc StopCode) Reason() StopCodeReason {
	return StopCodeReason(sc & StopCodeMaskReason)
}

func (sc StopCode) Description() string {
	switch sc.Reason() {
	case StopCodeReasonOK:
		return "shutdown"
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

// StopCodeReason is the reason component of StopCode.
type StopCodeReason uint8

const (
	// 0 - Success
	StopCodeReasonOK StopCodeReason = iota

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

// stopCodeReasons is the list of stop code reason string representations.
var stopCodeReasons = []string{
	"OK",
	"EXP",
	"MATH",
	"INVLOP",
	"PGFAULT",
	"SEGFAULT",
	"HWERR",
	"SECERR",
}

func (r StopCodeReason) String() string {
	if int(r) < len(stopCodeReasons) {
		return stopCodeReasons[r]
	}
	return fmt.Sprintf("reason(%d)", uint8(r))
}

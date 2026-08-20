// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

// StopReason is a bitmask that describes why an instance stopped.
type StopReason uint32

func (sr StopReason) Forced() bool {
	return (sr & StopReasonForced) != 0
}

// Origin returns the primary origin of the stop.
func (sr StopReason) Origin() string {
	if sr == 0 {
		return "unknown"
	}

	switch {
	case sr&StopReasonPlatform == StopReasonPlatform && sr&StopReasonUser != StopReasonUser:
		return "platform"
	case sr&StopReasonUser == StopReasonUser:
		return "user"
	case sr&StopReasonApplication == StopReasonApplication:
		return "app"
	case sr&StopReasonKernel == StopReasonKernel:
		return "kernel"
	default:
		return "unknown"
	}
}

// OriginCode returns a short-code representation of the stop reason.
//
// Format: "FUPAK", using '-' for absent bits.
func (sr StopReason) OriginCode() string {
	code := [5]byte{'-', '-', '-', '-', '-'}
	if sr&StopReasonForced != 0 {
		code[0] = 'F'
	}
	if sr&StopReasonUser != 0 {
		code[1] = 'U'
	}
	if sr&StopReasonPlatform != 0 {
		code[2] = 'P'
	}
	if sr&StopReasonApplication != 0 {
		code[3] = 'A'
	}
	if sr&StopReasonKernel != 0 {
		code[4] = 'K'
	}
	return string(code[:])
}

// String returns a human-readable description of the stop reason.
func (sr StopReason) String() string {
	switch StopReason(sr) {
	case StopReasonUnknown:
		return "unknown stop"
	case StopReasonKernelCrash:
		return "kernel crash"
	case StopReasonAppExit:
		return "app exit"
	case StopReasonPlatformShutdown:
		return "platform shutdown"
	case StopReasonUserShutdownIncomplete:
		return "user shutdown incomplete"
	case StopReasonUserShutdownComplete:
		return "user shutdown complete"
	case StopReasonForcedUserShutdown:
		return "forced user shutdown"
	default:
		origin := sr.Origin()
		if origin == "unknown" {
			return "unknown stop"
		}
		if sr.Forced() {
			return "forced " + origin + " stop"
		}
		return origin + " stop"
	}
}

const (
	// StopReason.OriginCode() renders these bits as "FUPAK".
	// ----K
	StopReasonKernel StopReason = 1 << iota
	// ---A-
	StopReasonApplication
	// --P--
	StopReasonPlatform
	// -U---
	StopReasonUser
	// F----
	StopReasonForced

	// Common stop reason scenarios.
	StopReasonUnknown StopReason = 0
	// ----K
	StopReasonKernelCrash StopReason = StopReasonKernel
	// ---AK
	StopReasonAppExit StopReason = StopReasonApplication | StopReasonKernel
	// --PAK
	StopReasonPlatformShutdown StopReason = StopReasonPlatform | StopReasonApplication | StopReasonKernel
	// -UP-K
	StopReasonUserShutdownIncomplete StopReason = StopReasonUser | StopReasonPlatform | StopReasonKernel
	// -UPAK
	StopReasonUserShutdownComplete StopReason = StopReasonUser | StopReasonPlatform | StopReasonApplication | StopReasonKernel
	// FUP--
	StopReasonForcedUserShutdown StopReason = StopReasonForced | StopReasonUser | StopReasonPlatform
)

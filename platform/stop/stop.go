// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import "fmt"

// Stop represents the state of a stopped instance.
//
// The main reason for the stop is represented by the StopReason, while the
// StopCode provides additional details, dependent on the reason.
type Stop struct {
	StopReason
	StopCode *uint32
}

func (s Stop) String() string {
	var codeString string
	if s.StopCode != nil {
		if s.StopReason == StopReasonPlatform {
			stopCode := FlatStopCode(*s.StopCode)
			codeString = stopCode.String()
		} else if s.StopReason&StopReasonKernel != 0 {
			stopCode := KernelStopCode(*s.StopCode)
			codeString = stopCode.String()
		} else {
			codeString = fmt.Sprintf("code(%d)", *s.StopCode)
		}
	}

	result := s.StopReason.String()
	if codeString != "" {
		result += ": " + codeString
	}
	return result
}

func (s Stop) FlatStopCode() *FlatStopCode {
	if s.StopCode == nil {
		return nil
	}
	if s.StopReason != StopReasonPlatform {
		return nil
	}
	stopCode := FlatStopCode(*s.StopCode)
	return &stopCode
}

func (s Stop) KernelStopCode() *KernelStopCode {
	if s.StopCode == nil {
		return nil
	}
	if s.StopReason&StopReasonKernel == 0 {
		return nil
	}
	stopCode := KernelStopCode(*s.StopCode)
	return &stopCode
}

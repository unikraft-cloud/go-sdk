// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"fmt"
	"math"
	"strings"
	"time"

	"unikraft.com/cloud/sdk/platform/stop"
)

func (instance *Instance) stopReasonValue() stop.StopReason {
	if instance.StopReason == nil {
		return 0
	}
	return stop.StopReason(*instance.StopReason)
}

func (instance *Instance) stopCodeValue() stop.StopCode {
	if instance.StopCode == nil {
		return 0
	}
	return stop.StopCode(*instance.StopCode)
}

// StopCodeErrno returns the application errno, using Linux's errno.h values.
func (instance *Instance) StopCodeErrno() uint8 {
	return uint8(instance.stopCodeValue().Errno())
}

// StopCodeShutdownTable returns whether the stop originated from the inittable
// (0) or from the termtable (1).
func (instance *Instance) StopCodeShutdownTable() uint8 {
	return instance.stopCodeValue().ShutdownTable()
}

// StopCodeInitLevel returns the initlevel at the time of the stop.
func (instance *Instance) StopCodeInitLevel() uint8 {
	return instance.stopCodeValue().InitLevel()
}

// StopCodeReason provides the identity value for the reason for the stop.
func (instance *Instance) StopCodeReason() uint8 {
	return uint8(instance.stopCodeValue().Reason())
}

// DescribeStopOrigin provides a human-readable interpretation of the stop
// reason.
func (instance *Instance) DescribeStopOrigin() string {
	sr := instance.stopReasonValue()
	if sr == 0 {
		return "unknown"
	}

	var ret strings.Builder
	if sr.Forced() {
		ret.WriteString("force ")
	}
	ret.WriteString("initiated by ")
	ret.WriteString(sr.Origin())
	return ret.String()
}

// StopOriginCode provides a human-readable interpretation of the stop reason in
// the form of a short-code.
func (instance *Instance) StopOriginCode() string {
	return instance.stopReasonValue().OriginCode()
}

// DescribeStopReason provides a human-readable description of the stop reason.
func (instance *Instance) DescribeStopReason() string {
	return instance.stopCodeValue().Description()
}

// StopReasonCode returns a human-readable short-code representation of the stop
// reason.
func (instance *Instance) StopReasonCode() string {
	sc := instance.stopCodeValue()
	if sc == 0 {
		return ""
	}

	var ret strings.Builder
	if sc.ShutdownTable() == 0 {
		ret.WriteString("i")
	} else {
		ret.WriteString("t")
	}

	fmt.Fprintf(&ret, "%d", sc.InitLevel())
	ret.WriteString(" ")
	ret.WriteString(sc.Reason().String())

	if sc.Errno() != 0 {
		ret.WriteString(" ")
		ret.WriteString(sc.ErrnoString())
	}

	return ret.String()
}

// DescribeStatus returns a human-readable description of the instance's status.
func (instance *Instance) DescribeStatus() string {
	if instance.State == nil {
		return ""
	}

	switch *instance.State {
	case InstanceStateRunning:
		dur := time.Since(*instance.StartedAt)
		days := int64(dur.Hours() / 24)
		hours := int64(math.Mod(dur.Hours(), 24))
		minutes := int64(math.Mod(dur.Minutes(), 60))
		seconds := int64(math.Mod(dur.Seconds(), 60))

		chunks := []struct {
			singularName string
			amount       int64
		}{
			{"day", days},
			{"hr", hours},
			{"min", minutes},
			{"sec", seconds},
		}

		parts := []string{}

		for i, chunk := range chunks {
			if len(parts) > 0 && i+1 == len(chunks) { // Skip secs if greater than 1m
				continue
			}
			switch chunk.amount {
			case 0:
				continue
			case 1:
				parts = append(parts, fmt.Sprintf("%d%s", chunk.amount, chunk.singularName))
			default:
				parts = append(parts, fmt.Sprintf("%d%ss", chunk.amount, chunk.singularName))
			}
		}

		return fmt.Sprintf("since %s", strings.Join(parts, " "))
	case InstanceStateStopped:
		reason := instance.DescribeStopReason()
		if reason == "shutdown" {
			return ""
		}
		return reason
	default:
		return string(*instance.State)
	}
}

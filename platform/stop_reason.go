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

// Stop returns a structured representation of the stop reason for the
// instance, if it is stopped.
func (instance *Instance) Stop() *stop.Stop {
	if instance.StopReason == nil {
		return nil
	}

	return &stop.Stop{
		StopReason: stop.StopReason(*instance.StopReason),
		StopCode:   instance.StopCode,
	}
}

// DescribeStop provides a human-readable description of the stop reason.
func (instance *Instance) DescribeStop() string {
	stop := instance.Stop()
	if stop == nil {
		return ""
	}
	return stop.String()
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
		return instance.DescribeStop()
	default:
		return string(*instance.State)
	}
}

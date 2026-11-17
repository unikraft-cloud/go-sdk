// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
	"unikraft.com/cloud/sdk/platform/stop"
)

func TestStopInfoNilWhenNoStopReason(t *testing.T) {
	instance := &Instance{State: gen.InstanceStateRunning}

	assert.Nil(t, instance.StopInfo())
	assert.Empty(t, instance.DescribeStop())
}

func TestStopInfoPopulatedFromStopReasonAndCode(t *testing.T) {
	reason := uint32(stop.StopReasonAppExit)
	code := uint32(42)
	instance := &Instance{
		State:      gen.InstanceStateStopped,
		StopReason: &reason,
		StopCode:   &code,
	}

	info := instance.StopInfo()
	require.NotNil(t, info)
	assert.Equal(t, stop.StopReasonAppExit, info.StopReason)
	require.NotNil(t, info.StopCode)
	assert.Equal(t, uint32(42), *info.StopCode)
}

func TestDescribeStopMatchesStopInfoString(t *testing.T) {
	reason := uint32(stop.StopReasonKernelCrash)
	instance := &Instance{
		State:      gen.InstanceStateStopped,
		StopReason: &reason,
	}

	assert.Equal(t, "kernel crash", instance.DescribeStop())
}

func TestDescribeStatusEmptyWhenStateEmpty(t *testing.T) {
	instance := &Instance{}
	assert.Empty(t, instance.DescribeStatus())
}

func TestDescribeStatusRunningShowsUptime(t *testing.T) {
	started := time.Now().Add(-90 * time.Second)
	instance := &Instance{
		State:     gen.InstanceStateRunning,
		StartedAt: &started,
	}

	// Seconds are dropped whenever a larger unit (minutes here) is present.
	assert.Equal(t, "since 1min", instance.DescribeStatus())
}

func TestDescribeStatusStoppedDelegatesToDescribeStop(t *testing.T) {
	reason := uint32(stop.StopReasonAppExit)
	instance := &Instance{
		State:      gen.InstanceStateStopped,
		StopReason: &reason,
	}

	assert.Equal(t, instance.DescribeStop(), instance.DescribeStatus())
	assert.Equal(t, "app exit", instance.DescribeStatus())
}

func TestDescribeStatusOtherStateReturnsRawState(t *testing.T) {
	instance := &Instance{State: gen.InstanceState("draining")}
	assert.Equal(t, "draining", instance.DescribeStatus())
}

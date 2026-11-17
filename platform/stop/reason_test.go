// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStopReasonForced(t *testing.T) {
	assert.False(t, StopReasonUnknown.Forced())
	assert.False(t, StopReasonAppExit.Forced())
	assert.True(t, StopReasonForcedUserShutdown.Forced())
	assert.True(t, (StopReasonForced | StopReasonKernel).Forced())
}

func TestStopReasonOrigin(t *testing.T) {
	tests := []struct {
		name string
		sr   StopReason
		want string
	}{
		{"zero", StopReasonUnknown, "unknown"},
		{"kernel only", StopReasonKernel, "kernel"},
		{"app+kernel", StopReasonAppExit, "app"},
		{"platform+app+kernel", StopReasonPlatformShutdown, "platform"},
		{"user+platform+app+kernel", StopReasonUserShutdownComplete, "user"},
		{"user+platform+incomplete", StopReasonUserShutdownIncomplete, "user"},
		{"forced+user+platform", StopReasonForcedUserShutdown, "user"},
		{"unrecognized bit combo", StopReasonPlatform | StopReasonUser, "user"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sr.Origin())
		})
	}
}

func TestStopReasonOriginCode(t *testing.T) {
	tests := []struct {
		name string
		sr   StopReason
		want string
	}{
		{"none set", StopReasonUnknown, "-----"},
		{"kernel only", StopReasonKernel, "----K"},
		{"app exit", StopReasonAppExit, "---AK"},
		{"platform shutdown", StopReasonPlatformShutdown, "--PAK"},
		{"user shutdown complete", StopReasonUserShutdownComplete, "-UPAK"},
		{"user shutdown incomplete", StopReasonUserShutdownIncomplete, "-UP-K"},
		{"forced user shutdown", StopReasonForcedUserShutdown, "FUP--"},
		{"all bits", StopReasonForced | StopReasonUser | StopReasonPlatform | StopReasonApplication | StopReasonKernel, "FUPAK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sr.OriginCode())
		})
	}
}

func TestStopReasonString(t *testing.T) {
	tests := []struct {
		name string
		sr   StopReason
		want string
	}{
		{"unknown", StopReasonUnknown, "unknown stop"},
		{"kernel crash", StopReasonKernelCrash, "kernel crash"},
		{"app exit", StopReasonAppExit, "app exit"},
		{"platform shutdown", StopReasonPlatformShutdown, "platform shutdown"},
		{"user shutdown incomplete", StopReasonUserShutdownIncomplete, "user shutdown incomplete"},
		{"user shutdown complete", StopReasonUserShutdownComplete, "user shutdown complete"},
		{"forced user shutdown", StopReasonForcedUserShutdown, "forced user shutdown"},
		{"forced unnamed combo", StopReasonForced | StopReasonKernel, "forced kernel stop"},
		{"unnamed unforced combo", StopReasonPlatform | StopReasonUser, "user stop"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sr.String())
		})
	}
}

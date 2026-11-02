// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package timeouts

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"unikraft.com/cloud/sdk/platform"
)

// timedOutResponse builds the response a metro returns when a create call was
// accepted but the wait for the target state ran out
func timedOutResponse(instances ...platform.Instance) *platform.Response[platform.CreateInstanceResponseData] {
	return &platform.Response[platform.CreateInstanceResponseData]{
		Status:  platform.ResponseStatusError,
		Message: "Operation timed out",
		Data:    &platform.CreateInstanceResponseData{Instances: instances},
	}
}

// rejectedResponse builds the response a metro returns when it refused the
// operation outright.
func rejectedResponse(code platform.APIHTTPError, message string) *platform.Response[platform.CreateInstanceResponseData] {
	return &platform.Response[platform.CreateInstanceResponseData]{
		Status:  platform.ResponseStatusError,
		Message: "Failed to perform all operations",
		Data: &platform.CreateInstanceResponseData{
			Instances: []platform.Instance{{
				Uuid:    "4a1f0c9e-0000-0000-0000-000000000001",
				Name:    "demo",
				Status:  new(platform.ResponseStatusError),
				Message: &message,
				Error:   new(int32(code)),
			}},
		},
	}
}

// TestTolerateOnlyClearsTimeouts is the guard against a real failure being
// reported as a success with a warning.
func TestTolerateOnlyClearsTimeouts(t *testing.T) {
	timedOut := timedOutResponse(platform.Instance{
		Uuid:   "4a1f0c9e-0000-0000-0000-000000000001",
		Name:   "demo",
		State:  platform.InstanceStateStarting,
		Status: new(platform.ResponseStatusError),
	})
	rejected := rejectedResponse(platform.APIHTTPErrorNotAllowed,
		"Deletion protection enabled")

	for _, tt := range []struct {
		name string
		resp *platform.Response[platform.CreateInstanceResponseData]
		err  error
		want bool
	}{
		{
			name: "timeout",
			resp: timedOut,
			err:  platform.NewFromResponse(timedOut),
			want: true,
		},
		{
			name: "rejected operation",
			resp: rejected,
			err:  platform.NewFromResponse(rejected),
		},
		{
			name: "no error",
			resp: &platform.Response[platform.CreateInstanceResponseData]{
				Status: platform.ResponseStatusSuccess,
				Data:   &platform.CreateInstanceResponseData{},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			timedOut, err := Tolerate(t.Context(), tt.resp, tt.err, "did not finish in time")
			assert.Equal(t, tt.want, timedOut)
			if tt.want {
				assert.NoError(t, err, "a timeout is cleared so the caller can carry on")
			} else {
				assert.Equal(t, tt.err, err, "every other error is returned untouched")
			}
		})
	}
}

// TestDowngrade pins the field names the compatibility retry reaches by
// reflection.
func TestDowngrade(t *testing.T) {
	create := platform.CreateInstanceRequest{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&create).Elem()))
	assert.Nil(t, create.TimeoutS)
	assert.Equal(t, new(int64(-1)), create.WaitTimeoutMs) //nolint:staticcheck // The deprecated field is what is being set.

	starts := []platform.StartInstancesRequestItem{
		{TimeoutS: new(int64(-1))},
		{TimeoutS: new(int64(-1))},
	}
	assert.True(t, downgrade(reflect.ValueOf(&starts).Elem()))
	for i, start := range starts {
		assert.Nil(t, start.TimeoutS, "item %d", i)
		assert.Equal(t, new(int64(-1)), start.WaitTimeoutMs, "item %d", i) //nolint:staticcheck // The deprecated field is what is being set.
	}

	mixed := []platform.StartInstancesRequestItem{
		{TimeoutS: new(int64(5))},
		{},
	}
	assert.True(t, downgrade(reflect.ValueOf(&mixed).Elem()))
	assert.Equal(t, new(int64(5000)), mixed[0].WaitTimeoutMs, "seconds become milliseconds") //nolint:staticcheck // The deprecated field is what is being set.
	assert.Nil(t, mixed[1].WaitTimeoutMs, "an unset wait stays unset")                       //nolint:staticcheck // The deprecated field is what is being set.

	deletes := []platform.DeleteInstanceRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&deletes).Elem()))
	assert.Nil(t, deletes[0].TimeoutS)

	templates := []platform.CreateTemplateInstancesRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&templates).Elem()))
	assert.Nil(t, templates[0].TimeoutS)

	checkpoints := []platform.CreateCheckpointInstancesRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&checkpoints).Elem()))
	assert.Nil(t, checkpoints[0].TimeoutS)

	wait := platform.WaitInstanceByUUIDRequestBody{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&wait).Elem()))
	assert.Nil(t, wait.TimeoutS)
	assert.Equal(t, new(int64(-1)), wait.TimeoutMs) //nolint:staticcheck // The deprecated field is what is being set.

	del := platform.DeleteInstanceByUUIDRequestBody{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&del).Elem()), "a body with no deprecated field just drops timeout_s")
	assert.Nil(t, del.TimeoutS)

	stops := []platform.StopInstancesRequestItem{{}}
	assert.False(t, downgrade(reflect.ValueOf(&stops).Elem()))
	assert.False(t, downgrade(reflect.ValueOf(new("not a request")).Elem()))
}

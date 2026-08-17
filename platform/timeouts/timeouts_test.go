// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package timeouts

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"unikraft.com/cloud/sdk/platform/gen"
)

// timedOutResponse builds the response a metro returns when a create call was
// accepted but the wait for the target state ran out
func timedOutResponse(instances ...gen.Instance) *gen.Response[gen.CreateInstanceResponseData] {
	return &gen.Response[gen.CreateInstanceResponseData]{
		Status:  gen.ResponseStatusError,
		Message: "Operation timed out",
		Data:    &gen.CreateInstanceResponseData{Instances: instances},
	}
}

// rejectedResponse builds the response a metro returns when it refused the
// operation outright.
func rejectedResponse(code gen.APIHTTPError, message string) *gen.Response[gen.CreateInstanceResponseData] {
	return &gen.Response[gen.CreateInstanceResponseData]{
		Status:  gen.ResponseStatusError,
		Message: "Failed to perform all operations",
		Data: &gen.CreateInstanceResponseData{
			Instances: []gen.Instance{{
				Uuid:    "4a1f0c9e-0000-0000-0000-000000000001",
				Name:    "demo",
				Status:  new(gen.ResponseStatusError),
				Message: &message,
				Error:   new(int32(code)),
			}},
		},
	}
}

// TestTolerateOnlyClearsTimeouts is the guard against a real failure being
// reported as a success with a warning.
func TestTolerateOnlyClearsTimeouts(t *testing.T) {
	timedOut := timedOutResponse(gen.Instance{
		Uuid:   "4a1f0c9e-0000-0000-0000-000000000001",
		Name:   "demo",
		State:  gen.InstanceStateStarting,
		Status: new(gen.ResponseStatusError),
	})
	rejected := rejectedResponse(gen.APIHTTPErrorNotAllowed,
		"Deletion protection enabled")

	for _, tt := range []struct {
		name string
		resp *gen.Response[gen.CreateInstanceResponseData]
		err  error
		want bool
	}{
		{
			name: "timeout",
			resp: timedOut,
			err:  gen.NewFromResponse(timedOut),
			want: true,
		},
		{
			name: "rejected operation",
			resp: rejected,
			err:  gen.NewFromResponse(rejected),
		},
		{
			name: "no error",
			resp: &gen.Response[gen.CreateInstanceResponseData]{
				Status: gen.ResponseStatusSuccess,
				Data:   &gen.CreateInstanceResponseData{},
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
	create := gen.CreateInstanceRequest{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&create).Elem()))
	assert.Nil(t, create.TimeoutS)
	assert.Equal(t, new(int64(-1)), create.WaitTimeoutMs) //nolint:staticcheck // The deprecated field is what is being set.

	starts := []gen.StartInstancesRequestItem{
		{TimeoutS: new(int64(-1))},
		{TimeoutS: new(int64(-1))},
	}
	assert.True(t, downgrade(reflect.ValueOf(&starts).Elem()))
	for i, start := range starts {
		assert.Nil(t, start.TimeoutS, "item %d", i)
		assert.Equal(t, new(int64(-1)), start.WaitTimeoutMs, "item %d", i) //nolint:staticcheck // The deprecated field is what is being set.
	}

	mixed := []gen.StartInstancesRequestItem{
		{TimeoutS: new(int64(5))},
		{},
	}
	assert.True(t, downgrade(reflect.ValueOf(&mixed).Elem()))
	assert.Equal(t, new(int64(5000)), mixed[0].WaitTimeoutMs, "seconds become milliseconds") //nolint:staticcheck // The deprecated field is what is being set.
	assert.Nil(t, mixed[1].WaitTimeoutMs, "an unset wait stays unset")                       //nolint:staticcheck // The deprecated field is what is being set.

	deletes := []gen.DeleteInstanceRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&deletes).Elem()))
	assert.Nil(t, deletes[0].TimeoutS)

	templates := []gen.CreateTemplateInstancesRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&templates).Elem()))
	assert.Nil(t, templates[0].TimeoutS)

	checkpoints := []gen.CreateCheckpointInstancesRequestItem{{TimeoutS: new(int64(-1))}}
	assert.True(t, downgrade(reflect.ValueOf(&checkpoints).Elem()))
	assert.Nil(t, checkpoints[0].TimeoutS)

	wait := gen.WaitInstanceByUUIDRequestBody{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&wait).Elem()))
	assert.Nil(t, wait.TimeoutS)
	assert.Equal(t, new(int64(-1)), wait.TimeoutMs) //nolint:staticcheck // The deprecated field is what is being set.

	del := gen.DeleteInstanceByUUIDRequestBody{TimeoutS: new(int64(-1))}
	assert.True(t, downgrade(reflect.ValueOf(&del).Elem()), "a body with no deprecated field just drops timeout_s")
	assert.Nil(t, del.TimeoutS)

	stops := []gen.StopInstancesRequestItem{{}}
	assert.False(t, downgrade(reflect.ValueOf(&stops).Elem()))
	assert.False(t, downgrade(reflect.ValueOf(new("not a request")).Elem()))
}

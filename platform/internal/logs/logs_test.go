// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package logs

import (
	"context"
	"encoding/base64"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform/gen"
)

// fakeGenClient embeds the (nil) gen.Client interface so it satisfies the
// full interface, and overrides only GetInstanceLogs, which is the sole
// method the logs reader calls.
type fakeGenClient struct {
	gen.Client
	getInstanceLogs func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error)
}

func (f fakeGenClient) GetInstanceLogs(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
	return f.getInstanceLogs(ctx, request, opts)
}

// contentServer answers GetInstanceLogs by slicing full according to the
// requested offset/limit, mimicking the real API's paging behavior closely
// enough for the reader's offset/rewind logic to exercise realistically.
func contentServer(full string) func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
	return func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
		item := request[0]
		start := int64(0)
		if item.Offset != nil {
			start = *item.Offset
		}
		if start < 0 {
			start = 0
		}
		if start > int64(len(full)) {
			start = int64(len(full))
		}
		end := int64(len(full))
		if item.Limit != nil && start+*item.Limit < end {
			end = start + *item.Limit
		}

		data := full[start:end]
		return &gen.Response[gen.GetInstancesLogsResponseData]{
			Status: gen.ResponseStatusSuccess,
			Data: &gen.GetInstancesLogsResponseData{
				Instances: []gen.GetInstancesLogsResponseLoggedInstance{
					{
						Uuid:      "abc",
						Output:    base64.StdEncoding.EncodeToString([]byte(data)),
						Available: gen.GetInstancesLogsResponseAvailable{Start: 0, End: int64(len(full))},
						Range:     gen.GetInstancesLogsResponseRange{Start: start, End: end},
					},
				},
			},
		}, nil
	}
}

func TestInstanceLogsReadsAllAvailableContentWithoutTail(t *testing.T) {
	client := fakeGenClient{getInstanceLogs: contentServer("hello world")}

	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, nil, false)
	require.NoError(t, err)

	got, err := io.ReadAll(r)
	require.NoError(t, err)
	assert.Equal(t, "hello world", string(got))
}

func TestInstanceLogsTailRewindsToLastLines(t *testing.T) {
	client := fakeGenClient{getInstanceLogs: contentServer("line1\nline2\nline3")}

	tail := 1
	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, &tail, false)
	require.NoError(t, err)

	got, err := io.ReadAll(r)
	require.NoError(t, err)
	// tail counts trailing newlines, so tail=1 rewinds to the start of the
	// second-to-last newline-terminated segment.
	assert.Equal(t, "line2\nline3", string(got))
}

func TestInstanceLogsTailZeroReadsFromLastNewline(t *testing.T) {
	client := fakeGenClient{getInstanceLogs: contentServer("line1\nline2\nline3")}

	tail := 0
	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, &tail, false)
	require.NoError(t, err)

	got, err := io.ReadAll(r)
	require.NoError(t, err)
	assert.Equal(t, "line3", string(got))
}

func TestInstanceLogsTailLargerThanContentReadsFromStart(t *testing.T) {
	client := fakeGenClient{getInstanceLogs: contentServer("line1\nline2\nline3")}

	tail := 100
	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, &tail, false)
	require.NoError(t, err)

	got, err := io.ReadAll(r)
	require.NoError(t, err)
	assert.Equal(t, "line1\nline2\nline3", string(got))
}

func TestInstanceLogsPropagatesNoDataError(t *testing.T) {
	client := fakeGenClient{
		getInstanceLogs: func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
			return &gen.Response[gen.GetInstancesLogsResponseData]{
				Status: gen.ResponseStatusSuccess,
				Data:   &gen.GetInstancesLogsResponseData{Instances: nil},
			}, nil
		},
	}

	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, nil, false)
	require.NoError(t, err) // Reader() itself doesn't fetch when tail is nil.

	_, err = io.ReadAll(r)
	assert.ErrorContains(t, err, "no data returned")
}

func TestInstanceLogsDetectsShortReadMismatch(t *testing.T) {
	client := fakeGenClient{
		getInstanceLogs: func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
			return &gen.Response[gen.GetInstancesLogsResponseData]{
				Status: gen.ResponseStatusSuccess,
				Data: &gen.GetInstancesLogsResponseData{
					Instances: []gen.GetInstancesLogsResponseLoggedInstance{
						{
							Uuid:      "abc",
							Output:    base64.StdEncoding.EncodeToString([]byte("hello")),
							Available: gen.GetInstancesLogsResponseAvailable{Start: 0, End: 5},
							// Range claims more bytes than were actually encoded.
							Range: gen.GetInstancesLogsResponseRange{Start: 0, End: 999},
						},
					},
				},
			}, nil
		},
	}

	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: new(string)}, nil, false)
	require.NoError(t, err)

	_, err = io.ReadAll(r)
	assert.ErrorContains(t, err, "expected to read")
}

func TestInstanceLogsFollowBlocksThenRespectsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := fakeGenClient{getInstanceLogs: contentServer("hi")}
	r, err := InstanceLogs(ctx, client).Reader(gen.NameOrUUID{Uuid: new(string)}, nil, true)
	require.NoError(t, err)

	errCh := make(chan error, 1)
	go func() {
		_, err := io.ReadAll(r)
		errCh <- err
	}()

	// Give the reader time to drain the available "hi" and settle into the
	// backoff-and-retry loop waiting for more (since Follow is set).
	time.Sleep(50 * time.Millisecond)
	cancel()

	select {
	case err := <-errCh:
		assert.ErrorIs(t, err, context.Canceled)
	case <-time.After(3 * time.Second):
		t.Fatal("cancelling the context did not unblock the follow read")
	}
}

func TestInstanceLogsSendsUuidAndNameFromRef(t *testing.T) {
	var gotUUID, gotName *string
	client := fakeGenClient{
		getInstanceLogs: func(ctx context.Context, request []gen.GetInstancesLogsRequestItem, opts gen.GetInstanceLogsOpts) (*gen.Response[gen.GetInstancesLogsResponseData], error) {
			gotUUID = request[0].Uuid
			gotName = request[0].Name
			return &gen.Response[gen.GetInstancesLogsResponseData]{
				Status: gen.ResponseStatusSuccess,
				Data: &gen.GetInstancesLogsResponseData{
					Instances: []gen.GetInstancesLogsResponseLoggedInstance{{
						Output:    base64.StdEncoding.EncodeToString(nil),
						Available: gen.GetInstancesLogsResponseAvailable{Start: 0, End: 0},
						Range:     gen.GetInstancesLogsResponseRange{Start: 0, End: 0},
					}},
				},
			}, nil
		},
	}

	uuid := "the-uuid"
	r, err := InstanceLogs(context.Background(), client).Reader(gen.NameOrUUID{Uuid: &uuid}, nil, false)
	require.NoError(t, err)
	_, _ = io.ReadAll(r)

	require.NotNil(t, gotUUID)
	assert.Equal(t, "the-uuid", *gotUUID)
	assert.Nil(t, gotName)
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"reflect"
	"slices"
	"time"

	"unikraft.com/cloud/sdk/platform/gen"
	"unikraft.com/cloud/sdk/platform/internal/logs"
)

// Instance is a live, resource-oriented view of an instance. It embeds
// gen.Instance so its fields are accessed directly (inst.MemoryMb,
// inst.State, ...), and adds behavior (Refresh, Start, Stop, Suspend,
// Delete, Update, Logs) on top of the flat generated client.
type Instance struct {
	gen.Instance

	client gen.Client
	orig   gen.Instance // snapshot as of the last fetch/refresh, for diffing in Update
}

// UUID returns the UUID of the instance.
func (i *Instance) UUID() string { return i.Uuid }

// cloneInstance deep-clones a gen.Instance via a JSON round-trip, so that
// slice/map/pointer fields aren't aliased between an Instance and its
// snapshot. in was itself just unmarshaled from the API response, so a
// marshal/unmarshal round-trip of it failing indicates a broken invariant,
// not a normal error condition.
func cloneInstance(in gen.Instance) gen.Instance {
	b, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("platform: cloning instance: %v", err))
	}
	var out gen.Instance
	if err := json.Unmarshal(b, &out); err != nil {
		panic(fmt.Sprintf("platform: cloning instance: %v", err))
	}
	return out
}

func wrapInstance(client gen.Client, data gen.Instance) *Instance {
	return &Instance{
		Instance: cloneInstance(data),
		client:   client,
		orig:     cloneInstance(data),
	}
}

// Refresh re-fetches the instance from the API and resets the diffing
// snapshot used by Update.
func (i *Instance) Refresh(ctx context.Context) error {
	resp, err := i.client.GetInstances(ctx, []gen.NameOrUUID{{Uuid: &i.Uuid}}, gen.GetInstancesOpts{Details: new(true)})
	if err != nil {
		return err
	}
	if len(resp.Data.Instances) == 0 {
		return fmt.Errorf("instance %s not found", i.Uuid)
	}
	fresh := resp.Data.Instances[0]
	i.Instance, i.orig = cloneInstance(fresh), cloneInstance(fresh)
	return nil
}

// StartOption configures a call to Instance.Start.
type StartOption func(*gen.StartInstanceByUUIDRequestBody)

// WithStartTimeout blocks Start until the instance reaches the running
// state, or until d elapses.
func WithStartTimeout(d time.Duration) StartOption {
	return func(b *gen.StartInstanceByUUIDRequestBody) {
		b.TimeoutS = new(int64(d / time.Second))
	}
}

// Start starts the instance, then refreshes it to pick up the resulting state.
func (i *Instance) Start(ctx context.Context, opts ...StartOption) error {
	var body gen.StartInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	if _, err := i.client.StartInstanceByUUID(ctx, i.Uuid, body); err != nil {
		return err
	}
	return i.Refresh(ctx)
}

// StopOption configures a call to Instance.Stop.
type StopOption func(*gen.StopInstanceByUUIDRequestBody)

// WithStopForce immediately force-stops the instance instead of draining connections.
func WithStopForce() StopOption {
	return func(b *gen.StopInstanceByUUIDRequestBody) { b.Force = new(true) }
}

// WithStopQuick performs a quick shutdown. Overridden by WithStopForce.
func WithStopQuick() StopOption {
	return func(b *gen.StopInstanceByUUIDRequestBody) { b.Quick = new(true) }
}

// WithStopDrainTimeout drains connections for up to d before stopping the
// instance. Ignored if WithStopForce is set.
func WithStopDrainTimeout(d time.Duration) StopOption {
	return func(b *gen.StopInstanceByUUIDRequestBody) {
		b.DrainTimeoutMs = new(uint64(d / time.Millisecond))
	}
}

// WithStopIfState only stops the instance if it is currently in state.
func WithStopIfState(state gen.InstanceState) StopOption {
	return func(b *gen.StopInstanceByUUIDRequestBody) { b.Ifstate = new(string(state)) }
}

// WithStopCoredump forces the VMM to shut down immediately and generate a
// coredump. Only takes effect combined with WithStopForce.
func WithStopCoredump() StopOption {
	return func(b *gen.StopInstanceByUUIDRequestBody) { b.Dump = new(true) }
}

// Stop stops the instance, then refreshes it to pick up the resulting state.
//
// For the informational "why did this instance stop" accessor, see StopInfo.
func (i *Instance) Stop(ctx context.Context, opts ...StopOption) error {
	var body gen.StopInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	if _, err := i.client.StopInstanceByUUID(ctx, i.Uuid, body); err != nil {
		return err
	}
	return i.Refresh(ctx)
}

// SuspendOption configures a call to Instance.Suspend.
type SuspendOption func(*gen.SuspendInstanceByUUIDRequestBody)

// WithSuspendDrainTimeout drains connections for up to d before suspending
// the instance.
func WithSuspendDrainTimeout(d time.Duration) SuspendOption {
	return func(b *gen.SuspendInstanceByUUIDRequestBody) {
		b.DrainTimeoutMs = new(uint64(d / time.Millisecond))
	}
}

// Suspend suspends the instance, then refreshes it to pick up the resulting state.
func (i *Instance) Suspend(ctx context.Context, opts ...SuspendOption) error {
	var body gen.SuspendInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	if _, err := i.client.SuspendInstanceByUUID(ctx, i.Uuid, body); err != nil {
		return err
	}
	return i.Refresh(ctx)
}

// DeleteOption configures a call to Instance.Delete.
type DeleteOption func(*deleteConfig)

type deleteConfig struct {
	body      gen.DeleteInstanceByUUIDRequestBody
	missingOK bool
	retryBusy time.Duration
}

// WithDeleteTimeout blocks Delete until the instance is deleted, or until d elapses.
func WithDeleteTimeout(d time.Duration) DeleteOption {
	return func(c *deleteConfig) {
		c.body.TimeoutS = new(int64(d / time.Second))
	}
}

// WithoutRetention deletes the instance immediately, without retention, even
// if it's currently being retained.
func WithoutRetention() DeleteOption {
	return func(c *deleteConfig) { c.body.DontRetain = new(true) }
}

// WithDeleteMissingOK makes Delete return nil when the instance does not
// exist.
func WithDeleteMissingOK() DeleteOption {
	return func(c *deleteConfig) { c.missingOK = true }
}

// WithDeleteRetryBusy retries Delete with backoff for up to d while the API
// reports the instance as busy.
func WithDeleteRetryBusy(d time.Duration) DeleteOption {
	return func(c *deleteConfig) { c.retryBusy = d }
}

// Delete deletes the instance. Unlike the other lifecycle methods, it does
// not refresh afterwards: the instance is gone, and re-fetching it would
// just 404.
func (i *Instance) Delete(ctx context.Context, opts ...DeleteOption) error {
	var cfg deleteConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	err := whileBusy(ctx, cfg.retryBusy, func() error {
		_, err := i.client.DeleteInstanceByUUID(ctx, i.Uuid, cfg.body)
		return err
	})
	if cfg.missingOK && gen.ErrorContains(err, gen.APIHTTPErrorNotFound) {
		return nil
	}
	return err
}

// WaitOption configures a call to Instance.Wait.
type WaitOption func(*gen.WaitInstanceByUUIDRequestBody)

// WithWaitState waits for state instead of the running state.
func WithWaitState(state gen.InstanceState) WaitOption {
	return func(b *gen.WaitInstanceByUUIDRequestBody) { b.State = &state }
}

// WithWaitTimeout stops waiting after d, rounded up to a full second. A
// negative d waits for as long as the platform allows.
func WithWaitTimeout(d time.Duration) WaitOption {
	return func(b *gen.WaitInstanceByUUIDRequestBody) {
		if d < 0 {
			b.TimeoutS = new(int64(-1))
			return
		}
		b.TimeoutS = new(int64((d + time.Second - 1) / time.Second))
	}
}

// Wait blocks until the instance reaches a state, then refreshes it.
func (i *Instance) Wait(ctx context.Context, opts ...WaitOption) error {
	var body gen.WaitInstanceByUUIDRequestBody
	for _, opt := range opts {
		opt(&body)
	}
	if _, err := i.client.WaitInstanceByUUID(ctx, i.Uuid, body); err != nil {
		return err
	}
	return i.Refresh(ctx)
}

// Update diffs the instance's current field values against the snapshot
// taken at the last fetch/refresh, sends the minimal patch for whatever
// changed, and always refreshes afterwards: UpdateInstances's response is
// ack/status-only and some property changes apply asynchronously
// (gen.Instance.Updates), so only a real Refresh reflects the new state
// correctly.
//
// Only 13 of the 15 MutableInstanceProperty values are diffed here:
// Hostname and Dependencies are write-only from the SDK's read side (no
// field on gen.Instance to snapshot or compare), so use Patch to change
// them. Image is
// diffed as a URL-only change; credentials/pull-policy for the image aren't
// tracked on Instance and can't be expressed here either.
func (i *Instance) Update(ctx context.Context) error {
	items := i.diff()
	if len(items) == 0 {
		return nil
	}
	_, err := i.client.UpdateInstances(ctx, items)
	if rerr := i.Refresh(ctx); rerr != nil {
		err = errors.Join(err, rerr)
	}
	return err
}

func (i *Instance) diff() []gen.UpdateInstancesRequestItem {
	var items []gen.UpdateInstancesRequestItem
	set := func(prop gen.MutableInstanceProperty, value any) {
		items = append(items, gen.UpdateInstancesRequestItem{
			Uuid: &i.Uuid, Prop: prop, Op: gen.MutableInstanceOperationSet, Value: &value,
		})
	}

	if i.MemoryMb != i.orig.MemoryMb {
		set(gen.MutableInstancePropertyMemoryMb, i.MemoryMb)
	}
	if i.Vcpus != i.orig.Vcpus {
		set(gen.MutableInstancePropertyVcpus, i.Vcpus)
	}
	if i.Image != i.orig.Image {
		set(gen.MutableInstancePropertyImage, gen.ImageSpec{Url: i.Image})
	}

	if !slices.Equal(i.Args, i.orig.Args) {
		set(gen.MutableInstancePropertyArgs, i.Args)
	}
	if !maps.Equal(i.Env, i.orig.Env) {
		set(gen.MutableInstancePropertyEnv, i.Env)
	}
	if !slices.Equal(i.Tags, i.orig.Tags) {
		set(gen.MutableInstancePropertyTags, i.Tags)
	}

	if !reflect.DeepEqual(i.ScaleToZero, i.orig.ScaleToZero) {
		set(gen.MutableInstancePropertyScaleToZero, i.ScaleToZero)
	}
	if !reflect.DeepEqual(i.Autokill, i.orig.Autokill) {
		set(gen.MutableInstancePropertyAutokill, i.Autokill)
	}
	if !reflect.DeepEqual(i.DeleteLock, i.orig.DeleteLock) {
		set(gen.MutableInstancePropertyDeleteLock, i.DeleteLock)
	}
	if !reflect.DeepEqual(i.SchedPriority, i.orig.SchedPriority) {
		set(gen.MutableInstancePropertySchedPriority, i.SchedPriority)
	}
	if !reflect.DeepEqual(i.Schedules, i.orig.Schedules) {
		set(gen.MutableInstancePropertySchedules, i.Schedules)
	}
	if !reflect.DeepEqual(i.Roms, i.orig.Roms) {
		set(gen.MutableInstancePropertyRoms, i.Roms)
	}
	if !reflect.DeepEqual(i.Plugins, i.orig.Plugins) {
		set(gen.MutableInstancePropertyPlugins, i.Plugins)
	}

	// Hostname, Dependencies intentionally not diffed: gen.Instance has no
	// field to snapshot/compare them against.

	return items
}

// LogsOption configures a call to Instance.Logs.
type LogsOption func(*logsConfig)

type logsConfig struct {
	tail   *int
	follow bool
}

// WithLogsTail starts the read n lines back from the end.
func WithLogsTail(n int) LogsOption {
	return func(c *logsConfig) { c.tail = &n }
}

// WithLogsFollow keeps the reader open, blocking for new output as it arrives.
func WithLogsFollow() LogsOption {
	return func(c *logsConfig) { c.follow = true }
}

// Logs returns a reader over the instance's log output. The returned
// io.ReadCloser's Close cancels the underlying request, which is what
// unblocks a WithLogsFollow read.
func (i *Instance) Logs(ctx context.Context, opts ...LogsOption) (io.ReadCloser, error) {
	var cfg logsConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	logCtx, cancel := context.WithCancel(ctx)
	r, err := logs.InstanceLogs(logCtx, i.client).Reader(gen.NameOrUUID{Uuid: &i.Uuid}, cfg.tail, cfg.follow)
	if err != nil {
		cancel()
		return nil, err
	}
	return &logsReadCloser{Reader: r, cancel: cancel}, nil
}

// Metrics returns the instance's current resource-usage and networking
// metrics.
func (i *Instance) Metrics(ctx context.Context) (*gen.GetInstancesMetricsResponseInstanceMetrics, error) {
	resp, err := i.client.GetInstanceMetricsByUUID(ctx, i.Uuid)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Instances) == 0 {
		return nil, fmt.Errorf("instance %s not found", i.Uuid)
	}
	return &resp.Data.Instances[0], nil
}

// History returns the instance's checkpoint history, oldest first.
func (i *Instance) History(ctx context.Context) ([]gen.CheckpointHistoryEntry, error) {
	resp, err := i.client.GetInstanceHistoryByUUID(ctx, i.Uuid)
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Instances) == 0 {
		return nil, fmt.Errorf("instance %s not found", i.Uuid)
	}
	return resp.Data.Instances[0].History, nil
}

// LogsPage reads up to limit bytes of the instance's log, starting at
// offset. A negative offset counts back from the end of the log.
func (i *Instance) LogsPage(ctx context.Context, offset, limit int64) (*gen.GetInstancesLogsResponseLoggedInstance, error) {
	resp, err := i.client.GetInstanceLogsByUUID(ctx, i.Uuid, gen.GetInstanceLogsByUUIDRequestBody{
		Offset: &offset,
		Limit:  &limit,
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Instances) == 0 {
		return nil, fmt.Errorf("instance %s not found", i.Uuid)
	}
	return &resp.Data.Instances[0], nil
}

type logsReadCloser struct {
	io.Reader
	cancel context.CancelFunc
}

func (rc *logsReadCloser) Close() error {
	rc.cancel()
	return nil
}

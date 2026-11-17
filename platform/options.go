// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"time"

	"unikraft.com/cloud/sdk/platform/gen"
)

// CreateInstanceOption configures a call to [Client.CreateInstance].
type CreateInstanceOption func(*gen.CreateInstanceRequest)

// WithCreateName sets the instance's name. If not set, a random name is generated.
func WithCreateName(name string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Name = new(name) }
}

// WithCreateImage sets the image to use for the instance, as a
// gen.ImageReference or a gen.ImageSpec. Mutually exclusive
// with WithCreateTemplate, WithCreateBranchFrom, and WithCreateCheckpoint.
func WithCreateImage(image gen.ImageSource) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Image = image }
}

// WithCreateArgs sets the arguments to pass to the instance when it starts.
func WithCreateArgs(args ...string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Args = args }
}

// WithCreateEnv sets environment variables for the instance.
func WithCreateEnv(env map[string]string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Env = env }
}

// WithCreateMemoryMB sets how much memory, in MB, to allocate for the
// instance. Defaults to 128.
func WithCreateMemoryMB(mb int64) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.MemoryMb = new(mb) }
}

// WithCreateServiceGroup configures the (persistent or ephemeral) service
// group for the instance. Not used by template instances.
func WithCreateServiceGroup(sg gen.CreateInstanceRequestServiceGroup) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.ServiceGroup = &sg }
}

// WithCreateVolumes attaches existing or new volumes to the instance.
func WithCreateVolumes(volumes ...gen.CreateInstanceRequestVolume) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Volumes = volumes }
}

// WithCreateAutostart starts the instance automatically on creation. Must be
// set when WithCreateTimeout is used.
func WithCreateAutostart() CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Autostart = new(true) }
}

// WithCreateReplicas requests n additional replicas, for a total of n+1
// created instances. Defaults to 0.
func WithCreateReplicas(n int64) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Replicas = new(n) }
}

// WithCreateRestartPolicy sets how the instance behaves when it stops or
// crashes. Cannot be combined with the delete-on-stop feature.
func WithCreateRestartPolicy(policy gen.InstanceRestartPolicy) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.RestartPolicy = &policy }
}

// WithCreateScaleToZero configures scale-to-zero for the instance. Requires
// WithCreateServiceGroup; cannot combine with the delete-on-stop feature.
func WithCreateScaleToZero(cfg gen.CreateInstanceScaleToZero) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.ScaleToZero = &cfg }
}

// WithCreateVCPUs sets the number of vCPUs to allocate for the instance.
// Defaults to 1.
func WithCreateVCPUs(n int32) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Vcpus = new(n) }
}

// WithCreateFeatures enables features for the instance. scale-to-zero and
// delete-on-stop are mutually exclusive.
func WithCreateFeatures(features ...gen.InstanceFeature) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Features = features }
}

// WithCreateTimeout blocks instance creation until all new instances reach
// the running state, or until d elapses. Requires WithCreateAutostart.
func WithCreateTimeout(d time.Duration) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.TimeoutS = new(int64(d / time.Second)) }
}

// WithCreateROMs attaches Read-Only Memory (ROM) blobs to the instance.
func WithCreateROMs(roms ...gen.CreateInstanceRequestRom) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Roms = roms }
}

// WithCreatePlugins attaches plugins to the instance. At most 8 may be attached.
func WithCreatePlugins(plugins ...gen.CreateInstanceRequestPlugin) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Plugins = plugins }
}

// WithCreateTags associates tags with the instance.
func WithCreateTags(tags ...string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Tags = tags }
}

// WithCreateTemplate creates the instance from a template. Mutually
// exclusive with WithCreateImage, WithCreateBranchFrom, and WithCreateCheckpoint.
func WithCreateTemplate(tmpl gen.CreateInstanceRequestTemplate) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Template = &tmpl }
}

// WithCreateSchedPriority sets the scheduling priority for the instance.
// Only settable by users with scheduling priority override permissions.
func WithCreateSchedPriority(p gen.SchedPriority) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.SchedPriority = &p }
}

// WithCreateSchedules sets calendar-based schedules for the instance, to
// automatically start, stop, delete, or exec a command in it.
func WithCreateSchedules(schedules ...gen.Schedule) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Schedules = schedules }
}

// WithCreateAutokill configures automatic delete-on-idle/request-limit
// behavior for the instance. Not used for template instances.
func WithCreateAutokill(cfg gen.CreateInstanceRequestAutokill) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Autokill = &cfg }
}

// WithCreateHostname sets the instance's hostname, used for internal DNS
// resolution. If not set, the hostname defaults to the instance name.
func WithCreateHostname(hostname string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Hostname = new(hostname) }
}

// WithCreateDependencies sets instances that must be running before this
// instance starts.
func WithCreateDependencies(deps ...gen.NameOrUUID) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Dependencies = deps }
}

// WithCreateBranchFrom creates the instance as a branch of an existing
// instance. Mutually exclusive with WithCreateImage and WithCreateTemplate.
func WithCreateBranchFrom(ref gen.NameOrUUID) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.BranchFrom = &ref }
}

// WithCreateCheckpoint creates the instance from an existing checkpoint.
// Mutually exclusive with WithCreateImage, WithCreateTemplate, and WithCreateBranchFrom.
func WithCreateCheckpoint(ref gen.NameOrUUID) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Checkpoint = &ref }
}

// WithCreateGateway sets the default gateway to configure inside the guest.
func WithCreateGateway(gateway string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Gateway = new(gateway) }
}

// WithCreateNameserver sets the DNS resolver to configure inside the guest.
func WithCreateNameserver(nameserver string) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Nameserver = new(nameserver) }
}

// WithCreateNetworkInterfaces attaches one to four network interfaces to
// the instance.
func WithCreateNetworkInterfaces(ifaces ...gen.CreateInstanceRequestNetworkInterface) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.NetworkInterfaces = ifaces }
}

// WithCreateType sets the VM type. Defaults to micro (Firecracker); full
// (QEMU) is required for GPU passthrough but drops scale-to-zero/template/branch/checkpoint support.
func WithCreateType(t gen.InstanceType) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Type = &t }
}

// WithCreateGPUs sets the number of GPUs to attach to the instance.
// Currently restricted to at most 1, and requires WithCreateType(gen.InstanceTypeFull).
func WithCreateGPUs(n int32) CreateInstanceOption {
	return func(r *gen.CreateInstanceRequest) { r.Gpus = new(n) }
}

// PaginationOrder controls the sort order used by [Client.Instances].
type PaginationOrder string

const (
	PaginationOrderAsc  PaginationOrder = "asc"
	PaginationOrderDesc PaginationOrder = "desc"
)

// InstancesOption configures a call to [Client.Instances].
type InstancesOption func(*gen.GetInstancesOpts)

// WithDetails includes full details for each returned instance.
func WithDetails() InstancesOption {
	return func(o *gen.GetInstancesOpts) { o.Details = new(true) }
}

// WithLimit caps the number of instances returned.
func WithLimit(n uint32) InstancesOption {
	return func(o *gen.GetInstancesOpts) { o.Count = new(n) }
}

// WithPage continues listing from the pagination cursor returned by a
// previous call.
func WithPage(from string) InstancesOption {
	return func(o *gen.GetInstancesOpts) { o.From = new(from) }
}

// WithTags filters instances to those matching all of the given tags.
func WithTags(tags ...string) InstancesOption {
	return func(o *gen.GetInstancesOpts) { o.Tags = tags }
}

// WithOrder sets the sort order for the returned instances.
func WithOrder(order PaginationOrder) InstancesOption {
	return func(o *gen.GetInstancesOpts) { o.Order = new(gen.PaginationOrder(order)) }
}

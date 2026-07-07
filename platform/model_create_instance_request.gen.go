// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Features to enable for the instance.  Features are specific
// configurations or capabilities that can be enabled for the
// instance.  The `scale-to-zero` and `delete-on-stop` features are
// mutually exclusive.
type CreateInstanceRequestFeatures string

const (
	CreateInstanceRequestFeaturesDeleteOnStop CreateInstanceRequestFeatures = "delete-on-stop"
)

// The request message for creating a new instance.
type CreateInstanceRequest struct {
	// (Optional).  The name of the instance.
	//
	// If not provided, a random name will be generated.  The name must be unique.
	Name *string `json:"name,omitempty"`
	// (Optional).  The image to use for the instance.
	//
	// Either an image or a template must be specified.
	Image *ImageSpec `json:"image,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// (Optional).  The arguments to pass to the instance when it starts.
	Args []string `json:"args,omitempty"`
	// (Optional).  Environment variables to set for the instance.
	Env map[string]string `json:"env,omitempty"`
	// (Optional).  Memory in MB to allocate for the instance.  Default is 128.
	MemoryMb *int64 `json:"memory_mb,omitempty"`
	// (Optional).  The service group configuration when creating an instance.
	//
	// When creating an instance, either a previously created (persistent) service
	// group can be referenced (either through its name or UUID), or a new
	// (ephemeral) service group can be created for the instance by specifying the
	// list of services it should expose and optionally the domains it should use.
	// Not used by template instances.
	ServiceGroup *CreateInstanceRequestServiceGroup `json:"service_group,omitempty"`
	// Volumes to attach to the instance.
	//
	// This list can contain both existing and new volumes to create as part of
	// the instance creation.  Existing volumes can be referenced by their name or
	// UUID.  New volumes can be created by specifying a name, size in MiB, and
	// mount point in the instance.  The mount point is the directory in the
	// instance where the volume will be mounted.
	Volumes []CreateInstanceRequestVolume `json:"volumes,omitempty"`
	// Whether the instance should start automatically on creation.
	// Must be set to true when `timeout_s` is specified.
	Autostart *bool `json:"autostart,omitempty"`
	// (Optional).  Number of additional replicas to create.  The total
	// number of instances created is `replicas + 1`.  Defaults to 0.
	Replicas *int64 `json:"replicas,omitempty"`
	// Restart policy for the instance.  This defines how the instance
	// should behave when it stops or crashes.  Cannot be combined with
	// the `delete-on-stop` feature.
	RestartPolicy *InstanceRestartPolicy `json:"restart_policy,omitempty"`
	// Scale-to-zero configuration for the instance.  Requires
	// `service_group` to be set.  Cannot be combined with the
	// `delete-on-stop` feature.
	ScaleToZero *CreateInstanceScaleToZero `json:"scale_to_zero,omitempty"`
	// (Optional).  Number of vCPUs to allocate for the instance.
	// Defaults to 1.
	Vcpus *int32 `json:"vcpus,omitempty"`
	// Deprecated: Use `timeout_s` instead.  Timeout in milliseconds to
	// wait for all new instances to reach running state.  Requires
	// `autostart` to be set.  If `timeout_s` is not set, this value is
	// converted by rounding up to the next full second.  No wait
	// performed for a value of 0.
	WaitTimeoutMs *int64 `json:"wait_timeout_ms,omitempty"`
	// Features to enable for the instance.  Features are specific
	// configurations or capabilities that can be enabled for the
	// instance.  The `scale-to-zero` and `delete-on-stop` features are
	// mutually exclusive.
	Features []InstanceFeature `json:"features,omitempty"`
	// Timeout in seconds to wait for all new instances to reach running
	// state.  Requires `autostart` to be set.  If you autostart your
	// new instance, you can wait for it to finish starting with a
	// blocking API call if you specify a wait timeout greater than
	// zero.  No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitempty"`
	// Read-Only Memory (ROM) blobs to attach to the instance.
	// Unikraft Cloud supports the ability to attach Read-Only Memory (ROM) blobs
	// to instances. It allows you to create a general-purpose base image and
	// then customize individual instances by attaching code or data as separate
	// ROM blobs.
	Roms []CreateInstanceRequestRom `json:"roms,omitempty"`
	// (Optional).  Tags to associate with the instance.
	Tags []string `json:"tags,omitempty"`
	// Template instances.
	// An existing instance can be saved as a template. This template is then
	// used to create new instances that inherit the exact configuration and
	// state the original instance had when the template was created.
	Template *CreateInstanceRequestTemplate `json:"template,omitempty"`
	// The scheduling priority for the instance. Only settable by
	// users with scheduling priority override permissions.
	SchedPriority *SchedPriority `json:"sched_priority,omitempty"`
	// (Optional).  Schedules for the instance.  Scheduled operations let you
	// automatically start, stop, delete, or exec a command in the instance on
	// a calendar-based schedule.  For `exec` schedules, set the `args` field
	// to the command and its arguments.  Each instance stores its own
	// schedules, and cloning preserves them.
	Schedules []Schedule `json:"schedules,omitempty"`
	// (Optional).  Automatic delete-on-idle/request-limit configuration.
	// Not used for template instances.
	Autokill *CreateInstanceRequestAutokill `json:"autokill,omitempty"`
	// (Optional).  The hostname of the instance.
	//
	// If not provided, the hostname will be set to the instance name.  The
	// hostname must be a valid DNS label (e.g., "my-instance") and is used for
	// internal DNS resolution within the Unikraft Cloud network.
	Hostname *string `json:"hostname,omitempty"`
	// (Optional).  Dependencies of the instance.
	//
	// A list of instance identifiers (name or UUID) that this instance depends
	// on.  Dependencies define startup ordering and can be used to ensure that
	// prerequisite instances are running before this instance starts.
	Dependencies []NameOrUUID `json:"dependencies,omitempty"`
	// (Optional).  Reference to an existing instance to branch from.
	// The instance can be running, stopped, or a template.  If the source
	// instance is running, a snapshot will be taken asynchronously and the
	// new instance will wait for it to complete before starting.
	// Mutually exclusive with `image` and `template`.
	BranchFrom *NameOrUUID `json:"branch_from,omitempty"`
	// (Optional).  Reference to an existing checkpoint to create the instance
	// from.  The checkpoint must be in the `checkpoint` state.  The new instance
	// will be created with the same configuration and state as the checkpoint.
	// Mutually exclusive with `image`, `template`, and `branch_from`.
	Checkpoint *NameOrUUID `json:"checkpoint,omitempty"`
	// The default gateway to configure inside the guest.
	Gateway *string `json:"gateway,omitempty"`
	// The DNS resolver to configure inside the guest.
	Nameserver *string `json:"nameserver,omitempty"`
	// A list of one to four interfaces to attach
	NetworkInterfaces []CreateInstanceRequestNetworkInterface `json:"network_interfaces,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CreateInstanceRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequest
	return json.Marshal((Alias)(m))
}

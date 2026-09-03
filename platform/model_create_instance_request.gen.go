// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// The request message for creating a new instance.
type CreateInstanceRequest struct {
	// (Optional).  The name of the instance.
	//
	// If not provided, a random name will be generated.  The name must be unique.
	Name *string `json:"name,omitzero"`
	// (Optional).  The image to use for the instance.
	//
	// Either an image or a template must be specified.  Accepts either a plain
	// image reference string (`"nginx:latest"`) or an object carrying additional
	// pull configuration (`{"url": "nginx:latest", "pull_policy": "always"}`).
	Image ImageSource `json:"image,omitzero"`
	// (Optional).  The arguments to pass to the instance when it starts.
	Args []string `json:"args,omitzero"`
	// (Optional).  Environment variables to set for the instance.
	Env map[string]string `json:"env,omitzero"`
	// (Optional).  Memory in MB to allocate for the instance.  Default is 128.
	MemoryMb *int64 `json:"memory_mb,omitzero"`
	// (Optional).  The service group configuration when creating an instance.
	//
	// When creating an instance, either a previously created (persistent) service
	// group can be referenced (either through its name or UUID), or a new
	// (ephemeral) service group can be created for the instance by specifying the
	// list of services it should expose and optionally the domains it should use.
	// Not used by template instances.
	ServiceGroup *CreateInstanceRequestServiceGroup `json:"service_group,omitzero"`
	// Volumes to attach to the instance.
	//
	// This list can contain both existing and new volumes to create as part of
	// the instance creation.  Existing volumes can be referenced by their name or
	// UUID.  New volumes can be created by specifying a name, size in MiB, and
	// mount point in the instance.  The mount point is the directory in the
	// instance where the volume will be mounted.
	Volumes []CreateInstanceRequestVolume `json:"volumes,omitzero"`
	// Whether the instance should start automatically on creation.
	// Must be set to true when `timeout_s` is specified.
	Autostart *bool `json:"autostart,omitzero"`
	// (Optional).  Number of additional replicas to create.  The total
	// number of instances created is `replicas + 1`.  Defaults to 0.
	Replicas *int64 `json:"replicas,omitzero"`
	// Restart policy for the instance.  This defines how the instance
	// should behave when it stops or crashes.  Cannot be combined with
	// the `delete-on-stop` feature.
	RestartPolicy *InstanceRestartPolicy `json:"restart_policy,omitzero"`
	// Scale-to-zero configuration for the instance.  Requires
	// `service_group` to be set.  Cannot be combined with the
	// `delete-on-stop` feature.
	ScaleToZero *CreateInstanceScaleToZero `json:"scale_to_zero,omitzero"`
	// (Optional).  Number of vCPUs to allocate for the instance.
	// Defaults to 1.
	Vcpus *int32 `json:"vcpus,omitzero"`
	// Deprecated: Use `timeout_s` instead.  Timeout in milliseconds to
	// wait for all new instances to reach running state.  Requires
	// `autostart` to be set.  If `timeout_s` is not set, this value is
	// converted by rounding up to the next full second.  No wait
	// performed for a value of 0.
	WaitTimeoutMs *int64 `json:"wait_timeout_ms,omitzero"`
	// Features to enable for the instance.  Features are specific
	// configurations or capabilities that can be enabled for the
	// instance.  The `scale-to-zero` and `delete-on-stop` features are
	// mutually exclusive.
	Features []InstanceFeature `json:"features,omitzero"`
	// Timeout in seconds to wait for all new instances to reach running
	// state.  Requires `autostart` to be set.  If you autostart your
	// new instance, you can wait for it to finish starting with a
	// blocking API call if you specify a wait timeout greater than
	// zero.  No wait performed for a value of 0.
	TimeoutS *int64 `json:"timeout_s,omitzero"`
	// Read-Only Memory (ROM) blobs to attach to the instance.
	// Unikraft Cloud supports the ability to attach Read-Only Memory (ROM) blobs
	// to instances. It allows you to create a general-purpose base image and
	// then customize individual instances by attaching code or data as separate
	// ROM blobs.
	Roms []CreateInstanceRequestRom `json:"roms,omitzero"`
	// (Optional).  Plugins to attach to the instance.  Plugins let you attach
	// small helper programs to an instance and reach each one over a direct,
	// authenticated HTTP endpoint.  Each plugin loads from its own ROM image,
	// mounts at `/uk/plugins/<plugin_name>`, and is reachable at
	// `.../v1/instances/<uuid>/plugins/<plugin_name>/<path>`.  At most 8 plugins
	// may be attached to an instance.
	Plugins []CreateInstanceRequestPlugin `json:"plugins,omitzero"`
	// (Optional).  Tags to associate with the instance.
	Tags []string `json:"tags,omitzero"`
	// (Optional).  Annotations to associate with the instance.
	//
	// Unlike tags, annotations also reach the guest: they are included in the
	// instance's startdata, and selected keys can be injected into the console
	// log output.
	//
	// Keys follow the Kubernetes annotation key syntax, `[<prefix>/]<name>`: the
	// optional prefix is a non-wildcard DNS subdomain of at most 253 characters,
	// and the name is at most 63 characters of `[-_.a-zA-Z0-9]` starting and
	// ending with an alphanumeric.  Values are unconstrained apart from ASCII
	// control characters.  An instance holds at most 256 annotations.
	//
	// When the instance inherits annotations from a template, branch, or
	// checkpoint, the given annotations are merged into them rather than
	// replacing them.  On a key clash the value given here wins.
	Annotations map[string]string `json:"annotations,omitzero"`
	// Template instances.
	// An existing instance can be saved as a template. This template is then
	// used to create new instances that inherit the exact configuration and
	// state the original instance had when the template was created.
	Template *CreateInstanceRequestTemplate `json:"template,omitzero"`
	// The scheduling priority for the instance. Only settable by
	// users with scheduling priority override permissions.
	SchedPriority *SchedPriority `json:"sched_priority,omitzero"`
	// (Optional).  Schedules for the instance.  Scheduled operations let you
	// automatically start, stop, delete, or exec a command in the instance on
	// a calendar-based schedule.  For `exec` schedules, set the `args` field
	// to the command and its arguments.  Each instance stores its own
	// schedules, and cloning preserves them.
	Schedules []Schedule `json:"schedules,omitzero"`
	// (Optional).  Automatic delete-on-idle/request-limit configuration.
	// Not used for template instances.
	Autokill *CreateInstanceRequestAutokill `json:"autokill,omitzero"`
	// (Optional).  The hostname of the instance.
	//
	// If not provided, the hostname will be set to the instance name.  The
	// hostname must be a valid DNS label (e.g., "my-instance") and is used for
	// internal DNS resolution within the Unikraft Cloud network.
	Hostname *string `json:"hostname,omitzero"`
	// (Optional).  Dependencies of the instance.
	//
	// A list of instance identifiers (name or UUID) that this instance depends
	// on.  Dependencies define startup ordering and can be used to ensure that
	// prerequisite instances are running before this instance starts.
	Dependencies []NameOrUUID `json:"dependencies,omitzero"`
	// (Optional).  Reference to an existing instance to branch from.
	// The instance can be running or stopped, If the source
	// instance is running, a snapshot will be taken asynchronously and the
	// new instance will wait for it to complete before starting.
	// Mutually exclusive with `image` and `template`.
	BranchFrom *NameOrUUID `json:"branch_from,omitzero"`
	// (Optional).  Reference to an existing checkpoint to create the instance
	// from.  The checkpoint must be in the `checkpoint` state.  The new instance
	// will be created with the same configuration and state as the checkpoint.
	// Mutually exclusive with `image`, `template`, and `branch_from`.
	Checkpoint *NameOrUUID `json:"checkpoint,omitzero"`
	// The default gateway to configure inside the guest.
	Gateway *string `json:"gateway,omitzero"`
	// The DNS resolver to configure inside the guest.
	Nameserver *string `json:"nameserver,omitzero"`
	// A list of one to four interfaces to attach
	NetworkInterfaces []CreateInstanceRequestNetworkInterface `json:"network_interfaces,omitzero"`
	// (Optional).  The type of virtual machine to use for the instance.
	// Defaults to `micro`, which runs on Firecracker.  `full` runs on QEMU
	// instead and is required for GPU passthrough (see `gpus`) and, in the
	// future, Windows VMs.  QEMU-backed instances currently do not support
	// scale-to-zero, templates, branching, or checkpointing, and only
	// support block-based volumes (no virtiofs).  Requires a plan with full
	// VM support and cannot be combined with `template`, `branch_from`, or
	// `checkpoint`.
	Type *InstanceType `json:"type,omitzero"`
	// (Optional).  Number of GPUs to attach to the instance.  Currently
	// restricted to at most 1.  Requires `type` to be `full` and a plan
	// with GPU support.  A GPU stays assigned to the instance, even while
	// stopped, until the instance is deleted.  Cannot be combined with
	// `template`, `branch_from`, or `checkpoint`.
	Gpus *int32 `json:"gpus,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequest
	// Union members are decoded in a second step: a nil interface cannot be
	// decoded into directly.  Holding them as raw JSON ahead of the embedded
	// alias shadows the alias' own members of the same name.  An absent member
	// leaves the current value in place, whereas an explicit null clears it.
	aux := struct {
		Image jsontext.Value `json:"image,omitzero"`
		*Alias
	}{Alias: (*Alias)(m)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Image) > 0 {
		if aux.Image.Kind() == 'n' {
			m.Image = nil
		} else {
			value, err := UnmarshalImageSource(aux.Image)
			if err != nil {
				return err
			}
			m.Image = value
		}
	}
	return nil
}

func (m CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequest
	return json.Marshal((Alias)(m))
}

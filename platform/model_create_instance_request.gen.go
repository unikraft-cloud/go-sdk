// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The request message for creating a new instance.
// Restart policy for the instance.  This defines how the instance
// should behave when it stops or crashes.  Cannot be combined with
// the `delete-on-stop` feature.
type CreateInstanceRequestRestartPolicy string

const (
	CreateInstanceRequestRestartPolicyNever      CreateInstanceRequestRestartPolicy = "never"
	CreateInstanceRequestRestartPolicyAlways     CreateInstanceRequestRestartPolicy = "always"
	CreateInstanceRequestRestartPolicyOn_failure CreateInstanceRequestRestartPolicy = "on-failure"
)

// Features to enable for the instance.  Features are specific
// configurations or capabilities that can be enabled for the
// instance.  The `scale-to-zero` and `delete-on-stop` features are
// mutually exclusive.
type CreateInstanceRequestFeatures string

const (
	CreateInstanceRequestFeaturesDelete_on_stop CreateInstanceRequestFeatures = "delete-on-stop"
)

type CreateInstanceRequest struct {
	// (Optional).  The name of the instance.
	//
	// If not provided, a random name will be generated.  The name must be unique.
	Name *string `json:"name,omitempty"`
	// The image to use for the instance.
	Image *string `json:"image,omitempty"`
	// (Only applies when using global control plane).
	// The metro to route the request to.
	Metro *string `json:"metro,omitempty"`
	// (Optional).  The arguments to pass to the instance when it starts.
	Args []string `json:"args,omitempty"`
	// (Optional).  Environment variables to set for the instance.
	Env map[string]string `json:"env,omitempty"`
	// (Optional).  Memory in MB to allocate for the instance.  Default is 128.
	MemoryMb     *int64                             `json:"memory_mb,omitempty"`
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
	RestartPolicy *CreateInstanceRequestRestartPolicy `json:"restart_policy,omitempty"`
	ScaleToZero   *CreateInstanceRequestScaleToZero   `json:"scale_to_zero,omitempty"`
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
	Features []CreateInstanceRequestFeatures `json:"features,omitempty"`
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
	Tags     []string                       `json:"tags,omitempty"`
	Template *CreateInstanceRequestTemplate `json:"template,omitempty"`
	// (Optional).  The scheduling priority for the instance.  Higher values
	// indicate higher priority.
	SchedPriority *int32 `json:"sched_priority,omitempty"`
	// (Optional).  Schedules for the instance.  Scheduled operations let you
	// automatically start, stop, delete, or exec a command in the instance on
	// a calendar-based schedule.  For `exec` schedules, set the `args` field
	// to the command and its arguments.  Each instance stores its own
	// schedules, and cloning preserves them.
	Schedules []Schedule                     `json:"schedules,omitempty"`
	Autokill  *CreateInstanceRequestAutokill `json:"autokill,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateInstanceRequest) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":            {},
		"image":           {},
		"metro":           {},
		"args":            {},
		"env":             {},
		"memory_mb":       {},
		"service_group":   {},
		"volumes":         {},
		"autostart":       {},
		"replicas":        {},
		"restart_policy":  {},
		"scale_to_zero":   {},
		"vcpus":           {},
		"wait_timeout_ms": {},
		"features":        {},
		"timeout_s":       {},
		"roms":            {},
		"tags":            {},
		"template":        {},
		"sched_priority":  {},
		"schedules":       {},
		"autokill":        {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequest
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

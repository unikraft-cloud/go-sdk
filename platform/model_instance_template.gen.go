// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json"
	"time"
)

// An instance template describes the settable configuration for a template
// instance, excluding runtime-only fields.
// The state of the template instance.
type InstanceTemplateState string

const (
	InstanceTemplateStateStopped  InstanceTemplateState = "stopped"
	InstanceTemplateStateStarting InstanceTemplateState = "starting"
	InstanceTemplateStateRunning  InstanceTemplateState = "running"
	InstanceTemplateStateDraining InstanceTemplateState = "draining"
	InstanceTemplateStateStopping InstanceTemplateState = "stopping"
	InstanceTemplateStateTemplate InstanceTemplateState = "template"
	InstanceTemplateStateStandby  InstanceTemplateState = "standby"
)

// The restart policy for instances created from this template.
type InstanceTemplateRestartPolicy string

const (
	InstanceTemplateRestartPolicyNever      InstanceTemplateRestartPolicy = "never"
	InstanceTemplateRestartPolicyAlways     InstanceTemplateRestartPolicy = "always"
	InstanceTemplateRestartPolicyOn_failure InstanceTemplateRestartPolicy = "on-failure"
)

type InstanceTemplate struct {
	// The UUID of the template instance.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the template instance.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// Where the template instance is located.
	Metro *string `json:"metro,omitempty"`
	// The time the template instance was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The state of the template instance.
	State *InstanceTemplateState `json:"state,omitempty"`
	// The image used to create the template instance.
	Image *string `json:"image,omitempty"`
	// The amount of memory in megabytes allocated for the template instance.
	MemoryMb *uint64 `json:"memory_mb,omitempty"`
	// The number of vCPUs allocated for the template instance.
	Vcpus *uint32 `json:"vcpus,omitempty"`
	// The arguments passed to the template instance when it was created.
	Args []string `json:"args,omitempty"`
	// Environment variables set for the template instance.
	Env map[string]string `json:"env,omitempty"`
	// The restart policy for instances created from this template.
	RestartPolicy *InstanceTemplateRestartPolicy `json:"restart_policy,omitempty"`
	ScaleToZero   *InstanceTemplateScaleToZero   `json:"scale_to_zero,omitempty"`
	// The list of volumes attached to the template instance.
	Volumes []InstanceVolume `json:"volumes,omitempty"`
	// The tags associated with the template instance.
	Tags []string `json:"tags,omitempty"`
	// If set to true, the template instance cannot be deleted until the lock is removed.
	DeleteLock *bool `json:"delete_lock,omitempty"`
	// Scheduled operations for instances created from this template.
	Schedules        []Schedule                        `json:"schedules,omitempty"`
	TemplateAutokill *InstanceTemplateTemplateAutokill `json:"template_autokill,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstanceTemplate) UnmarshalJSON(data []byte) error {
	type Alias InstanceTemplate
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":              {},
		"name":              {},
		"metro":             {},
		"created_at":        {},
		"state":             {},
		"image":             {},
		"memory_mb":         {},
		"vcpus":             {},
		"args":              {},
		"env":               {},
		"restart_policy":    {},
		"scale_to_zero":     {},
		"volumes":           {},
		"tags":              {},
		"delete_lock":       {},
		"schedules":         {},
		"template_autokill": {},
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

func (m InstanceTemplate) MarshalJSON() ([]byte, error) {
	type Alias InstanceTemplate
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

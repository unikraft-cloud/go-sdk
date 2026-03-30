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

// A volume represents a storage device that can be attached to an instance.
// Current state of the volume.
type VolumeState string

const (
	VolumeStateUninitialized VolumeState = "uninitialized"
	VolumeStateInitializing  VolumeState = "initializing"
	VolumeStateAvailable     VolumeState = "available"
	VolumeStateIdle          VolumeState = "idle"
	VolumeStateMounted       VolumeState = "mounted"
	VolumeStateBusy          VolumeState = "busy"
	VolumeStateError         VolumeState = "error"
	VolumeStateTemplate      VolumeState = "template"
)

type Volume struct {
	// The UUID of the volume.
	//
	// This is a unique identifier for the volume that is generated when the
	// volume is created.  The UUID is used to reference the volume in
	// API calls and can be used to identify the volume in all API calls that
	// require an identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  The name can
	// also be used to identify the volume in API calls.
	Name *string `json:"name,omitempty"`
	// The time the volume was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Current state of the volume.
	State *VolumeState `json:"state,omitempty"`
	// The size of the volume in megabytes.
	SizeMb *uint64 `json:"size_mb,omitempty"`
	// Indicates if the volume will stay alive when the last instance is deleted
	// that this volume is attached to.
	Persistent *bool `json:"persistent,omitempty"`
	// List of instances that this volume is attached to.
	AttachedTo []VolumeInstanceID `json:"attached_to,omitempty"`
	// List of instances that have this volume mounted.
	// This does not apply to template volumes.
	MountedBy []VolumeInstanceMount `json:"mounted_by,omitempty"`
	// The tags associated with the volume.
	// Maximum 16 tags are allowed, and each tag may not be longer than 256 characters.
	Tags []string `json:"tags,omitempty"`
	// An optional field representing the status of the request.  This field is
	// only set when this message object is used as a response message.
	Status *ResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Message *string `json:"message,omitempty"`
	// An optional error code providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Error *int32 `json:"error,omitempty"`
	// Either static or dynamic reservation.
	QuotaPolicy *string `json:"quota_policy,omitempty"`
	// If set to true, the volume cannot be deleted.
	DeleteLock *bool `json:"delete_lock,omitempty"`
	// The amount of free space in the volume in megabytes.
	FreeMb *uint32 `json:"free_mb,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *Volume) UnmarshalJSON(data []byte) error {
	type Alias Volume
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":         {},
		"name":         {},
		"created_at":   {},
		"state":        {},
		"size_mb":      {},
		"persistent":   {},
		"attached_to":  {},
		"mounted_by":   {},
		"tags":         {},
		"status":       {},
		"message":      {},
		"error":        {},
		"quota_policy": {},
		"delete_lock":  {},
		"free_mb":      {},
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

func (m Volume) MarshalJSON() ([]byte, error) {
	type Alias Volume
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

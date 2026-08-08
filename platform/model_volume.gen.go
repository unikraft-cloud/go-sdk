// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"time"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// A volume represents a storage device that can be attached to an instance.
type Volume struct {
	// The UUID of the volume.
	//
	// This is a unique identifier for the volume that is generated when the
	// volume is created.  The UUID is used to reference the volume in
	// API calls and can be used to identify the volume in all API calls that
	// require an identifier.
	Uuid string `json:"uuid"`
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  The name can
	// also be used to identify the volume in API calls.
	Name string `json:"name"`
	// The time the volume was created.
	CreatedAt time.Time `json:"created_at"`
	// Current state of the volume.
	State VolumeState `json:"state"`
	// The size of the volume in megabytes.
	SizeMb uint64 `json:"size_mb"`
	// Indicates if the volume will stay alive when the last instance is deleted
	// that this volume is attached to.
	Persistent bool `json:"persistent"`
	// List of instances that this volume is attached to.
	AttachedTo []VolumeInstanceID `json:"attached_to,omitzero"`
	// List of instances that have this volume mounted.
	// This does not apply to template volumes.
	MountedBy []VolumeInstanceMount `json:"mounted_by,omitzero"`
	// The tags associated with the volume.
	// Maximum 16 tags are allowed, and each tag may not be longer than 256 characters.
	Tags []string `json:"tags,omitzero"`
	// An optional field representing the status of the request.  This field is
	// only set when this message object is used as a response message.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is only set when this message object is used as a response
	// message, and is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`
	// Either static or dynamic reservation.
	QuotaPolicy VolumeQuotaPolicy `json:"quota_policy"`
	// If set to true, the volume cannot be deleted.
	DeleteLock *bool `json:"delete_lock,omitzero"`
	// The amount of free space in the volume in megabytes.
	FreeMb *uint32 `json:"free_mb,omitzero"`
	// The filesystem type of this volume.
	// Without custom configuration, this is either `ext4` or `virtiofs`.
	Filesystem *string `json:"filesystem,omitzero"`
	// Host path backing this managed volume.
	// This field is only available for managed volumes and users with
	// appropriate permissions.
	HostPath *string `json:"host_path,omitzero"`
	// Optional script arguments that were applied to the custom volume filesystem
	// initialization scripts.
	Args map[string]string `json:"args,omitzero"`
	// The access mode of the volume, controlling volume sharing behavior.
	// Defaults to `rwo` if not specified.
	AccessMode *VolumeAccessMode `json:"access_mode,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Volume) UnmarshalJSON(data []byte) error {
	type Alias Volume
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Volume) MarshalJSON() ([]byte, error) {
	type Alias Volume
	return json.Marshal((Alias)(m))
}

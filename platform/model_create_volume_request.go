// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Creates one or more volumes with the given configuration. The volumes are
// automatically initialized with an empty file system. After initialization
// the volumes are in the available state and can be attached to an instance
// with the `PUT /v1/volumes/attach` endpoint. Note that, the size of a volume
// cannot be changed after creation.

type CreateVolumeRequest struct {
	// The size of the volume in megabytes.
	SizeMb uint64 `json:"size_mb"`
	// The name of the volume.
	//
	// This is a human-readable name that can be used to identify the volume.
	// The name must be unique within the context of your account.  If no name is
	// specified, a random name of the form `vol-X` is generated for you, where
	// `X` is a 5 character long random alphanumeric suffix..  The name can also
	// be used to identify the volume in API calls.
	Name *string `json:"name,omitempty"`
}

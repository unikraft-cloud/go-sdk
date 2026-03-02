// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type CloneVolumesRequest struct {
	// The UUID of the volume to clone. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to clone. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string `json:"name,omitempty"`
	// The name of the new cloned volume.
	VolName *string `json:"vol_name,omitempty"`
	// The tags associated with the volume.
	// Maximum 16 tags are allowed, and each tag may not be longer than 256 characters.
	Tags []string `json:"tags,omitempty"`
}

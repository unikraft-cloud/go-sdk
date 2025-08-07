// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Detaches a volume from instances. If no particular instance is specified the
// volume is detached from all instances. The instances from which to detach
// must not have the volume mounted. The API returns an error for each instance
// from which it was unable to detach the volume. If the volume has been
// created together with an instance, detaching the volume will make it
// persistent (i.e., it survives the deletion of the instance).

type DetachVolumesRequest struct {
	// The UUID of the volume to detach. Mutually exclusive with name.
	// Exactly one of uuid or name must be provided.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the volume to detach. Mutually exclusive with UUID.
	// Exactly one of uuid or name must be provided.
	Name *string                   `json:"name,omitempty"`
	From *DetachVolumesRequestFrom `json:"from,omitempty"`
}

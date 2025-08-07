// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// An identifier for the volume(s) to detach.
type DetachVolumesRequestID struct {
	// The UUID of the volume to detach.  Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// The name of the volume to detach.  Mutually exclusive with UUID.
	Name string `json:"name"`
}

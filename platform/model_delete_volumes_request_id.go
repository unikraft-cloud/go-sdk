// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// An identifier for the volume(s) to delete.

type DeleteVolumesRequestID struct {
	// The UUID of the volume to delete.  Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// The name of the volume to delete.  Mutually exclusive with UUID.
	Name string `json:"name"`
}

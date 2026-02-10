// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Read-Only Memory (ROM) blob to attach to the instance.

type CreateInstanceRequestRom struct {
	// The name of the ROM to use for the instance configuration.
	Name *string `json:"name,omitempty"`
	// The image of the ROM to use for the instance configuration.
	Image string `json:"image"`
}

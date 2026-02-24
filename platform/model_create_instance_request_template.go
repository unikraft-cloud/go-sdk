// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// (Optional).  The UUID of a template instance to create the instance from.

type CreateInstanceRequestTemplate struct {
	// Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// Mutually exclusive with UUID.
	Name string `json:"name"`
}

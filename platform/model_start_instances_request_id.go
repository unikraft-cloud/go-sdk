// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// An identifier for the instance(s) to start.

type StartInstancesRequestID struct {
	// The UUID of the instance to start.  Mutually exclusive with name.
	Uuid string `json:"uuid"`
	// The name of the instance to start.  Mutually exclusive with UUID.
	Name string `json:"name"`
}

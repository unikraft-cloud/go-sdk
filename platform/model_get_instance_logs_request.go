// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request message for getting the logs of an instance by their UUID or
// name.

type GetInstanceLogsRequest struct {
	// The UUID of the instance to retrieve logs for.  Mutually exclusive with
	// name.
	Uuid string `json:"uuid"`
	// The name of the instance to retrieve logs for.  Mutually exclusive with
	// UUID.
	Name string `json:"name"`
	// The byte offset of the log output to receive.  A negative sign makes the
	// offset relative to the end of the log.
	Offset *uint64 `json:"offset,omitempty"`
	// The amount of bytes to return at most.
	Limit *int64 `json:"limit,omitempty"`
}

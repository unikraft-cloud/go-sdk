// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request message for waiting for an instance to reach a certain state by
// its UUID.
type WaitInstanceByUUIDRequest struct {
	// The UUID of the instance to delete.
	Uuid *string `json:"uuid,omitempty"`
	// The desired state to wait for.  Default is `running`.
	State *string `json:"state,omitempty"`
	// Timeout in milliseconds to wait for the instance to reach the desired
	// state.  If the timeout is reached, the request will fail with an error.
	// A value of -1 means to wait indefinitely until the instance reaches the
	// desired state.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type WaitInstanceByUUIDRequestAllOf struct {
	// The UUID of the instance to delete.
	Uuid *string `json:"uuid,omitempty"`
	// The desired state to wait for.  Default is `running`.
	State *string `json:"state,omitempty"`
}

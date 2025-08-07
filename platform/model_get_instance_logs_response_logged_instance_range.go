// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Description of the range that was returned. Useful for requests with
// offset relative to end.

type GetInstanceLogsResponseLoggedInstanceRange struct {
	// The first retrieved byte.
	Start *int32 `json:"start,omitempty"`
	// The last retrieved byte.
	End *int32 `json:"end,omitempty"`
}

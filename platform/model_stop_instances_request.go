// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request message for stopping one or more instance(s) by their UUID(s) or
// name(s).

type StopInstancesRequest struct {
	// The list of IDs of the instance to stop.
	Ids []StopInstancesRequestID `json:"ids,omitempty"`
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The request message for checking authorization given a request ID.

type CheckAuthorizationRequest struct {
	// The request ID is a unique identifier for the request.  This is used to
	// track the request in the system and should be provided by the client.
	RequestId *string `json:"request_id,omitempty"`
}

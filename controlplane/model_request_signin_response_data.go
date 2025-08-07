// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The response data for this request.

type RequestSigninResponseData struct {
	// The full URL which can be used to remotely confirm the signin.
	AuthorizationUrl *string `json:"authorization_url,omitempty"`
	// A unique identifier for the request.  This can be used to track the
	// request in the system.
	RequestId *string `json:"request_id,omitempty"`
}

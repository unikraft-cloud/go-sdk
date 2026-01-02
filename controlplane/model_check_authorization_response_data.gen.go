// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The response data for this request.

type CheckAuthorizationResponseData struct {
	// The authorization token which can be used to authenticate requests.
	Token *string `json:"token,omitempty"`
	// The organization name the token is associated with.
	OrganizationName *string `json:"organization_name,omitempty"`
	// The display name of the organization the token is associated with.
	OrganizationDisplayName *string `json:"organization_display_name,omitempty"`
}

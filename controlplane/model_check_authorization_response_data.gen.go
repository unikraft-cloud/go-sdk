// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

type CheckAuthorizationResponseData struct {
	// The authorization token which can be used to authenticate requests.
	Token string `json:"token"`
	// The organization name the token is associated with.
	OrganizationName string `json:"organization_name"`
	// The display name of the organization the token is associated with.
	OrganizationDisplayName string `json:"organization_display_name"`
	// The global image registry.
	Registry string `json:"registry"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CheckAuthorizationResponseData) UnmarshalJSON(data []byte) error {
	type Alias CheckAuthorizationResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CheckAuthorizationResponseData) MarshalJSON() ([]byte, error) {
	type Alias CheckAuthorizationResponseData
	return json.Marshal((Alias)(m))
}

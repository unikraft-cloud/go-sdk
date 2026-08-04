// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type GetAuthorizationResponseData struct {
	// The organization name the token is associated with.
	OrganizationName string `json:"organization_name,omitzero"`
	// The display name of the organization the token is associated with.
	OrganizationDisplayName string `json:"organization_display_name,omitzero"`
	// The global image registry.
	Registry string `json:"registry,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *GetAuthorizationResponseData) UnmarshalJSON(data []byte) error {
	type Alias GetAuthorizationResponseData
	return json.Unmarshal(data, (*Alias)(m))
}

func (m GetAuthorizationResponseData) MarshalJSON() ([]byte, error) {
	type Alias GetAuthorizationResponseData
	return json.Marshal((Alias)(m))
}

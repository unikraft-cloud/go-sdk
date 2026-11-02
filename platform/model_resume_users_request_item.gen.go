// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A single request item to resume a user account.
type ResumeUsersRequestItem struct {
	// The UUID of the user.
	Uuid string `json:"uuid"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ResumeUsersRequestItem) UnmarshalJSON(data []byte) error {
	type Alias ResumeUsersRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ResumeUsersRequestItem) MarshalJSON() ([]byte, error) {
	type Alias ResumeUsersRequestItem
	return json.Marshal((Alias)(m))
}

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

// Per-item result for a resume users operation.
type ResumeUsersResponseResumedUser struct {
	// Indicates whether the operation was successful for this item.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitzero"`
	// Whether the user account is still suspended after this request, for
	// example because it was suspended again concurrently.
	Suspended bool `json:"suspended"`
	// The instances of the user that were restored, as recorded when the user
	// was suspended. Instances that had nothing to restore are not listed. An
	// instance that failed to start keeps its record, so a later resume tries
	// again.
	Instances []ResumeUsersResponseResumedInstance `json:"instances,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ResumeUsersResponseResumedUser) UnmarshalJSON(data []byte) error {
	type Alias ResumeUsersResponseResumedUser
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ResumeUsersResponseResumedUser) MarshalJSON() ([]byte, error) {
	type Alias ResumeUsersResponseResumedUser
	return json.Marshal((Alias)(m))
}

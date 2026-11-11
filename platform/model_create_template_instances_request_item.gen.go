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

// A single template instance to be created.
type CreateTemplateInstancesRequestItem struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`
	// Timeout in seconds to wait for the template instances to be created.
	// No wait is performed for a value of 0, which is the default. The wait is
	// capped at the platform's maximum API timeout (60 seconds by default);
	// larger values are rejected, and a value of -1 waits for that maximum. If
	// several instances are converted in one request, the shortest positive
	// timeout applies to all of them.
	//
	// Reaching the timeout is not an HTTP error: the response carries the
	// message `Operation timed out`, and each instance that has not become a
	// template yet is reported with its current `state` and an `error` status.
	TimeoutS *int64 `json:"timeout_s,omitzero"`
	// Automatic delete-on-idle configuration for the new template.
	Autokill *ItemAutokill `json:"autokill,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateTemplateInstancesRequestItem) UnmarshalJSON(data []byte) error {
	type Alias CreateTemplateInstancesRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateTemplateInstancesRequestItem) MarshalJSON() ([]byte, error) {
	type Alias CreateTemplateInstancesRequestItem
	return json.Marshal((Alias)(m))
}

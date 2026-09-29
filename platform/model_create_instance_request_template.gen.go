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

// Template configuration when creating an instance.
type CreateInstanceRequestTemplate struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name *string `json:"name,omitzero"`
	// Whether the instance needs to run in order to reach template state
	Prepare *bool `json:"prepare,omitzero"`
	// Configuration parameters to apply when building the new instance from the
	// source template.
	CreateArgs *Instance `json:"create_args,omitzero"`
	// Timeout in seconds for preparing the template before the preparation is
	// aborted. Only applies when `prepare` is set. A value of 0 means no
	// timeout.
	PrepareTimeoutS *int64 `json:"prepare_timeout_s,omitzero"`
	// Automatic delete-on-idle configuration for the template. Only applies
	// when `prepare` is set.
	Autokill *TemplateAutokill `json:"autokill,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestTemplate) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestTemplate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestTemplate) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestTemplate
	return json.Marshal((Alias)(m))
}

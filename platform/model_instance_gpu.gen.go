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

// A GPU attached to the instance.
type InstanceGpu struct {
	// The UUID of the GPU.
	Uuid string `json:"uuid"`
	// The GPU model, given as its PCI vendor and device ID in the form
	// `<vendor>:<device>`.
	Model string `json:"model"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceGpu) UnmarshalJSON(data []byte) error {
	type Alias InstanceGpu
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceGpu) MarshalJSON() ([]byte, error) {
	type Alias InstanceGpu
	return json.Marshal((Alias)(m))
}

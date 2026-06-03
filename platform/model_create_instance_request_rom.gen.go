// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Read-Only Memory (ROM) blob to attach to the instance.

type CreateInstanceRequestRom struct {
	// The name of the ROM to use for the instance configuration.
	Name string `json:"name"`
	// (Optional).  The image of the ROM to use for the instance configuration.
	// Mutually exclusive with `files`.
	Image *string `json:"image,omitempty"`
	// (Optional).  Inline files to use as the ROM content.  When specified,
	// the platform creates an EROFS image from the provided files.
	// Mutually exclusive with `image`.
	Files []InlineFile `json:"files,omitempty"`
	// (Optional).  The path at which the ROM should be automatically mounted
	// inside the instance.  When set, the platform mounts the ROM device at
	// the specified path so the guest does not need to mount it manually.
	// When omitted, the ROM is exposed as a raw block device and the guest is
	// responsible for mounting it.
	At *string `json:"at,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CreateInstanceRequestRom) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestRom
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestRom) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestRom
	return json.Marshal((Alias)(m))
}

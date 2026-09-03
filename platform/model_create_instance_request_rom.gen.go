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

// Read-Only Memory (ROM) blob to attach to the instance.
type CreateInstanceRequestRom struct {
	// The name of the ROM to use for the instance configuration.
	Name string `json:"name"`
	// (Optional).  The image of the ROM to use for the instance configuration.
	// Mutually exclusive with `files`.  Accepts either a plain image reference
	// string (`"nginx:latest"`) or an object carrying additional pull
	// configuration (`{"url": "nginx:latest", "pull_policy": "always"}`).
	Image ImageSource `json:"image,omitzero"`
	// (Optional).  Inline files to use as the ROM content.  When specified,
	// the platform creates an EROFS image from the provided files.
	// Mutually exclusive with `image`.
	Files []InlineFile `json:"files,omitzero"`
	// (Optional).  The path at which the ROM should be automatically mounted
	// inside the instance.  When set, the platform mounts the ROM device at
	// the specified path so the guest does not need to mount it manually.
	// When omitted, the ROM is exposed as a raw block device and the guest is
	// responsible for mounting it.
	At *string `json:"at,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestRom) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestRom
	// Union members are decoded in a second step: a nil interface cannot be
	// decoded into directly.  Holding them as raw JSON ahead of the embedded
	// alias shadows the alias' own members of the same name.  An absent member
	// leaves the current value in place, whereas an explicit null clears it.
	aux := struct {
		Image jsontext.Value `json:"image,omitzero"`
		*Alias
	}{Alias: (*Alias)(m)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Image) > 0 {
		if aux.Image.Kind() == 'n' {
			m.Image = nil
		} else {
			value, err := UnmarshalImageSource(aux.Image)
			if err != nil {
				return err
			}
			m.Image = value
		}
	}
	return nil
}

func (m CreateInstanceRequestRom) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestRom
	return json.Marshal((Alias)(m))
}

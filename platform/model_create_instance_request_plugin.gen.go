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

// A helper program attached to the instance and reachable over a direct,
// authenticated HTTP endpoint.  A plugin runs inside the instance next to the
// main application, loads from its own ROM image, and answers requests that
// the Unikraft Cloud API forwards to it.
type CreateInstanceRequestPlugin struct {
	// The plugin name.  It becomes the `<plugin_name>` segment in the plugin
	// endpoint (`.../plugins/<plugin_name>/<path>`).  A plugin name has a
	// maximum length of 63 characters and contains only letters (`a`-`z`,
	// `A`-`Z`), digits (`0`-`9`), hyphen (`-`), and underscore (`_`).
	Name string `json:"name"`
	// The plugin's ROM image.  The platform loads the image, mounts it at
	// `/uk/plugins/<plugin_name>`, and runs its `init` program when the plugin
	// starts.  Accepts either a plain image reference string
	// (`"user/myplugin:latest"`) or an object carrying additional pull
	// configuration (`{"url": "user/myplugin:latest", "pull_policy": "always"}`).
	Rom ImageSource `json:"rom"`
	// (Optional).  Arbitrary JSON configuration that the platform passes to the
	// plugin's `init` program on `STDIN`.  Any JSON value works, including a
	// string, a number, or an object.
	Config *interface{} `json:"config,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestPlugin) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestPlugin
	// Union members are decoded in a second step: a nil interface cannot be
	// decoded into directly.  Holding them as raw JSON ahead of the embedded
	// alias shadows the alias' own members of the same name.  An absent member
	// leaves the current value in place, whereas an explicit null clears it.
	aux := struct {
		Rom jsontext.Value `json:"rom,omitzero"`
		*Alias
	}{Alias: (*Alias)(m)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Rom) > 0 {
		if aux.Rom.Kind() == 'n' {
			m.Rom = nil
		} else {
			value, err := UnmarshalImageSource(aux.Rom)
			if err != nil {
				return err
			}
			m.Rom = value
		}
	}
	return nil
}

func (m CreateInstanceRequestPlugin) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestPlugin
	return json.Marshal((Alias)(m))
}

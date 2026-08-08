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
	// The plugin's ROM image, given as an image reference string such as
	// `user/myplugin:latest`.  The platform loads the image, mounts it at
	// `/uk/plugins/<plugin_name>`, and runs its `init` program when the plugin
	// starts.
	Rom string `json:"rom"`
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
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestPlugin) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestPlugin
	return json.Marshal((Alias)(m))
}

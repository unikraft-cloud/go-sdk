// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

import (
	"encoding/json"
	"fmt"
)

// BuildContexts are the contexts of a build, by the names of buildkit's
// Dockerfile frontend. In JSON, they are one object, by name, which is why
// they are written by hand rather than rendered from the spec.
type BuildContexts struct {
	// Context is the build context, where the paths of the Kraftfile resolve.
	Context Source

	// Dockerfile holds the Dockerfile of the rootfs. When nil, Context holds it.
	Dockerfile *Source

	// Named are the named contexts, for FROM <name> and COPY --from=<name>.
	Named map[string]Source
}

// MarshalJSON encodes the contexts as one object. It refuses a named context
// with the name context or dockerfile.
func (c BuildContexts) MarshalJSON() ([]byte, error) {
	obj := make(map[string]Source, len(c.Named)+2)
	for name, src := range c.Named {
		if name == "context" || name == "dockerfile" {
			return nil, fmt.Errorf("a named context cannot have the name %q", name)
		}
		obj[name] = src
	}
	obj["context"] = c.Context
	if c.Dockerfile != nil {
		obj["dockerfile"] = *c.Dockerfile
	}

	return json.Marshal(obj)
}

// UnmarshalJSON decodes the contexts from one object. A key other than
// context and dockerfile is a named context, and null is no context.
func (c *BuildContexts) UnmarshalJSON(b []byte) error {
	var obj map[string]*Source
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	*c = BuildContexts{Dockerfile: obj["dockerfile"]}
	if src := obj["context"]; src != nil {
		c.Context = *src
	}
	for name, src := range obj {
		if src == nil || name == "context" || name == "dockerfile" {
			continue
		}
		if c.Named == nil {
			c.Named = make(map[string]Source)
		}
		c.Named[name] = *src
	}

	return nil
}

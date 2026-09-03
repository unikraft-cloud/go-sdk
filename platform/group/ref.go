// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"fmt"

	"unikraft.com/cloud/sdk/platform"
)

// Ref represents a resource that may be identified by name or UUID,
// optionally scoped to a metro and a node within that metro.
type Ref struct {
	// Metro is the metro in which to look for the resource. If empty, the
	// resource will be looked for in all metros.
	Metro string

	// Node is the node within Metro on which to look for the resource. It is
	// only meaningful when Metro is set.
	Node string

	// Name is the name of the resource, from the platform API. It is not
	// guaranteed to be unique across metros.
	Name string
	// UUID is the UUID of the resource, from the platform API. It should not
	// be, but is not guaranteed to be unique across metros.
	UUID string

	// Display is a human-friendly string to use when displaying this Ref.
	Display string
}

func (r Ref) NameOrUUID() platform.NameOrUUID {
	if r.UUID != "" {
		return platform.NameOrUUID{Uuid: &r.UUID}
	}
	if r.Name != "" {
		return platform.NameOrUUID{Name: &r.Name}
	}
	return platform.NameOrUUID{}
}

func (r Ref) IsZero() bool {
	return r.Name == "" && r.UUID == ""
}

func (r Ref) String() string {
	if r.Display != "" {
		return r.Display
	}
	if r.Name != "" {
		return r.Name
	}
	return r.UUID
}

// variants returns all ref shapes that could have been used to request the
// resource identified by r: every combination of the identifiers present
// (name, UUID, or both) across every scope level (metro/node, metro-only,
// unscoped).
func (r Ref) variants() []Ref {
	ids := make([]Ref, 0, 3)
	if r.Name != "" {
		ids = append(ids, Ref{Name: r.Name})
	}
	if r.UUID != "" {
		ids = append(ids, Ref{UUID: r.UUID})
	}
	if r.Name != "" && r.UUID != "" {
		ids = append(ids, Ref{Name: r.Name, UUID: r.UUID})
	}

	var rs []Ref
	for _, id := range ids {
		if r.Metro != "" && r.Node != "" {
			rs = append(rs, Ref{Metro: r.Metro, Node: r.Node, Name: id.Name, UUID: id.UUID})
		}
		if r.Metro != "" {
			rs = append(rs, Ref{Metro: r.Metro, Name: id.Name, UUID: id.UUID})
		}
		rs = append(rs, id)
	}
	return rs
}

type Refs []Ref

func (rs Refs) NameOrUUIDs() []platform.NameOrUUID {
	results := make([]platform.NameOrUUID, len(rs))
	for i, r := range rs {
		results[i] = r.NameOrUUID()
	}
	return results
}

func (rs Refs) Strings() []string {
	results := make([]string, len(rs))
	for i, r := range rs {
		results[i] = r.String()
	}
	return results
}

type ErrRefNotFound struct {
	Refs Refs
}

func (e ErrRefNotFound) Error() string {
	return fmt.Sprintf("references not found: %v", e.Refs.Strings())
}

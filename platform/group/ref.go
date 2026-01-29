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
// optionally scoped to a metro.
type Ref struct {
	Name string
	UUID string

	Metro string
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

func (r Ref) NameOrUUIDString() string {
	if r.Name != "" {
		return r.Name
	}
	return r.UUID
}

func (r Ref) variants() []Ref {
	var rs []Ref
	if r.Metro != "" {
		if r.Name != "" {
			rs = append(rs, Ref{Metro: r.Metro, Name: r.Name})
		}
		if r.UUID != "" {
			rs = append(rs, Ref{Metro: r.Metro, UUID: r.UUID})
		}
		if r.Name != "" && r.UUID != "" {
			rs = append(rs, Ref{Metro: r.Metro, Name: r.Name, UUID: r.UUID})
		}
	}
	if r.Name != "" {
		rs = append(rs, Ref{Name: r.Name})
	}
	if r.UUID != "" {
		rs = append(rs, Ref{UUID: r.UUID})
	}
	if r.Name != "" && r.UUID != "" {
		rs = append(rs, Ref{Name: r.Name, UUID: r.UUID})
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

func (rs Refs) NameOrUUIDStrings() []string {
	results := make([]string, len(rs))
	for i, r := range rs {
		results[i] = r.NameOrUUIDString()
	}
	return results
}

type ErrRefNotFound struct {
	Refs Refs
}

func (e ErrRefNotFound) Error() string {
	return fmt.Sprintf("references not found: %v", e.Refs.NameOrUUIDStrings())
}

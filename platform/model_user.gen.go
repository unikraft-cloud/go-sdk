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

type User struct {
	// The UUID of the user.
	Uuid string `json:"uuid"`
	// The name of the user.
	Name string `json:"name"`
	// Authentication token(s) associated with the user.
	AuthToken []string `json:"auth_token"`
	// The permission level of the user.
	Permissions []UserPermission `json:"permissions,omitempty"`
	// The user ID (UID) on the host system.
	Uid *uint32 `json:"uid,omitempty"`
	// Whether the user account is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Per-VM Configuration limits for the user.
	Vmdb *UserVmdb `json:"vmdb,omitempty"`
	// Network configuration limits for the user.
	Net *UserNet `json:"net,omitempty"`
	// Global VM configuration limits for the user.
	Vmm *UserVmm `json:"vmm,omitempty"`
	// Storage configuration limits for the user.
	Stor *UserStor `json:"stor,omitempty"`
	// Autoscale configuration limits for the user.
	Autoscale *UserAutoscale `json:"autoscale,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *User) UnmarshalJSON(data []byte) error {
	type Alias User
	return json.Unmarshal(data, (*Alias)(m))
}

func (m User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal((Alias)(m))
}

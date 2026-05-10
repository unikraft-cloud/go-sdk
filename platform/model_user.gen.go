// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// The permission level of the user.
type UserPermissions string

const (
	UserPermissionsRoot                    UserPermissions = "root"
	UserPermissionsOverride_edns_blacklist UserPermissions = "override_edns_blacklist"
	UserPermissionsDeveloper               UserPermissions = "developer"
	UserPermissionsVolume_manager          UserPermissions = "volume_manager"
	UserPermissionsOverride_vm_priority    UserPermissions = "override_vm_priority"
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

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *User) UnmarshalJSON(data []byte) error {
	type Alias User
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":        {},
		"name":        {},
		"auth_token":  {},
		"permissions": {},
		"uid":         {},
		"disabled":    {},
		"vmdb":        {},
		"net":         {},
		"vmm":         {},
		"stor":        {},
		"autoscale":   {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m User) MarshalJSON() ([]byte, error) {
	type Alias User
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

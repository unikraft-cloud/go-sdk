// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// SSHKey represents an SSH public key for authentication.

type SSHKey struct {
	// A name or label for this SSH key.
	Name string `json:"name"`
	// The SSH public key in OpenSSH format (e.g., "ssh-rsa AAAA... user@host"
	// or "ssh-ed25519 AAAA... user@host").
	PublicKey string `json:"public_key"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *SSHKey) UnmarshalJSON(data []byte) error {
	type Alias SSHKey
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":       {},
		"public_key": {},
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

func (m SSHKey) MarshalJSON() ([]byte, error) {
	type Alias SSHKey
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

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// SSHKey represents an SSH public key for authentication.
type SSHKey struct {
	// A name or label for this SSH key.
	Name string `json:"name"`
	// The SSH public key in OpenSSH format (e.g., "ssh-rsa AAAA... user@host"
	// or "ssh-ed25519 AAAA... user@host").
	PublicKey string `json:"public_key"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *SSHKey) UnmarshalJSON(data []byte) error {
	type Alias SSHKey
	return json.Unmarshal(data, (*Alias)(m))
}

func (m SSHKey) MarshalJSON() ([]byte, error) {
	type Alias SSHKey
	return json.Marshal((Alias)(m))
}

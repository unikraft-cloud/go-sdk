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

// An object that routes all traffic through another interface.
type NetworkInterfaceRelay struct {
	// Whether the relay forwards DNS requests. Set to false
	// to let the default DNS server handle them instead.
	// Defaults to true.
	RelayDns *bool `json:"relay_dns,omitempty"`
	// UUID of the existing interface to relay through.
	// Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// Name of the existing interface to relay through.
	// Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *NetworkInterfaceRelay) UnmarshalJSON(data []byte) error {
	type Alias NetworkInterfaceRelay
	return json.Unmarshal(data, (*Alias)(m))
}

func (m NetworkInterfaceRelay) MarshalJSON() ([]byte, error) {
	type Alias NetworkInterfaceRelay
	return json.Marshal((Alias)(m))
}

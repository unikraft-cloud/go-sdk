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

// An object that routes all traffic through another interface.
type InstanceNetworkInterfaceRelay struct {
	// UUID of the relay interface.
	Uuid string `json:"uuid"`
	// Name of the relay interface.
	Name string `json:"name"`
	// Whether DNS traffic is relayed through this interface.
	// Defaults to true.
	RelayDns bool `json:"relay_dns"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceNetworkInterfaceRelay) UnmarshalJSON(data []byte) error {
	type Alias InstanceNetworkInterfaceRelay
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceNetworkInterfaceRelay) MarshalJSON() ([]byte, error) {
	type Alias InstanceNetworkInterfaceRelay
	return json.Marshal((Alias)(m))
}

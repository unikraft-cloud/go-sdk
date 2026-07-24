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

// An instance network interface.
type InstanceNetworkInterface struct {
	// The UUID of the network interface. This is a unique identifier for the
	// network interface that is generated when the instance is created.
	Uuid string `json:"uuid"`
	// The private IP address of the network interface. This is the internal IP
	// address that is used for communication between instances within the same
	// network.
	PrivateIp string `json:"private_ip"`
	// The MAC address of the network interface.
	Mac string `json:"mac"`
	// The interface name. If omitted, Unikraft Cloud generates one as
	// <instance-name>-ethX, falling back to eth-<suffix> when the
	// instance name is too long.
	Name *string `json:"name,omitempty"`
	// The TAP device to attach the interface. Provide it together with ip.
	TapName *string `json:"tap_name,omitempty"`
	// Whether the interface is automatically configured inside the guest
	// (IP address, routes, etc.).  When absent or true, autoconfiguration
	// is enabled.  Present and false when the guest is expected to
	// configure the interface manually.
	Autoconfig *bool `json:"autoconfig,omitempty"`
	// Relay configuration for this interface.
	Relay *InstanceNetworkInterfaceRelay `json:"relay,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *InstanceNetworkInterface) UnmarshalJSON(data []byte) error {
	type Alias InstanceNetworkInterface
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceNetworkInterface) MarshalJSON() ([]byte, error) {
	type Alias InstanceNetworkInterface
	return json.Marshal((Alias)(m))
}

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
	// Amount of bytes received from interface.
	RxBytes uint64 `json:"rx_bytes"`
	// Count of packets received from interface
	RxPackets uint64 `json:"rx_packets"`
	// Amount of bytes sent to interface.
	TxBytes uint64 `json:"tx_bytes"`
	// Count of packets sent to interface
	TxPackets uint64 `json:"tx_packets"`

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

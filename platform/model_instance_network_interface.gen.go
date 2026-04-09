// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// An instance network interface.

type InstanceNetworkInterface struct {
	// The UUID of the network interface. This is a unique identifier for the
	// network interface that is generated when the instance is created.
	Uuid *string `json:"uuid,omitempty"`
	// The private IP address of the network interface. This is the internal IP
	// address that is used for communication between instances within the same
	// network.
	PrivateIp *string `json:"private_ip,omitempty"`
	// The MAC address of the network interface.
	Mac *string `json:"mac,omitempty"`
	// Amount of bytes received from interface.
	RxBytes *uint64 `json:"rx_bytes,omitempty"`
	// Count of packets received from interface
	RxPackets *uint64 `json:"rx_packets,omitempty"`
	// Amount of bytes sent to interface.
	TxBytes *uint64 `json:"tx_bytes,omitempty"`
	// Count of packets sent to interface
	TxPackets *uint64 `json:"tx_packets,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InstanceNetworkInterface) UnmarshalJSON(data []byte) error {
	type Alias InstanceNetworkInterface
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":       {},
		"private_ip": {},
		"mac":        {},
		"rx_bytes":   {},
		"rx_packets": {},
		"tx_bytes":   {},
		"tx_packets": {},
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

func (m InstanceNetworkInterface) MarshalJSON() ([]byte, error) {
	type Alias InstanceNetworkInterface
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

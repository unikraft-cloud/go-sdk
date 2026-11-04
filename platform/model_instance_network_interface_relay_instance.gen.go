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

// The instance that a relay interface is attached to.
type InstanceNetworkInterfaceRelayInstance struct {
	// UUID of the instance the relay interface is attached to.
	Uuid string `json:"uuid"`
	// Name of the instance the relay interface is attached to.
	Name string `json:"name"`
	// The name of the node the resource is placed on. Only present on metros
	// that span multiple nodes.
	Node *string `json:"node,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceNetworkInterfaceRelayInstance) UnmarshalJSON(data []byte) error {
	type Alias InstanceNetworkInterfaceRelayInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceNetworkInterfaceRelayInstance) MarshalJSON() ([]byte, error) {
	type Alias InstanceNetworkInterfaceRelayInstance
	return json.Marshal((Alias)(m))
}

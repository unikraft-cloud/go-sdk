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

// The service group configuration when creating an instance. If no existing
// (persistent) service group is specified via its identifier, a new
// (ephemeral) service group can be created by specifying the services it
// should expose. A service defines the configuration settings of an exposed
// port by the instance. A service is a combination of a public port, an
// internal port, and a set of handlers that define how the service will handle
// incoming connections.
type CreateInstanceRequestServiceGroup struct {
	// The UUID of the resource.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the resource.
	Name      *string                       `json:"name,omitzero"`
	Domains   []CreateInstanceRequestDomain `json:"domains,omitzero"`
	Services  []Service                     `json:"services,omitzero"`
	SoftLimit *uint32                       `json:"soft_limit,omitzero"`
	HardLimit *uint32                       `json:"hard_limit,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *CreateInstanceRequestServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CreateInstanceRequestServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestServiceGroup
	return json.Marshal((Alias)(m))
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

// The service group configuration when creating an instance.
//
// If no existing (persistent) service group is specified via its identifier,
// a new (ephemeral) service group can be created by specifying the services
// it should expose.  A service defines the configuration settings of an
// exposed port by the instance.  A service is a combination of a public port,
// an internal port, and a set of handlers that define how the service will
// handle incoming connections.
type CreateInstanceRequestServiceGroup struct {
	// If no existing (persistent) service group is specified via its
	// identifier, a new (ephemeral) service group can be created.  In addition
	// to the services it must expose, you can specify which domains it should
	// use too.
	Domains []CreateInstanceRequestDomain `json:"domains,omitzero"`
	// If no existing service group identifier is provided, one or more new
	// (ephemeral, non-persistent) service(s) can be created with the following
	// definitions.
	Services []Service `json:"services,omitzero"`
	// The soft limit for the number of services that can be created in this
	// service group.
	SoftLimit *uint32 `json:"soft_limit,omitzero"`
	// The hard limit for the number of services that can be created in this
	// service group.
	HardLimit *uint32 `json:"hard_limit,omitzero"`
	// (Optional).  Reference an existing (persistent) service group by its
	// UUID.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitzero"`
	// (Optional).  Reference an existing (persistent) service group by its
	// name.  Mutually exclusive with UUID.
	Name *string `json:"name,omitzero"`

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

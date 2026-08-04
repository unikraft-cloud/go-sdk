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

// The service group configuration for the instance.
//
// This is a reference to the service group that the instance is part of.  The
// service group defines the services (e.g. ports, connection handling) that
// the instance exposes and how they are configured.
type InstanceServiceGroup struct {
	// The UUID of the service group.
	//
	// This is a unique identifier for the service group that is generated when
	// the service is created.  The UUID is used to reference the service group
	// in API calls and can be used to identify the service in all API calls
	// that require an service identifier.
	Uuid string `json:"uuid,omitzero"`
	// The name of the service group.
	//
	// This is a human-readable name that can be used to identify the service
	// group.  The name is unique within the context of your account.  The name
	// can also be used to identify the service group in API calls.
	Name string `json:"name,omitzero"`
	// The domain configuration for the service group.
	Domains []ServiceGroupInstanceDomain `json:"domains,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InstanceServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias InstanceServiceGroup
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InstanceServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias InstanceServiceGroup
	return json.Marshal((Alias)(m))
}

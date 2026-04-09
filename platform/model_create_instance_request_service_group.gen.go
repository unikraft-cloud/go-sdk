// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// (Optional).  The service group configuration when creating an instance.
//
// When creating an instance, either a previously created (persistent) service
// group can be referenced (either through its name or UUID), or a new
// (ephemeral) service group can be created for the instance by specifying the
// list of services it should expose and optionally the domains it should use.
// Not used by template instances.

type CreateInstanceRequestServiceGroup struct {
	// If no existing (persistent) service group is specified via its
	// identifier, a new (ephemeral) service group can be created.  In addition
	// to the services it must expose, you can specify which domains it should
	// use too.
	Domains []CreateInstanceRequestDomain `json:"domains,omitempty"`
	// If no existing service group identifier is provided, one or more new
	// (ephemeral, non-persistent) service(s) can be created with the following
	// definitions.
	Services []Service `json:"services,omitempty"`
	// The soft limit for the number of services that can be created in this
	// service group.
	SoftLimit *uint32 `json:"soft_limit,omitempty"`
	// The hard limit for the number of services that can be created in this
	// service group.
	HardLimit *uint32 `json:"hard_limit,omitempty"`
	// (Optional).  Reference an existing (persistent) service group by its
	// UUID.  Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// (Optional).  Reference an existing (persistent) service group by its
	// name.  Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *CreateInstanceRequestServiceGroup) UnmarshalJSON(data []byte) error {
	type Alias CreateInstanceRequestServiceGroup
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"domains":    {},
		"services":   {},
		"soft_limit": {},
		"hard_limit": {},
		"uuid":       {},
		"name":       {},
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

func (m CreateInstanceRequestServiceGroup) MarshalJSON() ([]byte, error) {
	type Alias CreateInstanceRequestServiceGroup
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

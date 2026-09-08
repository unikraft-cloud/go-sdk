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

// The service group configuration for the instance. This is a reference to the
// service group that the instance is part of. The service group defines the
// services (e.g. ports, connection handling) that the instance exposes and how
// they are configured.
type InstanceServiceGroup struct {
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
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

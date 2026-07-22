// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Information about an available region.
type Region struct {
	// The region identifier (e.g., "us-east-1", "us-central1").
	Name *string `json:"name,omitempty"`
	// Human-readable display name.
	DisplayName *string `json:"display_name,omitempty"`
	// The country code where this region is located.
	Country *string `json:"country,omitempty"`
	// Geographic coordinates of the region.
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	// Availability zones within this region.
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// Whether this region is currently available for provisioning.
	Available *bool `json:"available,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *Region) UnmarshalJSON(data []byte) error {
	type Alias Region
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Region) MarshalJSON() ([]byte, error) {
	type Alias Region
	return json.Marshal((Alias)(m))
}

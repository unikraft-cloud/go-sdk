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

// CloudProviderConfig contains provider-specific configuration for node
// provisioning.
type CloudProviderConfig struct {
	// AWS-specific configuration.
	Aws *AWSConfig `json:"aws,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *CloudProviderConfig) UnmarshalJSON(data []byte) error {
	type Alias CloudProviderConfig
	return json.Unmarshal(data, (*Alias)(m))
}

func (m CloudProviderConfig) MarshalJSON() ([]byte, error) {
	type Alias CloudProviderConfig
	return json.Marshal((Alias)(m))
}

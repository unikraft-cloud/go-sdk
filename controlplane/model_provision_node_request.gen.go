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

// Request message for creating a new node.
type ProvisionNodeRequest struct {
	// Optional name for the node. If not provided, a name will be
	// auto-generated.
	Name *string `json:"name,omitempty"`
	// The cloud provider where the machine should be provisioned.
	Provider ProvisionNodeRequestProvider `json:"provider"`
	// The provider region where the machine should be provisioned.
	Region string `json:"region"`
	// The machine type to provision. This is provider-specific.
	MachineType string `json:"machine_type"`
	// SSH keys for accessing the node. At least one key is required.
	SshKeys []SSHKey `json:"ssh_keys"`
	// Optional user-defined tags.
	Tags map[string]string `json:"tags,omitempty"`
	// Optional provider-specific configuration for advanced customization.
	ProviderConfig *NodeProviderConfig `json:"provider_config,omitempty"`
	// The Unikraft Cloud metro to associate this machine with.
	Metro *string `json:"metro,omitempty"`
	// Optional user overrides for platform configuration. Keys should be
	// kebab-case (e.g., "vm-user-min-memory").
	PlatformConfig map[string]string `json:"platform_config,omitempty"`
	// Optional image references to pre-pull on the provisioned machine.
	ImagePulls []string `json:"image_pulls,omitempty"`
	// Optional network count override. If not specified, defaults to 1000.
	NetCount *uint32 `json:"net_count,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *ProvisionNodeRequest) UnmarshalJSON(data []byte) error {
	type Alias ProvisionNodeRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ProvisionNodeRequest) MarshalJSON() ([]byte, error) {
	type Alias ProvisionNodeRequest
	return json.Marshal((Alias)(m))
}

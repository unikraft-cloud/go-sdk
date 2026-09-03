// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// Request message for creating a new node.
type ProvisionNodeRequest struct {
	// Optional name for the node. If not provided, a name will be
	// auto-generated.
	Name *string `json:"name,omitzero"`
	// The cloud provider where the machine should be provisioned.
	Cloudprovider *CloudProvider `json:"cloudprovider,omitzero"`
	// SSH keys for accessing the node. At least one key is required.
	SshKeys []SSHKey `json:"ssh_keys"`
	// Optional user-defined tags.
	Tags map[string]string `json:"tags,omitzero"`
	// Optional provider-specific configuration for advanced customization.
	CloudproviderConfig *CloudProviderConfig `json:"cloudprovider_config,omitzero"`
	// The Unikraft Cloud metro to associate this machine with.
	Metro *string `json:"metro,omitzero"`
	// Optional user overrides for platform configuration. Keys should be
	// kebab-case (e.g., "vm-user-min-memory").
	PlatformConfig map[string]string `json:"platform_config,omitzero"`
	// Optional image references to pre-pull on the provisioned machine.
	ImagePulls []string `json:"image_pulls,omitzero"`
	// Optional network count override. If not specified, defaults to 1000.
	NetCount *uint32 `json:"net_count,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ProvisionNodeRequest) UnmarshalJSON(data []byte) error {
	type Alias ProvisionNodeRequest
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ProvisionNodeRequest) MarshalJSON() ([]byte, error) {
	type Alias ProvisionNodeRequest
	return json.Marshal((Alias)(m))
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"time"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Node represents a physical or virtual compute node provisioned on a
// cloud provider for use as part of the Unikraft Cloud infrastructure.
//
// Node nodes are the underlying compute resources where Unikraft Cloud
// instances run.  They are provisioned on-demand across multiple cloud
// providers and managed through a unified API.
type Node struct {
	// The UUID of the machine.
	//
	// This is a unique identifier for the machine that is generated when the
	// machine is created. The UUID is used to reference the machine in API calls and
	// can be used to identify the machine in all API calls that require a machine
	// identifier.
	Uuid string `json:"uuid"`
	// The name of the machine.
	//
	// This is a human-readable name that can be used to identify the machine.
	// The name must be unique within the context of your account. The name can
	// also be used to identify the machine in API calls.
	Name string `json:"name"`
	// The time the machine was created.
	CreatedAt time.Time `json:"created_at"`
	// The time the machine was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The current state of the machine.
	State NodeState `json:"state"`
	// An optional message providing additional information about the current
	// state, particularly useful for error states.
	StateMessage *string `json:"state_message,omitempty"`
	// The cloud provider where this machine is provisioned.
	Cloudprovider *CloudProvider `json:"cloudprovider,omitempty"`
	// The number of vCPUs available on this machine.
	Vcpus uint32 `json:"vcpus"`
	// The amount of memory in MiB available on this machine.
	MemoryMib uint64 `json:"memory_mib"`
	// The SSH keys configured for access to this machine.
	SshKeys []SSHKey `json:"ssh_keys,omitempty"`
	// The public IPv4 address of the machine, if assigned.
	PublicIpv4 *string `json:"public_ipv4,omitempty"`
	// The public IPv6 address of the machine, if assigned.
	PublicIpv6 *string `json:"public_ipv6,omitempty"`
	// The private IPv4 address of the machine within the provider's network.
	PrivateIpv4 *string `json:"private_ipv4,omitempty"`
	// The Unikraft Cloud metro this machine is associated with.
	Metro *string `json:"metro,omitempty"`
	// The provider-specific instance ID or resource identifier.
	ProviderInstanceId *string `json:"provider_instance_id,omitempty"`
	// Provider-specific configuration that was used to provision this node.
	ProviderConfig *CloudProviderConfig `json:"provider_config,omitempty"`
	// User-defined tags for organizing and filtering nodes.
	Tags map[string]string `json:"tags,omitempty"`
	// The time when the machine became ready (entered READY state).
	ReadyAt *time.Time `json:"ready_at,omitempty"`
	// The total uptime of the machine in seconds since it became ready.
	UptimeSeconds uint64 `json:"uptime_seconds"`
	// Whether the machine is protected from deletion. When true, delete operations
	// will fail until this is set to false.
	DeleteLock bool `json:"delete_lock"`
	// The API endpoint for the node, used to connect to the node's platform SDK.
	ApiEndpoint *string `json:"api_endpoint,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *Node) UnmarshalJSON(data []byte) error {
	type Alias Node
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Node) MarshalJSON() ([]byte, error) {
	type Alias Node
	return json.Marshal((Alias)(m))
}

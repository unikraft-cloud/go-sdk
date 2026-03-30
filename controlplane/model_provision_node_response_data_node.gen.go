// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"encoding/json"
	"time"
)

// The created node.
// The current state of the machine.
type ProvisionNodeResponseDataNodeState string

const (
	ProvisionNodeResponseDataNodeStateUnspecified    ProvisionNodeResponseDataNodeState = "unspecified"
	ProvisionNodeResponseDataNodeStatePending        ProvisionNodeResponseDataNodeState = "pending"
	ProvisionNodeResponseDataNodeStateProvisioning   ProvisionNodeResponseDataNodeState = "provisioning"
	ProvisionNodeResponseDataNodeStateConfiguring    ProvisionNodeResponseDataNodeState = "configuring"
	ProvisionNodeResponseDataNodeStateReady          ProvisionNodeResponseDataNodeState = "ready"
	ProvisionNodeResponseDataNodeStateDeprovisioning ProvisionNodeResponseDataNodeState = "deprovisioning"
	ProvisionNodeResponseDataNodeStateDeprovisioned  ProvisionNodeResponseDataNodeState = "deprovisioned"
	ProvisionNodeResponseDataNodeStateError          ProvisionNodeResponseDataNodeState = "error"
	ProvisionNodeResponseDataNodeStateMaintenance    ProvisionNodeResponseDataNodeState = "maintenance"
)

// The cloud provider where this machine is provisioned.
type ProvisionNodeResponseDataNodeProvider string

const (
	ProvisionNodeResponseDataNodeProviderUnspecified ProvisionNodeResponseDataNodeProvider = "unspecified"
	ProvisionNodeResponseDataNodeProviderAws         ProvisionNodeResponseDataNodeProvider = "aws"
)

type ProvisionNodeResponseDataNode struct {
	// The UUID of the machine.
	//
	// This is a unique identifier for the machine that is generated when the
	// machine is created. The UUID is used to reference the machine in API calls and
	// can be used to identify the machine in all API calls that require a machine
	// identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the machine.
	//
	// This is a human-readable name that can be used to identify the machine.
	// The name must be unique within the context of your account. The name can
	// also be used to identify the machine in API calls.
	Name *string `json:"name,omitempty"`
	// The time the machine was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time the machine was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The current state of the machine.
	State *ProvisionNodeResponseDataNodeState `json:"state,omitempty"`
	// An optional message providing additional information about the current
	// state, particularly useful for error states.
	StateMessage *string `json:"state_message,omitempty"`
	// The cloud provider where this machine is provisioned.
	Provider *ProvisionNodeResponseDataNodeProvider `json:"provider,omitempty"`
	// The provider's region where the machine is located (e.g., "us-east-1" for
	// AWS, "us-central1" for GCP, "westeurope" for Azure).
	Region *string `json:"region,omitempty"`
	// The machine type as defined by the provider (e.g., "m5.xlarge" for AWS,
	// "n2-standard-4" for GCP, "Standard_D4s_v3" for Azure).
	//
	// This determines the compute resources (CPU, memory, etc.) available on
	// the machine. The valid values depend on the chosen provider.
	MachineType *string `json:"machine_type,omitempty"`
	// The number of vCPUs available on this machine.
	Vcpus *uint32 `json:"vcpus,omitempty"`
	// The amount of memory in MiB available on this machine.
	MemoryMib *uint64 `json:"memory_mib,omitempty"`
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
	ProviderInstanceId *string             `json:"provider_instance_id,omitempty"`
	ProviderConfig     *NodeProviderConfig `json:"provider_config,omitempty"`
	// User-defined tags for organizing and filtering nodes.
	Tags map[string]string `json:"tags,omitempty"`
	// The time when the machine became ready (entered READY state).
	ReadyAt *time.Time `json:"ready_at,omitempty"`
	// The total uptime of the machine in seconds since it became ready.
	UptimeSeconds *uint64 `json:"uptime_seconds,omitempty"`
	// Whether the machine is protected from deletion. When true, delete operations
	// will fail until this is set to false.
	DeleteLock *bool `json:"delete_lock,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *ProvisionNodeResponseDataNode) UnmarshalJSON(data []byte) error {
	type Alias ProvisionNodeResponseDataNode
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"uuid":                 {},
		"name":                 {},
		"created_at":           {},
		"updated_at":           {},
		"state":                {},
		"state_message":        {},
		"provider":             {},
		"region":               {},
		"machine_type":         {},
		"vcpus":                {},
		"memory_mib":           {},
		"ssh_keys":             {},
		"public_ipv4":          {},
		"public_ipv6":          {},
		"private_ipv4":         {},
		"metro":                {},
		"provider_instance_id": {},
		"provider_config":      {},
		"tags":                 {},
		"ready_at":             {},
		"uptime_seconds":       {},
		"delete_lock":          {},
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

func (m ProvisionNodeResponseDataNode) MarshalJSON() ([]byte, error) {
	type Alias ProvisionNodeResponseDataNode
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

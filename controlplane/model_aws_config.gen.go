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

// AWSConfig contains AWS-specific configuration for node provisioning.
type AWSConfig struct {
	// The VPC ID where the instance will be launched. If not specified, the
	// default VPC for the region will be used.
	VpcId *string `json:"vpc_id,omitzero"`
	// The subnet ID where the instance will be launched. If not specified,
	// a subnet will be selected from the VPC.
	SubnetId *string `json:"subnet_id,omitzero"`
	// Security group IDs to attach to the instance. If not specified, the
	// default security group for the VPC will be used.
	SecurityGroupIds []string `json:"security_group_ids,omitzero"`
	// The IAM instance profile ARN or name to attach to the instance.
	IamInstanceProfile *string `json:"iam_instance_profile,omitzero"`
	// Root EBS volume configuration. If not specified, provider defaults
	// will be used.
	RootVolume *AWSEBSConfig `json:"root_volume,omitzero"`
	// Additional EBS volumes to attach to the instance.
	AdditionalVolumes []AWSEBSConfig `json:"additional_volumes,omitzero"`
	// Whether to use a dedicated host. Dedicated hosts provide visibility and
	// control over how instances are placed on physical servers.
	DedicatedHost bool `json:"dedicated_host"`
	// Placement group name for the instance. Placement groups influence how
	// instances are placed on underlying hardware.
	PlacementGroup *string `json:"placement_group,omitzero"`
	// The AWS region where the machine is located.
	Region string `json:"region"`
	// The AWS machine type. This determines the compute resources
	// (CPU, memory, etc.) available on the machine. The valid values depend
	// on the chosen provider.
	MachineType string `json:"machine_type"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AWSConfig) UnmarshalJSON(data []byte) error {
	type Alias AWSConfig
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AWSConfig) MarshalJSON() ([]byte, error) {
	type Alias AWSConfig
	return json.Marshal((Alias)(m))
}

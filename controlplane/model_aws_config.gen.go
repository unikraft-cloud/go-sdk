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

// AWSConfig contains AWS-specific configuration for node provisioning.
type AWSConfig struct {
	// The AWS availability zone (e.g., "us-east-1a"). If not specified, AWS
	// will select an availability zone within the region.
	AvailabilityZone *string `json:"availability_zone,omitempty"`
	// The VPC ID where the instance will be launched. If not specified, the
	// default VPC for the region will be used.
	VpcId *string `json:"vpc_id,omitempty"`
	// The subnet ID where the instance will be launched. If not specified,
	// a subnet will be selected from the VPC.
	SubnetId *string `json:"subnet_id,omitempty"`
	// Security group IDs to attach to the instance. If not specified, the
	// default security group for the VPC will be used.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// The IAM instance profile ARN or name to attach to the instance.
	IamInstanceProfile *string `json:"iam_instance_profile,omitempty"`
	// Root EBS volume configuration. If not specified, provider defaults
	// will be used.
	RootVolume *AWSEBSConfig `json:"root_volume,omitempty"`
	// Additional EBS volumes to attach to the instance.
	AdditionalVolumes []AWSEBSConfig `json:"additional_volumes,omitempty"`
	// Whether to use a dedicated host. Dedicated hosts provide visibility and
	// control over how instances are placed on physical servers.
	DedicatedHost *bool `json:"dedicated_host,omitempty"`
	// Placement group name for the instance. Placement groups influence how
	// instances are placed on underlying hardware.
	PlacementGroup *string `json:"placement_group,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *AWSConfig) UnmarshalJSON(data []byte) error {
	type Alias AWSConfig
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AWSConfig) MarshalJSON() ([]byte, error) {
	type Alias AWSConfig
	return json.Marshal((Alias)(m))
}

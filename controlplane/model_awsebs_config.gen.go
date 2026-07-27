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

// AWSEBSConfig defines the configuration for an AWS EBS volume.
type AWSEBSConfig struct {
	// The device name (e.g., "/dev/sda1", "/dev/xvdf"). Required for
	// additional volumes.
	DeviceName *string `json:"device_name,omitempty"`
	// Size of the volume in GiB.
	SizeGib uint32 `json:"size_gib"`
	// The volume type (gp3, gp2, io1, io2, st1, sc1, standard).
	VolumeType string `json:"volume_type"`
	// The number of IOPS. Only valid for io1, io2, and gp3 volumes.
	Iops *uint32 `json:"iops,omitempty"`
	// The throughput in MiB/s. Only valid for gp3 volumes.
	ThroughputMibps *uint32 `json:"throughput_mibps,omitempty"`
	// Whether the volume should be encrypted.
	Encrypted bool `json:"encrypted"`
	// The KMS key ID for encryption. If encrypted is true and this is not
	// specified, the default EBS encryption key will be used.
	KmsKeyId *string `json:"kms_key_id,omitempty"`
	// Whether to delete the volume when the instance is terminated.
	DeleteOnTermination bool `json:"delete_on_termination"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *AWSEBSConfig) UnmarshalJSON(data []byte) error {
	type Alias AWSEBSConfig
	return json.Unmarshal(data, (*Alias)(m))
}

func (m AWSEBSConfig) MarshalJSON() ([]byte, error) {
	type Alias AWSEBSConfig
	return json.Marshal((Alias)(m))
}

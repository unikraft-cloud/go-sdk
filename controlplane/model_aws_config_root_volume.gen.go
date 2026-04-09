// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// Root EBS volume configuration. If not specified, provider defaults
// will be used.

type AWSConfigRootVolume struct {
	// The device name (e.g., "/dev/sda1", "/dev/xvdf"). Required for
	// additional volumes.
	DeviceName *string `json:"device_name,omitempty"`
	// Size of the volume in GiB.
	SizeGib *uint32 `json:"size_gib,omitempty"`
	// The volume type (gp3, gp2, io1, io2, st1, sc1, standard).
	VolumeType *string `json:"volume_type,omitempty"`
	// The number of IOPS. Only valid for io1, io2, and gp3 volumes.
	Iops *uint32 `json:"iops,omitempty"`
	// The throughput in MiB/s. Only valid for gp3 volumes.
	ThroughputMibps *uint32 `json:"throughput_mibps,omitempty"`
	// Whether the volume should be encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The KMS key ID for encryption. If encrypted is true and this is not
	// specified, the default EBS encryption key will be used.
	KmsKeyId *string `json:"kms_key_id,omitempty"`
	// Whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *AWSConfigRootVolume) UnmarshalJSON(data []byte) error {
	type Alias AWSConfigRootVolume
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"device_name":           {},
		"size_gib":              {},
		"volume_type":           {},
		"iops":                  {},
		"throughput_mibps":      {},
		"encrypted":             {},
		"kms_key_id":            {},
		"delete_on_termination": {},
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

func (m AWSConfigRootVolume) MarshalJSON() ([]byte, error) {
	type Alias AWSConfigRootVolume
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

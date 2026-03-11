// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

type Image struct {
	// The digest of the image is a unique identifier of the image manifest which
	// is a string representation including the hashing algorithm and the hash
	// value separated by a colon.
	Digest *string `json:"digest,omitempty"`
	// The canonical name of the image is known as the "tag".
	Tags        []string          `json:"tags,omitempty"`
	Initrd      *bool             `json:"initrd,omitempty"`
	SizeInBytes *int64            `json:"size_in_bytes,omitempty"`
	Args        *string           `json:"args,omitempty"`
	KernelArgs  *string           `json:"kernel_args,omitempty"`
	Users       []string          `json:"users,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *Image) UnmarshalJSON(data []byte) error {
	type Alias Image
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"digest":        {},
		"tags":          {},
		"initrd":        {},
		"size_in_bytes": {},
		"args":          {},
		"kernel_args":   {},
		"users":         {},
		"labels":        {},
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

func (m Image) MarshalJSON() ([]byte, error) {
	type Alias Image
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

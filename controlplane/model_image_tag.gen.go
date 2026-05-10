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

type ImageTag struct {
	// The tag name.
	Name string `json:"name"`
	// The digest for the tag.
	Digest string `json:"digest"`
	// The size of the image in bytes.
	Size uint64 `json:"size"`
	// The push time of the image.
	PushTime time.Time `json:"push_time"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *ImageTag) UnmarshalJSON(data []byte) error {
	type Alias ImageTag
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":      {},
		"digest":    {},
		"size":      {},
		"push_time": {},
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

func (m ImageTag) MarshalJSON() ([]byte, error) {
	type Alias ImageTag
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

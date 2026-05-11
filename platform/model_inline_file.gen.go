// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// An inline file entry represents a single file within an image.

type InlineFile struct {
	// The file path within the image.
	Path string `json:"path"`
	// (Optional).  The encoding of the data field.  Defaults to "text".
	Encoding *InlineDataEncoding `json:"encoding,omitempty"`
	// The file data, encoded according to the encoding field.
	Data string `json:"data"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *InlineFile) UnmarshalJSON(data []byte) error {
	type Alias InlineFile
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"path":     {},
		"encoding": {},
		"data":     {},
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

func (m InlineFile) MarshalJSON() ([]byte, error) {
	type Alias InlineFile
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

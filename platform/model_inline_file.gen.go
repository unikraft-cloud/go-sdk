// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// An inline file entry represents a single file within an image.
type InlineFile struct {
	// The file path within the image.
	Path string `json:"path"`
	// (Optional).  The encoding of the data field.  Defaults to "text".
	Encoding *InlineDataEncoding `json:"encoding,omitempty"`
	// The file data, encoded according to the encoding field.
	Data string `json:"data"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *InlineFile) UnmarshalJSON(data []byte) error {
	type Alias InlineFile
	return json.Unmarshal(data, (*Alias)(m))
}

func (m InlineFile) MarshalJSON() ([]byte, error) {
	type Alias InlineFile
	return json.Marshal((Alias)(m))
}

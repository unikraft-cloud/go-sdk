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

type ImageTag struct {
	// The tag name.
	Name string `json:"name"`
	// The digest for the tag.
	Digest string `json:"digest"`
	// The size of the image in bytes.
	Size uint64 `json:"size"`
	// The push time of the image.
	PushTime time.Time `json:"push_time"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *ImageTag) UnmarshalJSON(data []byte) error {
	type Alias ImageTag
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ImageTag) MarshalJSON() ([]byte, error) {
	type Alias ImageTag
	return json.Marshal((Alias)(m))
}

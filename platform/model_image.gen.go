// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// An image representing a VM which can be deployed on Unikraft Cloud.
type Image struct {
	// The image URL.
	Url string `json:"url"`
	// The time the image was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether the image is an initrd or ROM.
	InitrdOrRom bool `json:"initrd_or_rom"`
	// The size of the image in bytes.
	SizeInBytes int64 `json:"size_in_bytes"`
	// Command-line arguments for the image.
	Args []string `json:"args,omitzero"`
	// Environment variables for the image.
	Env map[string]string `json:"env,omitzero"`
	// Tags associated with the image.
	Tags []string `json:"tags,omitzero"`
	// Users associated with the image.
	Users []string `json:"users,omitzero"`
	// Whether the image is pinned and exempt from cache eviction. Only
	// populated (and only ever `true`) for callers with image manager
	// permissions; omitted otherwise, including when the image is not pinned.
	Persistent *bool `json:"persistent,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Image) UnmarshalJSON(data []byte) error {
	type Alias Image
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Image) MarshalJSON() ([]byte, error) {
	type Alias Image
	return json.Marshal((Alias)(m))
}

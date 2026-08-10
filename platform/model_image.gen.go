// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"time"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type Image struct {
	Url string `json:"url"`
	// The time the volume was created.
	CreatedAt   time.Time         `json:"created_at"`
	InitrdOrRom bool              `json:"initrd_or_rom"`
	SizeInBytes int64             `json:"size_in_bytes"`
	Args        []string          `json:"args,omitzero"`
	Env         map[string]string `json:"env,omitzero"`
	Tags        []string          `json:"tags,omitzero"`
	Users       []string          `json:"users,omitzero"`

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

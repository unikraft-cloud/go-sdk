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

// ImageSpec describes an image to use for an instance.
type ImageSpec struct {
	// The image URL
	Url string `json:"url"`
	// Optional credentials for authenticating to an OCI registry.
	// Only valid for OCI registry URLs; the platform rejects this
	// field for non-OCI schemes.
	Credentials *string `json:"credentials,omitzero"`
	// Optional HTTP headers to send when fetching the image.
	Headers map[string]string `json:"headers,omitzero"`
	// Controls when the image is pulled relative to what is already cached on
	// the node.
	PullPolicy *PullPolicy `json:"pull_policy,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ImageSpec) UnmarshalJSON(data []byte) error {
	type Alias ImageSpec
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ImageSpec) MarshalJSON() ([]byte, error) {
	type Alias ImageSpec
	return json.Marshal((Alias)(m))
}

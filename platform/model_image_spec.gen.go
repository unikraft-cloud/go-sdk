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

// ImageSpec describes an image to use for an instance.
type ImageSpec struct {
	// The image URL
	Url string `json:"url"`
	// Optional credentials for authenticating to an OCI registry.
	// Only valid for OCI registry URLs; the platform rejects this
	// field for non-OCI schemes.
	Credentials *string `json:"credentials,omitempty"`
	// Optional HTTP headers to send when fetching the image.
	Headers map[string]string `json:"headers,omitempty"`
	// Controls when the image is pulled relative to what is already cached on
	// the node.
	PullPolicy *PullPolicy `json:"pull_policy,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *ImageSpec) UnmarshalJSON(data []byte) error {
	type Alias ImageSpec
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ImageSpec) MarshalJSON() ([]byte, error) {
	type Alias ImageSpec
	return json.Marshal((Alias)(m))
}

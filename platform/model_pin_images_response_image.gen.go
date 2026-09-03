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

// The result of pinning a single image.  On success, `uuid` through `tags`
// are set; on failure, only `message` and `error` are set (the image
// being pulled is not otherwise identified in the response).
type PinImagesResponseImage struct {
	// Indicates whether this image was pulled and pinned successfully.
	Status ResponseStatus `json:"status"`
	// The UUID of the image.  Only set on success.
	Uuid *string `json:"uuid,omitzero"`
	// The name of the image.  Only set on success.
	Name *string `json:"name,omitzero"`
	// The time the image was created.  Only set on success.
	CreatedAt *time.Time `json:"created_at,omitzero"`
	// The current state of the image (e.g. `ready`).  Only set on success.
	State *string `json:"state,omitzero"`
	// The image URL.  Only set on success.
	Url *string `json:"url,omitzero"`
	// Whether the image is pinned and exempt from cache eviction.  Only set
	// on success, where it is always `true`.
	Persistent *bool `json:"persistent,omitzero"`
	// The tags associated with the image.  Only set on success.
	Tags []string `json:"tags,omitzero"`
	// Set when the image could not be pulled; carries the agent's error, if
	// any (e.g. a registry connectivity issue).
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the
	// status.  This field is only set when the status is not `success`.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *PinImagesResponseImage) UnmarshalJSON(data []byte) error {
	type Alias PinImagesResponseImage
	return json.Unmarshal(data, (*Alias)(m))
}

func (m PinImagesResponseImage) MarshalJSON() ([]byte, error) {
	type Alias PinImagesResponseImage
	return json.Marshal((Alias)(m))
}

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

// The result of unpinning a single image.
type UnpinImagesResponseImage struct {
	// Indicates whether the operation was successful for this item.
	Status ResponseStatus `json:"status"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the image.
	Uuid string `json:"uuid"`
	// The name of the image. Only set on success, and only if the image
	// has a name.
	Name *string `json:"name,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UnpinImagesResponseImage) UnmarshalJSON(data []byte) error {
	type Alias UnpinImagesResponseImage
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UnpinImagesResponseImage) MarshalJSON() ([]byte, error) {
	type Alias UnpinImagesResponseImage
	return json.Marshal((Alias)(m))
}

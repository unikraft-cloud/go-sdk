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

// The request item for unpinning a single image. Select the image by exactly
// one of `uuid` or `url`.
type UnpinImageRequestItem struct {
	// The UUID of the image to unpin.
	Uuid *string `json:"uuid,omitzero"`
	// The URL of the image to unpin, by tag or by digest. A URL without a
	// digest selects the image the tag currently resolves to on the node.
	Url *string `json:"url,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UnpinImageRequestItem) UnmarshalJSON(data []byte) error {
	type Alias UnpinImageRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UnpinImageRequestItem) MarshalJSON() ([]byte, error) {
	type Alias UnpinImageRequestItem
	return json.Marshal((Alias)(m))
}

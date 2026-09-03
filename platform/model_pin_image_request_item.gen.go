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

// The request item for pinning a single image.
type PinImageRequestItem struct {
	// The image URL to pull and pin.
	Url string `json:"url"`
	// Optional credentials for authenticating to an OCI registry.
	// Only valid for OCI registry URLs; the platform rejects this
	// field for non-OCI schemes.
	Credentials *string `json:"credentials,omitzero"`
	// Optional HTTP headers to send when fetching the image.
	Headers map[string]string `json:"headers,omitzero"`
	// Controls when the image is pulled relative to what is already cached on
	// the node.  If unset, this is inferred from the URL.
	PullPolicy *PullPolicy `json:"pull_policy,omitzero"`
	// Number of seconds to wait for the pull to complete.  Required and must
	// be non-zero; `-1` waits up to the platform's maximum timeout.
	TimeoutS int64 `json:"timeout_s"`
	// Avoid duplicate pulls by merging with any in-flight request for the
	// same image.  Defaults to `true`.
	MergeRequests *bool `json:"merge_requests,omitzero"`
	// (Optional).  Automatically unpin the image after a period of inactivity.
	Autokill *PinImageRequestItemAutokill `json:"autokill,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *PinImageRequestItem) UnmarshalJSON(data []byte) error {
	type Alias PinImageRequestItem
	return json.Unmarshal(data, (*Alias)(m))
}

func (m PinImageRequestItem) MarshalJSON() ([]byte, error) {
	type Alias PinImageRequestItem
	return json.Marshal((Alias)(m))
}

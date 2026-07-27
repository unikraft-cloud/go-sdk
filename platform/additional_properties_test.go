// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-json-experiment/json/jsontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform"
)

func TestImage_AdditionalProperties_Unmarshal(t *testing.T) {
	data := []byte(`{
		"url": "index.unikraft.io/example:latest",
		"created_at": "2026-01-01T00:00:00Z",
		"initrd_or_rom": false,
		"size_in_bytes": 12345,
		"region": "fra0",
		"extra": {"nested": true}
	}`)

	var img platform.Image
	require.NoError(t, json.Unmarshal(data, &img))

	assert.Equal(t, "index.unikraft.io/example:latest", img.Url)
	assert.Equal(t, int64(12345), img.SizeInBytes)

	require.NotEmpty(t, img.AdditionalProperties, "unknown fields were not captured")

	wantProps := map[string]string{
		"region": `"fra0"`,
		"extra":  `{"nested": true}`,
	}
	for key, want := range wantProps {
		got, ok := img.AdditionalProperties[key]
		require.True(t, ok, "AdditionalProperties missing key %q", key)
		assert.True(t, jsontext.Value(got).IsValid(), "AdditionalProperties[%q] is not valid JSON: %s", key, got)
		assert.JSONEq(t, want, string(got), "AdditionalProperties[%q] mismatch", key)
	}
}

func TestImage_AdditionalProperties_Marshal(t *testing.T) {
	img := platform.Image{
		Url:         "index.unikraft.io/example:latest",
		CreatedAt:   time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		SizeInBytes: 42,
		AdditionalProperties: map[string]jsontext.Value{
			"region": jsontext.Value(`"fra0"`),
			"extra":  jsontext.Value(`{"nested":true}`),
		},
	}

	out, err := json.Marshal(img)
	require.NoError(t, err)

	var raw map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(out, &raw))

	assert.NotContains(t, raw, "AdditionalProperties", `marshalled JSON contains a literal "AdditionalProperties" key; fields must be inlined`)

	for _, key := range []string{"region", "extra"} {
		assert.Contains(t, raw, key, "marshalled JSON is missing top-level key %q", key)
	}
}

func TestImage_AdditionalProperties_RoundTrip(t *testing.T) {
	original := platform.Image{
		Url:         "index.unikraft.io/example:latest",
		CreatedAt:   time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		SizeInBytes: 99,
		AdditionalProperties: map[string]jsontext.Value{
			"region": jsontext.Value(`"fra0"`),
		},
	}

	out, err := json.Marshal(original)
	require.NoError(t, err)

	var decoded platform.Image
	require.NoError(t, json.Unmarshal(out, &decoded))

	assert.Equal(t, original.Url, decoded.Url)
	assert.Equal(t, original.SizeInBytes, decoded.SizeInBytes)
	got, ok := decoded.AdditionalProperties["region"]
	require.True(t, ok, `AdditionalProperties missing "region" after round-trip`)
	assert.Equal(t, `"fra0"`, string(got))
}

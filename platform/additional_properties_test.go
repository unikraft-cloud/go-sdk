// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-json-experiment/json/jsontext"

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
	if err := json.Unmarshal(data, &img); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if img.Url != "index.unikraft.io/example:latest" {
		t.Errorf("Url: got %q, want %q", img.Url, "index.unikraft.io/example:latest")
	}
	if img.SizeInBytes != 12345 {
		t.Errorf("SizeInBytes: got %d, want 12345", img.SizeInBytes)
	}

	if len(img.AdditionalProperties) == 0 {
		t.Fatal("AdditionalProperties is empty; unknown fields were not captured")
	}

	wantProps := map[string]string{
		"region": `"fra0"`,
		"extra":  `{"nested": true}`,
	}
	for key, want := range wantProps {
		got, ok := img.AdditionalProperties[key]
		if !ok {
			t.Errorf("AdditionalProperties missing key %q", key)
			continue
		}
		if !jsontext.Value(got).IsValid() {
			t.Errorf("AdditionalProperties[%q] is not valid JSON: %s", key, got)
		}
		_ = want
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
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(out, &raw); err != nil {
		t.Fatalf("could not decode marshalled output: %v", err)
	}

	if _, bad := raw["AdditionalProperties"]; bad {
		t.Error(`marshalled JSON contains a literal "AdditionalProperties" key; fields must be inlined`)
	}

	for _, key := range []string{"region", "extra"} {
		if _, ok := raw[key]; !ok {
			t.Errorf("marshalled JSON is missing top-level key %q", key)
		}
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
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var decoded platform.Image
	if err := json.Unmarshal(out, &decoded); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if decoded.Url != original.Url {
		t.Errorf("Url: got %q, want %q", decoded.Url, original.Url)
	}
	if decoded.SizeInBytes != original.SizeInBytes {
		t.Errorf("SizeInBytes: got %d, want %d", decoded.SizeInBytes, original.SizeInBytes)
	}
	if got, ok := decoded.AdditionalProperties["region"]; !ok {
		t.Error(`AdditionalProperties missing "region" after round-trip`)
	} else if string(got) != `"fra0"` {
		t.Errorf(`AdditionalProperties["region"]: got %s, want "fra0"`, got)
	}
}

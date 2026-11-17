// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseImageRef(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ref     string
		want    imageRef
		wantErr bool
	}{
		{ref: "nginx", want: imageRef{namespace: "official", repository: "official/nginx", tag: "latest"}},
		{ref: "org/app:1.2.3", want: imageRef{namespace: "org", repository: "org/app", tag: "1.2.3"}},
		{ref: "index.unikraft.io/org/app", want: imageRef{namespace: "org", repository: "org/app", tag: "latest"}},
		{ref: "https://localhost:5000/org/app:v1", want: imageRef{namespace: "org", repository: "org/app", tag: "v1"}},
		{ref: "org/team/app:v1", want: imageRef{namespace: "org", repository: "org/team/app", tag: "v1"}},
		{ref: "org/app@sha256:ABCDEF", want: imageRef{namespace: "org", repository: "org/app", digest: "sha256:abcdef"}},
		{ref: "org/", wantErr: true},
		{ref: "org/app:", wantErr: true},
		{ref: "org/app@nodigest", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			t.Parallel()

			got, err := parseImageRef(tt.ref)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClientFindImage(t *testing.T) {
	t.Parallel()

	images := map[string]any{
		"status": "success",
		"data": map[string]any{"images": []map[string]any{
			{"name": "org/other", "tags": []map[string]any{{"name": "v1", "digest": "sha256:aa"}}},
			{"name": "org/app", "tags": []map[string]any{
				{"name": "v1", "digest": "sha256:aa", "size": 10, "push_time": "2026-01-02T03:04:05Z"},
				{"name": "v2", "digest": "sha256:bb", "size": 20, "push_time": "2026-01-02T03:04:05Z"},
			}},
		}},
	}

	tests := []struct {
		ref     string
		wantTag string
	}{
		{ref: "org/app:v2", wantTag: "v2"},
		{ref: "org/app@sha256:AA", wantTag: "v1"},
		{ref: "org/app:v3"},
		{ref: "org/app"},
		{ref: "org/missing:v1"},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, images)
			})
			c := newClientAgainst(srv)

			image, err := c.FindImage(t.Context(), tt.ref)
			require.NoError(t, err)
			assert.Equal(t, "/v1/images", (*reqs)[0].Path)
			assert.Equal(t, "details=true&namespace=org", (*reqs)[0].Query)

			exists, err := c.ImageExists(t.Context(), tt.ref)
			require.NoError(t, err)
			assert.Equal(t, tt.wantTag != "", exists)

			if tt.wantTag == "" {
				assert.Nil(t, image)
				return
			}
			require.NotNil(t, image)
			assert.Equal(t, "org/app:"+tt.wantTag, image.Reference())
		})
	}
}

func TestClientFindImageRejectsBadRef(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		t.Fatalf("unexpected request %s", r.Path)
	})

	_, err := newClientAgainst(srv).FindImage(t.Context(), "org/app:")
	assert.Error(t, err)
	assert.Empty(t, *reqs)
}

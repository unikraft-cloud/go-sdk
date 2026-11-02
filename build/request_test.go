// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/build"
)

// Every form of a Kraftfile is a Kraftfile.
var (
	_ build.Kraftfile = build.KraftfileInline("")
	_ build.Kraftfile = build.Source{}
	_ build.Kraftfile = build.KraftfileObject{}
)

func TestSource(t *testing.T) {
	t.Parallel()

	t.Run("encodes", func(t *testing.T) {
		t.Parallel()

		for name, tc := range map[string]struct {
			source build.Source
			want   string
		}{
			"a repository without a ref is its default branch": {
				source: build.Git("https://github.com/acme/app.git", ""),
				want:   `{"git": "https://github.com/acme/app.git"}`,
			},
			"a repository narrowed to a directory and a file": {
				source: build.Git("https://github.com/acme/app.git", "v1.2.3").WithSubdir("app").WithFilename("Kraftfile.prod"),
				want:   `{"git": "https://github.com/acme/app.git", "ref": "v1.2.3", "subdir": "app", "filename": "Kraftfile.prod"}`,
			},
			"a tarball narrowed to a directory": {
				source: build.URL("https://example.com/task.tar.gz").WithSubdir("environment"),
				want:   `{"url": "https://example.com/task.tar.gz", "subdir": "environment"}`,
			},
		} {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				b, err := json.Marshal(tc.source)
				require.NoError(t, err)
				assert.JSONEq(t, tc.want, string(b))
			})
		}
	})
}

func TestKraftfile(t *testing.T) {
	t.Parallel()

	t.Run("an inline file is a string", func(t *testing.T) {
		t.Parallel()

		b, err := json.Marshal(build.KraftfileInline("spec: v0.7\nrootfs: ./Dockerfile\n"))
		require.NoError(t, err)
		assert.JSONEq(t, `"spec: v0.7\nrootfs: ./Dockerfile\n"`, string(b))
	})

	t.Run("a source is an object with its file name", func(t *testing.T) {
		t.Parallel()

		b, err := json.Marshal(build.KraftfileFrom(build.URL("https://example.com/app.tar.gz").WithFilename("Kraftfile.prod")))
		require.NoError(t, err)
		assert.JSONEq(t, `{"url": "https://example.com/app.tar.gz", "filename": "Kraftfile.prod"}`, string(b))
	})

	t.Run("a parsed file is its JSON form", func(t *testing.T) {
		t.Parallel()

		kf, err := build.KraftfileObjectOf(struct {
			Runtime string   `json:"runtime"`
			Cmd     []string `json:"cmd"`
		}{
			Runtime: "unikraft.org/base-compat:latest",
			Cmd:     []string{"/usr/bin/node", "server.js"},
		})
		require.NoError(t, err)

		b, err := json.Marshal(kf)
		require.NoError(t, err)
		assert.JSONEq(t, `{"runtime": "unikraft.org/base-compat:latest", "cmd": ["/usr/bin/node", "server.js"]}`, string(b))
	})

	t.Run("a value that is not an object is refused", func(t *testing.T) {
		t.Parallel()

		_, err := build.KraftfileObjectOf("spec: v0.7")
		require.Error(t, err)
	})

	t.Run("nothing is refused", func(t *testing.T) {
		t.Parallel()

		_, err := build.KraftfileObjectOf(nil)
		require.Error(t, err)

		_, err = build.KraftfileObjectOf((*struct{ Runtime string })(nil))
		require.Error(t, err)
	})

	t.Run("an object with a source key is refused", func(t *testing.T) {
		t.Parallel()

		_, err := build.KraftfileObjectOf(map[string]any{"git": "https://github.com/acme/app.git"})
		require.Error(t, err)

		_, err = build.KraftfileObjectOf(map[string]any{"url": "https://example.com/app.tar.gz", "runtime": "base:latest"})
		require.Error(t, err)

		_, err = build.KraftfileObjectOf(map[string]any{"session": "context"})
		require.Error(t, err)
	})
}

func TestCreateBuildRequestJSON(t *testing.T) {
	t.Parallel()

	src := build.Git("https://github.com/acme/app.git", "main").WithSubdir("app")
	body := build.CreateBuildRequest{
		Contexts: build.BuildContexts{Context: src},
		Output:   []build.BuildOutput{build.Image("index.unikraft.io/acme/app:1.2.3")},
	}

	for name, kf := range map[string]build.Kraftfile{
		"an inline file":   build.KraftfileInline("spec: v0.7\nruntime: base:latest\n"),
		"a source":         build.KraftfileFrom(src.WithFilename("Kraftfile.prod")),
		"a session source": build.KraftfileFrom(build.Source{Session: new("context")}),
		"a parsed file":    build.KraftfileObject{"runtime": "base:latest", "cmd": []any{"/server"}},
		"no Kraftfile":     nil,
	} {
		t.Run(name+" round-trips", func(t *testing.T) {
			t.Parallel()

			want := body
			want.Kraftfile = kf

			b, err := json.Marshal(want)
			require.NoError(t, err)

			var got build.CreateBuildRequest
			require.NoError(t, json.Unmarshal(b, &got))
			assert.Equal(t, want, got)
		})
	}

	t.Run("a source with a Kraftfile key is refused", func(t *testing.T) {
		t.Parallel()

		var got build.CreateBuildRequest
		err := json.Unmarshal([]byte(`{"kraftfile": {"url": "https://example.com/app.tar.gz", "runtime": "base:latest"}}`), &got)
		require.ErrorContains(t, err, "kraftfile")
	})

	t.Run("a Kraftfile that is neither a string nor an object is refused", func(t *testing.T) {
		t.Parallel()

		var got build.CreateBuildRequest
		err := json.Unmarshal([]byte(`{"kraftfile": 42}`), &got)
		require.ErrorContains(t, err, "kraftfile")
	})
}

func TestBuildContexts(t *testing.T) {
	t.Parallel()

	t.Run("every context is a key of one object, and round-trips", func(t *testing.T) {
		t.Parallel()

		want := build.BuildContexts{
			Context:    build.Git("https://github.com/acme/app.git", "main"),
			Dockerfile: new(build.Git("https://github.com/acme/app.git", "main").WithSubdir("docker").WithFilename("Dockerfile.prod")),
			Named:      map[string]build.Source{"base": build.URL("https://example.com/base.tar.gz")},
		}

		b, err := json.Marshal(want)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"context": {"git": "https://github.com/acme/app.git", "ref": "main"},
			"dockerfile": {"git": "https://github.com/acme/app.git", "ref": "main", "subdir": "docker", "filename": "Dockerfile.prod"},
			"base": {"url": "https://example.com/base.tar.gz"}
		}`, string(b))

		var got build.BuildContexts
		require.NoError(t, json.Unmarshal(b, &got))
		assert.Equal(t, want, got)
	})

	t.Run("a named context with the name of a context of the frontend is refused", func(t *testing.T) {
		t.Parallel()

		for _, name := range []string{"context", "dockerfile"} {
			_, err := json.Marshal(build.BuildContexts{
				Context: build.URL("https://example.com/app.tar.gz"),
				Named:   map[string]build.Source{name: build.URL("https://example.com/other.tar.gz")},
			})
			require.ErrorContains(t, err, name)
		}
	})
}

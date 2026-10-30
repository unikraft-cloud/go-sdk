// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/build"
)

// testTimeout bounds every test, so that a regression that hangs fails here
// rather than at go test's own timeout.
const testTimeout = 10 * time.Second

// testContext returns a context that ends with the test, or after
// testTimeout, whichever comes first.
func testContext(t *testing.T) context.Context {
	t.Helper()

	ctx, cancel := context.WithTimeout(t.Context(), testTimeout)
	t.Cleanup(cancel)

	return ctx
}

// request is what a test server saw of a request.
type request struct {
	method  string
	path    string
	query   string
	headers http.Header
	body    string
}

// answer is what a test server answers every request with.
type answer struct {
	status      int
	body        string
	contentType string // application/json when body is set, unless given
	location    string // the Location header, for a redirect
}

// serve starts a server that records every request it gets into requests and
// answers it with a, and returns a client of it.
func serve(t *testing.T, requests chan<- request, a answer) *build.Client {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		requests <- request{
			method:  r.Method,
			path:    r.URL.Path,
			query:   r.URL.RawQuery,
			headers: r.Header.Clone(),
			body:    string(b),
		}

		switch {
		case a.contentType != "":
			w.Header().Set("Content-Type", a.contentType)
		case a.body != "":
			w.Header().Set("Content-Type", "application/json")
		}
		if a.location != "" {
			w.Header().Set("Location", a.location)
		}
		w.WriteHeader(a.status)
		fmt.Fprint(w, a.body)
	}))
	t.Cleanup(srv.Close)

	c, err := build.New(srv.URL, build.WithToken("tok"))
	require.NoError(t, err)

	return c
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("refuses a server without a scheme", func(t *testing.T) {
		t.Parallel()

		_, err := build.New("build.unikraft.cloud")
		require.Error(t, err)
	})

	t.Run("refuses a server that does not parse", func(t *testing.T) {
		t.Parallel()

		_, err := build.New("://nope")
		require.Error(t, err)
	})

	t.Run("refuses a nil HTTP client", func(t *testing.T) {
		t.Parallel()

		_, err := build.New("https://build.unikraft.cloud", build.WithHTTPClient(nil))
		require.Error(t, err)
	})
}

func TestBuilds(t *testing.T) {
	t.Parallel()

	t.Run("create sends the request and returns the new build", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{
			status: http.StatusCreated,
			body:   `{"id":"0b9c6f1e-6a52-4d1f-9a43-2f0d9b7f5c11","status":"running","created_at":"2026-10-09T12:00:00Z"}`,
		})

		src := build.Git("https://github.com/acme/app.git", "main")
		b, err := c.CreateBuild(testContext(t), build.CreateBuildRequest{
			Kraftfile: build.KraftfileFrom(src),
			Contexts:  build.BuildContexts{Context: src},
			Output:    []build.BuildOutput{build.Image("index.unikraft.io/acme/app:1.2.3")},
		})
		require.NoError(t, err)

		req := <-requests
		assert.Equal(t, http.MethodPost, req.method)
		assert.Equal(t, "/v1/builds", req.path)
		assert.Equal(t, "Bearer tok", req.headers.Get("Authorization"))
		assert.Equal(t, "application/json", req.headers.Get("Content-Type"))
		assert.Equal(t, "application/json", req.headers.Get("Accept"))
		assert.JSONEq(t, `{
			"kraftfile": {"git": "https://github.com/acme/app.git", "ref": "main"},
			"contexts": {"context": {"git": "https://github.com/acme/app.git", "ref": "main"}},
			"output": [{"image": "index.unikraft.io/acme/app:1.2.3"}]
		}`, req.body)

		assert.Equal(t, "0b9c6f1e-6a52-4d1f-9a43-2f0d9b7f5c11", b.ID)
		assert.Equal(t, build.BuildStatusRunning, b.Status)
		assert.Equal(t, time.Date(2026, 10, 9, 12, 0, 0, 0, time.UTC), b.CreatedAt)
		assert.Nil(t, b.StartedAt)
		assert.Nil(t, b.Result)
		assert.Nil(t, b.Error)
	})

	t.Run("get returns the build with its result", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusOK, body: `{
			"id": "0b9c6f1e-6a52-4d1f-9a43-2f0d9b7f5c11",
			"status": "succeeded",
			"created_at": "2026-10-09T12:00:00Z",
			"started_at": "2026-10-09T12:00:01Z",
			"finished_at": "2026-10-09T12:03:40Z",
			"result": {"images": [{"ref": "index.unikraft.io/acme/app:1.2.3", "digest": "sha256:4c1a", "platforms": ["linux/amd64"]}]}
		}`})

		b, err := c.GetBuild(testContext(t), "0b9c6f1e-6a52-4d1f-9a43-2f0d9b7f5c11")
		require.NoError(t, err)

		req := <-requests
		assert.Equal(t, http.MethodGet, req.method)
		assert.Equal(t, "/v1/builds/0b9c6f1e-6a52-4d1f-9a43-2f0d9b7f5c11", req.path)
		assert.Empty(t, req.body)

		assert.Equal(t, build.BuildStatusSucceeded, b.Status)
		require.NotNil(t, b.StartedAt)
		assert.Equal(t, time.Date(2026, 10, 9, 12, 0, 1, 0, time.UTC), *b.StartedAt)
		require.NotNil(t, b.Result)
		require.Len(t, b.Result.Images, 1)
		assert.Equal(t, "sha256:4c1a", b.Result.Images[0].Digest)
		assert.Equal(t, []string{"linux/amd64"}, b.Result.Images[0].Platforms)
	})

	t.Run("list filters by status and limit", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusOK, body: `{"builds": [
			{"id": "b2", "status": "running", "created_at": "2026-10-09T12:00:00Z"},
			{"id": "b1", "status": "failed", "created_at": "2026-10-09T11:00:00Z", "error": {"code": "failed", "message": "RUN npm ci exited with 1"}}
		]}`})

		list, err := c.ListBuilds(testContext(t), build.ListBuildsOptions{
			Status: new(build.BuildStatusRunning),
			Limit:  new(int32(10)),
		})
		require.NoError(t, err)

		req := <-requests
		assert.Equal(t, http.MethodGet, req.method)
		assert.Equal(t, "/v1/builds", req.path)
		assert.Equal(t, "limit=10&status=running", req.query)

		require.Len(t, list.Builds, 2)
		assert.Equal(t, "b2", list.Builds[0].ID)
		require.NotNil(t, list.Builds[1].Error)
		assert.Equal(t, build.ErrorCodeFailed, list.Builds[1].Error.Code)
		require.EqualError(t, list.Builds[1].Error, "RUN npm ci exited with 1 (failed)")
	})

	t.Run("list sends no query when nothing is set", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusOK, body: `{"builds": []}`})

		list, err := c.ListBuilds(testContext(t), build.ListBuildsOptions{})
		require.NoError(t, err)

		assert.Empty(t, (<-requests).query)
		assert.Empty(t, list.Builds)
	})

	t.Run("cancel returns the build", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusAccepted, body: `{"id": "b1", "status": "running", "created_at": "2026-10-09T12:00:00Z"}`})

		b, err := c.CancelBuild(testContext(t), "b1")
		require.NoError(t, err)

		req := <-requests
		assert.Equal(t, http.MethodPost, req.method)
		assert.Equal(t, "/v1/builds/b1/cancel", req.path)
		assert.Equal(t, build.BuildStatusRunning, b.Status)
	})

	t.Run("delete sends no body and expects none", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusNoContent})

		require.NoError(t, c.DeleteBuild(testContext(t), "b1"))

		req := <-requests
		assert.Equal(t, http.MethodDelete, req.method)
		assert.Equal(t, "/v1/builds/b1", req.path)
		assert.Empty(t, req.body)
	})

	t.Run("an ID that is not a path segment is refused before any request", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 8)
		c := serve(t, requests, answer{status: http.StatusOK, body: `{"builds": []}`})

		for _, id := range []string{"", ".", "..", "b1/logs"} {
			_, err := c.GetBuild(testContext(t), id)
			require.Error(t, err, "get %q", id)

			_, err = c.CancelBuild(testContext(t), id)
			require.Error(t, err, "cancel %q", id)

			require.Error(t, c.DeleteBuild(testContext(t), id), "delete %q", id)

			for _, err := range c.Logs(testContext(t), id) {
				require.Error(t, err, "logs %q", id)
			}
		}

		assert.Empty(t, requests)
	})
}

func TestAPIError(t *testing.T) {
	t.Parallel()

	t.Run("an error body becomes the code and the message", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusConflict, body: `{"error": {"code": "conflict", "message": "build b1 is still running"}}`})

		err := c.DeleteBuild(testContext(t), "b1")

		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, http.StatusConflict, apiErr.StatusCode)
		assert.Equal(t, build.ErrorCodeConflict, apiErr.Code)
		assert.Equal(t, "build b1 is still running", apiErr.Message)
		require.EqualError(t, err, "build b1 is still running (conflict)")
	})

	t.Run("any other body becomes the message", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusBadGateway, body: "upstream unavailable\n", contentType: "text/plain"})

		_, err := c.GetBuild(testContext(t), "b1")

		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, http.StatusBadGateway, apiErr.StatusCode)
		assert.Empty(t, apiErr.Code)
		assert.Equal(t, "upstream unavailable", apiErr.Message)
		require.EqualError(t, err, "Bad Gateway: upstream unavailable")
	})

	t.Run("no body is the status text", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusNotFound})

		_, err := c.GetBuild(testContext(t), "b1")

		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		require.EqualError(t, err, "Not Found")
	})
}

func TestAnswers(t *testing.T) {
	t.Parallel()

	t.Run("a redirect is an error and is not followed", func(t *testing.T) {
		t.Parallel()

		// A relative Location points back at the test server, so a redirect
		// that is followed shows up as a second, recorded request.
		requests := make(chan request, 2)
		c := serve(t, requests, answer{status: http.StatusFound, location: "/v1/builds?redirected"})

		_, err := c.CreateBuild(testContext(t), build.CreateBuildRequest{
			Kraftfile: build.KraftfileInline("spec: v0.7\n"),
			Contexts:  build.BuildContexts{Context: build.URL("https://example.com/app.tar.gz")},
			Output:    []build.BuildOutput{build.Image("index.unikraft.io/acme/app:1.2.3")},
		})
		require.ErrorContains(t, err, "redirected to")
		require.ErrorContains(t, err, "/v1/builds?redirected")

		assert.Equal(t, http.MethodPost, (<-requests).method)
		assert.Empty(t, requests, "the redirect was followed")
	})

	t.Run("an answer that is not JSON is an error", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusOK, body: "<html>maintenance</html>", contentType: "text/html"})

		_, err := c.GetBuild(testContext(t), "b1")
		require.ErrorContains(t, err, "not application/json")
	})

	t.Run("a body over the limit is an error", func(t *testing.T) {
		t.Parallel()

		requests := make(chan request, 1)
		c := serve(t, requests, answer{status: http.StatusOK, body: `{"id": "b1", "pad": "` + strings.Repeat("x", 8<<20) + `"}`})

		_, err := c.GetBuild(testContext(t), "b1")
		require.ErrorContains(t, err, "larger than the limit")
	})
}

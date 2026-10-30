// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build_test

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/build"
)

// A progress stream, line by line.
const (
	lineKernel      = `{"vertexes":[{"digest":"sha256:a1","name":"[kernel] pull unikraft.org/base-compat:latest","started":"2026-10-09T12:00:01Z","completed":"2026-10-09T12:00:02Z"}]}`
	lineRootfs      = `{"vertexes":[{"digest":"sha256:b2","name":"[rootfs 2/4] RUN npm ci","started":"2026-10-09T12:00:03Z"}]}`
	lineLog         = `{"logs":[{"vertex":"sha256:b2","stream":1,"data":"YWRkZWQgMjEyIHBhY2thZ2Vz","timestamp":"2026-10-09T12:00:20Z"}]}`
	lineResult      = `{"result":{"images":[{"ref":"index.unikraft.io/acme/app:1.2.3","digest":"sha256:4c1a","platforms":["linux/amd64"]}]}}`
	lineError       = `{"error":{"code":"failed","message":"process \"npm ci\" did not complete successfully: exit code: 1"}}`
	lineInterrupted = `{"error":{"code":"interrupted","message":"the service restarted while the build was running"}}`
	lineNullVertex  = `{"vertexes":[null]}`
	lineNullWarning = `{"warnings":[null],"statuses":[null],"logs":[null]}`
	lineBadTime     = `{"vertexes":[{"digest":"sha256:a1","name":"[kernel] pull","started":"yesterday"}]}`
	lineNotJSON     = `not json`
)

// allLines is the whole stream of a succeeded build.
var allLines = []string{lineKernel, lineRootfs, lineLog, lineResult}

// reply is what a stream server answers one request with: a status with a
// JSON body, or the lines of a stream and then a partial line.
type reply struct {
	status      int
	body        string
	lines       []string
	partial     string
	contentType string // application/x-ndjson unless given
}

// streamServer answers every request for the progress stream of build b1
// with the reply that answer returns for the number of the request, and
// returns a client of it along with the count of requests it got.
func streamServer(t *testing.T, answer func(n int32) reply) (*build.Client, *atomic.Int32) {
	t.Helper()

	var requests atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := requests.Add(1)

		assert.Equal(t, "/v1/builds/b1/logs", r.URL.Path)
		assert.Equal(t, "application/x-ndjson", r.Header.Get("Accept"))

		rep := answer(n)
		if rep.status != 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(rep.status)
			fmt.Fprint(w, rep.body)
			return
		}

		w.Header().Set("Content-Type", cmp.Or(rep.contentType, "application/x-ndjson"))
		for _, line := range rep.lines {
			fmt.Fprintln(w, line)
		}
		fmt.Fprint(w, rep.partial)
	}))
	t.Cleanup(srv.Close)

	c, err := build.New(srv.URL, build.WithToken("tok"), build.WithRetry(time.Millisecond, 4*time.Millisecond))
	require.NoError(t, err)

	return c, &requests
}

// holdingServer sends one line of the stream of build b1 and then holds it
// open until the client goes away, which closes the returned channel.
func holdingServer(t *testing.T) (*build.Client, <-chan struct{}) {
	t.Helper()

	gone := make(chan struct{})
	var once sync.Once
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ndjson")
		fmt.Fprintln(w, lineKernel)
		w.(http.Flusher).Flush()
		<-r.Context().Done()
		once.Do(func() { close(gone) })
	}))
	t.Cleanup(srv.Close)

	c, err := build.New(srv.URL, build.WithRetry(time.Millisecond, 4*time.Millisecond))
	require.NoError(t, err)

	return c, gone
}

// follow drains the stream of build b1 into its events and the error it
// ended with, if any.
func follow(ctx context.Context, c *build.Client) ([]build.Event, error) {
	var events []build.Event
	for ev, err := range c.Logs(ctx, "b1") {
		if err != nil {
			return events, err
		}
		events = append(events, ev)
	}

	return events, nil
}

// requireStream checks that events are the whole stream of a succeeded
// build, in order and once each.
func requireStream(t *testing.T, events []build.Event) {
	t.Helper()

	require.Len(t, events, len(allLines))

	require.NotNil(t, events[0].SolveStatus)
	require.Len(t, events[0].Vertexes, 1)
	assert.Equal(t, "[kernel] pull unikraft.org/base-compat:latest", events[0].Vertexes[0].Name)
	assert.False(t, events[0].Done())

	require.NotNil(t, events[1].SolveStatus)
	require.Len(t, events[1].Vertexes, 1)
	assert.Equal(t, "[rootfs 2/4] RUN npm ci", events[1].Vertexes[0].Name)

	require.NotNil(t, events[2].SolveStatus)
	require.Len(t, events[2].Logs, 1)
	assert.Equal(t, []byte("added 212 packages"), events[2].Logs[0].Data)

	assert.Nil(t, events[3].SolveStatus)
	require.NotNil(t, events[3].Result)
	require.Len(t, events[3].Result.Images, 1)
	assert.Equal(t, "index.unikraft.io/acme/app:1.2.3", events[3].Result.Images[0].Ref)
	assert.True(t, events[3].Done())
}

func TestLogs(t *testing.T) {
	t.Parallel()

	t.Run("yields every event once after the connection drops", func(t *testing.T) {
		t.Parallel()

		// The first answer ends after two lines; the second replays the whole
		// stream, and only its last two lines are new.
		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{lines: allLines[:2]}
			}
			return reply{lines: allLines}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("skips the right lines over several reconnects", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			switch n {
			case 1:
				return reply{lines: allLines[:1]}
			case 2:
				return reply{lines: allLines[:2]}
			default:
				return reply{lines: allLines}
			}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(3), requests.Load())
	})

	t.Run("a stream cut mid-line is asked for again", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{lines: allLines[:1], partial: lineRootfs[:20]}
			}
			return reply{lines: allLines}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("a stream whose read fails is asked for again", func(t *testing.T) {
		t.Parallel()

		// The first answer sends two lines and then goes quiet for longer
		// than the HTTP client's timeout, so the read of the body fails
		// rather than ending.
		var requests atomic.Int32
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/x-ndjson")
			if requests.Add(1) == 1 {
				fmt.Fprintln(w, lineKernel)
				fmt.Fprintln(w, lineRootfs)
				w.(http.Flusher).Flush()
				<-r.Context().Done()
				return
			}
			for _, line := range allLines {
				fmt.Fprintln(w, line)
			}
		}))
		t.Cleanup(srv.Close)

		c, err := build.New(srv.URL,
			build.WithHTTPClient(&http.Client{Timeout: 100 * time.Millisecond}),
			build.WithRetry(time.Millisecond, 4*time.Millisecond))
		require.NoError(t, err)

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("a shorter replay that ends in an error is not skipped", func(t *testing.T) {
		t.Parallel()

		// The service restarted: the build is lost, and its stream is now
		// one line that says so, at a position already yielded.
		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{lines: allLines[:3]}
			}
			return reply{lines: []string{lineInterrupted}}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		require.Len(t, events, 4)
		require.NotNil(t, events[3].Error)
		assert.Equal(t, build.ErrorCodeInterrupted, events[3].Error.Code)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("a 5xx while following is retried", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			switch n {
			case 1:
				return reply{lines: allLines[:2]}
			case 2:
				return reply{status: http.StatusServiceUnavailable, body: "upstream unavailable"}
			default:
				return reply{lines: allLines}
			}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(3), requests.Load())
	})

	t.Run("a 502 before the stream is retried", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{status: http.StatusBadGateway, body: `{"error": {"code": "auth_unavailable", "message": "could not validate token"}}`}
			}
			return reply{lines: allLines}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		requireStream(t, events)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("a 4xx while following is not retried", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{lines: allLines[:2]}
			}
			return reply{status: http.StatusNotFound, body: `{"error": {"code": "not_found", "message": "no build b1"}}`}
		})

		events, err := follow(testContext(t), c)
		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, build.ErrorCodeNotFound, apiErr.Code)
		assert.Len(t, events, 2)
		assert.Equal(t, int32(2), requests.Load())
	})

	t.Run("a refused build is an API error", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			return reply{status: http.StatusNotFound, body: `{"error": {"code": "not_found", "message": "no build b1"}}`}
		})

		events, err := follow(testContext(t), c)
		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, build.ErrorCodeNotFound, apiErr.Code)
		assert.Empty(t, events)
		assert.Equal(t, int32(1), requests.Load())
	})

	t.Run("a server that never answered is not retried", func(t *testing.T) {
		t.Parallel()

		c, err := build.New("http://127.0.0.1:1", build.WithRetry(time.Millisecond, 4*time.Millisecond))
		require.NoError(t, err)

		events, err := follow(testContext(t), c)
		require.Error(t, err)
		require.NotContains(t, err.Error(), "giving up")
		assert.Empty(t, events)
	})

	t.Run("a service that goes away while following is given up on", func(t *testing.T) {
		t.Parallel()

		// After the first answer, every connection is closed without one.
		var requests atomic.Int32
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			if requests.Add(1) == 1 {
				w.Header().Set("Content-Type", "application/x-ndjson")
				fmt.Fprintln(w, lineKernel)
				fmt.Fprintln(w, lineRootfs)
				return
			}

			conn, _, err := w.(http.Hijacker).Hijack()
			if assert.NoError(t, err) {
				conn.Close()
			}
		}))
		t.Cleanup(srv.Close)

		c, err := build.New(srv.URL, build.WithRetry(time.Millisecond, 4*time.Millisecond))
		require.NoError(t, err)

		events, err := follow(testContext(t), c)
		require.ErrorContains(t, err, "giving up on the progress stream")
		assert.Len(t, events, 2)
		// One stream, then the attempts of the budget, some of which the
		// transport itself retried on a fresh connection.
		assert.GreaterOrEqual(t, requests.Load(), int32(7))
	})

	t.Run("a service that keeps failing is given up on", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(n int32) reply {
			if n == 1 {
				return reply{lines: allLines[:2]}
			}
			return reply{status: http.StatusServiceUnavailable, body: "upstream unavailable"}
		})

		events, err := follow(testContext(t), c)
		require.ErrorContains(t, err, "giving up on the progress stream")
		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, http.StatusServiceUnavailable, apiErr.StatusCode)
		assert.Len(t, events, 2)
		// The stream, then one attempt plus the five more of the budget.
		assert.Equal(t, int32(7), requests.Load())
	})

	t.Run("an answer that is not a stream ends it", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			return reply{lines: []string{"<html>maintenance</html>"}, contentType: "text/html"}
		})

		events, err := follow(testContext(t), c)
		require.ErrorContains(t, err, "not a progress stream")
		assert.Empty(t, events)
		assert.Equal(t, int32(1), requests.Load())
	})

	t.Run("a line that is not JSON ends the stream", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			return reply{lines: []string{lineKernel, lineNotJSON, lineResult}}
		})

		events, err := follow(testContext(t), c)
		var syntaxErr *json.SyntaxError
		require.ErrorAs(t, err, &syntaxErr)
		require.ErrorContains(t, err, "line 2")
		assert.Len(t, events, 1)
		assert.Equal(t, int32(1), requests.Load())
	})

	t.Run("a line with a bad timestamp ends the stream", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			return reply{lines: []string{lineBadTime, lineResult}}
		})

		events, err := follow(testContext(t), c)
		var parseErr *time.ParseError
		require.ErrorAs(t, err, &parseErr)
		assert.Empty(t, events)
		assert.Equal(t, int32(1), requests.Load())
	})

	t.Run("a line longer than the limit ends the stream", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			long := `{"logs":[{"vertex":"sha256:b2","data":"` + strings.Repeat("x", build.MaxLineSize) + `"}]}`
			return reply{lines: []string{lineKernel, long, lineResult}}
		})

		events, err := follow(testContext(t), c)
		require.ErrorContains(t, err, "longer than")
		assert.Len(t, events, 1)
		assert.Equal(t, int32(1), requests.Load())
	})

	t.Run("a null entry is dropped", func(t *testing.T) {
		t.Parallel()

		c, _ := streamServer(t, func(int32) reply {
			return reply{lines: []string{lineNullVertex, lineNullWarning, lineKernel, lineResult}}
		})

		events, err := follow(testContext(t), c)
		require.NoError(t, err)
		require.Len(t, events, 4)
		require.NotNil(t, events[0].SolveStatus)
		assert.Empty(t, events[0].Vertexes)
		require.NotNil(t, events[1].SolveStatus)
		assert.Empty(t, events[1].Warnings)
		assert.Empty(t, events[1].Statuses)
		assert.Empty(t, events[1].Logs)
		require.NotNil(t, events[3].Result)
	})

	t.Run("cancelling the context ends the stream with its error", func(t *testing.T) {
		t.Parallel()

		c, gone := holdingServer(t)
		ctx, cancel := context.WithCancel(testContext(t))
		defer cancel()

		var events int
		var last error
		for _, err := range c.Logs(ctx, "b1") {
			if err != nil {
				last = err
				break
			}
			events++
			cancel()
		}

		require.ErrorIs(t, last, context.Canceled)
		assert.Equal(t, 1, events)

		select {
		case <-gone:
		case <-time.After(testTimeout):
			t.Fatal("the server still sees the client after the context ended")
		}
	})

	t.Run("stops when the caller breaks out", func(t *testing.T) {
		t.Parallel()

		c, requests := streamServer(t, func(int32) reply {
			return reply{lines: allLines}
		})

		var events int
		for _, err := range c.Logs(testContext(t), "b1") {
			require.NoError(t, err)
			events++
			break
		}

		assert.Equal(t, 1, events)
		assert.Equal(t, int32(1), requests.Load())
	})
}

func TestPlainLogs(t *testing.T) {
	t.Parallel()

	t.Run("renders a succeeded build to the end", func(t *testing.T) {
		t.Parallel()

		c, _ := streamServer(t, func(int32) reply {
			return reply{lines: allLines}
		})

		logs := c.PlainLogs(testContext(t), "b1")
		defer logs.Close()

		var out bytes.Buffer
		_, err := io.Copy(&out, logs)
		require.NoError(t, err)
		assert.Contains(t, out.String(), "[kernel] pull unikraft.org/base-compat:latest")
		assert.Contains(t, out.String(), "[rootfs 2/4] RUN npm ci")
		assert.Contains(t, out.String(), "added 212 packages")
	})

	t.Run("a failed build is the build's error", func(t *testing.T) {
		t.Parallel()

		c, _ := streamServer(t, func(int32) reply {
			return reply{lines: []string{lineKernel, lineRootfs, lineError}}
		})

		logs := c.PlainLogs(testContext(t), "b1")
		defer logs.Close()

		var out bytes.Buffer
		_, err := io.Copy(&out, logs)

		var buildErr *build.Error
		require.ErrorAs(t, err, &buildErr)
		assert.Equal(t, build.ErrorCodeFailed, buildErr.Code)
		assert.Contains(t, out.String(), "[rootfs 2/4] RUN npm ci")
	})

	t.Run("a refused build is an API error", func(t *testing.T) {
		t.Parallel()

		c, _ := streamServer(t, func(int32) reply {
			return reply{status: http.StatusNotFound, body: `{"error": {"code": "not_found", "message": "no build b1"}}`}
		})

		logs := c.PlainLogs(testContext(t), "b1")
		defer logs.Close()

		_, err := io.Copy(io.Discard, logs)

		var apiErr *build.APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
	})

	t.Run("cancelling the context ends the copy with its error", func(t *testing.T) {
		t.Parallel()

		c, gone := holdingServer(t)
		ctx, cancel := context.WithCancel(testContext(t))
		defer cancel()

		logs := c.PlainLogs(ctx, "b1")
		defer logs.Close()

		buf := make([]byte, 1)
		_, err := logs.Read(buf)
		require.NoError(t, err)

		cancel()

		_, err = io.Copy(io.Discard, logs)
		require.ErrorIs(t, err, context.Canceled)

		select {
		case <-gone:
		case <-time.After(testTimeout):
			t.Fatal("the server still sees the client after the context ended")
		}
	})

	t.Run("closing stops following the build", func(t *testing.T) {
		t.Parallel()

		c, gone := holdingServer(t)

		logs := c.PlainLogs(testContext(t), "b1")

		buf := make([]byte, 1)
		_, err := logs.Read(buf)
		require.NoError(t, err)

		require.NoError(t, logs.Close())

		_, err = logs.Read(buf)
		require.ErrorIs(t, err, io.ErrClosedPipe)

		select {
		case <-gone:
		case <-time.After(testTimeout):
			t.Fatal("the server still sees the client after Close")
		}
	})
}

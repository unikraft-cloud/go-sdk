// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"slices"
	"time"

	bkclient "github.com/moby/buildkit/client"
	"github.com/moby/buildkit/util/progress/progressui"
)

// Event is one line of the progress stream of a build: a buildkit SolveStatus
// in the shape of `docker buildx build --progress=rawjson`, which is nil on
// the last line, plus the terminal Result or Error that only the last line
// carries. It embeds buildkit's own type rather than a copy of it, so that
// anything that already renders buildkit progress renders this stream as-is.
type Event struct {
	*bkclient.SolveStatus

	// Result is what the build produced. It is only set on the last event of
	// a succeeded build.
	Result *BuildResult `json:"result,omitempty"`

	// Error is why the build did not succeed. It is only set on the last
	// event of a failed or cancelled build.
	Error *Error `json:"error,omitempty"`
}

// Done reports whether e is the last event of the stream.
func (e Event) Done() bool {
	return e.Result != nil || e.Error != nil
}

const (
	// logsRetryMin and logsRetryMax bound the wait before the progress
	// stream is asked for again, which doubles on every attempt in a row and
	// starts over once an attempt brings a new event.
	logsRetryMin = 250 * time.Millisecond
	logsRetryMax = 5 * time.Second

	// logsRetryBudget is how many more times in a row the stream is asked
	// for after an attempt got no stream at all, because the service could
	// not be reached or answered with a 5xx, before Logs gives up.
	logsRetryBudget = 5

	// initialLineSize is the buffer a stream starts with, which grows as
	// needed up to maxLineSize, the bound on one line of the stream.
	initialLineSize = 64 << 10
	maxLineSize     = 16 << 20
)

// errStreamEnded means the progress stream ended before its last event: it
// was cut by the connection, or by something between the service and the
// caller.
var errStreamEnded = errors.New("progress stream ended early")

// errConnect wraps a request for the progress stream that got no answer at
// all.
var errConnect = errors.New("requesting progress stream")

// Logs follows the progress of the build with the given ID, yielding every
// event of its stream exactly once, in order, until the last one: the Result
// or Error that says how the build went.
//
// The service replays the stream from the start on every request. A stream
// that ends before its last event, cut by the connection or by something
// between the service and the caller, is asked for again after a short and
// growing wait, and the events already yielded are skipped. That goes on for
// as long as ctx allows: a stream is the service saying that the build is
// alive, however long it stays silent. A service that cannot be reached, or
// answers with a 5xx such as auth_unavailable, is asked again up to five more
// times in a row, then given up on with the last error.
//
// Logs yields an error and stops when the build cannot be followed: the
// service refused the request (an *APIError with a 4xx, such as a build that
// does not exist), the request got no answer before the service ever gave
// one (so a wrong server fails straight away), the answer was not a progress
// stream, a line could not be decoded, or ctx is done. A failed build is not
// an error here but an Event whose Error is set; PlainLogs turns it into one.
func (c *Client) Logs(ctx context.Context, id string) iter.Seq2[Event, error] {
	return func(yield func(Event, error) bool) {
		var (
			seen     int  // the lines yielded so far, which a replay skips
			answered bool // the service has answered once, so it is there
			failed   int  // the attempts in a row that got no stream at all
			idle     int  // the attempts in a row that brought no new event
		)

		for {
			yielded, stop, err := c.followLogs(ctx, id, seen, yield)
			seen += yielded
			if stop {
				return
			}
			if ctx.Err() != nil {
				yield(Event{}, ctx.Err())
				return
			}

			if !errors.Is(err, errConnect) {
				answered = true
			}
			if yielded > 0 {
				idle = 0
			}

			switch {
			case errors.Is(err, errStreamEnded):
				failed = 0
			case noStream(err, answered):
				failed++
				if failed > logsRetryBudget {
					yield(Event{}, fmt.Errorf("giving up on the progress stream after %d attempts: %w", failed, err))
					return
				}
			default:
				yield(Event{}, err)
				return
			}

			wait := min(c.retryMin<<idle, c.retryMax)
			if wait < c.retryMax {
				idle++
			}

			select {
			case <-time.After(wait):
			case <-ctx.Done():
				yield(Event{}, ctx.Err())
				return
			}
		}
	}
}

// noStream reports whether err means that an attempt got no stream at all
// from a service that is known to be there: the request got no answer after
// the service had answered before, or the answer was a 5xx.
func noStream(err error, answered bool) bool {
	var apiErr *APIError

	switch {
	case errors.As(err, &apiErr):
		return apiErr.StatusCode >= http.StatusInternalServerError
	case errors.Is(err, errConnect):
		return answered
	default:
		return false
	}
}

// followLogs requests the progress stream once, skips its first seen lines
// unless one is the last event, and yields the rest, returning how many it
// yielded. stop is true when the last event was yielded or yield returned
// false. Otherwise err says why the stream ended: errStreamEnded when it was
// cut short, errConnect when the request got no answer, and the error itself
// when the service refused the request, the answer was not a stream, or a
// line could not be decoded.
func (c *Client) followLogs(ctx context.Context, id string, seen int, yield func(Event, error) bool) (yielded int, stop bool, err error) {
	segment, err := pathSegment("build ID", id)
	if err != nil {
		return 0, false, err
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/v1/builds/"+segment+"/logs", nil, nil)
	if err != nil {
		return 0, false, err
	}
	req.Header.Set("Accept", "application/x-ndjson")

	resp, err := c.hc.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return 0, false, ctx.Err()
		}
		return 0, false, fmt.Errorf("%w: %w", errConnect, err)
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= http.StatusBadRequest:
		return 0, false, readError(resp)
	case resp.StatusCode >= http.StatusMultipleChoices:
		return 0, false, redirected(resp)
	}
	if err := checkMediaType(resp, "application/x-ndjson"); err != nil {
		return 0, false, fmt.Errorf("not a progress stream: %w", err)
	}

	// Lines are read rather than decoded straight off the body, so that a
	// read error, which means the stream was cut, is told apart from a line
	// that does not decode, which a replay would only repeat.
	sc := bufio.NewScanner(resp.Body)
	sc.Buffer(make([]byte, 0, initialLineSize), maxLineSize)
	sc.Split(scanLines)

	line := 0
	for sc.Scan() {
		raw := bytes.TrimSpace(sc.Bytes())
		if len(raw) == 0 {
			continue
		}
		line++

		var ev Event
		if err := json.Unmarshal(raw, &ev); err != nil {
			return yielded, false, fmt.Errorf("decoding line %d of the progress stream: %w", line, err)
		}
		dropNil(ev.SolveStatus)

		// A replay is skipped up to the lines yielded before, except for a
		// last event: a service that restarted replays a shorter stream that
		// ends in an error, which must not be skipped.
		if line <= seen && !ev.Done() {
			continue
		}
		yielded++

		if !yield(ev, nil) || ev.Done() {
			return yielded, true, nil
		}
	}

	err = sc.Err()
	switch {
	case ctx.Err() != nil:
		return yielded, false, ctx.Err()
	case errors.Is(err, bufio.ErrTooLong):
		return yielded, false, fmt.Errorf("line %d of the progress stream is longer than %d bytes", line+1, maxLineSize)
	case err == nil:
		// The stream ended, before its last event.
		return yielded, false, errStreamEnded
	default:
		// The stream was cut mid-line, or the connection failed.
		return yielded, false, fmt.Errorf("%w: %w", errStreamEnded, err)
	}
}

// scanLines splits the stream into lines. Unlike bufio.ScanLines, it does
// not hand out a last line that has no newline: a stream that ends mid-line
// was cut, and is reported as io.ErrUnexpectedEOF.
func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[:i], nil
	}
	if atEOF && len(data) > 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}

	return 0, nil, nil
}

// dropNil removes the nil entries of ss, which a null on a line decodes to
// and which progressui dereferences as they are.
func dropNil(ss *bkclient.SolveStatus) {
	if ss == nil {
		return
	}

	ss.Vertexes = slices.DeleteFunc(ss.Vertexes, isNil)
	ss.Statuses = slices.DeleteFunc(ss.Statuses, isNil)
	ss.Logs = slices.DeleteFunc(ss.Logs, isNil)
	ss.Warnings = slices.DeleteFunc(ss.Warnings, isNil)
}

func isNil[T any](p *T) bool { return p == nil }

// PlainLogs renders the progress of the build with the given ID the way
// `docker buildx build --progress=plain` does, as a reader to copy wherever
// the progress should go. It follows the build to the end, as Logs does, and
// reads drive it: nothing is rendered while the reader is not read. Once the
// build has succeeded, Read returns io.EOF; once it has failed or been
// cancelled, Read returns the build's *Error instead, so that a failed build
// is an error from io.Copy. Close stops following the build, as does ctx.
func (c *Client) PlainLogs(ctx context.Context, id string) io.ReadCloser {
	ctx, cancel := context.WithCancel(ctx)
	pr, pw := io.Pipe()

	// A write to the pipe waits for a read, so a reader that stops reading
	// without closing would hold the display, and the stream with it, for
	// ever. Once ctx is done, the write end is closed with its cause, so that
	// a waiting write returns and the display can stop. The first close
	// wins, so a build that ended before ctx did is still reported as it
	// went.
	stop := context.AfterFunc(ctx, func() { pw.CloseWithError(context.Cause(ctx)) })

	go func() {
		err := c.renderPlain(ctx, id, pw)
		stop()
		pw.CloseWithError(err)
		cancel()
	}()

	return &plainLogs{PipeReader: pr, cancel: cancel}
}

// plainLogs is the reader PlainLogs returns.
type plainLogs struct {
	*io.PipeReader
	cancel context.CancelFunc
}

func (l *plainLogs) Close() error {
	l.cancel()
	return l.PipeReader.Close()
}

// renderPlain feeds the progress stream of the build through progressui's
// plain mode into w. It returns the build's *Error when the build did not
// succeed, the error that stopped the stream when it could not be followed,
// and nil once a succeeded build has been rendered to the end.
func (c *Client) renderPlain(ctx context.Context, id string, w io.Writer) error {
	display, err := progressui.NewDisplay(w, progressui.PlainMode)
	if err != nil {
		return err
	}

	// The display returns once ch is closed or ctx is done. It never stops
	// because of w: the plain printer ignores write errors.
	ch := make(chan *bkclient.SolveStatus)
	rendered := make(chan error, 1)
	go func() {
		_, err := display.UpdateFrom(ctx, ch)
		rendered <- err
	}()

	var buildErr, streamErr error

events:
	for ev, err := range c.Logs(ctx, id) {
		if err != nil {
			streamErr = err
			break
		}
		if ev.Error != nil {
			buildErr = ev.Error
		}
		if ev.SolveStatus == nil {
			continue
		}

		select {
		case ch <- ev.SolveStatus:
		case <-ctx.Done():
			// The display stops taking events once ctx is done.
			streamErr = ctx.Err()
			break events
		}
	}
	close(ch)

	renderErr := <-rendered
	switch {
	case streamErr != nil:
		return streamErr
	case buildErr != nil:
		return buildErr
	default:
		return renderErr
	}
}

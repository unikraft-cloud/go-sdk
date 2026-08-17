// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sse

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadEventBasicFields(t *testing.T) {
	r := NewReader(strings.NewReader("id: 1\nevent: message\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, "1", event.ID)
	assert.Equal(t, "message", event.Event)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventMultipleDataLinesAreJoinedWithNewline(t *testing.T) {
	r := NewReader(strings.NewReader("data: line1\ndata: line2\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("line1\nline2"), event.Data)
}

func TestReadEventIgnoresCommentLines(t *testing.T) {
	r := NewReader(strings.NewReader(":keep-alive\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventIgnoresMalformedFieldLines(t *testing.T) {
	r := NewReader(strings.NewReader("nocolon\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventRetryField(t *testing.T) {
	r := NewReader(strings.NewReader("retry: 5000\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, "5000", event.Retry)
}

func TestReadEventUnknownFieldIsIgnored(t *testing.T) {
	r := NewReader(strings.NewReader("foo: bar\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventTrimsLeadingSpaceInValue(t *testing.T) {
	r := NewReader(strings.NewReader("event:  spaced\ndata: value\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	// Only a single leading space is trimmed, matching the SSE spec.
	assert.Equal(t, " spaced", event.Event)
}

func TestReadEventMultipleEventsInSequence(t *testing.T) {
	r := NewReader(strings.NewReader("data: first\n\ndata: second\n\n"))

	first, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("first"), first.Data)

	second, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("second"), second.Data)

	_, err = r.ReadEvent()
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadEventEmptyStreamReturnsEOF(t *testing.T) {
	r := NewReader(strings.NewReader(""))

	_, err := r.ReadEvent()
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadEventWithoutTrailingBlankLineStillReturnsData(t *testing.T) {
	// No trailing "\n\n" - the stream just ends after a data line.
	r := NewReader(strings.NewReader("data: hello"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)

	_, err = r.ReadEvent()
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadEventHandlesCRLFLineEndings(t *testing.T) {
	r := NewReader(strings.NewReader("data: hello\r\n\r\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventHandlesMixedLFAndCRLF(t *testing.T) {
	r := NewReader(strings.NewReader("event: a\r\ndata: hello\n\r\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, "a", event.Event)
	assert.Equal(t, []byte("hello"), event.Data)
}

func TestReadEventEmptyDataField(t *testing.T) {
	r := NewReader(strings.NewReader("event: ping\ndata:\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, "ping", event.Event)
	assert.Equal(t, []byte(""), event.Data)
}

func TestReadEventConsecutiveBlankLinesAreSkippedOutsideEvent(t *testing.T) {
	r := NewReader(strings.NewReader("\n\n\ndata: hello\n\n"))

	event, err := r.ReadEvent()
	require.NoError(t, err)
	assert.Equal(t, []byte("hello"), event.Data)
}

type errReader struct{ err error }

func (r errReader) Read([]byte) (int, error) { return 0, r.err }

func TestReadEventPropagatesScannerError(t *testing.T) {
	boom := assert.AnError
	r := NewReader(errReader{err: boom})

	_, err := r.ReadEvent()
	assert.ErrorIs(t, err, boom)
}

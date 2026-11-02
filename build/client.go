// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// userAgent identifies the client to the service.
const userAgent = "unikraft-cloud-build-client"

const (
	// maxBodySize bounds how much of a successful response is read for its
	// JSON body.
	maxBodySize = 8 << 20

	// maxErrorBody bounds how much of a failed response is read for its
	// message.
	maxErrorBody = 1 << 20

	// logsRetryMin and logsRetryMax bound the wait before the progress
	// stream is asked for again, which doubles on every attempt in a row and
	// starts over once an attempt brings a new event.
	logsRetryMin = 250 * time.Millisecond
	logsRetryMax = 5 * time.Second
)

// Client talks to one build service with one token.
type Client struct {
	server *url.URL
	token  string
	hc     *http.Client

	// retryMin and retryMax bound the wait between two attempts at the
	// progress stream of a build; see Logs.
	retryMin, retryMax time.Duration
}

// Option configures a Client.
type Option func(*Client)

// WithToken sets the Unikraft Cloud token sent as the bearer token of every
// request. Without it, requests are sent unauthenticated and refused.
func WithToken(token string) Option {
	return func(c *Client) { c.token = token }
}

// WithHTTPClient sets the HTTP client requests are sent with, in place of
// http.DefaultClient; nil is an error from New. It is used through a copy
// that never follows a redirect; see New. A timeout on it applies to every
// request, including the progress stream of a build, which lasts for as long
// as the build does, so bound that with the context instead.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.hc = hc }
}

// withRetry sets the wait between two attempts at the progress stream, which
// starts at base, doubles on every attempt in a row, and never exceeds limit.
// The tests shorten it.
func withRetry(base, limit time.Duration) Option {
	return func(c *Client) { c.retryMin, c.retryMax = base, limit }
}

// New returns a client of the build service at server, e.g.
// https://build.unikraft.cloud. The paths of the API are appended to it.
//
// A redirect is never followed: an answer that redirects, say from http to
// https, is returned as an error naming where to, since following it would
// send the token again and turn a POST into a GET. Give New the final URL.
func New(server string, opts ...Option) (*Client, error) {
	u, err := url.Parse(server)
	if err != nil {
		return nil, fmt.Errorf("parsing server URL: %w", err)
	}
	if u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("server URL %q has no scheme or host", server)
	}

	c := &Client{
		server:   u,
		hc:       http.DefaultClient,
		retryMin: logsRetryMin,
		retryMax: logsRetryMax,
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.hc == nil {
		return nil, errors.New("the HTTP client is nil")
	}
	c.hc = noRedirects(c.hc)

	return c, nil
}

// noRedirects returns a copy of hc that returns every redirecting answer as
// it is, instead of following it.
func noRedirects(hc *http.Client) *http.Client {
	c := *hc
	c.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &c
}

// pathSegment returns value escaped as one segment of a request path. It
// refuses a value that is empty, "." or "..", which the path would resolve
// away, or that contains a slash, which would add a segment; name says what
// the value is in the error.
func pathSegment(name, value string) (string, error) {
	switch {
	case value == "":
		return "", fmt.Errorf("%s is empty", name)
	case value == "." || value == "..":
		return "", fmt.Errorf("%s %q is not a path segment", name, value)
	case strings.Contains(value, "/"):
		return "", fmt.Errorf("%s %q contains a slash", name, value)
	}

	return url.PathEscape(value), nil
}

// newRequest builds a request for path, which is relative to the server and
// already escaped, with the token and the JSON encoding of body, if any.
func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values, body any) (*http.Request, error) {
	u := c.server.JoinPath(path)
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// do sends a request and decodes the JSON body of a successful response into
// out, unless out is nil or the response has no content. A response with a
// status of 400 or above is returned as an *APIError, and a redirect as an
// error naming its target.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body, out any) error {
	req, err := c.newRequest(ctx, method, path, query, body)
	if err != nil {
		return err
	}
	if out != nil {
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= http.StatusBadRequest:
		return readError(resp)
	case resp.StatusCode >= http.StatusMultipleChoices:
		return redirected(resp)
	case out == nil || resp.StatusCode == http.StatusNoContent:
		return nil
	}

	if err := checkMediaType(resp, "application/json"); err != nil {
		return err
	}
	b, err := readBody(resp.Body, maxBodySize)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}
	if err := json.Unmarshal(b, out); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}

// redirected is the error for an answer that redirects, which the client
// never follows.
func redirected(resp *http.Response) error {
	return fmt.Errorf("%s %s was redirected to %q: use the final server URL",
		resp.Request.Method, resp.Request.URL.Path, resp.Header.Get("Location"))
}

// checkMediaType returns an error unless the response carries the media type
// want, so that a page from something in front of the service is never
// decoded as if the service had answered.
func checkMediaType(resp *http.Response, want string) error {
	ct := resp.Header.Get("Content-Type")
	mt, _, err := mime.ParseMediaType(ct)
	if err != nil || mt != want {
		return fmt.Errorf("response is %q, not %s", ct, want)
	}

	return nil
}

// errBodyTooLarge means a response body went past the limit put on it.
var errBodyTooLarge = errors.New("response body is larger than the limit")

// readBody reads r to its end, or returns errBodyTooLarge when it holds more
// than limit bytes.
func readBody(r io.Reader, limit int64) ([]byte, error) {
	b, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(b)) > limit {
		return nil, errBodyTooLarge
	}

	return b, nil
}

// APIError is a request the service refused. StatusCode is the HTTP status
// it answered with. Code and Message are the error it answered with, when the
// body was an ErrorBody; otherwise Code is empty and Message carries the body
// as it was, which happens when something in front of the service answered.
type APIError struct {
	StatusCode int
	Code       ErrorCode
	Message    string
}

func (e *APIError) Error() string {
	switch {
	case e.Code != "":
		return fmt.Sprintf("%s (%s)", e.Message, e.Code)
	case e.Message != "":
		return fmt.Sprintf("%s: %s", http.StatusText(e.StatusCode), e.Message)
	default:
		return http.StatusText(e.StatusCode)
	}
}

// readError turns a failed response into an *APIError.
func readError(resp *http.Response) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))

	apiErr := &APIError{StatusCode: resp.StatusCode}

	var eb ErrorBody
	if err := json.Unmarshal(body, &eb); err == nil && eb.Error.Code != "" {
		apiErr.Code = eb.Error.Code
		apiErr.Message = eb.Error.Message
	} else {
		apiErr.Message = strings.TrimSpace(string(body))
	}

	return apiErr
}

// Error reports the error as "<message> (<code>)", so that an *Error reads
// well on its own and still names its code, or as whichever of the two it
// has.
func (e *Error) Error() string {
	switch {
	case e.Message != "" && e.Code != "":
		return fmt.Sprintf("%s (%s)", e.Message, e.Code)
	case e.Message != "":
		return e.Message
	case e.Code != "":
		return string(e.Code)
	default:
		return "error with neither a code nor a message"
	}
}

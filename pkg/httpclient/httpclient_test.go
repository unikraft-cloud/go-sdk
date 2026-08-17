// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package httpclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPClientKeepsConnectionsAlive(t *testing.T) {
	t.Parallel()

	c := NewHTTPClient()
	require.NotNil(t, c)

	transport, ok := c.Transport.(*http.Transport)
	require.True(t, ok, "expected *http.Transport")
	assert.False(t, transport.DisableKeepAlives)
	assert.False(t, transport.TLSClientConfig != nil && transport.TLSClientConfig.InsecureSkipVerify)
}

func TestNewInsecureHTTPClientSkipsCertVerification(t *testing.T) {
	t.Parallel()

	c := NewInsecureHTTPClient()
	require.NotNil(t, c)

	transport, ok := c.Transport.(*http.Transport)
	require.True(t, ok, "expected *http.Transport")
	require.NotNil(t, transport.TLSClientConfig)
	assert.True(t, transport.TLSClientConfig.InsecureSkipVerify)
}

func TestHTTPClientInterfaceSatisfiedByStdlibClient(t *testing.T) {
	var _ HTTPClient = NewHTTPClient()
	var _ HTTPClient = NewInsecureHTTPClient()
}

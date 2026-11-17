// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupWithClientAndNames(t *testing.T) {
	g := New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"})

	assert.Equal(t, []string{"fra", "waw"}, g.Names())
}

func TestGroupGetByName(t *testing.T) {
	g := New[fakeClient]().WithClient("fra", fakeClient{id: "fra"})

	client, err := g.getByName("fra")
	require.NoError(t, err)
	assert.Equal(t, "fra", client.id)

	_, err = g.getByName("missing")
	require.ErrorContains(t, err, "missing")
}

func TestGroupFilterKeepsOnlyRequestedNames(t *testing.T) {
	g := New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"}).
		WithClient("ams", fakeClient{id: "ams"})

	filtered := g.Filter([]string{"ams", "fra"})
	assert.Equal(t, []string{"ams", "fra"}, filtered.Names())
}

func TestGroupFilterIgnoresUnknownNames(t *testing.T) {
	g := New[fakeClient]().WithClient("fra", fakeClient{id: "fra"})

	filtered := g.Filter([]string{"fra", "does-not-exist"})
	assert.Equal(t, []string{"fra"}, filtered.Names())
}

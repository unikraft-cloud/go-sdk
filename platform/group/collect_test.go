// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectMetroReturnsTheClientsResult(t *testing.T) {
	g := testGroup()

	result, err := CollectMetro(context.Background(), g, "waw", func(ctx context.Context, c fakeClient) (string, error) {
		return strings.ToUpper(c.id), nil
	})
	require.NoError(t, err)
	assert.Equal(t, "WAW", result)
}

func TestCollectMetroPropagatesError(t *testing.T) {
	g := testGroup()
	boom := errors.New("boom")

	_, err := CollectMetro(context.Background(), g, "fra", func(ctx context.Context, c fakeClient) (string, error) {
		return "", boom
	})
	require.ErrorIs(t, err, boom)
}

func TestCollectAllPreservesClientOrder(t *testing.T) {
	g := testGroup() // registered in order: fra, waw

	results, err := CollectAll(context.Background(), g, func(ctx context.Context, c fakeClient) (string, error) {
		return c.id, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"fra", "waw"}, results)
}

func TestCollectAllOnEmptyGroupReturnsEmptySlice(t *testing.T) {
	g := New[fakeClient]()
	results, err := CollectAll(context.Background(), g, func(ctx context.Context, c fakeClient) (string, error) {
		return c.id, nil
	})
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestCollectAllSlicesConcatenatesInClientOrder(t *testing.T) {
	g := testGroup()

	results, err := CollectAllSlices(context.Background(), g, func(ctx context.Context, c fakeClient) ([]string, error) {
		return []string{c.id + "-1", c.id + "-2"}, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"fra-1", "fra-2", "waw-1", "waw-2"}, results)
}

func TestCollectMetrosFiltersBeforeCollecting(t *testing.T) {
	g := New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"}).
		WithClient("ams", fakeClient{id: "ams"})

	results, err := CollectMetros(context.Background(), g, []string{"ams", "fra"}, func(ctx context.Context, c fakeClient) (string, error) {
		return c.id, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"ams", "fra"}, results)
}

func TestCollectMetrosSlicesFlattensPerMetroResults(t *testing.T) {
	g := New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"})

	results, err := CollectMetrosSlices(context.Background(), g, []string{"waw"}, func(ctx context.Context, c fakeClient) ([]int, error) {
		return []int{1, 2, 3}, nil
	})
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, results)
}

func TestCollectRefsReturnsPerClientResultsAndAggregatesFoundRefs(t *testing.T) {
	g := testGroup()

	refs := Refs{{Metro: "fra", Name: "foo"}, {Name: "bar"}}
	results, err := CollectRefs(context.Background(), g, refs, func(ctx context.Context, c fakeClient, refs Refs) (string, Refs, error) {
		return c.id, refs, nil
	})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"fra", "waw"}, results)
}

func TestCollectRefsReturnsErrRefNotFound(t *testing.T) {
	g := testGroup()

	_, err := CollectRefs(context.Background(), g, Refs{{Name: "missing"}}, func(ctx context.Context, c fakeClient, refs Refs) (string, Refs, error) {
		return c.id, Refs{}, nil
	})
	require.Error(t, err)
	var notFound ErrRefNotFound
	require.ErrorAs(t, err, &notFound)
}

func TestCollectRefsSlicesFlattensAcrossClients(t *testing.T) {
	g := testGroup()

	refs := Refs{{Metro: "fra", Name: "foo"}, {Metro: "waw", Name: "bar"}}
	results, err := CollectRefsSlices(context.Background(), g, refs, func(ctx context.Context, c fakeClient, refs Refs) ([]string, Refs, error) {
		var out []string
		for _, r := range refs {
			out = append(out, c.id+":"+r.Name)
		}
		return out, refs, nil
	})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"fra:foo", "waw:bar"}, results)
}

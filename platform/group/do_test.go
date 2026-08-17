// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"errors"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testGroup() *Group[fakeClient] {
	return New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"})
}

func TestDoMetroCallsOnlyTheNamedClient(t *testing.T) {
	g := testGroup()

	var seen string
	err := DoMetro(context.Background(), g, "waw", func(ctx context.Context, c fakeClient) error {
		seen = c.id
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, "waw", seen)
}

func TestDoMetroUnknownNameErrors(t *testing.T) {
	g := testGroup()

	err := DoMetro(context.Background(), g, "missing", func(ctx context.Context, c fakeClient) error {
		t.Fatal("fn must not be called for an unknown metro")
		return nil
	})
	assert.ErrorContains(t, err, "missing")
}

func TestDoMetroWrapsCallbackError(t *testing.T) {
	g := testGroup()
	boom := errors.New("boom")

	err := DoMetro(context.Background(), g, "fra", func(ctx context.Context, c fakeClient) error {
		return boom
	})
	require.Error(t, err)
	assert.ErrorIs(t, err, boom)
	assert.ErrorContains(t, err, `"fra"`)
}

func TestDoAllCallsEveryClient(t *testing.T) {
	g := testGroup()

	var mu sync.Mutex
	var seen []string
	err := DoAll(context.Background(), g, func(ctx context.Context, c fakeClient) error {
		mu.Lock()
		seen = append(seen, c.id)
		mu.Unlock()
		return nil
	})
	require.NoError(t, err)
	sort.Strings(seen)
	assert.Equal(t, []string{"fra", "waw"}, seen)
}

func TestDoAllJoinsAllErrors(t *testing.T) {
	g := testGroup()
	fraErr := errors.New("fra broke")
	wawErr := errors.New("waw broke")

	err := DoAll(context.Background(), g, func(ctx context.Context, c fakeClient) error {
		if c.id == "fra" {
			return fraErr
		}
		return wawErr
	})
	require.Error(t, err)
	assert.ErrorIs(t, err, fraErr)
	assert.ErrorIs(t, err, wawErr)
}

func TestDoAllOnEmptyGroupSucceeds(t *testing.T) {
	g := New[fakeClient]()
	err := DoAll(context.Background(), g, func(ctx context.Context, c fakeClient) error {
		t.Fatal("fn must not be called for an empty group")
		return nil
	})
	assert.NoError(t, err)
}

func TestDoMetrosCallsOnlyTheFilteredClients(t *testing.T) {
	g := New[fakeClient]().
		WithClient("fra", fakeClient{id: "fra"}).
		WithClient("waw", fakeClient{id: "waw"}).
		WithClient("ams", fakeClient{id: "ams"})

	var mu sync.Mutex
	var seen []string
	err := DoMetros(context.Background(), g, []string{"fra", "ams"}, func(ctx context.Context, c fakeClient) error {
		mu.Lock()
		seen = append(seen, c.id)
		mu.Unlock()
		return nil
	})
	require.NoError(t, err)
	sort.Strings(seen)
	assert.Equal(t, []string{"ams", "fra"}, seen)
}

func TestDoRefsRoutesScopedRefsToTheNamedClientOnly(t *testing.T) {
	g := testGroup()

	var mu sync.Mutex
	handled := map[string]Refs{}
	err := DoRefs(context.Background(), g, Refs{{Metro: "fra", Name: "foo"}}, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		mu.Lock()
		handled[c.id] = refs
		mu.Unlock()
		return refs, nil
	})
	require.NoError(t, err)
	assert.Equal(t, Refs{{Metro: "fra", Name: "foo"}}, handled["fra"])
	assert.Nil(t, handled["waw"], "waw must not be contacted for a ref scoped to fra")
}

func TestDoRefsBroadcastsUnscopedRefsToAllClients(t *testing.T) {
	g := testGroup()

	var mu sync.Mutex
	handledBy := map[string]bool{}
	err := DoRefs(context.Background(), g, Refs{{Name: "bar"}}, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		mu.Lock()
		handledBy[c.id] = len(refs) == 1
		mu.Unlock()
		return refs, nil
	})
	require.NoError(t, err)
	assert.True(t, handledBy["fra"])
	assert.True(t, handledBy["waw"])
}

func TestDoRefsSucceedsWhenAllRefsAreReportedFound(t *testing.T) {
	g := testGroup()

	refs := Refs{{Metro: "fra", Name: "foo"}, {Name: "bar"}}
	err := DoRefs(context.Background(), g, refs, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		return refs, nil // echo back everything as found
	})
	assert.NoError(t, err)
}

func TestDoRefsReturnsErrRefNotFoundWhenNoClientReportsIt(t *testing.T) {
	g := testGroup()

	refs := Refs{{Name: "missing"}}
	err := DoRefs(context.Background(), g, refs, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		return nil, nil // never found
	})

	require.Error(t, err)
	var notFound ErrRefNotFound
	require.ErrorAs(t, err, &notFound)
	assert.Equal(t, Refs{{Name: "missing"}}, notFound.Refs)
}

func TestDoRefsPartiallyFoundReportsOnlyMissingRefs(t *testing.T) {
	g := testGroup()

	refs := Refs{{Name: "found"}, {Name: "missing"}}
	err := DoRefs(context.Background(), g, refs, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		var found Refs
		for _, r := range refs {
			if r.Name == "found" {
				found = append(found, r)
			}
		}
		return found, nil
	})

	require.Error(t, err)
	var notFound ErrRefNotFound
	require.ErrorAs(t, err, &notFound)
	assert.Equal(t, Refs{{Name: "missing"}}, notFound.Refs)
}

func TestDoRefsScopedToUnknownMetroErrorsWithoutCallingAnyClient(t *testing.T) {
	g := testGroup()

	err := DoRefs(context.Background(), g, Refs{{Metro: "missing", Name: "x"}}, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		t.Fatal("fn must not be called when the ref's metro is unknown")
		return nil, nil
	})
	require.Error(t, err)
	assert.ErrorContains(t, err, "missing")

	var notFound ErrRefNotFound
	assert.False(t, errors.As(err, &notFound), "an unknown metro is a hard error, not a not-found result")
}

func TestDoRefsPropagatesCallbackError(t *testing.T) {
	g := testGroup()
	boom := errors.New("boom")

	err := DoRefs(context.Background(), g, Refs{{Metro: "fra", Name: "x"}}, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		return nil, boom
	})
	require.Error(t, err)
	assert.ErrorIs(t, err, boom)
}

func TestDoRefsEmptyRefsSucceedsTrivially(t *testing.T) {
	g := testGroup()

	err := DoRefs(context.Background(), g, Refs{}, func(ctx context.Context, c fakeClient, refs Refs) (Refs, error) {
		t.Fatal("fn must not be called when there are no refs")
		return Refs{}, nil
	})
	assert.NoError(t, err)
}

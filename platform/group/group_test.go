// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/cloud/sdk/platform/group"
)

type fakeClient struct {
	platform.Client
	name string
}

type notFoundErr struct {
	err group.ErrRefNotFound
}

func (e notFoundErr) Error() string { return "custom: " + e.err.Error() }
func (e notFoundErr) Unwrap() error { return e.err }

func newGroup() *group.Group[fakeClient] {
	return group.New[fakeClient]().
		WithClient("fra", fakeClient{name: "fra"}).
		WithClient("sfo", fakeClient{name: "sfo"})
}

func findOnly(names ...string) func(context.Context, fakeClient, group.Refs) (group.Refs, error) {
	return func(ctx context.Context, c fakeClient, refs group.Refs) (group.Refs, error) {
		var found group.Refs
		for _, ref := range refs {
			for _, name := range names {
				if ref.Name == name {
					found = append(found, ref)
				}
			}
		}
		return found, nil
	}
}

func TestDoRefsNotFound(t *testing.T) {
	refs := group.Refs{{Name: "a"}, {Name: "b"}}

	err := group.DoRefs(t.Context(), newGroup(), refs, findOnly("a"))
	notFound, ok := errors.AsType[group.ErrRefNotFound](err)
	require.True(t, ok)
	assert.Equal(t, group.Refs{{Name: "b"}}, notFound.Refs)

	require.NoError(t, group.DoRefs(t.Context(), newGroup(), refs, findOnly("a", "b")))
}

func TestDoRefsWithNotFound(t *testing.T) {
	refs := group.Refs{{Name: "a"}, {Name: "b"}}
	g := newGroup().WithNotFound(func(err group.ErrRefNotFound) error {
		return notFoundErr{err: err}
	})

	for name, g := range map[string]*group.Group[fakeClient]{
		"group":    g,
		"filtered": g.Filter([]string{"fra"}),
	} {
		t.Run(name, func(t *testing.T) {
			err := group.DoRefs(t.Context(), g, refs, findOnly("a"))
			custom, ok := errors.AsType[notFoundErr](err)
			require.True(t, ok)
			assert.Equal(t, group.Refs{{Name: "b"}}, custom.err.Refs)
			_, ok = errors.AsType[group.ErrRefNotFound](err)
			assert.True(t, ok)

			_, err = group.CollectRefs(t.Context(), g, refs, func(ctx context.Context, c fakeClient, refs group.Refs) (string, group.Refs, error) {
				found, err := findOnly("a")(ctx, c, refs)
				return c.name, found, err
			})
			_, ok = errors.AsType[notFoundErr](err)
			assert.True(t, ok)
		})
	}
}

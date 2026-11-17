// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeListing serves pages of items the way the platform API does: count
// items from the inclusive from cursor.
type fakeListing struct {
	items []string
	// removeAfter, when set, removes these items after the first fetch.
	removeAfter []string
	calls       int
}

func (f *fakeListing) fetch(_ context.Context, count uint32, from *string) ([]string, error) {
	f.calls++
	items := f.items
	if f.calls > 1 {
		items = slices.DeleteFunc(slices.Clone(items), func(s string) bool { return slices.Contains(f.removeAfter, s) })
	}
	start := 0
	if from != nil {
		start = slices.Index(items, *from)
		if start < 0 {
			return nil, nil
		}
	}
	end := min(start+int(count), len(items))
	return items[start:end], nil
}

func collectAll(seq func(func(string, error) bool)) ([]string, error) {
	var out []string
	for item, err := range seq {
		if err != nil {
			return out, err
		}
		out = append(out, item)
	}
	return out, nil
}

func TestPaginate(t *testing.T) {
	t.Parallel()

	letters := []string{"a", "b", "c", "d", "e", "f", "g"}

	tests := []struct {
		name     string
		listing  fakeListing
		pageSize int
		want     []string
		wantErr  error
	}{
		{name: "empty", listing: fakeListing{}, pageSize: 3},
		{name: "one short page", listing: fakeListing{items: letters[:2]}, pageSize: 3, want: letters[:2]},
		{name: "exact pages", listing: fakeListing{items: letters[:6]}, pageSize: 3, want: letters[:6]},
		{name: "several pages", listing: fakeListing{items: letters}, pageSize: 3, want: letters},
		{name: "default page size", listing: fakeListing{items: letters}, want: letters},
		{
			name:     "cursor deleted between pages",
			listing:  fakeListing{items: letters, removeAfter: []string{"c"}},
			pageSize: 3,
			want:     []string{"a", "b", "c", "d", "e", "f", "g"},
		},
		{
			name:     "every read item deleted",
			listing:  fakeListing{items: letters, removeAfter: []string{"a", "b", "c"}},
			pageSize: 3,
			want:     []string{"a", "b", "c"},
			wantErr:  errListingLost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := collectAll(paginate(t.Context(), tt.pageSize, tt.listing.fetch, func(s string) string { return s }))
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPaginateStopsWhenConsumerBreaks(t *testing.T) {
	t.Parallel()

	listing := fakeListing{items: []string{"a", "b", "c", "d"}}
	for item := range paginate(t.Context(), 2, listing.fetch, func(s string) string { return s }) {
		if item == "b" {
			break
		}
	}
	assert.Equal(t, 1, listing.calls)
}

func TestPaginateYieldsFetchError(t *testing.T) {
	t.Parallel()

	boom := errors.New("boom")
	fetch := func(context.Context, uint32, *string) ([]string, error) { return nil, boom }
	_, err := collectAll(paginate(t.Context(), 2, fetch, func(s string) string { return s }))
	assert.ErrorIs(t, err, boom)
}

func TestListSendsPagingQuery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		path  string
		key   string
		list  func(c Client) error
		query string
	}{
		{
			name: "instances", path: "/v1/instances", key: "instances",
			query: "count=2&details=true&tags=web",
			list: func(c Client) error {
				return drain(c.ListInstances(t.Context(), WithListPageSize(2), WithListDetails(), WithListTags("web")))
			},
		},
		{
			name: "volumes", path: "/v1/volumes", key: "volumes", query: "count=2",
			list: func(c Client) error { return drain(c.ListVolumes(t.Context(), WithListPageSize(2))) },
		},
		{
			name: "service groups", path: "/v1/services", key: "service_groups", query: "count=2",
			list: func(c Client) error {
				return drain(c.ListServiceGroups(t.Context(), WithListPageSize(2), WithListTags("ignored")))
			},
		},
		{
			name: "certificates", path: "/v1/certificates", key: "certificates", query: "count=2",
			list: func(c Client) error { return drain(c.ListCertificates(t.Context(), WithListPageSize(2))) },
		},
		{
			name: "templates", path: "/v1/instances/templates", key: "instances", query: "count=2",
			list: func(c Client) error { return drain(c.ListTemplates(t.Context(), WithListPageSize(2))) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
				writeJSON(t, w, http.StatusOK, okWithList(tt.key, map[string]any{"uuid": "u1"}))
			})

			require.NoError(t, tt.list(newClientAgainst(srv)))
			require.Len(t, *reqs, 1)
			assert.Equal(t, tt.path, (*reqs)[0].Path)
			assert.Equal(t, tt.query, (*reqs)[0].Query)
		})
	}
}

func TestListInstancesFollowsCursor(t *testing.T) {
	t.Parallel()

	srv, reqs := newTestServer(t, func(w http.ResponseWriter, r *recordedRequest) {
		page := []map[string]any{{"uuid": "u1"}, {"uuid": "u2"}}
		if r.Query != "count=2" {
			page = []map[string]any{{"uuid": "u2"}, {"uuid": "u3"}}
		}
		writeJSON(t, w, http.StatusOK, map[string]any{"status": "success", "data": map[string]any{"instances": page}})
	})

	var uuids []string
	for inst, err := range newClientAgainst(srv).ListInstances(t.Context(), WithListPageSize(2)) {
		require.NoError(t, err)
		uuids = append(uuids, inst.Uuid)
	}
	assert.Equal(t, []string{"u1", "u2", "u3"}, uuids)
	require.Len(t, *reqs, 2)
	assert.Equal(t, "count=3&from=u2", (*reqs)[1].Query)
}

func drain[T any](seq func(func(T, error) bool)) error {
	for _, err := range seq {
		if err != nil {
			return err
		}
	}
	return nil
}

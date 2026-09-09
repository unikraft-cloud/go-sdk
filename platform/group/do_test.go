// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"errors"
	"sort"
	"sync"
	"testing"

	"unikraft.com/cloud/sdk/platform"
)

// fakeClient satisfies platform.Client without implementing any behavior;
// the group package only routes to clients, it never calls them itself.
type fakeClient struct {
	platform.Client
	id string
}

func buildGroup(names []string, wildcards ...string) *Group[fakeClient] {
	g := New[fakeClient]()
	for _, name := range names {
		g = g.WithClient(name, fakeClient{id: name})
	}
	for _, name := range wildcards {
		g = g.WithWildcardClient(name, fakeClient{id: name})
	}
	return g
}

// runDoRefs performs DoRefs recording which refs each client received, with
// per-client callbacks to control the found refs that are reported back.
func runDoRefs(t *testing.T, g *Group[fakeClient], refs Refs, found map[string]Refs) (map[string]Refs, error) {
	t.Helper()
	var mu sync.Mutex
	received := make(map[string]Refs)
	err := DoRefs(context.Background(), g, refs, func(_ context.Context, c fakeClient, refs Refs) (Refs, error) {
		mu.Lock()
		received[c.id] = append(received[c.id], refs...)
		mu.Unlock()
		if found == nil {
			return refs, nil
		}
		return found[c.id], nil
	})
	return received, err
}

func names(rs Refs) []string {
	ss := rs.Strings()
	sort.Strings(ss)
	return ss
}

func TestDoRefsRouting(t *testing.T) {
	tests := []struct {
		name      string
		clients   []string
		wildcards []string
		refs      Refs
		expected  map[string][]string // client id -> ref strings routed to it
		wantErr   bool
	}{
		{
			name:    "metro-scoped ref routes to its metro client",
			clients: []string{"fra", "dus"},
			refs:    Refs{{Metro: "fra", Name: "foo"}},
			expected: map[string][]string{
				"fra": {"foo"},
			},
		},
		{
			name:    "unscoped ref broadcasts to all clients",
			clients: []string{"fra", "dus"},
			refs:    Refs{{Name: "foo"}},
			expected: map[string][]string{
				"fra": {"foo"},
				"dus": {"foo"},
			},
		},
		{
			name:    "unknown metro errors",
			clients: []string{"fra", "dus"},
			refs:    Refs{{Metro: "lhr", Name: "foo"}},
			wantErr: true,
		},
		{
			name:    "metro-scoped ref fans out to the metro's node clients",
			clients: []string{"fra/n1", "fra/n2", "dus/n1"},
			refs:    Refs{{Metro: "fra", Name: "foo"}},
			expected: map[string][]string{
				"fra/n1": {"foo"},
				"fra/n2": {"foo"},
			},
		},
		{
			name:    "node-scoped ref routes to the exact node client",
			clients: []string{"fra/n1", "fra/n2"},
			refs:    Refs{{Metro: "fra", Node: "n2", Name: "foo"}},
			expected: map[string][]string{
				"fra/n2": {"foo"},
			},
		},
		{
			name:    "node-scoped ref prefers the node client over the metro client",
			clients: []string{"fra", "fra/n1"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo"}},
			expected: map[string][]string{
				"fra/n1": {"foo"},
			},
		},
		{
			name:    "node-scoped ref falls back to the metro client",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo"}},
			expected: map[string][]string{
				"fra": {"foo"},
			},
		},
		{
			name:    "unknown node among node clients errors",
			clients: []string{"fra/n1", "fra/n2"},
			refs:    Refs{{Metro: "fra", Node: "n9", Name: "foo"}},
			wantErr: true,
		},
		{
			name:      "node-scoped miss falls back to the wildcard client",
			clients:   []string{"fra/n1"},
			wildcards: []string{"*"},
			refs:      Refs{{Metro: "fra", Node: "n9", Name: "foo"}},
			expected: map[string][]string{
				"*": {"foo"},
			},
		},
		{
			name:    "node-scoped ref never fans out to sibling node clients",
			clients: []string{"fra/n1", "fra/n2"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo"}},
			expected: map[string][]string{
				"fra/n1": {"foo"},
			},
		},
		{
			name:      "metro-scoped ref falls back to the wildcard client",
			wildcards: []string{"*"},
			refs:      Refs{{Metro: "fra", Name: "foo"}},
			expected: map[string][]string{
				"*": {"foo"},
			},
		},
		{
			name:      "specific client preferred over wildcard",
			clients:   []string{"fra"},
			wildcards: []string{"*"},
			refs:      Refs{{Metro: "fra", Name: "foo"}},
			expected: map[string][]string{
				"fra": {"foo"},
			},
		},
		{
			name:      "unscoped ref broadcasts to wildcard too",
			clients:   []string{"fra"},
			wildcards: []string{"*"},
			refs:      Refs{{Name: "foo"}},
			expected: map[string][]string{
				"fra": {"foo"},
				"*":   {"foo"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := buildGroup(tt.clients, tt.wildcards...)
			// Echo back the received refs so every requested ref is found.
			received, err := runDoRefs(t, g, tt.refs, nil)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got none (received: %v)", received)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(received) != len(tt.expected) {
				t.Fatalf("refs routed to clients %v, expected %v", received, tt.expected)
			}
			for id, want := range tt.expected {
				got := names(received[id])
				sort.Strings(want)
				if len(got) != len(want) {
					t.Fatalf("client %q received %v, expected %v", id, got, want)
				}
				for i := range got {
					if got[i] != want[i] {
						t.Fatalf("client %q received %v, expected %v", id, got, want)
					}
				}
			}
		})
	}
}

func TestDoRefsNotFound(t *testing.T) {
	tests := []struct {
		name      string
		clients   []string
		wildcards []string
		refs      Refs
		found     map[string]Refs
		notFound  []string // expected not-found ref strings, empty for success
	}{
		{
			name:    "found ref without metro is stamped with the client scope",
			clients: []string{"fra", "dus"},
			refs:    Refs{{Metro: "fra", Name: "foo"}},
			found: map[string]Refs{
				"fra": {{Name: "foo"}},
			},
		},
		{
			name:    "node-scoped request satisfied by unattributed response",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo"}},
			found: map[string]Refs{
				"fra": {{Name: "foo"}},
			},
		},
		{
			name:    "node-scoped request satisfied by node-attributed response",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo"}},
			found: map[string]Refs{
				"fra": {{Node: "n1", Name: "foo"}},
			},
		},
		{
			name:    "response attributed to a different node does not satisfy",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", Node: "n1", Name: "foo", Display: "fra/n1/foo"}},
			found: map[string]Refs{
				"fra": {{Node: "n2", Name: "foo"}},
			},
			notFound: []string{"fra/n1/foo"},
		},
		{
			name:    "result from a node client does not satisfy another node's request",
			clients: []string{"dus/n1", "dus/n2"},
			refs: Refs{
				{Metro: "dus", Node: "n1", Name: "foo", Display: "dus/n1/foo"},
				{Metro: "dus", Node: "n2", Name: "foo", Display: "dus/n2/foo"},
			},
			found: map[string]Refs{
				// n1's unattributed result is stamped with the client's node
				"dus/n1": {{Name: "foo"}},
			},
			notFound: []string{"dus/n2/foo"},
		},
		{
			name:    "metro-scoped request satisfied by node client",
			clients: []string{"fra/n1", "fra/n2"},
			refs:    Refs{{Metro: "fra", Name: "foo"}},
			found: map[string]Refs{
				"fra/n2": {{Name: "foo"}},
			},
		},
		{
			name:     "missing ref reported",
			clients:  []string{"fra"},
			refs:     Refs{{Metro: "fra", Name: "foo", Display: "fra/foo"}, {Metro: "fra", Name: "bar"}},
			found:    map[string]Refs{"fra": {{Name: "bar"}}},
			notFound: []string{"fra/foo"},
		},
		{
			name:      "wildcard response with metro attribution satisfies",
			wildcards: []string{"*"},
			refs:      Refs{{Metro: "fra", Name: "foo"}},
			found: map[string]Refs{
				"*": {{Metro: "fra", Name: "foo"}},
			},
		},
		{
			name:      "wildcard response without attribution does not satisfy scoped request",
			wildcards: []string{"*"},
			refs:      Refs{{Metro: "fra", Name: "foo", Display: "fra/foo"}},
			found: map[string]Refs{
				"*": {{Name: "foo"}},
			},
			notFound: []string{"fra/foo"},
		},
		{
			name:      "wildcard response satisfies unscoped request",
			wildcards: []string{"*"},
			refs:      Refs{{Name: "foo"}},
			found: map[string]Refs{
				"*": {{Metro: "fra", Node: "n1", Name: "foo"}},
			},
		},
		{
			name:    "uuid-only response satisfies uuid request",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", UUID: "6b2bdc53-6411-478b-83c8-62714631a766"}},
			found: map[string]Refs{
				"fra": {{UUID: "6b2bdc53-6411-478b-83c8-62714631a766"}},
			},
		},
		{
			name:    "response with both identifiers satisfies name request",
			clients: []string{"fra"},
			refs:    Refs{{Metro: "fra", Name: "foo"}},
			found: map[string]Refs{
				"fra": {{Name: "foo", UUID: "6b2bdc53-6411-478b-83c8-62714631a766"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := buildGroup(tt.clients, tt.wildcards...)
			_, err := runDoRefs(t, g, tt.refs, tt.found)
			if len(tt.notFound) == 0 {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			var refErr ErrRefNotFound
			if !errors.As(err, &refErr) {
				t.Fatalf("expected ErrRefNotFound, got %v", err)
			}
			got := names(refErr.Refs)
			want := append([]string{}, tt.notFound...)
			sort.Strings(want)
			if len(got) != len(want) {
				t.Fatalf("not found %v, expected %v", got, want)
			}
			for i := range got {
				if got[i] != want[i] {
					t.Fatalf("not found %v, expected %v", got, want)
				}
			}
		})
	}
}

func TestDoMetroResolution(t *testing.T) {
	tests := []struct {
		name      string
		clients   []string
		wildcards []string
		metro     string
		expected  string // client id expected to serve the call
		wantErr   bool
	}{
		{
			name:     "exact metro client",
			clients:  []string{"fra", "dus"},
			metro:    "fra",
			expected: "fra",
		},
		{
			name:     "exact node client",
			clients:  []string{"fra/n1", "fra/n2"},
			metro:    "fra/n2",
			expected: "fra/n2",
		},
		{
			name:     "single node client of the metro",
			clients:  []string{"fra/n1", "dus/n1"},
			metro:    "fra",
			expected: "fra/n1",
		},
		{
			name:    "multiple node clients are ambiguous",
			clients: []string{"fra/n1", "fra/n2"},
			metro:   "fra",
			wantErr: true,
		},
		{
			name:      "wildcard fallback",
			wildcards: []string{"*"},
			metro:     "fra",
			expected:  "*",
		},
		{
			name:    "unknown metro errors",
			clients: []string{"fra"},
			metro:   "dus",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := buildGroup(tt.clients, tt.wildcards...)
			var served string
			err := DoMetro(context.Background(), g, tt.metro, func(_ context.Context, c fakeClient) error {
				served = c.id
				return nil
			})
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got client %q", served)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if served != tt.expected {
				t.Fatalf("served by %q, expected %q", served, tt.expected)
			}
		})
	}
}

func TestFilterPreservesWildcard(t *testing.T) {
	g := buildGroup([]string{"fra"}, "*")
	filtered := g.Filter([]string{"*"})

	received, err := runDoRefs(t, filtered, Refs{{Metro: "dus", Name: "foo"}}, map[string]Refs{
		"*": {{Metro: "dus", Name: "foo"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(received["*"]) != 1 {
		t.Fatalf("wildcard client did not receive the scoped ref: %v", received)
	}
}

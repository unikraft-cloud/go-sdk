// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package group provides functionality to make grouped operations across
// multiple platform clients.
//
// The general pattern is to create a Group from multiple platform.Client, then
// call one of the (Do|Collect)(Metro|Refs|All)(Slices)? functions to perform
// an operation across them.
//
//   - The Do functions are used when the operation does not return a value,
//     while the Collect functions are used when the operation returns a value.
//   - The Slices variants of the Collect functions are used when the operation
//     returns a slice of values, and the results from all clients can be
//     concatenated into a single slice.
//   - The Metro variants of the functions are used when the operation is to be
//     performed on a single metro/client.
//   - The Metros variants of the functions are used when the operation is to be
//     performed on a specific subset of metros/clients.
//   - The Refs variants of the functions are used when the operation is to be
//     performed on a set of refs, which may be distributed across multiple
//     clients.
//   - The All variants of the functions are used when the operation is to be
//     performed on all clients in the group.
//
// Clients are named after the scope of the endpoint they talk to: a
// metro-level client is named after its metro (e.g. "fra"), while a client
// that talks directly to a single node within a metro is named
// "<metro>/<node>" (e.g. "fra/node1"). A wildcard client (added with
// WithWildcardClient) talks to an endpoint able to answer requests for any
// metro, such as a global API router, and receives the requests of every
// metro that has no more specific client.
package group

import (
	"fmt"
	"strings"

	"unikraft.com/cloud/sdk/platform"
)

// New creates a new empty Group of platform.Client.
func New[C platform.Client]() *Group[C] {
	return &Group[C]{
		clients:    []namedClient[C]{},
		clientsMap: map[string]int{},
	}
}

// WithClient adds a named client to the Group.
//
// The name used here should be unique within the Group. Clients scoped to a
// single node within a metro should be named "<metro>/<node>". Refs select
// clients through their separate Metro and Node fields — a node client named
// "fra/node1" is selected by Ref{Metro: "fra", Node: "node1"}, never by
// putting "fra/node1" into Metro. The Metro variants of Do and Collect
// accept either a metro or a "<metro>/<node>" name.
func (c *Group[C]) WithClient(name string, client C) *Group[C] {
	return c.withClient(name, client, false)
}

// WithWildcardClient adds a client able to answer requests for any metro,
// such as a global API router. Requests scoped to a metro (or node) that has
// no dedicated client in the Group are routed to the wildcard client.
func (c *Group[C]) WithWildcardClient(name string, client C) *Group[C] {
	return c.withClient(name, client, true)
}

func (c *Group[C]) withClient(name string, client C, wildcard bool) *Group[C] {
	c.clients = append(c.clients, namedClient[C]{
		Name:     name,
		Client:   client,
		Wildcard: wildcard,
	})
	c.clientsMap[name] = len(c.clients) - 1
	return c
}

// Names returns the names of all clients in the Group.
func (c *Group[C]) Names() []string {
	names := make([]string, len(c.clients))
	for i, nc := range c.clients {
		names[i] = nc.Name
	}
	return names
}

// Filter creates a new group containing only the clients whose names are in
// the provided list.
func (c *Group[C]) Filter(names []string) *Group[C] {
	g := New[C]()
	for _, name := range names {
		if idx, ok := c.clientsMap[name]; ok {
			nc := c.clients[idx]
			g = g.withClient(nc.Name, nc.Client, nc.Wildcard)
		}
	}
	return g
}

type Group[C platform.Client] struct {
	clients    []namedClient[C]
	clientsMap map[string]int
}

type namedClient[C platform.Client] struct {
	Name     string
	Client   C
	Wildcard bool
}

// scope returns the metro and node parts of the client's name. Both are
// empty for wildcard clients.
func (nc namedClient[C]) scope() (metro, node string) {
	if nc.Wildcard {
		return "", ""
	}
	metro, node, _ = strings.Cut(nc.Name, "/")
	return metro, node
}

// resolveMetro returns the index of the single client responsible for the
// given metro (or "metro/node") name: the exact client if one exists, else
// the metro's only node client, else a wildcard client.
func (c *Group[C]) resolveMetro(name string) (int, error) {
	if idx, ok := c.clientsMap[name]; ok {
		return idx, nil
	}

	matches := c.prefixMatches(name)
	if len(matches) == 1 {
		return matches[0], nil
	}
	if len(matches) > 1 {
		names := make([]string, len(matches))
		for i, idx := range matches {
			names[i] = c.clients[idx].Name
		}
		return 0, fmt.Errorf("metro %q has multiple clients %v, specify one", name, names)
	}

	for idx, nc := range c.clients {
		if nc.Wildcard {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("failed to find client %q", name)
}

// resolveRef returns the indices of the clients that should receive the
// given ref: the most specific clients able to answer for the ref's scope.
func (c *Group[C]) resolveRef(ref Ref) ([]int, error) {
	if ref.Metro == "" {
		indices := make([]int, len(c.clients))
		for i := range c.clients {
			indices[i] = i
		}
		return indices, nil
	}

	if ref.Node != "" {
		if idx, ok := c.clientsMap[ref.Metro+"/"+ref.Node]; ok {
			return []int{idx}, nil
		}
	}
	if idx, ok := c.clientsMap[ref.Metro]; ok {
		return []int{idx}, nil
	}

	// A ref scoped to a metro whose nodes have their own clients is sent to
	// all of them. A node-scoped ref never fans out to sibling node clients:
	// with no dedicated client, it falls through to the wildcard clients.
	if ref.Node == "" {
		if matches := c.prefixMatches(ref.Metro); len(matches) > 0 {
			return matches, nil
		}
	}

	var wildcards []int
	for idx, nc := range c.clients {
		if nc.Wildcard {
			wildcards = append(wildcards, idx)
		}
	}
	if len(wildcards) > 0 {
		return wildcards, nil
	}
	name := ref.Metro
	if ref.Node != "" {
		name += "/" + ref.Node
	}
	return nil, fmt.Errorf("failed to find client %q", name)
}

// prefixMatches returns the indices of all clients scoped to nodes of the
// given metro.
func (c *Group[C]) prefixMatches(metro string) []int {
	var matches []int
	for idx, nc := range c.clients {
		if !nc.Wildcard && strings.HasPrefix(nc.Name, metro+"/") {
			matches = append(matches, idx)
		}
	}
	return matches
}

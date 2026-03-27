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
//   - The Refs variants of the functions are used when the operation is to be
//     performed on a set of refs, which may be distributed across multiple
//     clients.
//   - The All variants of the functions are used when the operation is to be
//     performed on all clients in the group.
package group

import (
	"fmt"

	"unikraft.com/cloud/sdk/platform"
)

// New creates a new empty Group of platform.Client.
func New[C platform.Client]() *Group[C] {
	return &Group[C]{
		clients:    []namedClient[C]{},
		clientsMap: map[string]C{},
	}
}

// WithClient adds a named client to the Group.
//
// The metro name used here should be unique within the Group, and is used as
// the Metro parameter in the Refs and Metro variants of Do and Collect.
func (c *Group[C]) WithClient(metro string, client C) *Group[C] {
	c.clients = append(c.clients, namedClient[C]{
		Name:   metro,
		Client: client,
	})
	c.clientsMap[metro] = client
	return c
}

type Group[C platform.Client] struct {
	clients    []namedClient[C]
	clientsMap map[string]C
}

type namedClient[C platform.Client] struct {
	Name   string
	Client C
}

func (c *Group[C]) getByName(endpoint string) (C, error) {
	var zero C
	client, ok := c.clientsMap[endpoint]
	if !ok {
		return zero, fmt.Errorf("failed to find client %q", endpoint)
	}
	return client, nil
}

type comparableClient interface {
	platform.Client
	comparable
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"fmt"
	"sync"

	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/x/joinerrgroup"
	"unikraft.com/x/log"
)

// DoMetro performs the given function fn on the client in the group
// responsible for the given metro (or "metro/node") name: the exact client
// if one exists, else the metro's only node client, else a wildcard client.
func DoMetro[C platform.Client](ctx context.Context, c *Group[C], name string, fn func(context.Context, C) error) error {
	idx, err := c.resolveMetro(name)
	if err != nil {
		return err
	}
	metroClient := c.clients[idx].Client
	logger := log.G(ctx).
		With().
		Str("metro", name).
		Logger()
	ctx = log.WithLogger(ctx, &logger)
	err = fn(ctx, metroClient)
	if err != nil {
		return fmt.Errorf("failed on %q client: %w", c.clients[idx].Name, err)
	}
	return nil
}

// DoMetros performs the given function fn across the clients in the group
// corresponding to the given metros.
func DoMetros[C platform.Client](ctx context.Context, c *Group[C], names []string, fn func(context.Context, C) error) error {
	return DoAll(ctx, c.Filter(names), fn)
}

// DoAll performs the given function fn across all clients in the group.
func DoAll[C platform.Client](ctx context.Context, c *Group[C], fn func(context.Context, C) error) error {
	eg := joinerrgroup.Group{}
	for idx, client := range c.clients {
		eg.Go(func() error {
			logger := log.G(ctx).
				With().
				Str("metro", client.Name).
				Logger()
			ctx := log.WithLogger(ctx, &logger)
			err := fn(withIndexCtx(ctx, idx), client.Client)
			if err != nil {
				return fmt.Errorf("failed on %q client: %w", client.Name, err)
			}
			return nil
		})
	}
	return eg.Wait()
}

// DoRefs performs the given function fn across all clients in the group
// distributing the refs across the clients based on the Metro and Node fields
// of each Ref. If a Ref does not have the Metro field set, it is sent to all
// clients. A Ref scoped to a metro (and optionally a node) is sent to the
// most specific clients able to answer for it: the exact node client, else
// the metro client, else the wildcard clients — with a metro-scoped Ref also
// fanning out to all of the metro's node clients before falling back to the
// wildcards.
//
// Each callback must return the list of Refs that were found on that client,
// with the Metro and Node fields set to the true origin of the resource when
// the callback knows it (e.g. from the response); empty Metro and Node fields
// are filled in from the answering client's scope. Callbacks of wildcard
// clients should always set the Metro field themselves, or scoped requests
// answered by them will be reported as not found.
//
// After all callbacks have completed, DoRefs checks that all requested Refs
// were found across the clients, returning an error if any were not found. A
// node-scoped request is satisfied by a result whose node is unattributed
// (the serving endpoint is responsible for honoring the node scope), but
// never by a result attributed to a different node.
func DoRefs[C platform.Client](ctx context.Context, c *Group[C], refs Refs, fn func(context.Context, C, Refs) (Refs, error)) error {
	targets := make([]Refs, len(c.clients))
	for _, ref := range refs {
		indices, err := c.resolveRef(ref)
		if err != nil {
			return err
		}
		for _, idx := range indices {
			targets[idx] = append(targets[idx], ref)
		}
	}

	eg := joinerrgroup.Group{}
	refMap := make(map[Ref]struct{})
	// nodeless tracks the results whose node could not be attributed (from
	// neither the response nor the client's scope): only these may satisfy a
	// node-scoped request without matching its node.
	nodeless := make(map[Ref]struct{})
	var mu sync.Mutex

	for idx, client := range c.clients {
		refs := targets[idx]
		if len(refs) == 0 {
			continue
		}

		eg.Go(func() error {
			logger := log.G(ctx).
				With().
				Str("metro", client.Name).
				Strs("refs", refs.Strings()).
				Logger()
			ctx := log.WithLogger(ctx, &logger)

			refs, err := fn(withIndexCtx(ctx, idx), client.Client, refs)
			if err != nil {
				return fmt.Errorf("failed on %q client: %w", client.Name, err)
			}

			metro, node := client.scope()
			mu.Lock()
			for _, ref := range refs {
				// fill in the origin of the resource from the client's scope
				// when the callback did not report it
				if ref.Metro == "" {
					ref.Metro = metro
					if ref.Node == "" {
						ref.Node = node
					}
				}
				ref.Display = ""
				// track all possible ref permutations that could have been
				// used to fetch this resource
				for _, variant := range ref.variants() {
					refMap[variant] = struct{}{}
					if ref.Node == "" {
						nodeless[variant] = struct{}{}
					}
				}
			}
			mu.Unlock()

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	notFound := make([]Ref, 0)
	for _, ref := range refs {
		lookup := ref
		lookup.Display = ""
		if _, ok := refMap[lookup]; ok {
			continue
		}
		// A node-scoped request answered by a client whose response did not
		// attribute a node still satisfies the request: the serving endpoint
		// is responsible for honoring the node scope. A result attributed to
		// a different node never does.
		lookup.Node = ""
		if _, ok := nodeless[lookup]; ok {
			continue
		}
		notFound = append(notFound, ref)
	}
	if len(notFound) > 0 {
		return ErrRefNotFound{Refs: notFound}
	}

	return nil
}

type ctxIndexKey struct{}

func withIndexCtx(ctx context.Context, index int) context.Context {
	return context.WithValue(ctx, ctxIndexKey{}, index)
}

func mustGetIndexCtx(ctx context.Context) int {
	index, ok := ctx.Value(ctxIndexKey{}).(int)
	if !ok {
		panic("index not found in context")
	}
	return index
}

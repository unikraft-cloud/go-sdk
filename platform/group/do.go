// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"sync"

	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/x/joinerrgroup"
	"unikraft.com/x/log"
)

// DoMetro performs the given function fn on the client in the group
// corresponding to the given metro.
func DoMetro[C platform.Client](ctx context.Context, c *Group[C], name string, fn func(context.Context, C) error) error {
	metroClient, err := c.getByName(name)
	if err != nil {
		return err
	}
	logger := log.G(ctx).
		With().
		Str("metro", name).
		Logger()
	ctx = log.WithLogger(ctx, &logger)
	return fn(ctx, metroClient)
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
			return fn(withIndexCtx(ctx, idx), client.Client)
		})
	}
	return eg.Wait()
}

// DoRefs performs the given function fn across all clients in the group
// distributing the refs across the clients based on the Metro field of each
// Ref. If a Ref does not have the Metro field set, it is sent to all clients.
//
// Each callback must return the list of Refs that were found on that client.
// After all callbacks have completed, DoRefs checks that all requested Refs
// were found across the clients, returning an error if any were not found.
func DoRefs[C comparableClient](ctx context.Context, c *Group[C], refs Refs, fn func(context.Context, C, Refs) (Refs, error)) error {
	targets := make(map[C]Refs)
	for _, ref := range refs {
		if ref.Metro != "" {
			client, err := c.getByName(ref.Metro)
			if err != nil {
				return err
			}
			targets[client] = append(targets[client], ref)
		} else {
			for _, client := range c.clients {
				targets[client.Client] = append(targets[client.Client], ref)
			}
		}
	}

	eg := joinerrgroup.Group{}
	refMap := make(map[Ref]struct{})
	var mu sync.Mutex

	for idx, client := range c.clients {
		refs, ok := targets[client.Client]
		if !ok || len(refs) == 0 {
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
				return err
			}

			mu.Lock()
			for _, ref := range refs {
				// track all possible ref permutations that could have been used to
				// fetch this resource
				ref.Metro = client.Name
				ref.Display = ""
				for _, ref := range ref.variants() {
					refMap[ref] = struct{}{}
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
		ref.Display = ""
		if _, ok := refMap[ref]; !ok {
			notFound = append(notFound, ref)
		}
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

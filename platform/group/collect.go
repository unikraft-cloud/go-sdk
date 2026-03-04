// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"

	"unikraft.com/cloud/sdk/platform"
)

// CollectMetro performs the given function fn on the client in the group
// corresponding to the given metro, collecting and returning the result.
func CollectMetro[C platform.Client, T any](ctx context.Context, c *Group[C], metro string, fn func(context.Context, C) (T, error)) (T, error) {
	var result T
	err := DoMetro(ctx, c, metro, func(ctx context.Context, client C) error {
		var err error
		result, err = fn(ctx, client)
		return err
	})
	return result, err
}

// CollectAll performs the given function fn across all clients in the group,
// collecting and returning the results.
func CollectAll[C platform.Client, T any](ctx context.Context, c *Group[C], fn func(context.Context, C) (T, error)) ([]T, error) {
	results := make([]T, len(c.clients))
	err := DoAll(ctx, c, func(ctx context.Context, client C) error {
		idx := mustGetIndexCtx(ctx)

		result, err := fn(ctx, client)
		results[idx] = result
		return err
	})
	return results, err
}

// CollectAllSlices performs the same operation as CollectAll, but for
// functions that return slices. The resulting slices from all clients are
// concatenated into a single slice and returned.
func CollectAllSlices[C platform.Client, T any](ctx context.Context, c *Group[C], fn func(context.Context, C) ([]T, error)) ([]T, error) {
	slices, err := CollectAll(ctx, c, fn)
	return flatten(slices), err
}

// CollectRefs performs the given function fn across all clients in the group
// distributing the refs across the clients based on the Metro field of each
// Ref. If a Ref does not have the Metro field set, it is sent to all clients.
//
// Each callback must return the list of Refs that were found on that client.
// After all callbacks have completed, CollectRefs checks that all requested
// Refs were found across the clients, returning an error if any were not found.
func CollectRefs[C comparableClient, T any](ctx context.Context, c *Group[C], refs Refs, fn func(context.Context, C, Refs) (T, Refs, error)) ([]T, error) {
	results := make([]T, len(c.clients))
	err := DoRefs(ctx, c, refs, func(ctx context.Context, client C, refs Refs) (Refs, error) {
		idx := mustGetIndexCtx(ctx)
		result, foundRefs, err := fn(ctx, client, refs)
		results[idx] = result
		return foundRefs, err
	})
	return results, err
}

// CollectRefsSlices performs same operation as CollectRefs, but for functions
// that return slices. The resulting slices from all clients are concatenated
// into a single slice and returned.
func CollectRefsSlices[C comparableClient, T any](ctx context.Context, c *Group[C], refs Refs, fn func(context.Context, C, Refs) ([]T, Refs, error)) ([]T, error) {
	slices, err := CollectRefs(ctx, c, refs, fn)
	return flatten(slices), err
}

func flatten[T any](slices [][]T) []T {
	total := 0
	for _, slice := range slices {
		total += len(slice)
	}
	result := make([]T, 0, total)
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

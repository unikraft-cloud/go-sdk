// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"errors"
	"iter"
	"slices"
)

const (
	defaultPageSize = 100
	// trailLength is the number of recent cursors kept to resume a listing
	// whose last cursor was deleted.
	trailLength = 32
)

var errListingLost = errors.New("the listing lost its place: the items it read last are gone; list again to start over")

// fetchPage reads up to count items, starting at the item that from names.
// A nil from starts at the first item.
type fetchPage[T any] func(ctx context.Context, count uint32, from *string) ([]T, error)

// paginate yields every item of a list endpoint, one page at a time.
//
// The platform API pages with a count and an inclusive from cursor, and a
// page shorter than count is the last one. After the first page, paginate
// asks for one more item and drops the cursor item, which the previous page
// already yielded. If the cursor item was deleted between two pages,
// paginate continues from the most recent item that still exists.
func paginate[T any](ctx context.Context, pageSize int, fetch fetchPage[T], cursor func(T) string) iter.Seq2[T, error] {
	count := uint32(defaultPageSize)
	if pageSize > 0 {
		count = uint32(pageSize)
	}
	return func(yield func(T, error) bool) {
		var zero T
		var start *string
		var trail []string
		for {
			want := count
			if start != nil {
				want++
			}
			page, err := fetch(ctx, want, start)
			if err != nil {
				yield(zero, err)
				return
			}
			if start != nil && len(page) == 0 {
				page, start, err = resume(ctx, fetch, count+1, trail, *start)
				if err != nil {
					yield(zero, err)
					return
				}
			}
			if start != nil && len(page) > 0 && cursor(page[0]) == *start {
				page = page[1:]
			}
			for _, item := range page {
				if !yield(item, nil) {
					return
				}
			}
			if len(page) < int(count) {
				return
			}

			for _, item := range page {
				if c := cursor(item); c != "" {
					trail = append(trail, c)
				}
			}
			trail = trail[max(0, len(trail)-trailLength):]
			if len(trail) == 0 || (start != nil && trail[len(trail)-1] == *start) {
				return
			}
			start = new(trail[len(trail)-1])
		}
	}
}

// resume reads on from the most recent cursor in trail that still exists.
func resume[T any](ctx context.Context, fetch fetchPage[T], count uint32, trail []string, gone string) ([]T, *string, error) {
	for _, c := range slices.Backward(trail) {
		if c == gone {
			continue
		}
		page, err := fetch(ctx, count, &c)
		if err != nil {
			return nil, nil, err
		}
		if len(page) > 0 {
			return page, &c, nil
		}
	}
	return nil, nil, errListingLost
}

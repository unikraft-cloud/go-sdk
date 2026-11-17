// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"regexp"
	"time"
)

const (
	busyFirstInterval = 500 * time.Millisecond
	busyMaxInterval   = 2 * time.Second
)

// busyPattern matches how the API reports a resource that something else
// still holds.
var busyPattern = regexp.MustCompile(`\bEBUSY\b|Unknown error -16\b`)

func isBusy(err error) bool {
	return err != nil && busyPattern.MatchString(err.Error())
}

// whileBusy runs attempt, and retries it with backoff for up to limit while
// it reports busy.
func whileBusy(ctx context.Context, limit time.Duration, attempt func() error) error {
	err := attempt()
	if limit <= 0 {
		return err
	}
	deadline := time.Now().Add(limit)
	interval := busyFirstInterval
	for isBusy(err) && time.Now().Add(interval).Before(deadline) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
		interval = min(interval*2, busyMaxInterval)
		err = attempt()
	}
	return err
}

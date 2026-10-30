// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

// WithRetry lets the tests shorten the wait between two attempts at the
// progress stream, per client, so that they stay parallel and fast.
var WithRetry = withRetry

// MaxLineSize is the bound on a line of the progress stream, for the test
// that goes past it.
const MaxLineSize = maxLineSize

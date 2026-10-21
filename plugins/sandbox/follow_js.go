// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"errors"
	"fmt"
)

// A browser's WebSocket sends no Authorization header, so nothing is started
// that could not be followed.
var errNoFollow = fmt.Errorf("sandbox: commands cannot be followed from a browser, its WebSocket carries no credentials: %w", errors.ErrUnsupported)

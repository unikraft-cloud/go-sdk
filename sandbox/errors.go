// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import "unikraft.com/cloud/sdk/platform"

// APIHTTPError is the API-wide error code enumeration.  The Sandbox API is
// served by the same metros as the Platform API and reports the same codes, so
// the type is aliased rather than duplicated: the constants are declared once,
// in platform, and a code obtained from either package can be passed to either
// package's ErrorContains and ErrorContainsOnly.
type APIHTTPError = platform.APIHTTPError

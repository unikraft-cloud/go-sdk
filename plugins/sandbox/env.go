// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import "strings"

// EnvMap turns the NAME=value records a caller holds into the map the plugin
// takes; only the first = separates a name from its value.
func EnvMap(env []string) map[string]string {
	if len(env) == 0 {
		return nil
	}

	vars := make(map[string]string, len(env))
	for _, record := range env {
		name, value, _ := strings.Cut(record, "=")
		vars[name] = value
	}
	return vars
}

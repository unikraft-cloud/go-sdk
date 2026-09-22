// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvMap(t *testing.T) {
	assert.Nil(t, EnvMap(nil), "nothing to set is not an empty environment")
	assert.Equal(t, map[string]string{
		"HOME":  "/",
		"MULTI": "a=b",
		"EMPTY": "",
	}, EnvMap([]string{"HOME=/", "MULTI=a=b", "EMPTY="}))
}

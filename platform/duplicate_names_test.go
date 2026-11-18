// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"unikraft.com/cloud/sdk/platform"
)

func TestInstance_DuplicateNames_Unmarshal(t *testing.T) {
	data := []byte(`{
		"data": {"instances": [
			{"name": "a", "env": {"FOO": "1", "BAR": "x", "FOO": "2"}}
		]}
	}`)

	var resp platform.Response[platform.GetInstancesResponseData]
	require.NoError(t, json.Unmarshal(data, &resp))

	require.Len(t, resp.Data.Instances, 1)
	assert.Equal(t, map[string]string{"FOO": "2", "BAR": "x"}, resp.Data.Instances[0].Env)
}

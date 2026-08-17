// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNameRef(t *testing.T) {
	ref := Name("my-instance")

	assert.Equal(t, "my-instance", ref.String())

	nou := ref.toNameOrUUID()
	require.NotNil(t, nou.Name)
	assert.Equal(t, "my-instance", *nou.Name)
	assert.Nil(t, nou.Uuid)
}

func TestUUIDRef(t *testing.T) {
	ref := UUID("11111111-1111-1111-1111-111111111111")

	assert.Equal(t, "11111111-1111-1111-1111-111111111111", ref.String())

	nou := ref.toNameOrUUID()
	require.NotNil(t, nou.Uuid)
	assert.Equal(t, "11111111-1111-1111-1111-111111111111", *nou.Uuid)
	assert.Nil(t, nou.Name)
}

func TestRefsWithSameStringAreDistinctTypes(t *testing.T) {
	// Name and UUID produce distinct underlying types even for the same
	// string, so a Name("x") is never confused for a UUID("x").
	n := Name("x")
	u := UUID("x")

	assert.NotEqual(t, n.toNameOrUUID(), u.toNameOrUUID())
}

func TestEmptyRefs(t *testing.T) {
	assert.Equal(t, "", Name("").String())
	assert.Equal(t, "", UUID("").String())
}

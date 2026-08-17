// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefIsZero(t *testing.T) {
	assert.True(t, Ref{}.IsZero())
	assert.True(t, Ref{Metro: "fra"}.IsZero(), "Metro alone does not identify a resource")
	assert.False(t, Ref{Name: "foo"}.IsZero())
	assert.False(t, Ref{UUID: "abc"}.IsZero())
}

func TestRefString(t *testing.T) {
	assert.Equal(t, "custom", Ref{Name: "foo", Display: "custom"}.String())
	assert.Equal(t, "foo", Ref{Name: "foo"}.String())
	assert.Equal(t, "abc", Ref{UUID: "abc"}.String())
	assert.Equal(t, "", Ref{}.String())
}

func TestRefsStrings(t *testing.T) {
	refs := Refs{{Name: "a"}, {UUID: "b"}, {Name: "c", Display: "C"}}
	assert.Equal(t, []string{"a", "b", "C"}, refs.Strings())
}

func TestRefVariantsUnscoped(t *testing.T) {
	// A Ref with no Metro only ever produces unscoped variants.
	nameOnly := Ref{Name: "foo"}
	assert.ElementsMatch(t, []Ref{{Name: "foo"}}, nameOnly.variants())

	uuidOnly := Ref{UUID: "abc"}
	assert.ElementsMatch(t, []Ref{{UUID: "abc"}}, uuidOnly.variants())

	both := Ref{Name: "foo", UUID: "abc"}
	assert.ElementsMatch(t, []Ref{{Name: "foo"}, {UUID: "abc"}, {Name: "foo", UUID: "abc"}}, both.variants())
}

func TestRefVariantsScopedIncludesBothScopedAndUnscoped(t *testing.T) {
	// A metro-scoped Ref must also produce unscoped variants, since a
	// broadcast query (no Metro) needs to match against a scoped result.
	r := Ref{Metro: "fra", Name: "foo", UUID: "abc"}
	got := r.variants()

	assert.ElementsMatch(t, []Ref{
		{Metro: "fra", Name: "foo"},
		{Metro: "fra", UUID: "abc"},
		{Metro: "fra", Name: "foo", UUID: "abc"},
		{Name: "foo"},
		{UUID: "abc"},
		{Name: "foo", UUID: "abc"},
	}, got)
}

func TestRefVariantsZeroRefIsEmpty(t *testing.T) {
	assert.Empty(t, Ref{}.variants())
	assert.Empty(t, Ref{Metro: "fra"}.variants())
}

func TestErrRefNotFoundError(t *testing.T) {
	err := ErrRefNotFound{Refs: Refs{{Name: "foo"}, {UUID: "bar"}}}
	assert.Contains(t, err.Error(), "foo")
	assert.Contains(t, err.Error(), "bar")
}

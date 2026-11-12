// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func relationsFixture() Tree {
	return Tree{
		{Kind: KindInstance, UUID: "fork-c", Source: "ck3"},
		{Kind: KindCheckpoint, UUID: "ck3", Source: "fork-a"},
		{Kind: KindInstance, UUID: "fork-b", Source: "web"},
		{Kind: KindInstance, UUID: "fork-a", Source: "ck1"},
		{Kind: KindCheckpoint, UUID: "ck2", Source: "web"},
		{Kind: KindCheckpoint, UUID: "ck1", Source: "web"},
		{Kind: KindInstance, UUID: "web"},
	}
}

func uuidsOf(nodes []Node) []string {
	out := []string{}
	for _, n := range nodes {
		out = append(out, n.UUID)
	}
	return out
}

func TestFind(t *testing.T) {
	n, ok := relationsFixture().Find("ck2")
	assert.True(t, ok)
	assert.Equal(t, KindCheckpoint, n.Kind)

	_, ok = relationsFixture().Find("nope")
	assert.False(t, ok)
}

func TestCheckpointsOf(t *testing.T) {
	nodes := relationsFixture()
	assert.Equal(t, []string{"ck2", "ck1"}, uuidsOf(nodes.CheckpointsOf("web")))
	assert.Equal(t, []string{"ck3"}, uuidsOf(nodes.CheckpointsOf("fork-a")))
	assert.Empty(t, nodes.CheckpointsOf("fork-c"))
	assert.Empty(t, nodes.CheckpointsOf("ck1"), "a checkpoint owns no checkpoints")
}

func TestForksOf(t *testing.T) {
	nodes := relationsFixture()
	assert.Equal(t, []string{"fork-b", "fork-a"}, uuidsOf(nodes.ForksOf("web")),
		"forks of the live instance and of its own checkpoints, not of descendants")
	assert.Equal(t, []string{"fork-a"}, uuidsOf(nodes.ForksOf("ck1")))
	assert.Empty(t, nodes.ForksOf("ck2"))
	assert.Equal(t, []string{"fork-c"}, uuidsOf(nodes.ForksOf("fork-a")))
}

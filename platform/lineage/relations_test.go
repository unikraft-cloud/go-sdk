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
		{Kind: KindTemplate, UUID: "tmpl", Source: "fork-b"},
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
	assert.Equal(t, []string{"tmpl"}, uuidsOf(nodes.ForksOf("fork-b")), "a template cloned from an instance")
}

func TestMark(t *testing.T) {
	assert.Equal(t, MarkInstance, Node{Kind: KindInstance}.Mark())
	assert.Equal(t, MarkFork, Node{Kind: KindInstance, Source: "ck"}.Mark())
	assert.Equal(t, MarkFork, Node{Kind: KindInstance, Source: "root", Live: true}.Mark())
	assert.Equal(t, MarkCheckpoint, Node{Kind: KindCheckpoint, Source: "root"}.Mark())
	assert.Equal(t, MarkTemplate, Node{Kind: KindTemplate, Source: "ck"}.Mark())
}

func TestRelationOf(t *testing.T) {
	nodes := Tree{
		{Kind: KindInstance, UUID: "fork", Name: "fork", Source: "ck"},
		{Kind: KindInstance, UUID: "live", Name: "live", Source: "root", Live: true},
		{Kind: KindCheckpoint, UUID: "ck", Name: "ck", Source: "root"},
		{Kind: KindInstance, UUID: "root", Name: "root"},
		{Kind: KindInstance, UUID: DeletedPrefix + "x", Deleted: true, Source: "ck"},
		{Kind: KindCheckpoint, UUID: "ck-x", Source: DeletedPrefix + "x"},
		{Kind: KindInstance, UUID: "nameless"},
		{Kind: KindCheckpoint, UUID: "ck-nameless", Source: "nameless"},
		{Kind: KindTemplate, UUID: "tmpl", Name: "tmpl", Source: "ck"},
		{Kind: KindInstance, UUID: "from-tmpl", Name: "from-tmpl", Source: "tmpl"},
	}
	root, _ := nodes.Find("root")
	ck, _ := nodes.Find("ck")
	tmpl, _ := nodes.Find("tmpl")
	deleted, _ := nodes.Find(DeletedPrefix + "x")
	nameless, _ := nodes.Find("nameless")

	for _, tt := range []struct {
		uuid   string
		rel    Relation
		parent Node
	}{
		{"root", RelationRoot, Node{}},
		{"fork", RelationForkedFrom, ck},
		{"live", RelationForkedLiveFrom, root},
		{"ck", RelationCheckpointOf, root},
		{DeletedPrefix + "x", RelationDeleted, ck},
		{"ck-x", RelationCheckpointOfDeleted, deleted},
		{"ck-nameless", RelationCheckpointOfDeleted, nameless},
		{"from-tmpl", RelationCreatedFrom, tmpl},
	} {
		n, _ := nodes.Find(tt.uuid)
		rel, parent := nodes.RelationOf(n)
		assert.Equal(t, tt.rel, rel, tt.uuid)
		assert.Equal(t, tt.parent, parent, tt.uuid)
	}

	rel, parent := nodes.RelationOf(Node{Kind: KindInstance, Source: "missing"})
	assert.Equal(t, RelationRoot, rel, "a fork of a node outside the tree")
	assert.Zero(t, parent)
	rel, _ = nodes.RelationOf(Node{Kind: KindCheckpoint, Source: "missing"})
	assert.Equal(t, RelationCheckpointOfDeleted, rel, "a checkpoint of an instance outside the tree")
}

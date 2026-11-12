// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"unikraft.com/cloud/sdk/platform"
)

var epoch = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

func at(minute int) time.Time { return epoch.Add(time.Duration(minute) * time.Minute) }

func inst(uuid string, minute int) platform.Instance {
	return platform.Instance{
		Uuid:      uuid,
		Name:      uuid,
		CreatedAt: at(minute),
		State:     platform.InstanceStateStopped,
	}
}

func cloneOf(i platform.Instance, source string) platform.Instance {
	i.CloneSource = &platform.InstanceCloneSource{Uuid: source}
	return i
}

func ckpt(uuid string, minute int) platform.CheckpointHistoryEntry {
	return platform.CheckpointHistoryEntry{Uuid: uuid, Name: uuid, CreatedAt: at(minute)}
}

// history lists entries newest first, as the platform does.
func history(
	uuid string, entries ...platform.CheckpointHistoryEntry,
) platform.GetCheckpointHistoryResponseInstanceHistory {
	return platform.GetCheckpointHistoryResponseInstanceHistory{
		Uuid: uuid, Name: uuid, History: entries,
	}
}

func specFixture() ([]platform.Instance, []platform.GetCheckpointHistoryResponseInstanceHistory) {
	c1, c2, c3 := ckpt("ckpt-1", 10), ckpt("ckpt-2", 20), ckpt("ckpt-3", 40)
	f1c1 := ckpt("fork-1-ckpt-1", 50)
	return []platform.Instance{
		cloneOf(inst("http-gcc-fork-1-debug", 60), "fork-1-ckpt-1"),
		inst("other", 35),
		inst("http-gcc", 0),
		cloneOf(inst("http-gcc-fork-1", 30), "ckpt-2"),
	}, []platform.GetCheckpointHistoryResponseInstanceHistory{
		history("http-gcc", c3, c2, c1),
		history("http-gcc-fork-1", f1c1, c2, c1),
		history("http-gcc-fork-1-debug", f1c1, c2, c1),
		history("other"),
	}
}

func liveFixture() ([]platform.Instance, []platform.GetCheckpointHistoryResponseInstanceHistory) {
	instances, histories := specFixture()
	staging := cloneOf(inst("http-gcc-staging", 15), "http-gcc")
	return append(instances, staging), append(histories, history("http-gcc-staging"))
}

type edge struct{ uuid, source string }

func edges(nodes []Node) []edge {
	out := make([]edge, len(nodes))
	for i, n := range nodes {
		out[i] = edge{n.UUID, n.Source}
	}
	return out
}

func TestBuildSpecFixture(t *testing.T) {
	nodes := Build(specFixture())

	assert.Equal(t, []edge{
		{"http-gcc-fork-1-debug", "fork-1-ckpt-1"},
		{"fork-1-ckpt-1", "http-gcc-fork-1"},
		{"ckpt-3", "http-gcc"},
		{"other", ""},
		{"http-gcc-fork-1", "ckpt-2"},
		{"ckpt-2", "http-gcc"},
		{"ckpt-1", "http-gcc"},
		{"http-gcc", ""},
	}, edges(nodes))

	for _, n := range nodes {
		assert.False(t, n.Deleted, n.UUID)
		switch n.UUID {
		case "http-gcc", "http-gcc-fork-1", "http-gcc-fork-1-debug", "other":
			assert.Equal(t, KindInstance, n.Kind, n.UUID)
			assert.Equal(t, platform.InstanceStateStopped, n.State, n.UUID)
		default:
			assert.Equal(t, KindCheckpoint, n.Kind, n.UUID)
			assert.Empty(t, n.State, n.UUID)
		}
		assert.Equal(t, n.UUID, n.Name)
	}
}

func TestBuildRootOnly(t *testing.T) {
	nodes := Build([]platform.Instance{inst("solo", 0)}, nil)

	require.Len(t, nodes, 1)
	assert.Equal(t, Node{
		Kind:      KindInstance,
		UUID:      "solo",
		Name:      "solo",
		CreatedAt: at(0),
		State:     platform.InstanceStateStopped,
	}, nodes[0])
}

func TestBuildIgnoresHistoryOfUnknownInstance(t *testing.T) {
	nodes := Build(
		[]platform.Instance{inst("solo", 0)},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("gone", ckpt("c", 10)),
		},
	)

	assert.Equal(t, []edge{{"solo", ""}}, edges(nodes))
}

func TestBuildOldestOwnerWins(t *testing.T) {
	// Both instances predate the checkpoint, so both claim it as their own.
	nodes := Build(
		[]platform.Instance{inst("b", 0), inst("a", 5)},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("a", ckpt("c", 10)),
			history("b", ckpt("c", 10)),
		},
	)

	assert.Equal(t, []edge{{"c", "b"}, {"a", ""}, {"b", ""}}, edges(nodes))
}

func TestBuildTombstone(t *testing.T) {
	nodes := Build(
		[]platform.Instance{inst("root", 0), cloneOf(inst("fork", 30), "d2")},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("root", ckpt("r1", 5)),
			history("fork", ckpt("d2", 20), ckpt("d1", 10), ckpt("r1", 5)),
		},
	)

	assert.Equal(t, []edge{
		{"fork", "d2"},
		{"d2", "deleted:d2"},
		{"d1", "deleted:d2"},
		{"deleted:d2", "r1"},
		{"r1", "root"},
		{"root", ""},
	}, edges(nodes))

	tomb := nodes[3]
	assert.True(t, tomb.Deleted)
	assert.Equal(t, KindInstance, tomb.Kind)
	assert.Empty(t, tomb.Name)
	assert.Empty(t, tomb.State)
	assert.Equal(t, at(10).Add(-time.Nanosecond), tomb.CreatedAt)
}

func TestBuildTombstoneRoot(t *testing.T) {
	nodes := Build(
		[]platform.Instance{cloneOf(inst("fork", 30), "c2")},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("fork", ckpt("own", 40), ckpt("c2", 20), ckpt("c1", 10)),
		},
	)

	assert.Equal(t, []edge{
		{"own", "fork"},
		{"fork", "c2"},
		{"c2", "deleted:c2"},
		{"c1", "deleted:c2"},
		{"deleted:c2", ""},
	}, edges(nodes))
}

func TestBuildSameSecond(t *testing.T) {
	// The platform has second resolution: root takes c in the second it
	// was created, and fork is cloned from c in that second too.
	nodes := Build(
		[]platform.Instance{inst("root", 0), cloneOf(inst("fork", 0), "c")},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("root", ckpt("c", 0)),
			history("fork", ckpt("f", 0), ckpt("c", 0)),
		},
	)

	assert.Equal(t, []edge{
		{"c", "root"},
		{"f", "fork"},
		{"fork", "c"},
		{"root", ""},
	}, edges(nodes))
}

func uuids(nodes []Node) []string {
	out := make([]string, len(nodes))
	for i, n := range nodes {
		out[i] = n.UUID
	}
	return out
}

func TestScopeExcludesUnrelatedTree(t *testing.T) {
	nodes := Build(specFixture())

	for _, current := range []string{"http-gcc-fork-1-debug", "ckpt-1", "http-gcc"} {
		tree, truncated := nodes.Scope(current, 0)
		assert.False(t, truncated, current)
		assert.Equal(t, []string{
			"http-gcc-fork-1-debug",
			"fork-1-ckpt-1",
			"ckpt-3",
			"http-gcc-fork-1",
			"ckpt-2",
			"ckpt-1",
			"http-gcc",
		}, uuids(tree), current)
	}

	tree, truncated := nodes.Scope("other", 0)
	assert.False(t, truncated)
	assert.Equal(t, []string{"other"}, uuids(tree))
}

func TestScopeTruncation(t *testing.T) {
	nodes := Build(specFixture())

	// current and its nearest ancestors, when its chain is past the limit.
	tree, truncated := nodes.Scope("http-gcc-fork-1-debug", 3)
	assert.True(t, truncated)
	assert.Equal(t, []string{
		"http-gcc-fork-1-debug", "fork-1-ckpt-1", "http-gcc-fork-1",
	}, uuids(tree))

	// current is kept even when it is among the oldest, and a newer node only
	// with the ancestors it lacks.
	tree, truncated = nodes.Scope("http-gcc", 3)
	assert.True(t, truncated)
	assert.Equal(t, []string{"ckpt-3", "ckpt-2", "http-gcc"}, uuids(tree))

	// The newest nodes alone would leave ckpt-3 and http-gcc-fork-1 apart.
	tree, truncated = nodes.Scope("http-gcc-fork-1", 4)
	assert.True(t, truncated)
	assert.Equal(t, []string{
		"fork-1-ckpt-1", "http-gcc-fork-1", "ckpt-2", "http-gcc",
	}, uuids(tree))

	tree, truncated = nodes.Scope("http-gcc", 7)
	assert.False(t, truncated)
	assert.Len(t, tree, 7)
}

func TestScopeMissingCurrent(t *testing.T) {
	tree, truncated := Build(specFixture()).Scope("nope", 10)
	assert.Nil(t, tree)
	assert.False(t, truncated)
}

func TestScopeCycle(t *testing.T) {
	nodes := Tree{
		{Kind: KindCheckpoint, UUID: "c", Source: "i", CreatedAt: at(10)},
		{Kind: KindInstance, UUID: "i", Source: "c", CreatedAt: at(0)},
	}

	tree, truncated := nodes.Scope("i", 0)
	assert.False(t, truncated)
	assert.Equal(t, []string{"c", "i"}, uuids(tree))
}

func TestBuildLiveFork(t *testing.T) {
	nodes := Build(liveFixture())

	n, ok := nodes.Find("http-gcc-staging")
	require.True(t, ok)
	assert.Equal(t, "http-gcc", n.Source)
	assert.True(t, n.Live)

	for _, n := range nodes {
		if n.UUID != "http-gcc-staging" {
			assert.False(t, n.Live, n.UUID)
		}
	}
	assert.Equal(t, []string{"http-gcc-fork-1", "http-gcc-staging"}, uuids(nodes.ForksOf("http-gcc")))
}

func TestBuildRootsWithoutAKnownCloneSource(t *testing.T) {
	// The platform drops the clone source once the source is gone; an
	// inherited history does not stand in for it.
	gone := inst("gone-source", 30)
	template := cloneOf(inst("from-template", 40), "some-template")
	self := cloneOf(inst("self", 50), "self")
	nodes := Build(
		[]platform.Instance{inst("root", 0), gone, template, self},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("root", ckpt("c", 10)),
			history("gone-source", ckpt("c", 10)),
			history("from-template"),
		},
	)

	assert.Equal(t, []edge{
		{"self", ""},
		{"from-template", ""},
		{"gone-source", ""},
		{"c", "root"},
		{"root", ""},
	}, edges(nodes))
	for _, n := range nodes {
		assert.False(t, n.Live, n.UUID)
	}
}

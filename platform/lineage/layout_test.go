// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"unikraft.com/cloud/sdk/platform"
)

// assertLayout checks l against the testdata golden file name, and that it
// keeps the promises of Line and Branch with no crossings.
func assertLayout(t *testing.T, l Layout, name string) {
	t.Helper()
	golden.Assert(t, l.Text(func(n Node) string { return n.UUID }), name)
	assert.Empty(t, layoutFaults(l))
	assert.Zero(t, layoutCrossings(l))
}

// layoutFaults lists where l breaks the promises of Line and Branch, other
// than crossings.
func layoutFaults(l Layout) []string {
	var faults []string
	fault := func(format string, args ...any) {
		faults = append(faults, fmt.Sprintf(format, args...))
	}

	// held reports which color holds a lane in a row with a dot, a line or
	// the vertical part of a branch.
	held := map[[2]int]map[int]bool{}
	hold := func(row, lane, color int) {
		k := [2]int{row, lane}
		if held[k] == nil {
			held[k] = map[int]bool{}
		}
		held[k][color] = true
	}
	for i, n := range l.Rows {
		if l.RowOf[n.UUID] != i {
			fault("%s: RowOf %d, row %d", n.UUID, l.RowOf[n.UUID], i)
		}
		if lane := l.Lane[n.UUID]; lane < 0 || lane >= l.Columns {
			fault("%s: lane %d out of %d columns", n.UUID, lane, l.Columns)
		}
		hold(i, l.Lane[n.UUID], l.Color[n.UUID])
	}
	for _, line := range l.Lines {
		for r := line.Top; r <= line.Bottom; r++ {
			hold(r, line.Lane, line.Color)
		}
	}
	for _, b := range l.Branches {
		if b.Live {
			continue
		}
		top := b.ChildRow
		if b.ChildLane == b.ParentLane {
			top++
		}
		for r := top; r < b.ParentRow; r++ {
			hold(r, b.ChildLane, b.Color)
		}
		if b.ChildLane != b.ParentLane {
			hold(b.ParentRow, b.ChildLane, b.Color)
		}
	}
	for k, colors := range held {
		if len(colors) > 1 {
			fault("row %d, lane %d held by colors %v", k[0], k[1], colors)
		}
	}

	// from is the parent row of the branch that runs along a row from a lane
	// to the next. Only branches from one parent share those cells.
	from := map[[2]int]int{}
	for _, b := range l.Branches {
		row := b.ParentRow
		if b.Live {
			row = b.ChildRow
		}
		for c := min(b.ParentLane, b.ChildLane); c < max(b.ParentLane, b.ChildLane); c++ {
			k := [2]int{row, c}
			if p, ok := from[k]; ok && p != b.ParentRow {
				fault("row %d, lanes %d-%d: branches from rows %d and %d overlap", row, c, c+1, p, b.ParentRow)
			}
			from[k] = b.ParentRow
		}
	}

	for _, b := range l.Branches {
		parent := l.Rows[b.ParentRow]
		if !b.Live {
			continue
		}
		on := false
		for _, line := range l.Lines {
			if line.Lane == b.ParentLane && line.Color == l.Color[parent.UUID] &&
				line.Top <= b.ChildRow && b.ChildRow <= line.Bottom {
				on = true
			}
		}
		if !on {
			fault("live branch to row %d leaves no line of %s", b.ChildRow, parent.UUID)
		}
	}
	return faults
}

// layoutCrossings counts the lanes that branches cross along a row while a
// line or the vertical part of another branch runs through them.
func layoutCrossings(l Layout) int {
	crossings := 0
	for _, b := range l.Branches {
		row := b.ParentRow
		if b.Live {
			row = b.ChildRow
		}
		a, z := min(b.ParentLane, b.ChildLane), max(b.ParentLane, b.ChildLane)
		for c := a + 1; c < z; c++ {
			for _, line := range l.Lines {
				if line.Lane == c && line.Top <= row && row <= line.Bottom {
					crossings++
				}
			}
			for _, o := range l.Branches {
				if !o.Live && o.ChildLane == c && o.ChildRow <= row && row < o.ParentRow {
					crossings++
				}
			}
		}
	}
	return crossings
}

func specTree(t *testing.T) Tree {
	t.Helper()
	tree, _ := Build(specFixture()).Scope("http-gcc", 0)
	require.NotNil(t, tree)
	return tree
}

func TestLayoutSpec(t *testing.T) {
	l := specTree(t).Layout()
	assertLayout(t, l, "spec.golden")

	// Color follows the line, not the lane.
	assert.Equal(t, l.Lane["http-gcc-fork-1-debug"], l.Lane["ckpt-3"])
	assert.NotEqual(t, l.Color["http-gcc-fork-1-debug"], l.Color["ckpt-3"])
	assert.Equal(t, l.Color["http-gcc-fork-1"], l.Color["fork-1-ckpt-1"])
	assert.Equal(t, 2, l.LaneCount)
	assert.Equal(t, 2, l.Columns)
	assert.False(t, l.Overflows())
}

func TestLayoutSpecLiveFork(t *testing.T) {
	tree, _ := Build(liveFixture()).Scope("http-gcc", 0)
	assertLayout(t, tree.Layout(), "spec-live.golden")
}

func TestLayoutTemplate(t *testing.T) {
	assertLayout(t, Build(templateFixture()).Layout(), "template.golden")
}

func TestLayoutLiveTemplate(t *testing.T) {
	tmpl := cloneOf(inst("T", 20), "I")
	tmpl.State = platform.InstanceStateTemplate
	l := Build([]platform.Instance{
		inst("I", 0),
		cloneOf(inst("X", 10), "I"),
		tmpl,
		cloneOf(inst("C1", 30), "T"),
		cloneOf(inst("C2", 40), "T"),
		cloneOf(inst("Y", 50), "I"),
	}, nil).Layout()
	assertLayout(t, l, "live-template.golden")
}

func TestLayoutSingleRoot(t *testing.T) {
	l := Build([]platform.Instance{inst("solo", 0)}, nil).Layout()
	assertLayout(t, l, "single-root.golden")
	assert.Empty(t, l.Lines)
	assert.Empty(t, l.Branches)
	assert.Equal(t, 1, l.Columns)
}

func TestLayoutEmpty(t *testing.T) {
	l := Tree{}.Layout()
	assert.Empty(t, l.Grid())
	assert.Zero(t, l.Columns)
}

func TestLayoutLaneReuseNeedsGap(t *testing.T) {
	// c sits between a1 and b, so it cannot share a lane with either.
	l := Build(
		[]platform.Instance{inst("a", 0), inst("c", 15), inst("b", 20)},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("a", ckpt("a1", 10)),
			history("b", ckpt("b1", 30)),
		},
	).Layout()
	assertLayout(t, l, "lane-reuse.golden")
}

func TestLayoutLiftsChildAboveParent(t *testing.T) {
	// Clock skew: a checkpoint timestamped before its owner.
	l := Tree{
		{Kind: KindInstance, UUID: "x", CreatedAt: at(10)},
		{Kind: KindInstance, UUID: "y", CreatedAt: at(7)},
		{Kind: KindCheckpoint, UUID: "k", Source: "x", CreatedAt: at(5)},
	}.Layout()

	assert.Equal(t, []string{"k", "x", "y"}, uuids(l.Rows))
}

func TestLayoutCycleTerminates(t *testing.T) {
	l := Tree{
		{Kind: KindCheckpoint, UUID: "c", Source: "i", CreatedAt: at(0)},
		{Kind: KindInstance, UUID: "i", Source: "c", CreatedAt: at(10)},
	}.Layout()

	assert.Len(t, l.Rows, 2)
	assert.Len(t, l.Grid(), 2)
}

func TestLayoutSiblingForks(t *testing.T) {
	// Lane 0 continues root's line; any other shared lane would cross a sibling's dot.
	l := Tree{
		{Kind: KindInstance, UUID: "root", CreatedAt: at(0)},
		{Kind: KindCheckpoint, UUID: "ckpt", Source: "root", CreatedAt: at(10)},
		{Kind: KindInstance, UUID: "f1", Source: "ckpt", CreatedAt: at(20)},
		{Kind: KindInstance, UUID: "f2", Source: "ckpt", CreatedAt: at(30)},
		{Kind: KindInstance, UUID: "f3", Source: "ckpt", CreatedAt: at(40)},
		{Kind: KindInstance, UUID: "f4", Source: "ckpt", CreatedAt: at(50)},
	}.Layout()
	assertLayout(t, l, "sibling-forks.golden")
}

func TestLayoutBranchSkipsNewerCheckpoint(t *testing.T) {
	// g's branch from c1 would run through c2 in lane 0, and through f in lane 1.
	l := Tree{
		{Kind: KindInstance, UUID: "root", CreatedAt: at(0)},
		{Kind: KindCheckpoint, UUID: "c1", Source: "root", CreatedAt: at(10)},
		{Kind: KindCheckpoint, UUID: "c2", Source: "root", CreatedAt: at(20)},
		{Kind: KindInstance, UUID: "f", Source: "c2", CreatedAt: at(25)},
		{Kind: KindInstance, UUID: "g", Source: "c1", CreatedAt: at(30)},
	}.Layout()
	assertLayout(t, l, "skips-newer-checkpoint.golden")
}

func TestLayoutLiveForkExtendsParentLine(t *testing.T) {
	l := Tree{
		{Kind: KindInstance, UUID: "root", CreatedAt: at(0)},
		{Kind: KindInstance, UUID: "f", Source: "root", Live: true, CreatedAt: at(10)},
	}.Layout()
	assertLayout(t, l, "live-fork.golden")
	assert.Equal(t, []Line{{Lane: 0, Color: 0, Top: 0, Bottom: 1}}, l.Lines)
	assert.Equal(t, []Branch{{ParentRow: 1, ParentLane: 0, ChildRow: 0, ChildLane: 1, Color: 1, Live: true}}, l.Branches)
}

func TestLayoutAvoidsCrossing(t *testing.T) {
	// B needs a new lane; to the right of A's, its branch from r2 would
	// cross A's line.
	l := Tree{
		{Kind: KindInstance, UUID: "R", CreatedAt: at(0)},
		{Kind: KindCheckpoint, UUID: "r1", Source: "R", CreatedAt: at(10)},
		{Kind: KindInstance, UUID: "A", Source: "r1", CreatedAt: at(20)},
		{Kind: KindCheckpoint, UUID: "r2", Source: "R", CreatedAt: at(30)},
		{Kind: KindCheckpoint, UUID: "r3", Source: "R", CreatedAt: at(35)},
		{Kind: KindCheckpoint, UUID: "a1", Source: "A", CreatedAt: at(40)},
		{Kind: KindInstance, UUID: "B", Source: "r2", CreatedAt: at(50)},
	}.Layout()
	assertLayout(t, l, "avoids-crossing.golden")
}

func TestLayoutDeletedInstance(t *testing.T) {
	l := Build(
		[]platform.Instance{inst("root", 0), cloneOf(inst("fork", 30), "d2")},
		[]platform.GetCheckpointHistoryResponseInstanceHistory{
			history("root", ckpt("r1", 5)),
			history("fork", ckpt("d2", 20), ckpt("d1", 10), ckpt("r1", 5)),
		},
	).Layout()
	assertLayout(t, l, "deleted-instance.golden")
}

// wideTree has roots that all overlap, so that each needs a lane.
func wideTree(roots int) Tree {
	var t Tree
	for i := range roots {
		root := fmt.Sprintf("root-%d", i)
		t = append(t,
			Node{Kind: KindInstance, UUID: root, CreatedAt: at(i)},
			Node{Kind: KindCheckpoint, UUID: "ck-" + root, Source: root, CreatedAt: at(roots + i)},
		)
	}
	return t
}

func TestLayoutMaxLanes(t *testing.T) {
	l := wideTree(3).Layout(WithMaxLanes(3))
	assert.False(t, l.Overflows())
	assert.Equal(t, 3, l.Columns)
	assert.Len(t, l.Lines, 3)

	l = wideTree(5).Layout(WithMaxLanes(3))
	assertLayout(t, l, "max-lanes.golden")
	assert.True(t, l.Overflows())
	assert.Equal(t, 5, l.LaneCount)
	assert.Equal(t, 4, l.Columns, "one column holds the overflow")
	assert.Len(t, l.Lines, 3, "the lines past the cap are not drawn")
	assert.Equal(t, 3, l.Lane["root-4"])

	l = wideTree(5).Layout(WithMaxLanes(0))
	assert.False(t, l.Overflows())
	assert.Equal(t, 5, l.Columns)
}

func TestLayoutMaxLanesKeepsRoot(t *testing.T) {
	// The forks of ckpt-2 cross no line when their lanes go left of web's,
	// so they push web right, but never past the cap.
	tree := Tree{
		{Kind: KindInstance, UUID: "web", CreatedAt: at(0)},
		{Kind: KindCheckpoint, UUID: "ckpt-1", Source: "web", CreatedAt: at(10)},
		{Kind: KindCheckpoint, UUID: "ckpt-2", Source: "web", CreatedAt: at(20)},
		{Kind: KindCheckpoint, UUID: "ckpt-3", Source: "web", CreatedAt: at(30)},
		{Kind: KindInstance, UUID: "from-3", Source: "ckpt-3", CreatedAt: at(40)},
		{Kind: KindInstance, UUID: "from-1", Source: "ckpt-1", CreatedAt: at(50)},
	}
	for i := range 6 {
		tree = append(tree, Node{Kind: KindInstance, UUID: fmt.Sprintf("from-2-%d", i+1), Source: "ckpt-2", CreatedAt: at(60 + i)})
	}

	l := tree.Layout(WithMaxLanes(6))
	assertLayout(t, l, "max-lanes-branches.golden")
	assert.True(t, l.Overflows())
	assert.Less(t, l.Lane["web"], 6, "web is drawn")
	assert.Len(t, l.Lines, 1, "web's line is drawn")
	assert.Len(t, l.Branches, 5, "the branches into drawn lanes are drawn")
}

// randomTree grows a tree as the platform would: checkpoints of instances,
// and instances or templates cloned from checkpoints, from templates or live
// from instances.
func randomTree(r *rand.Rand, size int) Tree {
	var t Tree
	var instances, checkpoints, templates []string
	for i := range size {
		n := Node{Kind: KindInstance, UUID: fmt.Sprintf("n%d", i), CreatedAt: at(i)}
		switch x := r.IntN(11); {
		case len(instances) == 0 || x == 0:
		case x < 5:
			n.Kind, n.Source = KindCheckpoint, instances[r.IntN(len(instances))]
		case x < 8 && len(checkpoints) > 0:
			n.Source = checkpoints[r.IntN(len(checkpoints))]
		case x < 9 && len(templates) > 0:
			n.Source = templates[r.IntN(len(templates))]
		default:
			n.Source, n.Live = instances[r.IntN(len(instances))], true
		}
		switch {
		case n.Kind == KindCheckpoint:
			checkpoints = append(checkpoints, n.UUID)
		case r.IntN(10) == 0:
			n.Kind = KindTemplate
			templates = append(templates, n.UUID)
		default:
			instances = append(instances, n.UUID)
		}
		t = append(t, n)
	}
	return t
}

func TestLayoutRandomTrees(t *testing.T) {
	// Some trees have no layout free of crossings that the greedy placement
	// finds; they stay rare.
	const trees = 2000
	r := rand.New(rand.NewPCG(1, 2))
	crossed := 0
	for i := range trees {
		tree := randomTree(r, 2+r.IntN(40))
		l := tree.Layout()
		if !assert.Empty(t, layoutFaults(l), "tree %d:\n%s", i, l.Text(func(n Node) string { return n.UUID })) {
			return
		}
		capped := tree.Layout(WithMaxLanes(3))
		if !assert.Empty(t, layoutFaults(capped), "tree %d:\n%s", i, capped.Text(func(n Node) string { return n.UUID })) {
			return
		}
		if layoutCrossings(l) > 0 {
			crossed++
		}
	}
	assert.Less(t, crossed, trees/15)
}

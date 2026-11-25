// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridBridgesCrossing(t *testing.T) {
	// A layout the placement avoids, built by hand: c's branch from k has to
	// cross x's line.
	rows := Tree{
		{Kind: KindInstance, UUID: "c", Source: "k"},
		{Kind: KindCheckpoint, UUID: "x1", Source: "x"},
		{Kind: KindCheckpoint, UUID: "k", Source: "r"},
		{Kind: KindInstance, UUID: "x"},
		{Kind: KindInstance, UUID: "r"},
	}
	l := Layout{
		Rows:  rows,
		Lane:  map[string]int{"c": 2, "x1": 1, "k": 0, "x": 1, "r": 0},
		Color: map[string]int{"c": 2, "x1": 1, "k": 0, "x": 1, "r": 0},
		Lines: []Line{
			{Lane: 1, Color: 1, Top: 1, Bottom: 3},
			{Lane: 0, Color: 0, Top: 2, Bottom: 4},
		},
		Branches: []Branch{{ParentRow: 2, ParentLane: 0, ChildRow: 0, ChildLane: 2, Color: 2}},
		Columns:  3,
	}

	assert.Equal(t, ""+
		"    ●    c\n"+
		"  ○ │    x1\n"+
		"○─│─╯    k\n"+
		"│ ●      x\n"+
		"●        r\n",
		l.Text(func(n Node) string { return n.UUID }))
	assert.Equal(t, []Cell{
		{'○', 0}, {'─', 2}, {'│', 1}, {'─', 2}, {'╯', 2},
	}, l.Grid()[2], "the line keeps its color where the branch runs under it")
}

func TestGridNewestColorWins(t *testing.T) {
	// b and a branch from k along the same row, and b is the newer node.
	rows := Tree{
		{Kind: KindInstance, UUID: "b", Source: "k"},
		{Kind: KindInstance, UUID: "a", Source: "k"},
		{Kind: KindCheckpoint, UUID: "k", Source: "r"},
		{Kind: KindInstance, UUID: "r"},
	}
	l := Layout{
		Rows:  rows,
		Lane:  map[string]int{"b": 2, "a": 1, "k": 0, "r": 0},
		Color: map[string]int{"b": 2, "a": 1, "k": 0, "r": 0},
		Lines: []Line{{Lane: 0, Color: 0, Top: 2, Bottom: 3}},
		Branches: []Branch{
			{ParentRow: 2, ParentLane: 0, ChildRow: 1, ChildLane: 1, Color: 1},
			{ParentRow: 2, ParentLane: 0, ChildRow: 0, ChildLane: 2, Color: 2},
		},
		Columns: 3,
	}

	assert.Equal(t, []Cell{
		{'○', 0}, {'─', 2}, {'┴', 2}, {'─', 2}, {'╯', 2},
	}, l.Grid()[2])
}

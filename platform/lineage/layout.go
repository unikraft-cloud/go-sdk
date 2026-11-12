// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import "slices"

// Span is the row of an instance together with the rows of its own
// checkpoints, drawn as one vertical line.
type Span struct {
	Owner string
	Lane  int
	// Color is the index of the span in Layout.Spans. Unlike Lane, it is
	// unique.
	Color       int
	Top, Bottom int
	// End is the row where the line stops at the top: Top, or higher up at the
	// row of the topmost live fork of Owner, whose branch leaves the line just
	// below that row.
	End int
}

// Layout is the position of every node of a lineage diagram.
type Layout struct {
	// Rows are the nodes from top to bottom.
	Rows Tree
	// Lane, Color and RowOf are keyed by Node.UUID.
	Lane  map[string]int
	Color map[string]int
	RowOf map[string]int
	// Spans are the vertical lines, one per instance with its checkpoints.
	Spans     []Span
	LaneCount int
}

// Layout places the nodes of t in rows and lanes, newest first, with every
// child above its parent.
func (t Tree) Layout() Layout {
	rows := slices.Clone(t)

	// Bounded so that inconsistent input cannot loop.
	rows.sortNewestFirst()
	for range len(rows) {
		if !rows.liftChildren() {
			break
		}
	}

	l := Layout{
		Rows:  rows,
		Lane:  make(map[string]int, len(rows)),
		Color: make(map[string]int, len(rows)),
		RowOf: make(map[string]int, len(rows)),
	}
	for i, n := range rows {
		l.RowOf[n.UUID] = i
	}

	spanOf := map[string]int{}
	for i, n := range rows {
		o := n.owner()
		if k, ok := spanOf[o]; ok {
			l.Spans[k].Bottom = i
			continue
		}
		spanOf[o] = len(l.Spans)
		l.Spans = append(l.Spans, Span{Owner: o, Top: i, Bottom: i, End: i})
	}

	for i, n := range rows {
		if !n.Live {
			continue
		}
		if _, ok := l.RowOf[n.Source]; !ok {
			continue
		}
		s := &l.Spans[spanOf[n.Source]]
		s.End = min(s.End, i)
	}

	// A span claims rows down to its parent too, as its branch must cross no dot.
	slices.SortFunc(l.Spans, func(a, b Span) int { return b.Bottom - a.Bottom })
	byLane := [][]Span{}
	for k := range l.Spans {
		s := &l.Spans[k]
		parent, end := -1, s.Bottom
		if o, ok := l.RowOf[s.Owner]; ok && !rows[o].Live {
			if p, ok := l.RowOf[rows[o].Source]; ok && p > s.Bottom {
				parent, end = p, p-1
			}
		}
		top := s.End
		free := func(placed []Span) bool {
			for _, q := range placed {
				if q.Top == parent && q.End == q.Top {
					// The branch continues q's line straight up, as in git log.
					if q.Top < s.Bottom+2 {
						return false
					}
					continue
				}
				if q.End <= end+1 && q.Bottom >= top-1 {
					return false
				}
			}
			return true
		}
		lane := slices.IndexFunc(byLane, free)
		if lane < 0 {
			lane = len(byLane)
			byLane = append(byLane, nil)
		}
		s.Lane = lane
		s.Color = k
		byLane[lane] = append(byLane[lane], *s)
		spanOf[s.Owner] = k
	}
	l.LaneCount = len(byLane)

	for _, n := range rows {
		s := l.Spans[spanOf[n.owner()]]
		l.Lane[n.UUID] = s.Lane
		l.Color[n.UUID] = s.Color
	}
	return l
}

// owner returns the UUID of the instance whose span n is drawn on.
func (n Node) owner() string {
	if n.Kind == KindCheckpoint && n.Source != "" {
		return n.Source
	}
	return n.UUID
}

// liftChildren moves every node that sorts below its parent to directly above
// it, and reports whether it moved any.
func (t Tree) liftChildren() bool {
	moved := false
	for i := range t {
		src := t[i].Source
		if src == "" {
			continue
		}
		p := slices.IndexFunc(t[:i], func(n Node) bool { return n.UUID == src })
		if p < 0 {
			continue
		}
		child := t[i]
		copy(t[p+1:i+1], t[p:i])
		t[p] = child
		moved = true
	}
	return moved
}

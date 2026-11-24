// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import "slices"

// Layout is a lineage diagram: the row and lane of every node, and the lines
// and branches between them. Renderers only map rows and lanes to their
// medium; Grid draws it as text.
type Layout struct {
	// Rows are the nodes from top to bottom.
	Rows Tree
	// Lane, Color and RowOf are keyed by Node.UUID. Nodes in the lanes past
	// MaxLanes share lane MaxLanes.
	Lane  map[string]int
	Color map[string]int
	RowOf map[string]int
	// Lines are the vertical lines, one per instance with its checkpoints.
	Lines []Line
	// Branches join instances to their parents on other lines.
	Branches []Branch
	// Columns is the number of lanes drawn: LaneCount, or MaxLanes+1 when the
	// layout overflows.
	Columns int
	// LaneCount is the number of lanes the layout needs.
	LaneCount int
	// MaxLanes is the cap set with WithMaxLanes, or 0.
	MaxLanes int
}

// Line is the vertical line of an instance and its own checkpoints. It runs
// through every row from Top to Bottom in Lane, and no other line or branch
// shares those rows of the lane.
type Line struct {
	Lane int
	// Color is the color of the nodes on the line, as in Layout.Color.
	Color int
	// Top is the row of the newest node on the line or, higher up, the row of
	// the newest instance cloned live from it, whose branch leaves the line
	// there. Bottom is the row of the oldest node on the line.
	Top, Bottom int
}

// Branch joins an instance in ChildRow to its parent in ParentRow, which is
// on another line further down.
//
// A branch from a checkpoint or a template turns in ParentRow: it runs from
// the parent along the row to ChildLane, then up ChildLane to the child, which
// holds no other line or branch on the way. When ChildLane is ParentLane, the
// branch continues the parent's line straight up.
//
// A live branch turns in ChildRow instead, where the parent's line runs
// through ParentLane: it leaves the line along the row to the child.
//
// Where a branch runs along a row, it shares no cells with the branch of
// another parent. It crosses no other line in between, unless the layout had
// no lane free of crossings to give it.
type Branch struct {
	ParentRow, ParentLane int
	ChildRow, ChildLane   int
	// Color is the color of the child.
	Color int
	// Live is set when the child was cloned from the running parent rather
	// than from one of its checkpoints.
	Live bool
}

// Overflows reports whether the layout needs more lanes than MaxLanes.
func (l Layout) Overflows() bool {
	return l.MaxLanes > 0 && l.LaneCount > l.MaxLanes
}

// LayoutOption configures Tree.Layout.
type LayoutOption func(*layoutOptions)

type layoutOptions struct {
	maxLanes int
}

// WithMaxLanes caps the lanes drawn at n, so that a wide tree stays
// readable. The nodes of the lanes past it share one more lane, with no
// lines or branches. Lines are placed from the bottom up, and a line placed
// in a drawn lane stays in one, so the root's line is always drawn. A cap of
// 0 or less draws every lane.
func WithMaxLanes(n int) LayoutOption {
	return func(o *layoutOptions) {
		o.maxLanes = max(n, 0)
	}
}

// span is a line during layout: an instance together with the rows of its
// own checkpoints, and the branch to its parent.
type span struct {
	owner string
	// color is the index of the span once sorted. Unlike the lane, it is
	// unique.
	color       int
	top, bottom int
	// end is the row where the line stops at the top: top, or higher up at
	// the row of the topmost live fork of owner.
	end int
	// row is the row of owner, or -1 if the tree leaves it out.
	row int
	// parent is the row the branch of owner leaves from, or -1 if it has
	// none. fromOwner is the owner of the parent's span, and from its index.
	parent, from int
	fromOwner    string
	live         bool
	track        int
}

// claimBottom is the lowest row the span holds in its lane: down to just
// above its parent for a branch from a checkpoint, which turns in the
// parent's row.
func (s span) claimBottom() int {
	if s.parent >= 0 && !s.live {
		return s.parent - 1
	}
	return s.bottom
}

// turn is the row where the branch of s runs along to the lane of s.
func (s span) turn() int {
	if s.live {
		return s.row
	}
	return s.parent
}

// Layout places the nodes of t in rows and lanes, newest first, with every
// child above its parent.
func (t Tree) Layout(opts ...LayoutOption) Layout {
	var cfg layoutOptions
	for _, opt := range opts {
		opt(&cfg)
	}

	rows := slices.Clone(t)

	// Bounded so that inconsistent input cannot loop.
	rows.sortNewestFirst()
	for range len(rows) {
		if !rows.liftChildren() {
			break
		}
	}

	l := Layout{
		Rows:     rows,
		Lane:     make(map[string]int, len(rows)),
		Color:    make(map[string]int, len(rows)),
		RowOf:    make(map[string]int, len(rows)),
		MaxLanes: cfg.maxLanes,
	}
	for i, n := range rows {
		l.RowOf[n.UUID] = i
	}

	spanOf := map[string]int{}
	var spans []span
	for i, n := range rows {
		o := n.owner()
		if k, ok := spanOf[o]; ok {
			spans[k].bottom = i
			continue
		}
		spanOf[o] = len(spans)
		spans = append(spans, span{owner: o, top: i, bottom: i, end: i, row: -1, parent: -1})
	}

	for k := range spans {
		s := &spans[k]
		r, ok := l.RowOf[s.owner]
		if !ok {
			continue
		}
		s.row = r
		n := rows[r]
		p, ok := l.RowOf[n.Source]
		// Both kinds of branch need the parent below the whole span, so that
		// the parent's span is placed first.
		if !ok || n.Source == "" || p <= s.bottom {
			continue
		}
		s.parent, s.live, s.fromOwner = p, n.Live, rows[p].owner()
	}
	for _, s := range spans {
		if s.live {
			p := &spans[spanOf[s.fromOwner]]
			p.end = min(p.end, s.row)
		}
	}

	// Spans are placed bottom up, so that a parent's lane is known before its
	// children's.
	slices.SortFunc(spans, func(a, b span) int { return b.bottom - a.bottom })
	for k := range spans {
		spans[k].color = k
		spanOf[spans[k].owner] = k
	}
	for k := range spans {
		if spans[k].parent >= 0 {
			spans[k].from = spanOf[spans[k].fromOwner]
		}
	}

	pl := placer{spans: spans, maxLanes: cfg.maxLanes}
	for k := range spans {
		pl.place(k)
	}

	l.LaneCount = len(pl.order)
	l.Columns = l.LaneCount
	if cfg.maxLanes > 0 {
		l.Columns = min(l.LaneCount, cfg.maxLanes+1)
	}
	lane := func(s span) int {
		c := pl.col[s.track]
		if cfg.maxLanes > 0 {
			return min(c, cfg.maxLanes)
		}
		return c
	}
	drawn := func(s span) bool {
		return cfg.maxLanes == 0 || pl.col[s.track] < cfg.maxLanes
	}

	for _, n := range rows {
		s := spans[spanOf[n.owner()]]
		l.Lane[n.UUID] = lane(s)
		l.Color[n.UUID] = s.color
	}
	for _, s := range spans {
		if !drawn(s) {
			continue
		}
		if s.end < s.bottom {
			l.Lines = append(l.Lines, Line{
				Lane:   lane(s),
				Color:  s.color,
				Top:    s.end,
				Bottom: s.bottom,
			})
		}
		if s.parent >= 0 && drawn(spans[s.from]) {
			l.Branches = append(l.Branches, Branch{
				ParentRow:  s.parent,
				ParentLane: lane(spans[s.from]),
				ChildRow:   s.row,
				ChildLane:  lane(s),
				Color:      s.color,
				Live:       s.live,
			})
		}
	}
	return l
}

// placer assigns spans to tracks, which share a lane, and orders the tracks
// into columns, so that branches cross as few lines as it can.
type placer struct {
	spans []span
	// maxLanes is the cap set with WithMaxLanes, or 0.
	maxLanes int
	// tracks are the spans placed in each track.
	tracks [][]int
	// order is the tracks from left to right, and col the column of each.
	order, col []int
	// Scratch space for place.
	passing, acrossCol, acrossPos []int
}

// zeroed returns buf with n zeros, reusing its space.
func zeroed(buf []int, n int) []int {
	buf = slices.Grow(buf[:0], n)[:n]
	clear(buf)
	return buf
}

// free reports whether s fits in track without touching its spans, other
// than continuing the line of its parent straight up.
func (p *placer) free(track int, s span) bool {
	top, end := s.end, s.claimBottom()
	for _, k := range p.tracks[track] {
		q := p.spans[k]
		if s.parent >= 0 && !s.live && q.top == s.parent && q.end == q.top {
			// The branch continues q's line straight up, as in git log.
			if q.top < s.bottom+2 {
				return false
			}
			continue
		}
		if q.end <= end+1 && q.bottom >= top-1 {
			return false
		}
	}
	return true
}

// place puts the span k, with every span before it placed, in the track
// where its branch crosses the fewest lines, and the nearest to its parent
// among those. It opens a new track only when that crosses fewer lines: to
// the right of the others, or else next to its parent's. Under a cap, a new
// track never pushes a drawn one past it.
func (p *placer) place(k int) {
	s := p.spans[k]
	placed := p.spans[:k]
	n := len(p.order)

	// fromCol is the column that the live branch of the parent runs from
	// along the row where the branch of s turns, or -1.
	turn, parentCol, fromCol := -1, -1, -1
	if s.parent >= 0 {
		q := placed[s.from]
		turn, parentCol = s.turn(), p.col[q.track]
		if q.parent >= 0 && q.turn() == turn {
			fromCol = p.col[placed[q.from].track]
		}
	}
	// shares reports whether a branch to a drawn column, right of the
	// parent's or left of it, runs in the cells of the parent's live branch.
	shares := func(right bool, col int) bool {
		return fromCol >= 0 && right == (fromCol > parentCol) && (p.maxLanes == 0 || col < p.maxLanes)
	}

	// passing counts the columns before each whose lines run through the row
	// where the branch of s turns.
	p.passing = zeroed(p.passing, n+1)
	passing := p.passing
	if turn >= 0 {
		for _, q := range placed {
			if q.end <= turn && turn <= q.claimBottom() {
				passing[p.col[q.track]+1] = 1
			}
		}
		for c := range n {
			passing[c+1] += passing[c]
		}
	}
	// alongBetween counts the lines a branch along the row crosses between
	// two columns.
	alongBetween := func(a, b int) int {
		if turn < 0 || a == b {
			return 0
		}
		a, b = min(a, b), max(a, b)
		return passing[b] - passing[a+1]
	}

	// across counts the branches of placed spans that run along a row in
	// the lane of s, by column if s joins an existing track and by position
	// if it gets a new one.
	p.acrossCol = zeroed(p.acrossCol, n+1)
	p.acrossPos = zeroed(p.acrossPos, n+2)
	acrossCol, acrossPos := p.acrossCol, p.acrossPos
	for _, q := range placed {
		if q.parent < 0 || placed[q.from].track == q.track {
			continue
		}
		if r := q.turn(); r < s.end || r > s.claimBottom() {
			continue
		}
		a, b := p.col[placed[q.from].track], p.col[q.track]
		a, b = min(a, b), max(a, b)
		acrossCol[a+1]++
		acrossCol[b]--
		acrossPos[a+1]++
		acrossPos[b+1]--
	}
	for c := range n {
		acrossCol[c+1] += acrossCol[c]
	}
	for c := range n + 1 {
		acrossPos[c+1] += acrossPos[c]
	}

	// Once the columns reach the cap, a new one left of the cap would push
	// a drawn column past it.
	minPos := 0
	if p.maxLanes > 0 && n >= p.maxLanes {
		minPos = p.maxLanes
	}
	bestTrack, bestPos, bestCost, bestDist := -1, -1, -1, 0
	for t := range n {
		c := p.col[t]
		if !p.free(t, s) || (c != parentCol && shares(c > parentCol, c)) {
			continue
		}
		cost, dist := alongBetween(parentCol, c)+acrossCol[c], 0
		if turn >= 0 {
			dist = max(c-parentCol, parentCol-c)
		}
		if bestCost < 0 || cost < bestCost || (cost == bestCost && dist < bestDist) {
			bestTrack, bestCost, bestDist = t, cost, dist
		}
	}
	insert := func(pos int) {
		if shares(pos > parentCol, pos) {
			return
		}
		cost := acrossPos[pos]
		if turn >= 0 {
			// The new column sits just left of the column now at pos.
			if pos > parentCol {
				cost += passing[pos] - passing[parentCol+1]
			} else {
				cost += passing[parentCol] - passing[pos]
			}
		}
		if bestCost < 0 || cost < bestCost {
			bestTrack, bestPos, bestCost = -1, pos, cost
		}
	}
	insert(n)
	if turn >= 0 {
		for d := 1; bestCost != 0 && d <= n+1; d++ {
			for _, pos := range []int{parentCol + d, parentCol + 1 - d} {
				if pos >= minPos && pos < n {
					insert(pos)
				}
			}
		}
	}

	if bestTrack >= 0 {
		p.spans[k].track = bestTrack
		p.tracks[bestTrack] = append(p.tracks[bestTrack], k)
		return
	}
	t := len(p.tracks)
	p.spans[k].track = t
	p.tracks = append(p.tracks, []int{k})
	p.order = slices.Insert(p.order, bestPos, t)
	p.col = append(p.col, 0)
	for c, t := range p.order {
		p.col[t] = c
	}
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

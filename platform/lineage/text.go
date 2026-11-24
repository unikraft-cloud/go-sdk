// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import "strings"

// Cell is one character of a row of Grid.
type Cell struct {
	Rune rune
	// Color is the color of the line or node the character draws, as in
	// Layout.Color, or -1 for a blank.
	Color int
}

// LaneWidth is the characters a lane takes in Grid: its own, and one for
// the branches that run along the row to the next lane.
const LaneWidth = 2

// ink is the color of a cell and the row of the node that owns the color.
// Where lines and branches share a cell, the newest node, in the top row, wins.
type ink struct {
	color, row int
}

func (k *ink) paint(color, row int) {
	if k.color < 0 || row < k.row {
		k.color, k.row = color, row
	}
}

// joint is a character of Grid in a lane, by the directions it runs to.
type joint struct {
	up, down, left, right bool
	ink
}

// Grid draws l as text, a row of cells for each row of l. A branch that has
// to cross a line runs under it. A shared cell takes the newest node's color.
func (l Layout) Grid() [][]Cell {
	joints := make([][]joint, len(l.Rows))
	along := make([][]ink, len(l.Rows))
	for i := range joints {
		joints[i] = make([]joint, l.Columns)
		for c := range joints[i] {
			joints[i][c].color = -1
		}
		along[i] = make([]ink, max(l.Columns-1, 0))
		for g := range along[i] {
			along[i][g].color = -1
		}
	}

	vertical := func(lane, top, bottom, color, owner int) {
		for r := top; r <= bottom; r++ {
			j := &joints[r][lane]
			j.up = j.up || r > top
			j.down = j.down || r < bottom
			j.paint(color, owner)
		}
	}

	for _, line := range l.Lines {
		vertical(line.Lane, line.Top, line.Bottom, line.Color, line.Bottom)
	}
	for _, b := range l.Branches {
		if !b.Live {
			vertical(b.ChildLane, b.ChildRow, b.ParentRow, b.Color, b.ChildRow)
		}
	}
	for _, b := range l.Branches {
		if b.ChildLane == b.ParentLane {
			continue
		}
		row := b.ParentRow
		if b.Live {
			row = b.ChildRow
		}
		a, z := min(b.ParentLane, b.ChildLane), max(b.ParentLane, b.ChildLane)
		for c := a; c <= z; c++ {
			if c < z {
				along[row][c].paint(b.Color, b.ChildRow)
			}
			j := &joints[row][c]
			if c > a && c < z && j.up && j.down {
				continue // The branch runs under the line.
			}
			j.left = j.left || c > a
			j.right = j.right || c < z
			j.paint(b.Color, b.ChildRow)
		}
	}

	width := 0
	if l.Columns > 0 {
		width = (l.Columns-1)*LaneWidth + 1
	}
	grid := make([][]Cell, len(l.Rows))
	for i, n := range l.Rows {
		row := make([]Cell, width)
		for x := range row {
			row[x] = Cell{Rune: ' ', Color: -1}
		}
		for c, j := range joints[i] {
			if r := j.rune(); r != ' ' {
				row[c*LaneWidth] = Cell{Rune: r, Color: j.color}
			}
		}
		for g, k := range along[i] {
			if k.color >= 0 {
				for x := 1; x < LaneWidth; x++ {
					row[g*LaneWidth+x] = Cell{Rune: '─', Color: k.color}
				}
			}
		}
		row[l.Lane[n.UUID]*LaneWidth] = Cell{Rune: markRune(n), Color: l.Color[n.UUID]}
		grid[i] = row
	}
	return grid
}

func (j joint) rune() rune {
	switch {
	case j.up && j.down && j.left && j.right:
		return '┼'
	case j.up && j.down && j.right:
		return '├'
	case j.up && j.down && j.left:
		return '┤'
	case j.up && j.left && j.right:
		return '┴'
	case j.down && j.left && j.right:
		return '┬'
	case j.up && j.right:
		return '╰'
	case j.up && j.left:
		return '╯'
	case j.down && j.right:
		return '╭'
	case j.down && j.left:
		return '╮'
	case j.up || j.down:
		return '│'
	case j.left || j.right:
		return '─'
	default:
		return ' '
	}
}

func markRune(n Node) rune {
	switch {
	case n.Deleted:
		return '◌'
	case n.Mark() == MarkCheckpoint:
		return '○'
	case n.Mark() == MarkTemplate:
		return '◎'
	default:
		return '●'
	}
}

// Text draws l as text, each row of Grid followed by the label of its node.
func (l Layout) Text(label func(Node) string) string {
	var b strings.Builder
	var line strings.Builder
	for i, row := range l.Grid() {
		line.Reset()
		for _, c := range row {
			line.WriteRune(c.Rune)
		}
		if s := label(l.Rows[i]); s != "" {
			line.WriteString("    ")
			line.WriteString(s)
		}
		b.WriteString(strings.TrimRight(line.String(), " "))
		b.WriteByte('\n')
	}
	return b.String()
}

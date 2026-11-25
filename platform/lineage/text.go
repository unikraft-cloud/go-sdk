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

// laneWidth is the characters a lane takes in Grid: its own, and two for
// the branches that run along the row to the next lane.
const laneWidth = 3

// joint is a character of Grid in a lane, by the directions it runs to.
type joint struct {
	up, down, left, right bool
	color                 int
}

// Grid draws l as text, a row of cells for each row of l. A branch that has
// to cross a line runs under it.
func (l Layout) Grid() [][]Cell {
	joints := make([][]joint, len(l.Rows))
	along := make([][]int, len(l.Rows))
	for i := range joints {
		joints[i] = make([]joint, l.Columns)
		for c := range joints[i] {
			joints[i][c].color = -1
		}
		along[i] = make([]int, max(l.Columns-1, 0))
		for g := range along[i] {
			along[i][g] = -1
		}
	}

	paint := func(j *joint, color int) {
		if j.color < 0 {
			j.color = color
		}
	}
	vertical := func(lane, top, bottom, color int) {
		for r := top; r <= bottom; r++ {
			j := &joints[r][lane]
			j.up = j.up || r > top
			j.down = j.down || r < bottom
			paint(j, color)
		}
	}

	for _, line := range l.Lines {
		vertical(line.Lane, line.Top, line.Bottom, line.Color)
	}
	for _, b := range l.Branches {
		if !b.Live {
			vertical(b.ChildLane, b.ChildRow, b.ParentRow, b.Color)
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
				along[row][c] = b.Color
			}
			j := &joints[row][c]
			if c > a && c < z && j.up && j.down {
				continue // The branch runs under the line.
			}
			j.left = j.left || c > a
			j.right = j.right || c < z
			paint(j, b.Color)
		}
	}

	width := 0
	if l.Columns > 0 {
		width = (l.Columns-1)*laneWidth + 1
	}
	grid := make([][]Cell, len(l.Rows))
	for i, n := range l.Rows {
		row := make([]Cell, width)
		for x := range row {
			row[x] = Cell{Rune: ' ', Color: -1}
		}
		for c, j := range joints[i] {
			if r := j.rune(); r != ' ' {
				row[c*laneWidth] = Cell{Rune: r, Color: j.color}
			}
		}
		for g, color := range along[i] {
			if color >= 0 {
				row[g*laneWidth+1] = Cell{Rune: '─', Color: color}
				row[g*laneWidth+2] = Cell{Rune: '─', Color: color}
			}
		}
		row[l.Lane[n.UUID]*laneWidth] = Cell{Rune: markRune(n), Color: l.Color[n.UUID]}
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

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

// Find returns the node with the given UUID.
func (t Tree) Find(uuid string) (Node, bool) {
	for _, n := range t {
		if n.UUID == uuid {
			return n, true
		}
	}
	return Node{}, false
}

// CheckpointsOf returns the checkpoints that the instance with the given UUID
// took.
func (t Tree) CheckpointsOf(uuid string) Tree {
	var out Tree
	for _, n := range t {
		if n.Kind == KindCheckpoint && n.Source == uuid {
			out = append(out, n)
		}
	}
	return out
}

// ForksOf returns the instances cloned from the instance or checkpoint with
// the given UUID. For an instance, this includes the clones of its own
// checkpoints.
func (t Tree) ForksOf(uuid string) Tree {
	from := map[string]bool{uuid: true}
	for _, n := range t.CheckpointsOf(uuid) {
		from[n.UUID] = true
	}
	var out Tree
	for _, n := range t {
		if n.Kind == KindInstance && n.Source != "" && from[n.Source] {
			out = append(out, n)
		}
	}
	return out
}

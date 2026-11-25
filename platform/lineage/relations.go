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

// ForksOf returns the instances and templates cloned from the node with the
// given UUID. For an instance, this includes the clones of its own
// checkpoints.
func (t Tree) ForksOf(uuid string) Tree {
	from := map[string]bool{uuid: true}
	for _, n := range t.CheckpointsOf(uuid) {
		from[n.UUID] = true
	}
	var out Tree
	for _, n := range t {
		if n.Kind != KindCheckpoint && n.Source != "" && from[n.Source] {
			out = append(out, n)
		}
	}
	return out
}

// Mark is the shape a node is drawn with.
type Mark string

const (
	MarkInstance   Mark = "instance"
	MarkFork       Mark = "fork"
	MarkTemplate   Mark = "template"
	MarkCheckpoint Mark = "checkpoint"
)

// Mark returns the shape n is drawn with: a checkpoint, a template, an
// instance cloned from something, or any other instance.
func (n Node) Mark() Mark {
	switch {
	case n.Kind == KindCheckpoint:
		return MarkCheckpoint
	case n.Kind == KindTemplate:
		return MarkTemplate
	case n.Source != "":
		return MarkFork
	default:
		return MarkInstance
	}
}

// Relation is how a node relates to its parent.
type Relation string

const (
	// RelationRoot is an instance whose parent is not in the tree.
	RelationRoot Relation = "root"
	// RelationForkedFrom is an instance cloned from a checkpoint.
	RelationForkedFrom Relation = "forked-from"
	// RelationForkedLiveFrom is an instance cloned from a running instance.
	RelationForkedLiveFrom Relation = "forked-live-from"
	// RelationCreatedFrom is an instance or template created from a template.
	RelationCreatedFrom Relation = "created-from"
	// RelationCheckpointOf is a checkpoint of an instance with a name.
	RelationCheckpointOf Relation = "checkpoint-of"
	// RelationCheckpointOfDeleted is a checkpoint of a deleted instance, or of
	// one that is not in the tree or has no name.
	RelationCheckpointOfDeleted Relation = "checkpoint-of-deleted"
	// RelationDeleted is a placeholder for a deleted instance.
	RelationDeleted Relation = "deleted"
)

// RelationOf returns how n relates to its parent, and the parent if it is in
// t. Pass the whole tree rather than a scope of it, so that the parent of the
// oldest node in the scope is still found.
func (t Tree) RelationOf(n Node) (Relation, Node) {
	var src Node
	ok := false
	if n.Source != "" {
		src, ok = t.Find(n.Source)
	}
	switch {
	case n.Kind == KindCheckpoint:
		if !ok || src.Deleted || src.Name == "" {
			return RelationCheckpointOfDeleted, src
		}
		return RelationCheckpointOf, src
	case n.Deleted:
		return RelationDeleted, src
	case !ok:
		return RelationRoot, src
	case src.Kind == KindTemplate:
		return RelationCreatedFrom, src
	case n.Live:
		return RelationForkedLiveFrom, src
	default:
		return RelationForkedFrom, src
	}
}

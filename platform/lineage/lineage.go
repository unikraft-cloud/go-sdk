// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package lineage builds the family trees of instances: an instance is the
// child of the instance or checkpoint it was cloned from, and a checkpoint is
// the child of the instance that took it. FetchHistories reads the checkpoint
// histories that Build needs from the platform.
package lineage

import (
	"cmp"
	"slices"
	"time"

	"unikraft.com/cloud/sdk/platform"
)

// Kind tells instance nodes from checkpoint nodes.
type Kind string

const (
	KindInstance   Kind = "instance"
	KindCheckpoint Kind = "checkpoint"
)

// DeletedPrefix starts the UUID of a placeholder for a deleted instance. The
// UUID of the newest checkpoint the deleted instance left follows it.
const DeletedPrefix = "deleted:"

// Node is an instance or a checkpoint in a lineage tree.
type Node struct {
	Kind      Kind
	UUID      string
	Name      string
	CreatedAt time.Time
	State     platform.InstanceState
	// Source is the UUID of the parent: the instance that took a checkpoint,
	// or the instance or checkpoint an instance was cloned from.
	Source string
	// Live is set when an instance was cloned from the running instance in
	// Source rather than from one of its checkpoints.
	Live bool
	// Deleted is set on a placeholder for a deleted instance that is still
	// the parent of other nodes.
	Deleted bool
}

// Tree is a set of lineage trees as a flat list of nodes, each pointing at
// its parent through Node.Source.
type Tree []Node

// Build returns the lineage trees of instances, newest node first.
// histories are the checkpoint histories of instances, newest first as the
// platform lists them; a checkpoint that no listed instance took is
// attributed to a placeholder for a deleted instance.
func Build(instances []platform.Instance, histories []platform.GetCheckpointHistoryResponseInstanceHistory) Tree {
	insts := make(map[string]platform.Instance, len(instances))
	for _, inst := range instances {
		insts[inst.Uuid] = inst
	}

	// Oldest instance first, so every "first wins" below favours it deterministically.
	type hist struct {
		inst platform.Instance
		// entries are oldest first, and the instance took those from own on.
		entries []platform.CheckpointHistoryEntry
		own     int
	}
	var hs []hist
	for _, h := range histories {
		inst, ok := insts[h.Uuid]
		if !ok {
			continue
		}
		entries := slices.Clone(h.History)
		slices.Reverse(entries)
		hs = append(hs, hist{inst: inst, entries: entries, own: inherited(inst, entries)})
	}
	slices.SortStableFunc(hs, func(a, b hist) int {
		return cmp.Or(
			a.inst.CreatedAt.Compare(b.inst.CreatedAt),
			cmp.Compare(a.inst.Uuid, b.inst.Uuid),
		)
	})

	ckpts := map[string]platform.CheckpointHistoryEntry{}
	owner := map[string]string{}
	for _, h := range hs {
		for i, e := range h.entries {
			if _, ok := ckpts[e.Uuid]; !ok {
				ckpts[e.Uuid] = e
			}
			if _, ok := owner[e.Uuid]; !ok && i >= h.own {
				owner[e.Uuid] = h.inst.Uuid
			}
		}
	}

	// Orphans adjacent in some history belong to the same deleted instance.
	parent := map[string]string{}
	var find func(string) string
	find = func(x string) string {
		p, ok := parent[x]
		if !ok || p == x {
			return x
		}
		r := find(p)
		parent[x] = r
		return r
	}
	groupSource := map[string]string{}
	for _, h := range hs {
		prevOrphan, prevOwned := "", ""
		for _, e := range h.entries {
			if _, ok := owner[e.Uuid]; ok {
				prevOrphan, prevOwned = "", e.Uuid
				continue
			}
			if prevOrphan != "" {
				a, b := find(prevOrphan), find(e.Uuid)
				if a != b {
					parent[b] = a
					if s, ok := groupSource[b]; ok {
						if _, ok := groupSource[a]; !ok {
							groupSource[a] = s
						}
					}
				}
			} else if prevOwned != "" {
				if r := find(e.Uuid); groupSource[r] == "" {
					groupSource[r] = prevOwned
				}
			}
			prevOrphan = e.Uuid
		}
	}

	type group struct {
		oldest, newest platform.CheckpointHistoryEntry
	}
	groups := map[string]*group{}
	for id, e := range ckpts {
		if _, ok := owner[id]; ok {
			continue
		}
		r := find(id)
		g, ok := groups[r]
		if !ok {
			groups[r] = &group{oldest: e, newest: e}
			continue
		}
		if e.CreatedAt.Before(g.oldest.CreatedAt) ||
			(e.CreatedAt.Equal(g.oldest.CreatedAt) && e.Uuid < g.oldest.Uuid) {
			g.oldest = e
		}
		if e.CreatedAt.After(g.newest.CreatedAt) ||
			(e.CreatedAt.Equal(g.newest.CreatedAt) && e.Uuid > g.newest.Uuid) {
			g.newest = e
		}
	}

	nodes := make(Tree, 0, len(instances)+len(ckpts)+len(groups))
	for r, g := range groups {
		id := DeletedPrefix + g.newest.Uuid
		nodes = append(nodes, Node{
			Kind:      KindInstance,
			UUID:      id,
			CreatedAt: g.oldest.CreatedAt.Add(-time.Nanosecond),
			Source:    groupSource[r],
			Deleted:   true,
		})
		for cid := range ckpts {
			if _, ok := owner[cid]; !ok && find(cid) == r {
				owner[cid] = id
			}
		}
	}

	for id, e := range ckpts {
		nodes = append(nodes, Node{
			Kind:      KindCheckpoint,
			UUID:      id,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
			Source:    owner[id],
		})
	}

	// The platform reports the clone source only while it exists, so an
	// instance cloned from something gone, or from a template, is a root.
	for _, inst := range insts {
		n := Node{
			Kind:      KindInstance,
			UUID:      inst.Uuid,
			Name:      inst.Name,
			CreatedAt: inst.CreatedAt,
			State:     inst.State,
		}
		if src := inst.CloneSource; src != nil && src.Uuid != inst.Uuid {
			if _, ok := ckpts[src.Uuid]; ok {
				n.Source = src.Uuid
			} else if _, ok := insts[src.Uuid]; ok {
				n.Source, n.Live = src.Uuid, true
			}
		}
		nodes = append(nodes, n)
	}

	nodes.sortNewestFirst()
	return nodes
}

// inherited returns how many of entries, oldest first, inst inherited
// rather than took: every entry up to the checkpoint it was cloned from.
// Without that checkpoint in entries, it is every entry taken before inst
// was created. Timestamps have second resolution, so a tie counts as taken.
func inherited(inst platform.Instance, entries []platform.CheckpointHistoryEntry) int {
	if src := inst.CloneSource; src != nil {
		if i := slices.IndexFunc(entries, func(e platform.CheckpointHistoryEntry) bool { return e.Uuid == src.Uuid }); i >= 0 {
			return i + 1
		}
	}
	n := 0
	for n < len(entries) && entries[n].CreatedAt.Before(inst.CreatedAt) {
		n++
	}
	return n
}

// Scope returns the whole tree that contains current, newest node first. When
// the tree has more than limit nodes, Scope cuts it down to limit nodes that
// still form one tree, and reports truncated: current and as many of its
// ancestors as fit, then the newest nodes that fit with the ancestors they
// lack. A limit of 0 or less keeps every node. Scope returns nil when current
// is not in t.
func (t Tree) Scope(current string, limit int) (tree Tree, truncated bool) {
	byID := make(map[string]Node, len(t))
	children := map[string][]string{}
	for _, n := range t {
		byID[n.UUID] = n
		if n.Source != "" {
			children[n.Source] = append(children[n.Source], n.UUID)
		}
	}
	if _, ok := byID[current]; !ok {
		return nil, false
	}

	root := current
	seen := map[string]bool{root: true}
	for {
		src := byID[root].Source
		if _, ok := byID[src]; !ok || seen[src] {
			break
		}
		seen[src] = true
		root = src
	}

	in := map[string]bool{root: true}
	queue := []string{root}
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		tree = append(tree, byID[id])
		for _, c := range children[id] {
			if !in[c] {
				in[c] = true
				queue = append(queue, c)
			}
		}
	}
	// A cycle (inconsistent input) has no root; keep current's own chain anyway.
	for id := range seen {
		if !in[id] {
			in[id] = true
			tree = append(tree, byID[id])
		}
	}
	tree.sortNewestFirst()

	if limit <= 0 || len(tree) <= limit {
		return tree, false
	}
	keep := map[string]bool{}
	for id := current; in[id] && !keep[id] && len(keep) < limit; id = byID[id].Source {
		keep[id] = true
	}
	var path []string
	for _, n := range tree {
		if len(keep) == limit {
			break
		}
		path = path[:0]
		for id := n.UUID; in[id] && !keep[id] && !slices.Contains(path, id); id = byID[id].Source {
			path = append(path, id)
		}
		if len(keep)+len(path) <= limit {
			for _, id := range path {
				keep[id] = true
			}
		}
	}
	return slices.DeleteFunc(tree, func(n Node) bool { return !keep[n.UUID] }), true
}

// ScopeOf returns the tree that holds inst, cut down to limit as Scope does.
// When no tree holds inst, as when the listing missed it, found is false and
// the tree is inst alone, with the kind of its state.
func (t Tree) ScopeOf(inst platform.Instance, limit int) (tree Tree, truncated, found bool) {
	if tree, truncated = t.Scope(inst.Uuid, limit); tree != nil {
		return tree, truncated, true
	}
	kind := KindInstance
	if inst.State == platform.InstanceStateCheckpoint {
		kind = KindCheckpoint
	}
	return Tree{{
		Kind:      kind,
		UUID:      inst.Uuid,
		Name:      inst.Name,
		CreatedAt: inst.CreatedAt,
		State:     inst.State,
	}}, false, false
}

func (t Tree) sortNewestFirst() {
	slices.SortFunc(t, func(a, b Node) int {
		return cmp.Or(b.CreatedAt.Compare(a.CreatedAt), cmp.Compare(a.UUID, b.UUID))
	})
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"time"

	"unikraft.com/cloud/sdk/controlplane/gen"
)

// Node is a live, resource-oriented view of a node. It embeds gen.Node so
// its fields are accessed directly (node.State, node.Vcpus, ...), and adds
// behavior (Refresh, Update, Destroy, Wait) on top of the flat generated
// client.
type Node struct {
	gen.Node

	client gen.Client
	orig   gen.Node // snapshot as of the last fetch/refresh, for diffing in Update
}

// cloneNode deep-clones a gen.Node via a JSON round-trip, so that
// slice/map/pointer fields aren't aliased between a Node and its snapshot.
// in was itself just unmarshaled from the API response, so a
// marshal/unmarshal round-trip of it failing indicates a broken invariant,
// not a normal error condition.
func cloneNode(in gen.Node) gen.Node {
	b, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("controlplane: cloning node: %v", err))
	}
	var out gen.Node
	if err := json.Unmarshal(b, &out); err != nil {
		panic(fmt.Sprintf("controlplane: cloning node: %v", err))
	}
	return out
}

func wrapNode(client gen.Client, data gen.Node) *Node {
	return &Node{
		Node:   cloneNode(data),
		client: client,
		orig:   cloneNode(data),
	}
}

// Refresh re-fetches the node from the API and resets the diffing snapshot
// used by Update.
func (n *Node) Refresh(ctx context.Context) error {
	resp, err := n.client.GetNodeByUUID(ctx, n.Uuid)
	if err != nil {
		return err
	}
	if len(resp.Data.Nodes) == 0 {
		return fmt.Errorf("node %s not found", n.Uuid)
	}
	fresh := resp.Data.Nodes[0]
	n.Node, n.orig = cloneNode(fresh), cloneNode(fresh)
	return nil
}

// DestroyOption configures a call to Node.Destroy.
type DestroyOption func(*gen.DestroyNodeByUUIDOpts)

// WithDestroyForce destroys the node even if it is currently protected by DeleteLock.
func WithDestroyForce() DestroyOption {
	return func(o *gen.DestroyNodeByUUIDOpts) { o.Force = new(true) }
}

// Destroy destroys the node. Unlike Update and Wait, it does not refresh
// afterwards: the node is gone, and re-fetching it would just 404.
func (n *Node) Destroy(ctx context.Context, opts ...DestroyOption) error {
	var o gen.DestroyNodeByUUIDOpts
	for _, opt := range opts {
		opt(&o)
	}
	_, err := n.client.DestroyNodeByUUID(ctx, n.Uuid, o)
	return err
}

// WaitOption configures a call to Node.Wait.
type WaitOption func(*gen.WaitNodeByUUIDOpts)

// WithWaitStates blocks Wait until the node reaches one of the given states.
func WithWaitStates(states ...gen.NodeState) WaitOption {
	return func(o *gen.WaitNodeByUUIDOpts) { o.States = states }
}

// WithWaitTimeout blocks Wait for at most d before giving up.
func WithWaitTimeout(d time.Duration) WaitOption {
	return func(o *gen.WaitNodeByUUIDOpts) { o.TimeoutMs = new(int64(d / time.Millisecond)) }
}

// Wait blocks until the node reaches one of the desired states, or the
// timeout elapses, then refreshes the node to pick up the resulting state.
func (n *Node) Wait(ctx context.Context, opts ...WaitOption) error {
	var o gen.WaitNodeByUUIDOpts
	for _, opt := range opts {
		opt(&o)
	}
	resp, err := n.client.WaitNodeByUUID(ctx, n.Uuid, o)
	if err != nil {
		return err
	}
	if len(resp.Data.Nodes) == 0 {
		return fmt.Errorf("node %s not found", n.Uuid)
	}
	fresh := resp.Data.Nodes[0]
	n.Node, n.orig = cloneNode(fresh), cloneNode(fresh)
	return nil
}

// Update diffs the node's current field values against the snapshot taken
// at the last fetch/refresh, sends the minimal patch for whatever changed,
// and always refreshes afterwards.
//
// Only Tags and DeleteLock are diffed here: SshKeys updates use add/del
// semantics rather than set (gen.MutableNodePropertySshKeys), so there's no
// single before/after value to diff against — changing keys still requires
// going through Raw().UpdateNodeByUUID(...) directly.
func (n *Node) Update(ctx context.Context) error {
	items, err := n.diff()
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	_, err = n.client.UpdateNodeByUUID(ctx, n.Uuid, items)
	if rerr := n.Refresh(ctx); rerr != nil {
		err = errors.Join(err, rerr)
	}
	return err
}

func (n *Node) diff() ([]gen.UpdateNodePayload, error) {
	var items []gen.UpdateNodePayload
	set := func(prop gen.MutableNodeProperty, value any) error {
		b, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("marshalling %s: %w", prop, err)
		}
		v := string(b)
		items = append(items, gen.UpdateNodePayload{
			Property: prop, Operation: gen.MutableNodeOperationSet, Value: &v,
		})
		return nil
	}

	if !maps.Equal(n.Tags, n.orig.Tags) {
		if err := set(gen.MutableNodePropertyTags, n.Tags); err != nil {
			return nil, err
		}
	}
	if n.DeleteLock != n.orig.DeleteLock {
		if err := set(gen.MutableNodePropertyDeleteLock, n.DeleteLock); err != nil {
			return nil, err
		}
	}

	// SshKeys intentionally not diffed: see the Update doc comment.

	return items, nil
}

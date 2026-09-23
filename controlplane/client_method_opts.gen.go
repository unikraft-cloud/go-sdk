// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "time"

// CheckAuthorizationOpts holds options for [Client.CheckAuthorization].
type CheckAuthorizationOpts struct {
	// HeartbeatTimeout is how long the stream may be silent before the connection
	// is treated as dead and closed, so the caller reconnects rather than
	// waiting on a read which will never return.
	//
	// What this bounds is the absence of the server's heartbeat, not the
	// absence of events: a stream with nothing to report is normal. Zero keeps
	// DefaultStreamHeartbeatTimeout.
	HeartbeatTimeout time.Duration
}

// ListImagesOpts holds options for [Client.ListImages].
type ListImagesOpts struct {
	Details   *bool
	Namespace []string
}

// DestroyNodeOpts holds options for [Client.DestroyNode].
type DestroyNodeOpts struct {
	Force *bool
}

// DestroyNodeByUUIDOpts holds options for [Client.DestroyNodeByUUID].
type DestroyNodeByUUIDOpts struct {
	Force *bool
}

// ListMachineTypesOpts holds options for [Client.ListMachineTypes].
type ListMachineTypesOpts struct {
	Region *string
}

// ListNodesOpts holds options for [Client.ListNodes].
type ListNodesOpts struct {
	Uuid          []string
	Name          []string
	Cloudprovider *CloudProvider
	State         *NodeState
	Metro         *string
	Limit         *uint32
	Offset        *uint32
}

// UpdateNodesOpts holds options for [Client.UpdateNodes].
type UpdateNodesOpts struct {
	Property  []MutableNodeProperty
	Operation []MutableNodeOperation
	Value     []string
}

// WaitNodeByUUIDOpts holds options for [Client.WaitNodeByUUID].
type WaitNodeByUUIDOpts struct {
	States    []NodeState
	TimeoutMs *int64
}

// WaitNodesOpts holds options for [Client.WaitNodes].
type WaitNodesOpts struct {
	Uuid      []string
	Name      []string
	States    []NodeState
	TimeoutMs *int64
}

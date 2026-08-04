// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// ListImagesOpts holds query-parameter options for [Client.ListImages].
type ListImagesOpts struct {
	Details   *bool
	Namespace []string
}

// DestroyNodeOpts holds query-parameter options for [Client.DestroyNode].
type DestroyNodeOpts struct {
	Force *bool
}

// DestroyNodeByUUIDOpts holds query-parameter options for [Client.DestroyNodeByUUID].
type DestroyNodeByUUIDOpts struct {
	Force *bool
}

// ListMachineTypesOpts holds query-parameter options for [Client.ListMachineTypes].
type ListMachineTypesOpts struct {
	Region *string
}

// ListNodesOpts holds query-parameter options for [Client.ListNodes].
type ListNodesOpts struct {
	Uuid          []string
	Name          []string
	Cloudprovider *CloudProvider
	State         *NodeState
	Metro         *string
	Limit         *uint32
	Offset        *uint32
}

// UpdateNodesOpts holds query-parameter options for [Client.UpdateNodes].
type UpdateNodesOpts struct {
	Property  []MutableNodeProperty
	Operation []MutableNodeOperation
	Value     []string
}

// WaitNodeByUUIDOpts holds query-parameter options for [Client.WaitNodeByUUID].
type WaitNodeByUUIDOpts struct {
	States    []NodeState
	TimeoutMs *int64
}

// WaitNodesOpts holds query-parameter options for [Client.WaitNodes].
type WaitNodesOpts struct {
	Uuid      []string
	Name      []string
	States    []NodeState
	TimeoutMs *int64
}

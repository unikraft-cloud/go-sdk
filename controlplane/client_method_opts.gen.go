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
	Provider *string
	State    *string
	Metro    *string
	Region   *string
	Limit    *int32
	Offset   *int32
}

// WaitNodeByUUIDOpts holds query-parameter options for [Client.WaitNodeByUUID].
type WaitNodeByUUIDOpts struct {
	States    []NodeState
	TimeoutMs *int32
}

// WaitNodesOpts holds query-parameter options for [Client.WaitNodes].
type WaitNodesOpts struct {
	States    []NodeState
	TimeoutMs *int32
}

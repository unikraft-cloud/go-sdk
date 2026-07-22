// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

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
	Limit    *uint32
	Offset   *uint32
}

// WaitNodeByUUIDOpts holds query-parameter options for [Client.WaitNodeByUUID].
type WaitNodeByUUIDOpts struct {
	States    []string
	TimeoutMs *int64
}

// WaitNodesOpts holds query-parameter options for [Client.WaitNodes].
type WaitNodesOpts struct {
	States    []string
	TimeoutMs *int64
}

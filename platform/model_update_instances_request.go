// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request message for updating one or more instances.
// The property to modify.
type UpdateInstancesRequestProp string

const (
	UpdateInstancesRequestPropImage         UpdateInstancesRequestProp = "image"
	UpdateInstancesRequestPropArgs          UpdateInstancesRequestProp = "args"
	UpdateInstancesRequestPropEnv           UpdateInstancesRequestProp = "env"
	UpdateInstancesRequestPropMemory_mb     UpdateInstancesRequestProp = "memory_mb"
	UpdateInstancesRequestPropVcpus         UpdateInstancesRequestProp = "vcpus"
	UpdateInstancesRequestPropScale_to_zero UpdateInstancesRequestProp = "scale_to_zero"
	UpdateInstancesRequestPropTags          UpdateInstancesRequestProp = "tags"
	UpdateInstancesRequestPropDelete_lock   UpdateInstancesRequestProp = "delete_lock"
)

// The operation to perform on the property.
type UpdateInstancesRequestOp string

const (
	UpdateInstancesRequestOpSet UpdateInstancesRequestOp = "set"
	UpdateInstancesRequestOpAdd UpdateInstancesRequestOp = "add"
	UpdateInstancesRequestOpDel UpdateInstancesRequestOp = "del"
)

type UpdateInstancesRequest struct {
	// (Optional).  A client-provided identifier for tracking this operation in the response.
	Id *string `json:"id,omitempty"`
	// The UUID of the instance to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// The property to modify.
	Prop UpdateInstancesRequestProp `json:"prop"`
	// The operation to perform on the property.
	Op UpdateInstancesRequestOp `json:"op"`
	// The value for the update operation. The type depends on the property and operation:
	// - For "image": string
	// - For "args": string or array of strings
	// - For "env": object (for SET/ADD) or string/array of strings (for DEL)
	// - For "memory_mb": integer
	// - For "vcpus": integer
	// - For "scale_to_zero": object with cooldown_time_ms, policy, and stateful fields
	// - For "tags": array of strings
	// - For "delete_lock": boolean
	Value *interface{} `json:"value,omitempty"`
}

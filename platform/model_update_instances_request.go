// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// A single update operation to be applied to an instance.
// The property to modify. Must be one of the supported properties:
// - "image": Change the instance image (SET only)
// - "args": Update application arguments (SET only)
// - "env": Modify environment variables (SET, ADD, DEL)
// - "memory_mb": Change memory allocation (SET only)
// - "vcpus": Change CPU allocation (SET only)
// - "scale_to_zero": Configure scale-to-zero settings (SET only)
// - "tags": Manage instance tags (SET, ADD, DEL)
// - "delete_lock": Enable/disable deletion protection (SET only)
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

// The operation to perform on the property. Valid operations depend on the property:
// - "set": Supported by all properties
// - "add": Only supported by "env" and "tags"
// - "del": Only supported by "env" and "tags"
type UpdateInstancesRequestOp string

const (
	UpdateInstancesRequestOpSet UpdateInstancesRequestOp = "set"
	UpdateInstancesRequestOpAdd UpdateInstancesRequestOp = "add"
	UpdateInstancesRequestOpDel UpdateInstancesRequestOp = "del"
)

type UpdateInstancesRequest struct {
	// (Optional). A client-provided identifier for tracking this operation in the response.
	Id *interface{} `json:"id,omitempty"`
	// The UUID of the instance to update. Mutually exclusive with name.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance to update. Mutually exclusive with UUID.
	Name *string `json:"name,omitempty"`
	// The property to modify. Must be one of the supported properties:
	// - "image": Change the instance image (SET only)
	// - "args": Update application arguments (SET only)
	// - "env": Modify environment variables (SET, ADD, DEL)
	// - "memory_mb": Change memory allocation (SET only)
	// - "vcpus": Change CPU allocation (SET only)
	// - "scale_to_zero": Configure scale-to-zero settings (SET only)
	// - "tags": Manage instance tags (SET, ADD, DEL)
	// - "delete_lock": Enable/disable deletion protection (SET only)
	Prop UpdateInstancesRequestProp `json:"prop"`
	// The operation to perform on the property. Valid operations depend on the property:
	// - "set": Supported by all properties
	// - "add": Only supported by "env" and "tags"
	// - "del": Only supported by "env" and "tags"
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

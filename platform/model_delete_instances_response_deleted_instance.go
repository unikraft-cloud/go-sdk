// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Details of the instance which was deleted by this request.
// Indicates whether the start operation was successful or not for this
// instance.
type DeleteInstancesResponseDeletedInstanceStatus string

const (
	DeleteInstancesResponseDeletedInstanceStatusSuccess DeleteInstancesResponseDeletedInstanceStatus = "success"
	DeleteInstancesResponseDeletedInstanceStatusError   DeleteInstancesResponseDeletedInstanceStatus = "error"
)

type DeleteInstancesResponseDeletedInstance struct {
	// Indicates whether the start operation was successful or not for this
	// instance.
	Status *DeleteInstancesResponseDeletedInstanceStatus `json:"status,omitempty"`
	// The UUID of the instance which was deleted.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance which was deleted.
	Name *string `json:"name,omitempty"`
	// The previous state of the instance before it was deleted.
	PreviousState *string `json:"previous_state,omitempty"`
}

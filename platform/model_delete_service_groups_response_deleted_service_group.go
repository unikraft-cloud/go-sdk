// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Details of the service group which was deleted by this request.
// Indicates whether the delete operation was successful or not for this
// service group.
type DeleteServiceGroupsResponseDeletedServiceGroupStatus string

const (
	DeleteServiceGroupsResponseDeletedServiceGroupStatusSuccess DeleteServiceGroupsResponseDeletedServiceGroupStatus = "success"
	DeleteServiceGroupsResponseDeletedServiceGroupStatusError   DeleteServiceGroupsResponseDeletedServiceGroupStatus = "error"
)

type DeleteServiceGroupsResponseDeletedServiceGroup struct {
	// Indicates whether the delete operation was successful or not for this
	// service group.
	Status *DeleteServiceGroupsResponseDeletedServiceGroupStatus `json:"status,omitempty"`
	// The UUID of the service group which was deleted.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service group which was deleted.
	Name *string `json:"name,omitempty"`
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request message for getting one or more template instances.

type GetTemplateInstancesRequest struct {
	// The list of IDs of the template instances to retrieve.  If not set,
	// all template instances matching filters will be returned.
	Ids []NameOrUUID `json:"ids,omitempty"`
	// Whether to include details about the templates in the response.  By default
	// this is set to true, meaning that all information about the templates will
	// be included in the response.  If set to false, only the basic information
	// about the templates will be included, such as their name and UUID.
	Details *bool `json:"details,omitempty"`
	// If set, the listing starts from (but does not include) the template with
	// the given UUID.  This is useful for pagination.
	FromUuid *string `json:"from_uuid,omitempty"`
	// The maximum number of template instances to return.  This is useful for
	// pagination.  If not set, all the template instances matching filters will
	// be returned.  When filtering by IDs, this should not be set.
	Count *uint32 `json:"count,omitempty"`
	// A list of tags to filter the template instances by.
	Tags []string `json:"tags,omitempty"`
}

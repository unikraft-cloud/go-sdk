// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// GetAutoscaleConfigurationsOpts holds query-parameter options for [Client.GetAutoscaleConfigurations].
type GetAutoscaleConfigurationsOpts struct {
	Uuid []string
	Name []string
}

// GetCertificatesOpts holds query-parameter options for [Client.GetCertificates].
type GetCertificatesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetImageStoreOpts holds query-parameter options for [Client.GetImageStore].
type GetImageStoreOpts struct {
	Digest *string
	Tag    *string
}

// GetImagesOpts holds query-parameter options for [Client.GetImages].
type GetImagesOpts struct {
	Digest *string
	Tag    *string
}

// GetCheckpointHistoryOpts holds query-parameter options for [Client.GetCheckpointHistory].
type GetCheckpointHistoryOpts struct {
	Uuid []string
	Name []string
}

// GetCheckpointInstanceByUUIDOpts holds query-parameter options for [Client.GetCheckpointInstanceByUUID].
type GetCheckpointInstanceByUUIDOpts struct {
	Details *bool
}

// GetCheckpointInstancesOpts holds query-parameter options for [Client.GetCheckpointInstances].
type GetCheckpointInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	Tags    []string
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetInstanceByUUIDOpts holds query-parameter options for [Client.GetInstanceByUUID].
type GetInstanceByUUIDOpts struct {
	Details *bool
}

// GetInstanceHistoryOpts holds query-parameter options for [Client.GetInstanceHistory].
type GetInstanceHistoryOpts struct {
	Uuid []string
	Name []string
}

// GetInstanceLogsOpts holds query-parameter options for [Client.GetInstanceLogs].
type GetInstanceLogsOpts struct {
	Uuid   []string
	Name   []string
	Offset []int64
	Limit  []int64
}

// GetInstanceMetricsOpts holds query-parameter options for [Client.GetInstanceMetrics].
type GetInstanceMetricsOpts struct {
	Uuid []string
	Name []string
}

// GetInstancesOpts holds query-parameter options for [Client.GetInstances].
type GetInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Tags    []string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetTemplateInstanceByUUIDOpts holds query-parameter options for [Client.GetTemplateInstanceByUUID].
type GetTemplateInstanceByUUIDOpts struct {
	Details *bool
}

// GetTemplateInstancesOpts holds query-parameter options for [Client.GetTemplateInstances].
type GetTemplateInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	Tags    []string
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// WaitInstancesOpts holds query-parameter options for [Client.WaitInstances].
type WaitInstancesOpts struct {
	Uuid      []string
	Name      []string
	State     []InstanceState
	TimeoutMs []int64
	TimeoutS  []int64
}

// GetServiceGroupByUUIDOpts holds query-parameter options for [Client.GetServiceGroupByUUID].
type GetServiceGroupByUUIDOpts struct {
	Details *bool
}

// GetServiceGroupsOpts holds query-parameter options for [Client.GetServiceGroups].
type GetServiceGroupsOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetTemplateVolumeByUUIDOpts holds query-parameter options for [Client.GetTemplateVolumeByUUID].
type GetTemplateVolumeByUUIDOpts struct {
	Details *bool
}

// GetTemplateVolumesOpts holds query-parameter options for [Client.GetTemplateVolumes].
type GetTemplateVolumesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Tags    []string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetVolumeByUUIDOpts holds query-parameter options for [Client.GetVolumeByUUID].
type GetVolumeByUUIDOpts struct {
	Details *bool
}

// GetVolumesOpts holds query-parameter options for [Client.GetVolumes].
type GetVolumesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Tags    []string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "time"

// SubscribeAuditEventsOpts holds options for [Client.SubscribeAuditEvents].
type SubscribeAuditEventsOpts struct {
	Events []AuditEventType
	Uuid   []string
	Tags   []string

	// HeartbeatTimeout is how long the stream may be silent before the connection
	// is treated as dead and closed, so the caller reconnects rather than
	// waiting on a read which will never return.
	//
	// What this bounds is the absence of the server's heartbeat, not the
	// absence of events: a stream with nothing to report is normal. Zero keeps
	// DefaultStreamHeartbeatTimeout.
	HeartbeatTimeout time.Duration
}

// GetAutoscaleConfigurationsOpts holds options for [Client.GetAutoscaleConfigurations].
type GetAutoscaleConfigurationsOpts struct {
	Uuid []string
	Name []string
}

// GetCertificatesOpts holds options for [Client.GetCertificates].
type GetCertificatesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetImageStoreOpts holds options for [Client.GetImageStore].
type GetImageStoreOpts struct {
	Digest *string
	Tag    *string
}

// GetImagesOpts holds options for [Client.GetImages].
type GetImagesOpts struct {
	Digest *string
	Tag    *string
}

// GetCheckpointHistoryOpts holds options for [Client.GetCheckpointHistory].
type GetCheckpointHistoryOpts struct {
	Uuid []string
	Name []string
}

// GetCheckpointInstanceByUUIDOpts holds options for [Client.GetCheckpointInstanceByUUID].
type GetCheckpointInstanceByUUIDOpts struct {
	Details *bool
}

// GetCheckpointInstancesOpts holds options for [Client.GetCheckpointInstances].
type GetCheckpointInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
	Tags    []string
}

// GetInstanceByUUIDOpts holds options for [Client.GetInstanceByUUID].
type GetInstanceByUUIDOpts struct {
	Details *bool
}

// GetInstanceHistoryOpts holds options for [Client.GetInstanceHistory].
type GetInstanceHistoryOpts struct {
	Uuid []string
	Name []string
}

// GetInstanceLogsOpts holds options for [Client.GetInstanceLogs].
type GetInstanceLogsOpts struct {
	Uuid   []string
	Name   []string
	Offset []int64
	Limit  []int64
}

// GetInstanceMetricsOpts holds options for [Client.GetInstanceMetrics].
type GetInstanceMetricsOpts struct {
	Uuid []string
	Name []string
}

// GetInstancesOpts holds options for [Client.GetInstances].
type GetInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
	Tags    []string
}

// GetTemplateInstanceByUUIDOpts holds options for [Client.GetTemplateInstanceByUUID].
type GetTemplateInstanceByUUIDOpts struct {
	Details *bool
}

// GetTemplateInstancesOpts holds options for [Client.GetTemplateInstances].
type GetTemplateInstancesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
	Tags    []string
}

// WaitInstancesOpts holds options for [Client.WaitInstances].
type WaitInstancesOpts struct {
	Uuid      []string
	Name      []string
	State     []InstanceState
	TimeoutMs []int64
	TimeoutS  []int64
}

// GetServiceGroupByUUIDOpts holds options for [Client.GetServiceGroupByUUID].
type GetServiceGroupByUUIDOpts struct {
	Details *bool
}

// GetServiceGroupsOpts holds options for [Client.GetServiceGroups].
type GetServiceGroupsOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
}

// GetTemplateVolumeByUUIDOpts holds options for [Client.GetTemplateVolumeByUUID].
type GetTemplateVolumeByUUIDOpts struct {
	Details *bool
}

// GetTemplateVolumesOpts holds options for [Client.GetTemplateVolumes].
type GetTemplateVolumesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
	Tags    []string
}

// GetVolumeByUUIDOpts holds options for [Client.GetVolumeByUUID].
type GetVolumeByUUIDOpts struct {
	Details *bool
}

// GetVolumesOpts holds options for [Client.GetVolumes].
type GetVolumesOpts struct {
	Uuid    []string
	Name    []string
	Details *bool
	Count   *uint32
	From    *string
	Order   *PaginationOrder
	Sortby  *PaginationSortBy
	Tags    []string
}

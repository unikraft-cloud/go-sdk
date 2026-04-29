// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// GetCertificatesOpts holds query-parameter options for [Client.GetCertificates].
type GetCertificatesOpts struct {
	Details *bool
	Count   *int32
	From    *string
	Order   *string
	Sortby  *string
}

// GetImageStoreOpts holds query-parameter options for [Client.GetImageStore].
type GetImageStoreOpts struct {
	Metro  *string
	Digest *string
	Tag    *string
}

// GetImagesOpts holds query-parameter options for [Client.GetImages].
type GetImagesOpts struct {
	Metro  *string
	Digest *string
	Tag    *string
}

// GetInstanceByUUIDOpts holds query-parameter options for [Client.GetInstanceByUUID].
type GetInstanceByUUIDOpts struct {
	Details *bool
}

// GetInstancesOpts holds query-parameter options for [Client.GetInstances].
type GetInstancesOpts struct {
	Details *bool
	Count   *int32
	From    *string
	Tags    []string
	Order   *string
	Sortby  *string
}

// GetTemplateInstanceByUUIDOpts holds query-parameter options for [Client.GetTemplateInstanceByUUID].
type GetTemplateInstanceByUUIDOpts struct {
	Details *bool
}

// GetTemplateInstancesOpts holds query-parameter options for [Client.GetTemplateInstances].
type GetTemplateInstancesOpts struct {
	Details *bool
	Count   *int32
	Tags    []string
	From    *string
	Order   *string
	Sortby  *string
}

// GetServiceGroupByUUIDOpts holds query-parameter options for [Client.GetServiceGroupByUUID].
type GetServiceGroupByUUIDOpts struct {
	Details *bool
}

// GetServiceGroupsOpts holds query-parameter options for [Client.GetServiceGroups].
type GetServiceGroupsOpts struct {
	Details *bool
	Count   *int32
	From    *string
	Order   *string
	Sortby  *string
}

// GetTemplateVolumeByUUIDOpts holds query-parameter options for [Client.GetTemplateVolumeByUUID].
type GetTemplateVolumeByUUIDOpts struct {
	Details *bool
}

// GetTemplateVolumesOpts holds query-parameter options for [Client.GetTemplateVolumes].
type GetTemplateVolumesOpts struct {
	Details *bool
	Count   *int32
	From    *string
	Tags    []string
	Order   *string
	Sortby  *string
}

// GetVolumeByUUIDOpts holds query-parameter options for [Client.GetVolumeByUUID].
type GetVolumeByUUIDOpts struct {
	Details *bool
}

// GetVolumesOpts holds query-parameter options for [Client.GetVolumes].
type GetVolumesOpts struct {
	Details *bool
	Count   *int32
	From    *string
	Tags    []string
	Order   *string
	Sortby  *string
}

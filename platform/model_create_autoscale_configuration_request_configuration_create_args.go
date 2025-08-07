// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The arguments to use when creating the autoscale configuration.
type CreateAutoscaleConfigurationRequestConfigurationCreateArgs struct {
	// The ROM to use for the autoscale configuration.
	Roms *InstanceCreateArgsInstanceCreateRequestRoms `json:"roms,omitempty"`
	// The template to use for the autoscale configuration.
	Template *InstanceCreateArgsInstanceTemplateCreateRequest `json:"template,omitempty"`
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The request to create an autoscale configuration for a service.
type CreateAutoscaleConfigurationRequest struct {
	Configuration []CreateAutoscaleConfigurationRequestConfiguration `json:"configuration,omitempty"`
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type CreateAutoscaleConfigurationResponseConfigurationsResponseStatus string

const (
	CreateAutoscaleConfigurationResponseConfigurationsResponseStatusSuccess CreateAutoscaleConfigurationResponseConfigurationsResponseStatus = "success"
	CreateAutoscaleConfigurationResponseConfigurationsResponseStatusError   CreateAutoscaleConfigurationResponseConfigurationsResponseStatus = "error"
)

type CreateAutoscaleConfigurationResponseConfigurationsResponse struct {
	// The status of the response.
	Status *CreateAutoscaleConfigurationResponseConfigurationsResponseStatus `json:"status,omitempty"`
	// The UUID of the service where the configuration was created.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service where the configuration was created.
	Name *string `json:"name,omitempty"`
}

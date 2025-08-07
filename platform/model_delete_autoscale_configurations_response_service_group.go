// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type DeleteAutoscaleConfigurationsResponseServiceGroupStatus string

const (
	DeleteAutoscaleConfigurationsResponseServiceGroupStatusSuccess DeleteAutoscaleConfigurationsResponseServiceGroupStatus = "success"
	DeleteAutoscaleConfigurationsResponseServiceGroupStatusError   DeleteAutoscaleConfigurationsResponseServiceGroupStatus = "error"
)

type DeleteAutoscaleConfigurationsResponseServiceGroup struct {
	// The status of the response.
	Status *DeleteAutoscaleConfigurationsResponseServiceGroupStatus `json:"status,omitempty"`
	// The UUID of the service where the configuration was deleted.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the service where the configuration was deleted.
	Name *string `json:"name,omitempty"`
}

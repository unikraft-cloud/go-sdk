// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type DeleteAutoscaleConfigurationsResponseAllOfStatus string

const (
	DeleteAutoscaleConfigurationsResponseAllOfStatusSuccess DeleteAutoscaleConfigurationsResponseAllOfStatus = "success"
	DeleteAutoscaleConfigurationsResponseAllOfStatusError   DeleteAutoscaleConfigurationsResponseAllOfStatus = "error"
)

type DeleteAutoscaleConfigurationsResponseAllOf struct {
	// The status of the response.
	Status *DeleteAutoscaleConfigurationsResponseAllOfStatus `json:"status,omitempty"`
	// The response data for this request.
	Data *DeleteAutoscaleConfigurationsResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs *uint64 `json:"op_time_us,omitempty"`
}

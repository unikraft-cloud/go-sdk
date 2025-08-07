// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The response message for creating of a volume.
// The status of the response.
type CreateVolumeResponseStatus string

const (
	CreateVolumeResponseStatusSuccess CreateVolumeResponseStatus = "success"
	CreateVolumeResponseStatusError   CreateVolumeResponseStatus = "error"
)

type CreateVolumeResponse struct {
	// The status of the response.
	Status *CreateVolumeResponseStatus `json:"status,omitempty"`
	Data   *CreateVolumeResponseData   `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs *int32 `json:"op_time_us,omitempty"`
	// An optional message providing additional information about the response.
	Message *string `json:"message,omitempty"`
}

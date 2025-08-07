// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The response message for waiting for one or more instance(s) to reach a
// certain state given their UUID(s) or name(s).
// The status of the response.
type WaitInstanceResponseStatus string

const (
	WaitInstanceResponseStatusSuccess WaitInstanceResponseStatus = "success"
	WaitInstanceResponseStatusError   WaitInstanceResponseStatus = "error"
)

type WaitInstanceResponse struct {
	// The status of the response.
	Status *WaitInstanceResponseStatus `json:"status,omitempty"`
	// An optional message providing additional information about the response.
	Message *string                   `json:"message,omitempty"`
	Data    *WaitInstanceResponseData `json:"data,omitempty"`
	// A list of errors which may have occurred during the request.
	Errors []ResponseError `json:"errors,omitempty"`
	// The operation time in microseconds.  This is the time it took to process
	// the request and generate the response.
	OpTimeUs *int32 `json:"op_time_us,omitempty"`
}

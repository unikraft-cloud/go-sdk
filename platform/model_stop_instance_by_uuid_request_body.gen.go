// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Parameters for stopping the instance.
type StopInstanceByUUIDRequestBody struct {
	// Whether to immediately force stop the instance.
	Force *bool `json:"force,omitempty"`
	// Timeout for draining connections in milliseconds.
	// No draining will occur if set to 0.  The instance
	// does not receive new connections in the draining
	// phase.  The instance is stopped when the last
	// connection has been closed or the timeout expired.
	// The maximum timeout may vary.  Use -1 for the
	// largest possible value.  Ignored if force is set.
	//
	// Note: This endpoint does not block.  Use the wait
	// endpoint for the instance to reach the stopped
	// state.
	DrainTimeoutMs *uint64 `json:"drain_timeout_ms,omitempty"`
	// Whether to perform a quick shutdown.  This flag is
	// overridden by force.
	Quick *bool `json:"quick,omitempty"`
	// Only stop the instance if it is in this state.
	Ifstate *string `json:"ifstate,omitempty"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",inline"`
}

func (m *StopInstanceByUUIDRequestBody) UnmarshalJSON(data []byte) error {
	type Alias StopInstanceByUUIDRequestBody
	return json.Unmarshal(data, (*Alias)(m))
}

func (m StopInstanceByUUIDRequestBody) MarshalJSON() ([]byte, error) {
	type Alias StopInstanceByUUIDRequestBody
	return json.Marshal((Alias)(m))
}

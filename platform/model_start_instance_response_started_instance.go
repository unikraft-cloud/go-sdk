// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type StartInstanceResponseStartedInstance struct {
	// Indicates whether the start operation was successful or not for this
	// instance.
	Status *ResponseStatus `json:"status,omitempty"`
	// The UUID of the instance which was deleted.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the instance which was deleted.
	Name *string `json:"name,omitempty"`
	// The current state of the instance after this request.
	State *string `json:"state,omitempty"`
	// The previous state of the instance before it was deleted.
	PreviousState *string `json:"previous_state,omitempty"`
}

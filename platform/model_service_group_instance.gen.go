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

type ServiceGroupInstance struct {
	// The UUID of the instance.  This is a unique identifier for the instance
	// that is generated when the instance is created.  The UUID is used to
	// reference the instance in API calls and can be used to identify the
	// instance in all API calls that require an instance identifier.
	Uuid string `json:"uuid"`
	// The name of the instance.  This is a human-readable name that can be used
	// to identify the instance.  The name must be unique within the context of
	// your account.  If no name is specified, a random name is generated for
	// you.  The name can also be used to identify the instance in API calls.
	Name string `json:"name"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *ServiceGroupInstance) UnmarshalJSON(data []byte) error {
	type Alias ServiceGroupInstance
	return json.Unmarshal(data, (*Alias)(m))
}

func (m ServiceGroupInstance) MarshalJSON() ([]byte, error) {
	type Alias ServiceGroupInstance
	return json.Marshal((Alias)(m))
}

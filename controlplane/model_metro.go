// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

type Metro struct {
	// The UUID of the metro.
	Uuid *string `json:"uuid,omitempty"`
	// The API endpoint for the metro.
	Endpoint *string `json:"endpoint,omitempty"`
	// The name of the metro.
	Name *string `json:"name,omitempty"`
	// The IATA code of the metro.
	IataCode *string `json:"iata_code,omitempty"`
	// The country where the metro is located.
	Country *string `json:"country,omitempty"`
}

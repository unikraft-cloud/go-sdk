// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type CreateInstanceRequestDomainAllOf struct {
	// Publicly accessible domain name.
	//
	// If this name ends in a period `.` it must be a valid Full Qualified
	// Domain Name (FQDN), e.g. `example.com.`; otherwise it will become a
	// subdomain of the target metro, e.g. `example` becomes
	// `example.fra0.unikraft.app`.
	Name *string `json:"name,omitempty"`
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type CreateInstanceRequestServiceGroupAllOf struct {
	// Similarly, if no existing (persistent) service group is specified via its
	// identifier, a new (ephemeral) service group can be created.  In addition
	// to the services it must expose, you can specify which domains it should
	// use too.
	Domains []CreateInstanceRequestDomain `json:"domains,omitempty"`
}

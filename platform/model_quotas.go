// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type Quotas struct {
	// The UUID of the quota.
	Uuid   *string       `json:"uuid,omitempty"`
	Used   *QuotasUsed   `json:"used,omitempty"`
	Hard   *QuotasHard   `json:"hard,omitempty"`
	Limits *QuotasLimits `json:"limits,omitempty"`
}

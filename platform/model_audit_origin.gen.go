// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// What caused the operation.
type AuditOrigin string

const (
	AuditOriginUnknown     AuditOrigin = "unknown"
	AuditOriginApi         AuditOrigin = "api"
	AuditOriginGuest       AuditOrigin = "guest"
	AuditOriginProxy       AuditOrigin = "proxy"
	AuditOriginAutoscale   AuditOrigin = "autoscale"
	AuditOriginScaleToZero AuditOrigin = "scale-to-zero"
	AuditOriginScheduledOp AuditOrigin = "scheduled-op"
	AuditOriginRestart     AuditOrigin = "restart"
	AuditOriginUpdate      AuditOrigin = "update"
	AuditOriginSystem      AuditOrigin = "system"
	AuditOriginNetwork     AuditOrigin = "network"
	AuditOriginAutokill    AuditOrigin = "autokill"
	AuditOriginMtss        AuditOrigin = "mtss"
)

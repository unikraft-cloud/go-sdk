// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The type of an audit event.
//
// An enum rather than an open union because this is also a query parameter, and
// an open union generates as an interface that a client cannot serialise into a
// filter. Further event types are added here; a Go client decodes one it does
// not know as its plain string value rather than failing.
type AuditEventType string

const (
	AuditEventTypeVmStateChange AuditEventType = "vm.state_change"
	AuditEventTypeVmStartFailed AuditEventType = "vm.start_failed"
)

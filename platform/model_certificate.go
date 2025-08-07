// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "time"

type CertificateState string

const (
	CertificateStatePending CertificateState = "pending"
	CertificateStateValid   CertificateState = "valid"
	CertificateStateError   CertificateState = "error"
)

type Certificate struct {
	// The UUID of the certificate.
	//
	// This is a unique identifier for the certificate that is generated when the
	// certificate is created.  The UUID is used to reference the certificate in
	// API calls and can be used to identify the certificate in all API calls that
	// require an identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The name of the certificate.
	//
	// This is a human-readable name that can be used to identify the certificate.
	// The name must be unique within the context of your account.  The name can
	// also be used to identify the certificate in API calls.
	Name *string `json:"name,omitempty"`
	// The time the certificate was created.
	CreatedAt  *time.Time        `json:"created_at,omitempty"`
	CommonName *string           `json:"common_name,omitempty"`
	State      *CertificateState `json:"state,omitempty"`
}

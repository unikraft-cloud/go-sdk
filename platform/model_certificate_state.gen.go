// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The current state of a certificate.
//
// The certificate can be in one of the following states:
//
// | State     | Description |
// |-----------|-------------|
// | `pending` | The certificate request is pending while the certificate is being requested from the certification authority. During this phase any service using this certificate is not available if this is not a renewal. |
// | `valid`   | The certificate is valid and can be used by your services. |
// | `error`   | The certificate request failed after multiple attempts. This can happen, for example, if your DNS configuration is not correct, you run into Let’s Encrypt™ quota limits, or the domain validation process failed for some other reason. There won’t be any further automatic attempts. |
type CertificateState string

const (
	CertificateStatePending CertificateState = "pending"
	CertificateStateValid   CertificateState = "valid"
	CertificateStateError   CertificateState = "error"
)

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type CreateCertificateRequestAllOf struct {
	// The common name (CN) of the certificate.
	Cn *string `json:"cn,omitempty"`
	// The chain of the certificate.
	Chain *string `json:"chain,omitempty"`
	// The private key of the certificate.
	Pkey *string `json:"pkey,omitempty"`
}

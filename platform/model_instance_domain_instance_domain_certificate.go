// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The certificate associated with the domain.
//
// The certificate is used to secure the domain with TLS/SSL.  If no
// certificate is specified, Unikraft Cloud will automatically generate a
// new certificate for the domain based on Let's Encrypt and seek to
// accomplish a DNS-01 challenge.

type InstanceDomainInstanceDomainCertificate struct {
	// The UUID of the certificate.
	//
	// This is a unique identifier for the certificate that is generated
	// when the certificate is created.  The UUID is used to reference the
	// certificate in API calls and can be used to identify the certificate
	// in all API calls that require a certificate identifier.
	Uuid string `json:"uuid"`
	// The name of the certificate.
	//
	// This is a human-readable name that can be used to identify the
	// certificate.  The name is unique within the context of your account.
	// The name can also be used to identify the certificate in API calls.
	Name string `json:"name"`
}

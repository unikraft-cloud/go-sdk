// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// A reference to an existing certificate which can be used for the
// specified domain.  If unspecified, Unikraft Cloud will
// automatically generate a new certificate for the domain based on Let's
// Encrypt and seek to accomplish a DNS-01 challenge.

type CreateInstanceRequestCertificate struct {
	// (Optional).  The name of the existing certificate.  Mutually
	// exclusive with UUID.
	Name string `json:"name"`
	// (Optional).  The UUID for the existing certificate.  Mutually
	// exclusive with name.
	Uuid string `json:"uuid"`
}

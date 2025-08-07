// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Use an existing certificate for the domain.  If this field is
// specified, the domain must be associated with a valid certificate.

type CreateServiceGroupRequestDomainCertificate struct {
	// The UUID of the certificate to use for the service group.
	Uuid string `json:"uuid"`
	// The name of the certificate to use for the service group.
	Name string `json:"name"`
}

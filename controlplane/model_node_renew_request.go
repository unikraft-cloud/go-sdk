// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The request message for certificate renewal.

type NodeRenewRequest struct {
	// The csr is the certificate signing request (CSR) for the new certificate
	// in base64 URL encoded PEM format.
	//
	// The Organization(O) should be the organization UUID.
	// The OrganizationalUnit(OU) should have one item which is the machineID.
	// The CommonName(CN) should be the DNS hostname of the node.
	Csr *string `json:"csr,omitempty"`
	// The serial number of the certificate that ought to be renewed.
	PreviousSerial *string `json:"previous_serial,omitempty"`
}

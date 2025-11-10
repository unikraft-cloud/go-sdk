// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The request message for certificate activation.

type NodeActivateRequest struct {
	// The csr is the certificate signing request (CSR) for the new certificate
	// in base64 URL encoded PEM format.
	//
	// The Organization(O) should be the organization UUID.
	// The OrganizationalUnit(OU) should have one item which is the machineID.
	// The CommonName(CN) should be the DNS hostname of the node.
	Csr *string `json:"csr,omitempty"`
	// The hostname is the name of the machine making the request.
	Hostname *string `json:"hostname,omitempty"`
	// The operating system of the machine making the request.
	Os *string `json:"os,omitempty"`
	// The version of the operating system of the machine making the request, if
	// available.
	//
	// For Android, it's like "10", "11", "12", etc.  For iOS and macOS it's like
	// "15.6.1" or "12.4.0".  For Windows it's like "10.0.19044.1889". For FreeBSD
	// it's like "12.3-STABLE".  For Linux, this is simply the kernel version on
	// Linux, like "5.10.0-17-amd64".
	OsVersion *string `json:"os_version,omitempty"`
	// The OS distribution, if known.  E.g. "debian", "ubuntu", "nixos", ...
	Distro *string `json:"distro,omitempty"`
	// The OS distribution version if known.  E.g. "20.04", ...
	DistroVersion *string `json:"distro_version,omitempty"`
	// The OS distribution codename if known.  E.g. "jammy", "bullseye", ...
	DistroCodename *string `json:"distro_codename,omitempty"`
	// The features supported by the kernel.
	KernelFeatures []string `json:"kernel_features,omitempty"`
	// The region of the machine making the request.
	RegionHint *string `json:"region_hint,omitempty"`
}

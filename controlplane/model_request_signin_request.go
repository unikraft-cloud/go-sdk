// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// The request message for a sign-in request.

type RequestSigninRequest struct {
	// The hostname is the name of the machine making the request.  This is
	// mandatory as it consitutes a unique identifier for the machine.
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
	// A best-effort whether the client is running in a container.
	Container *bool `json:"container,omitempty"`
	// The OS distribution, if known.  E.g. "debian", "ubuntu", "nixos", ...
	Distro *string `json:"distro,omitempty"`
	// The OS distribution version if known.  E.g. "20.04", ...
	DistroVersion *string `json:"distro_version,omitempty"`
	// TThe OS distribution codename if known.  E.g. "jammy", "bullseye", ...
	DistroCodename *string `json:"distro_codename,omitempty"`
	// The CLI version is the version of the Unikraft CLI that is making the
	// request.
	CliVersion *string `json:"cli_version,omitempty"`
	// If available, the GOARCH value (of the built binary).
	Goarch *string `json:"goarch,omitempty"`
	// If available, the GOOS value (of the built binary)
	Goos *string `json:"goos,omitempty"`
	// if available, the Go version binary was built with.
	GoVersion *string `json:"go_version,omitempty"`
}

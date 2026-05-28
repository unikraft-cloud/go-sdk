// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import "encoding/json"

// The request message for a sign-in request.

type RequestSigninRequest struct {
	// The hostname is the name of the machine making the request.  This is
	// mandatory as it consitutes a unique identifier for the machine.
	Hostname string `json:"hostname"`
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

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *RequestSigninRequest) UnmarshalJSON(data []byte) error {
	type Alias RequestSigninRequest
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"hostname":        {},
		"os":              {},
		"os_version":      {},
		"container":       {},
		"distro":          {},
		"distro_version":  {},
		"distro_codename": {},
		"cli_version":     {},
		"goarch":          {},
		"goos":            {},
		"go_version":      {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m RequestSigninRequest) MarshalJSON() ([]byte, error) {
	type Alias RequestSigninRequest
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

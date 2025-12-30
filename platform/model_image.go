// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

type Image struct {
	// The digest of the image is a unique identifier of the image manifest which
	// is a string representation including the hashing algorithm and the hash
	// value separated by a colon.
	Digest *string `json:"digest,omitempty"`
	// The canonical name of the image is known as the "tag".
	Tags        []string          `json:"tags,omitempty"`
	Initrd      *bool             `json:"initrd,omitempty"`
	SizeInBytes *int64            `json:"size_in_bytes,omitempty"`
	Args        *string           `json:"args,omitempty"`
	KernelArgs  *string           `json:"kernel_args,omitempty"`
	Users       []string          `json:"users,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

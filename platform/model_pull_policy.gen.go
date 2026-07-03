// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// PullPolicy defines when an image should be pulled.
type PullPolicy string

const (
	PullPolicyAlways       PullPolicy = "always"
	PullPolicyIfNotPresent PullPolicy = "if_not_present"
	PullPolicyNever        PullPolicy = "never"
)

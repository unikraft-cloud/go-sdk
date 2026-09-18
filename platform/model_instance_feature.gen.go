// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// Features are specific configurations or capabilities that can be enabled for
// the instance.
//
// The list of available features to enable for the instance:
//
// | Feature            | Description |
// |--------------------|-------------|
// | `delete-on-stop`   | The instance will be deleted when it is stopped. This is useful for instances that are not needed after they are stopped, such as temporary or ephemeral instances. Cannot be combined with a `restart_policy` other than `never`. |
// | `nested-virt`      | Expose virtualization extensions to the guest, so that it can run virtual machines of its own. Requires the `nested_virt` permission. |
type InstanceFeature string

const (
	InstanceFeatureDeleteOnStop InstanceFeature = "delete-on-stop"
	InstanceFeatureNestedVirt   InstanceFeature = "nested-virt"
)

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

// NodeState defines the provisioning lifecycle state of a machine which will
// run the Unikraft Cloud Platform.
type NodeState string

const (
	NodeStateUnspecified    NodeState = "unspecified"
	NodeStatePending        NodeState = "pending"
	NodeStateProvisioning   NodeState = "provisioning"
	NodeStateConfiguring    NodeState = "configuring"
	NodeStateReady          NodeState = "ready"
	NodeStateDeprovisioning NodeState = "deprovisioning"
	NodeStateDeprovisioned  NodeState = "deprovisioned"
	NodeStateError          NodeState = "error"
	NodeStateMaintenance    NodeState = "maintenance"
)

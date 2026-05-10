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
	NodeStateUNSPECIFIED    NodeState = "unspecified"
	NodeStatePENDING        NodeState = "pending"
	NodeStatePROVISIONING   NodeState = "provisioning"
	NodeStateCONFIGURING    NodeState = "configuring"
	NodeStateREADY          NodeState = "ready"
	NodeStateDEPROVISIONING NodeState = "deprovisioning"
	NodeStateDEPROVISIONED  NodeState = "deprovisioned"
	NodeStateERROR          NodeState = "error"
	NodeStateMAINTENANCE    NodeState = "maintenance"
)

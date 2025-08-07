// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The status of the response.
type GetAutoscaleConfigurationPolicyResponsePolicyResponseStatus string

const (
	GetAutoscaleConfigurationPolicyResponsePolicyResponseStatusSuccess GetAutoscaleConfigurationPolicyResponsePolicyResponseStatus = "success"
	GetAutoscaleConfigurationPolicyResponsePolicyResponseStatusError   GetAutoscaleConfigurationPolicyResponsePolicyResponseStatus = "error"
)

type GetAutoscaleConfigurationPolicyResponsePolicyResponse struct {
	// The status of the response.
	Status *GetAutoscaleConfigurationPolicyResponsePolicyResponseStatus `json:"status,omitempty"`
	Policy *GetAutoscaleConfigurationPolicyResponsePolicyResponsePolicy `json:"policy,omitempty"`
}

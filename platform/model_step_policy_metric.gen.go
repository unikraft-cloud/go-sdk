// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The step policy is a type of autoscaling policy that scales the number of
// instances in a service by a fixed number of instances at each step.
// It uses a metric to determine when to scale up or down.
type StepPolicyMetric string

const (
	StepPolicyMetricCPU StepPolicyMetric = "cpu"
)

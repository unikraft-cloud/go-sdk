// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// AdjustmentType defines the type of adjustment to be made in an autoscaling
// step policy.
type AdjustmentType string

const (
	AdjustmentTypeChange     AdjustmentType = "change"
	AdjustmentTypeExact      AdjustmentType = "exact"
	AdjustmentTypePercentage AdjustmentType = "percentage"
)

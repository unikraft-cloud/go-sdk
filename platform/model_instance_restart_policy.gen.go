// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

// The restart policy of an instance.
//
// When an instance stops either because the application exits or the instance
// crashes, Unikraft Cloud can auto-restart your instance.  Auto-restarts are
// performed according to the restart policy configured for a particular
// instance.
//
// The policy can have the following values:
//
// | Policy       | Description |
// |--------------|-------------|
// | `never`      | Never restart the instance (default). |
// | `always`     | Always restart the instance when the stop is initiated from within the instance (i.e., the application exits or the instance crashes). |
// | `on-failure` | Only restart the instance if it crashes. |
//
// When an instance stops, the stop reason and the configured restart policy are
// evaluated to decide if a restart should be performed.  Unikraft Cloud uses an
// exponential back-off delay (immediate, 5s, 10s, 20s, 40s, ..., 5m) to slow
// down restarts in tight crash loops. If an instance runs without problems for
// 10s the back-off delay is reset and the restart sequence ends.
//
// The `restart.attempt` attribute reported in counts the number of restarts
// performed in the current sequence.  The `restart.next_at` field indicates
// when the next restart will take place if a back-off delay is in effect.
//
// A manual start or stop of the instance aborts the restart sequence and resets
// the back-off delay.
type InstanceRestartPolicy string

const (
	InstanceRestartPolicyNever      InstanceRestartPolicy = "never"
	InstanceRestartPolicyAlways     InstanceRestartPolicy = "always"
	InstanceRestartPolicyOn_failure InstanceRestartPolicy = "on-failure"
)

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"time"
)

var _ time.Time

// A schedule defines when an action should be performed on an instance.
//
// Each schedule specifies a name, a calendar expression following systemd
// calendar event syntax, and an action (start, stop, delete, or exec).
//
// Calendar expressions format: [weekday] [[year-]month-day] [hour:minute[:second]]
//
// Supported syntax:
// - `*` - Any value
// - `5` - Exact value
// - `1..5` - Range
// - `1..5/2` - Range with step
// - `1,2,5` - Comma-separated list
//
// Example: `*-*-* 09:00:00` - Every day at 09:00 UTC
// Example: `Sat,Sun *-*-* 20:00:00` - Every Saturday and Sunday at 20:00 UTC
type Schedule struct {
	// The name of the schedule.
	//
	// Must be unique within an instance.
	Name string `json:"name"`
	// The calendar expression specifying when the action should be performed.
	//
	// Uses systemd calendar event syntax.
	// See https://www.man7.org/linux/man-pages/man7/systemd.time.7.html
	When string `json:"when"`
	// The action to perform at the scheduled time.
	Action ScheduleAction `json:"action"`
	// The timestamp of when the next scheduled action will occur.
	//
	// This field is populated only in responses (not settable in requests).
	// Unix timestamp in seconds.  Omitted if no next execution is scheduled.
	NextAt *time.Time `json:"next_at,omitzero"`
	// The command to execute when the action is `exec`.
	//
	// Required when `action` is `SCHEDULE_ACTION_EXEC`, ignored otherwise.
	// Each element is a separate argument; the first element is the executable.
	Args []string `json:"args,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Schedule) UnmarshalJSON(data []byte) error {
	type Alias Schedule
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Schedule) MarshalJSON() ([]byte, error) {
	type Alias Schedule
	return json.Marshal((Alias)(m))
}

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "encoding/json"

// A schedule defines when an action should be performed on an instance.
//
// Each schedule specifies a name, a calendar expression following systemd
// calendar event syntax, and an action (start, stop, or delete).
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
// The action to perform at the scheduled time.
type ScheduleAction string

const (
	ScheduleActionStart  ScheduleAction = "start"
	ScheduleActionStop   ScheduleAction = "stop"
	ScheduleActionDelete ScheduleAction = "delete"
)

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
	NextAt *int64 `json:"next_at,omitempty"`

	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

func (m *Schedule) UnmarshalJSON(data []byte) error {
	type Alias Schedule
	if err := json.Unmarshal(data, (*Alias)(m)); err != nil {
		return err
	}

	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"name":    {},
		"when":    {},
		"action":  {},
		"next_at": {},
	}
	for key := range knownKeys {
		delete(extra, key)
	}
	if len(extra) == 0 {
		m.AdditionalProperties = nil
		return nil
	}
	m.AdditionalProperties = extra
	return nil
}

func (m Schedule) MarshalJSON() ([]byte, error) {
	type Alias Schedule
	base, err := json.Marshal((*Alias)(&m))
	if err != nil {
		return nil, err
	}
	if len(m.AdditionalProperties) == 0 {
		return base, nil
	}

	var out map[string]json.RawMessage
	if err := json.Unmarshal(base, &out); err != nil {
		return nil, err
	}
	for key, value := range m.AdditionalProperties {
		if _, exists := out[key]; !exists {
			out[key] = value
		}
	}
	return json.Marshal(out)
}

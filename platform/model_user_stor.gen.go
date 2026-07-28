// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type UserStor struct {
	// Maximum number of volumes the user can have at one moment.
	MaxVolumes *int32 `json:"max_volumes,omitzero"`
	// Minimum size of a volume in MB.
	MinVolumeMb *int32 `json:"min_volume_mb,omitzero"`
	// Maximum size of a volume in MB.
	MaxVolumeMb *int32 `json:"max_volume_mb,omitzero"`
	// Maximum total size of all volumes in MB.
	MaxTotalVolumeMb *int32 `json:"max_total_volume_mb,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *UserStor) UnmarshalJSON(data []byte) error {
	type Alias UserStor
	return json.Unmarshal(data, (*Alias)(m))
}

func (m UserStor) MarshalJSON() ([]byte, error) {
	type Alias UserStor
	return json.Marshal((Alias)(m))
}

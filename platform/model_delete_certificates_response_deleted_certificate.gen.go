// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"time"
)

var _ time.Time

// Details of the certificate which was deleted by this request.
type DeleteCertificatesResponseDeletedCertificate struct {
	// Indicates whether the delete operation was successful or not for this
	// certificate.
	Status ResponseStatus `json:"status"`
	// The UUID of the certificate which was deleted.
	Uuid string `json:"uuid"`
	// The name of the certificate which was deleted.
	Name string `json:"name"`
	// An optional message providing additional information about the status.
	// This field is useful when the status is not `success`.
	Message *string `json:"message,omitzero"`
	// An optional error code providing additional information about the status.
	// This field is useful when the status is not `success`.
	Error *int32 `json:"error,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *DeleteCertificatesResponseDeletedCertificate) UnmarshalJSON(data []byte) error {
	type Alias DeleteCertificatesResponseDeletedCertificate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m DeleteCertificatesResponseDeletedCertificate) MarshalJSON() ([]byte, error) {
	type Alias DeleteCertificatesResponseDeletedCertificate
	return json.Marshal((Alias)(m))
}

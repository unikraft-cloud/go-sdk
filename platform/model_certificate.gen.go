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

// Certificate with per-item response envelope fields merged in.
type Certificate struct {
	// Indicates whether the operation was successful for this item.
	Status *ResponseStatus `json:"status,omitzero"`
	// An optional message providing additional information.
	Message *string `json:"message,omitzero"`
	// An optional error code.
	Error *int32 `json:"error,omitzero"`
	// The UUID of the resource.
	Uuid string `json:"uuid"`
	// The human-readable name of the resource.
	Name string `json:"name"`
	// The time the certificate was created.
	CreatedAt time.Time `json:"created_at"`
	// The common name (CN) field from the certificate's subject.
	//
	// This is typically the primary domain name that the certificate is issued
	// for. It represents the main identity that the certificate validates.
	CommonName string `json:"common_name"`
	// The complete subject distinguished name (DN) of the certificate.
	//
	// This contains the full subject information from the certificate, including
	// the common name, organization, organizational unit, locality, state, and
	// country. The subject identifies the entity that the certificate is issued to.
	// Only present when the certificate has been issued.
	Subject *string `json:"subject,omitzero"`
	// The complete issuer distinguished name (DN) of the certificate.
	//
	// This identifies the Certificate Authority (CA) that issued the certificate.
	// It contains information about the CA including its common name, organization,
	// and country.
	// Only present when the certificate has been issued.
	Issuer *string `json:"issuer,omitzero"`
	// The unique serial number assigned to the certificate by the issuing CA.
	//
	// This is a unique identifier within the scope of the issuing CA that can be
	// used to identify and track the certificate. Serial numbers are typically
	// represented as hexadecimal strings.
	// Only present when the certificate has been issued.
	SerialNumber *string `json:"serial_number,omitzero"`
	// The date and time when the certificate becomes valid.
	//
	// The certificate should not be trusted before this date. This timestamp
	// marks the beginning of the certificate's validity period.
	// Only present when the certificate has been issued.
	NotBefore *time.Time `json:"not_before,omitzero"`
	// The date and time when the certificate expires.
	//
	// The certificate should not be trusted after this date. This timestamp
	// marks the end of the certificate's validity period. Certificates should
	// be renewed before this date to maintain service availability.
	// Only present when the certificate has been issued.
	NotAfter *time.Time `json:"not_after,omitzero"`
	// The current state of the certificate.
	//
	// This indicates whether the certificate is pending issuance, valid and
	// ready for use, or in an error state. See CertificateState enum for
	// detailed state descriptions.
	State CertificateState `json:"state"`
	// Validation status when state is pending.
	Validation *CertificateValidation `json:"validation,omitzero"`
	// Service groups using this certificate.
	ServiceGroups []ID `json:"service_groups,omitzero"`

	// AdditionalProperties captures any JSON object members that do not map to
	// an explicit field above.
	AdditionalProperties map[string]jsontext.Value `json:",embed"`
}

func (m *Certificate) UnmarshalJSON(data []byte) error {
	type Alias Certificate
	return json.Unmarshal(data, (*Alias)(m))
}

func (m Certificate) MarshalJSON() ([]byte, error) {
	type Alias Certificate
	return json.Marshal((Alias)(m))
}

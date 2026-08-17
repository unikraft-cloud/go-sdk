// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"unikraft.com/cloud/sdk/platform/gen"
)

// CreateCertificateOption configures a call to [Client.CreateCertificate].
type CreateCertificateOption func(*gen.CreateCertificateRequest)

// WithCreateCertificateName sets the certificate's name. If not set, a
// random name is generated.
func WithCreateCertificateName(name string) CreateCertificateOption {
	return func(r *gen.CreateCertificateRequest) { r.Name = new(name) }
}

// WithCreateCertificateCommonName sets the fully-qualified domain name the
// certificate is issued for.
func WithCreateCertificateCommonName(cn string) CreateCertificateOption {
	return func(r *gen.CreateCertificateRequest) { r.CommonName = new(cn) }
}

// WithCreateCertificateChain sets the certificate chain, in PEM format,
// including the certificate and any intermediates.
func WithCreateCertificateChain(chain string) CreateCertificateOption {
	return func(r *gen.CreateCertificateRequest) { r.Chain = chain }
}

// WithCreateCertificatePKey sets the private key, in PEM format, matching
// the public key in the certificate chain.
func WithCreateCertificatePKey(pkey string) CreateCertificateOption {
	return func(r *gen.CreateCertificateRequest) { r.Pkey = pkey }
}

// CertificatesOption configures a call to [Client.Certificates].
type CertificatesOption func(*gen.GetCertificatesOpts)

// WithCertificateDetails includes full details for each returned certificate.
func WithCertificateDetails() CertificatesOption {
	return func(o *gen.GetCertificatesOpts) { o.Details = new(true) }
}

// WithCertificateLimit caps the number of certificates returned.
func WithCertificateLimit(n uint32) CertificatesOption {
	return func(o *gen.GetCertificatesOpts) { o.Count = new(n) }
}

// WithCertificatePage continues listing from the pagination cursor returned
// by a previous call.
func WithCertificatePage(from string) CertificatesOption {
	return func(o *gen.GetCertificatesOpts) { o.From = new(from) }
}

// WithCertificateOrder sets the sort order for the returned certificates.
func WithCertificateOrder(order PaginationOrder) CertificatesOption {
	return func(o *gen.GetCertificatesOpts) { o.Order = new(gen.PaginationOrder(order)) }
}

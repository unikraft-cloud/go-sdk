// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"encoding/json"
	"fmt"

	"unikraft.com/cloud/sdk/platform/gen"
)

// Certificate is a live, resource-oriented view of a certificate. It embeds
// gen.Certificate so its fields are accessed directly (cert.CommonName,
// cert.State, ...), and adds behavior (Refresh, Update, Delete) on top of
// the flat generated client.
//
// Unlike Instance and Node, Certificate has no Update-by-diff: the chain and
// private key are write-only and never come back from the API, so there is
// nothing on the struct to compare a new value against.
type Certificate struct {
	gen.Certificate

	client gen.Client
}

// cloneCertificate deep-clones a gen.Certificate via a JSON round-trip, so
// that pointer fields aren't aliased between a Certificate and its
// snapshot. in was itself just unmarshaled from the API response, so a
// marshal/unmarshal round-trip of it failing indicates a broken invariant,
// not a normal error condition.
func cloneCertificate(in gen.Certificate) gen.Certificate {
	b, err := json.Marshal(in)
	if err != nil {
		panic(fmt.Sprintf("platform: cloning certificate: %v", err))
	}
	var out gen.Certificate
	if err := json.Unmarshal(b, &out); err != nil {
		panic(fmt.Sprintf("platform: cloning certificate: %v", err))
	}
	return out
}

func wrapCertificate(client gen.Client, data gen.Certificate) *Certificate {
	return &Certificate{Certificate: cloneCertificate(data), client: client}
}

// Refresh re-fetches the certificate from the API.
func (c *Certificate) Refresh(ctx context.Context) error {
	resp, err := c.client.GetCertificates(ctx, []gen.NameOrUUID{{Uuid: &c.Uuid}}, gen.GetCertificatesOpts{Details: new(true)})
	if err != nil {
		return err
	}
	if len(resp.Data.Certificates) == 0 {
		return fmt.Errorf("certificate %s not found", c.Uuid)
	}
	c.Certificate = cloneCertificate(resp.Data.Certificates[0])
	return nil
}

// Delete deletes the certificate. Unlike Update, it does not refresh
// afterwards: the certificate is gone, and re-fetching it would just 404.
func (c *Certificate) Delete(ctx context.Context) error {
	_, err := c.client.DeleteCertificateByUUID(ctx, c.Uuid)
	return err
}

// Update replaces the certificate's chain and private key, then refreshes it
// to pick up the resulting state.
func (c *Certificate) Update(ctx context.Context, chain, pkey string) error {
	_, err := c.client.UpdateCertificateByUUID(ctx, c.Uuid, gen.UpdateCertificateByUUIDRequestBody{
		Chain: chain,
		Pkey:  pkey,
	})
	if err != nil {
		return err
	}
	return c.Refresh(ctx)
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"fmt"

	"unikraft.com/cloud/sdk/platform/gen"
)

// Quotas reports the current user's resource usage and limits.
//
// Unlike Instance, ServiceGroup, Certificate, and Volume, Quotas is returned
// as the plain gen type rather than wrapped: it is read-only reporting data
// with no lifecycle of its own (no refresh/update/delete), so a wrapper
// would add fields and no behavior.
func (c *platformClient) Quotas(ctx context.Context) (*gen.Quotas, error) {
	resp, err := c.raw.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return firstQuota(resp.Data.Quotas)
}

// QuotasForUser reports resource usage and limits for the user identified
// by uuid.
func (c *platformClient) QuotasForUser(ctx context.Context, uuid string) (*gen.Quotas, error) {
	resp, err := c.raw.GetUserByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return firstQuota(resp.Data.Quotas)
}

func firstQuota(quotas []gen.Quotas) (*gen.Quotas, error) {
	if len(quotas) == 0 {
		return nil, fmt.Errorf("no quotas returned")
	}
	return &quotas[0], nil
}

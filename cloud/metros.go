// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/cloud/sdk/platform/group"
)

// Metro is a metro that the account can reach.
type Metro struct {
	// Code is the lower-case identity of the metro, such as "fra".
	Code string
	// Endpoint is the base URL of the platform API in the metro.
	Endpoint string
}

// AvailableMetros returns the metros that the account can reach. It asks the
// control plane once, and keeps a successful answer for the life of c.
func (c *Client) AvailableMetros(ctx context.Context) ([]Metro, error) {
	c.metrosMu.Lock()
	defer c.metrosMu.Unlock()
	if c.metros != nil {
		return slices.Clone(c.metros), nil
	}

	resp, err := c.ControlPlane().Raw().ListMetros(ctx)
	if err != nil {
		return nil, fmt.Errorf("discovering the available metros: %w", err)
	}
	var metros []Metro
	for _, m := range resp.Data.Metros {
		code := strings.ToLower(m.IataCode)
		if code == "" {
			code = strings.ToLower(m.Name)
		}
		if code == "" {
			continue
		}
		metros = append(metros, Metro{Code: code, Endpoint: metroBaseURL(m.Endpoint, code)})
	}
	if len(metros) == 0 {
		return nil, errors.New("the control plane reported no metros for this account")
	}
	c.metros = metros
	return slices.Clone(metros), nil
}

// metroBaseURL uses endpoint when it is a URL, and builds the URL from code
// when it is not.
func metroBaseURL(endpoint, code string) string {
	if !strings.Contains(endpoint, "://") {
		return "https://api." + code + ".unikraft.cloud"
	}
	endpoint = strings.TrimRight(endpoint, "/")
	return strings.TrimSuffix(endpoint, "/v1")
}

// Metros returns a group with one platform client for each metro in codes,
// for use with the functions of the group package. With no codes, the group
// has every metro that AvailableMetros returns.
func (c *Client) Metros(ctx context.Context, codes ...string) (*group.Group[platform.Client], error) {
	g := group.New[platform.Client]()
	if len(codes) > 0 {
		for _, code := range codes {
			g = g.WithClient(code, c.Platform(code))
		}
		return g, nil
	}

	metros, err := c.AvailableMetros(ctx)
	if err != nil {
		return nil, err
	}
	for _, m := range metros {
		g = g.WithClient(m.Code, c.Platform(m.Endpoint))
	}
	return g, nil
}

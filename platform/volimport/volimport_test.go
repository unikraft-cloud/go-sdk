// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volimport

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"unikraft.com/cloud/sdk/platform/gen"
)

// fakeClient gives instance on create, and records the instances that it
// deletes.
type fakeClient struct {
	gen.Client
	instance gen.Instance
	created  gen.CreateInstanceRequest
	deleted  []string
}

func (c *fakeClient) CreateInstance(_ context.Context, req gen.CreateInstanceRequest) (*gen.Response[gen.CreateInstanceResponseData], error) {
	c.created = req
	return &gen.Response[gen.CreateInstanceResponseData]{
		Data: &gen.CreateInstanceResponseData{Instances: []gen.Instance{c.instance}},
	}, nil
}

func (c *fakeClient) DeleteInstanceByUUID(_ context.Context, uuid string, _ gen.DeleteInstanceByUUIDRequestBody) (*gen.Response[gen.DeleteInstancesResponseData], error) {
	c.deleted = append(c.deleted, uuid)
	return &gen.Response[gen.DeleteInstancesResponseData]{}, nil
}

func TestStart(t *testing.T) {
	c := &fakeClient{instance: gen.Instance{
		Uuid: "inst-uuid",
		ServiceGroup: &gen.InstanceServiceGroup{
			Domains: []gen.ServiceGroupInstanceDomain{{Fqdn: "import.example.com."}},
		},
	}}

	uuid, fqdn, err := Start(t.Context(), c, DefaultImage, "vol-uuid", testAuth, 10, DefaultPort)
	require.NoError(t, err)
	require.Equal(t, "inst-uuid", uuid)
	require.Equal(t, "import.example.com", fqdn)
	require.Empty(t, c.deleted)

	req := c.created
	require.Equal(t, &gen.ImageSpec{Url: DefaultImage}, req.Image)
	require.Equal(t, []string{"volimport", "-p", "42069", "-a", testAuth, "-t", "10"}, req.Args)
	require.Equal(t, []gen.InstanceFeature{gen.InstanceFeatureDeleteOnStop}, req.Features)
	require.Len(t, req.Volumes, 1)
	require.Equal(t, "vol-uuid", *req.Volumes[0].Uuid)
	require.Equal(t, "/mnt", req.Volumes[0].At)
	require.Len(t, req.ServiceGroup.Services, 1)
	svc := req.ServiceGroup.Services[0]
	require.Equal(t, uint32(DefaultPort), svc.Port)
	require.Equal(t, internalPort, *svc.DestinationPort)
	require.Equal(t, []gen.ConnectionHandler{gen.ConnectionHandlerTls}, svc.Handlers)
}

func TestStartNoDomain(t *testing.T) {
	c := &fakeClient{instance: gen.Instance{Uuid: "inst-uuid"}}

	_, _, err := Start(t.Context(), c, DefaultImage, "vol-uuid", testAuth, 10, DefaultPort)
	require.EqualError(t, err, "import instance has no service group domain")
	require.Equal(t, []string{"inst-uuid"}, c.deleted)
}

func TestGenRandAuth(t *testing.T) {
	a, err := GenRandAuth()
	require.NoError(t, err)
	require.Regexp(t, `^[0-9A-Za-z]{32}$`, a)

	b, err := GenRandAuth()
	require.NoError(t, err)
	require.NotEqual(t, a, b)
}

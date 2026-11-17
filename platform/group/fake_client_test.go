// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package group

import (
	"context"
	"iter"
	"time"

	"unikraft.com/cloud/sdk/pkg/httpclient"
	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/cloud/sdk/platform/gen"
)

// fakeClient is a minimal, comparable stand-in for platform.Client, so group
// package tests can exercise fan-out/collection logic without any HTTP
// transport. Only WithMetro/WithHTTPClient/WithTimeout are ever exercised by
// the group package itself; the rest exist solely to satisfy the interface.
type fakeClient struct {
	id string
}

var _ platform.Client = fakeClient{}

func (f fakeClient) Raw() gen.Client { return nil }

func (f fakeClient) Instance(ctx context.Context, ref platform.Ref) (*platform.Instance, error) {
	return nil, nil
}

func (f fakeClient) Instances(ctx context.Context, refs []platform.Ref, opts ...platform.InstancesOption) ([]*platform.Instance, error) {
	return nil, nil
}

func (f fakeClient) CreateInstance(ctx context.Context, opts ...platform.CreateInstanceOption) ([]*platform.Instance, error) {
	return nil, nil
}

func (f fakeClient) ServiceGroup(ctx context.Context, ref platform.Ref) (*platform.ServiceGroup, error) {
	return nil, nil
}

func (f fakeClient) ServiceGroups(ctx context.Context, refs []platform.Ref, opts ...platform.ServiceGroupsOption) ([]*platform.ServiceGroup, error) {
	return nil, nil
}

func (f fakeClient) CreateServiceGroup(ctx context.Context, opts ...platform.CreateServiceGroupOption) (*platform.ServiceGroup, error) {
	return nil, nil
}

func (f fakeClient) Certificate(ctx context.Context, ref platform.Ref) (*platform.Certificate, error) {
	return nil, nil
}

func (f fakeClient) Certificates(ctx context.Context, refs []platform.Ref, opts ...platform.CertificatesOption) ([]*platform.Certificate, error) {
	return nil, nil
}

func (f fakeClient) CreateCertificate(ctx context.Context, opts ...platform.CreateCertificateOption) (*platform.Certificate, error) {
	return nil, nil
}

func (f fakeClient) Volume(ctx context.Context, ref platform.Ref) (*platform.Volume, error) {
	return nil, nil
}

func (f fakeClient) Volumes(ctx context.Context, refs []platform.Ref, opts ...platform.VolumesOption) ([]*platform.Volume, error) {
	return nil, nil
}

func (f fakeClient) CreateVolume(ctx context.Context, opts ...platform.CreateVolumeOption) (*platform.Volume, error) {
	return nil, nil
}

func (f fakeClient) Quotas(ctx context.Context) (*gen.Quotas, error) {
	return nil, nil
}

func (f fakeClient) QuotasForUser(ctx context.Context, uuid string) (*gen.Quotas, error) {
	return nil, nil
}

func (f fakeClient) WithMetro(metro string) platform.Client { return fakeClient{id: metro} }

func (f fakeClient) WithHTTPClient(hc httpclient.HTTPClient) platform.Client { return f }

func (f fakeClient) WithTimeout(to time.Duration) platform.Client { return f }

func (f fakeClient) Template(ctx context.Context, ref platform.Ref) (*platform.Template, error) {
	return nil, nil
}

func (f fakeClient) Templates(ctx context.Context, refs []platform.Ref, opts ...platform.TemplatesOption) ([]*platform.Template, error) {
	return nil, nil
}

func (f fakeClient) PrepareTemplate(ctx context.Context, name string, source gen.Instance, opts ...platform.PrepareTemplateOption) (*platform.Template, error) {
	return nil, nil
}

func (f fakeClient) Images(ctx context.Context, opts ...platform.ImagesOption) ([]gen.Image, error) {
	return nil, nil
}

func (f fakeClient) StartInstances(ctx context.Context, refs []platform.Ref, opts ...platform.StartOption) error {
	return nil
}

func (f fakeClient) StopInstances(ctx context.Context, refs []platform.Ref, opts ...platform.StopOption) error {
	return nil
}

func (f fakeClient) SuspendInstances(ctx context.Context, refs []platform.Ref, opts ...platform.SuspendOption) error {
	return nil
}

func (f fakeClient) DeleteInstances(ctx context.Context, refs []platform.Ref, opts ...platform.DeleteOption) error {
	return nil
}

func (f fakeClient) DeleteVolumes(ctx context.Context, refs []platform.Ref) error {
	return nil
}

func (f fakeClient) DeleteServiceGroups(ctx context.Context, refs []platform.Ref) error {
	return nil
}

func (f fakeClient) DeleteCertificates(ctx context.Context, refs []platform.Ref) error {
	return nil
}

func (f fakeClient) DeleteTemplates(ctx context.Context, refs []platform.Ref) error {
	return nil
}

func (f fakeClient) ListInstances(ctx context.Context, opts ...platform.ListOption) iter.Seq2[*platform.Instance, error] {
	return nil
}

func (f fakeClient) ListVolumes(ctx context.Context, opts ...platform.ListOption) iter.Seq2[*platform.Volume, error] {
	return nil
}

func (f fakeClient) ListServiceGroups(ctx context.Context, opts ...platform.ListOption) iter.Seq2[*platform.ServiceGroup, error] {
	return nil
}

func (f fakeClient) ListCertificates(ctx context.Context, opts ...platform.ListOption) iter.Seq2[*platform.Certificate, error] {
	return nil
}

func (f fakeClient) ListTemplates(ctx context.Context, opts ...platform.ListOption) iter.Seq2[*platform.Template, error] {
	return nil
}

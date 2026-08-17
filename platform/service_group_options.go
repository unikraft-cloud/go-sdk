// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"unikraft.com/cloud/sdk/platform/gen"
)

// CreateServiceGroupOption configures a call to [Client.CreateServiceGroup].
type CreateServiceGroupOption func(*gen.CreateServiceGroupRequest)

// WithCreateServiceGroupName sets the service group's name. If not set, a
// random name is generated.
func WithCreateServiceGroupName(name string) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.Name = new(name) }
}

// WithCreateServiceGroupServices sets the published ports and their handlers.
func WithCreateServiceGroupServices(services ...gen.Service) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.Services = services }
}

// WithCreateServiceGroupDomains associates domains with the service group.
func WithCreateServiceGroupDomains(domains ...gen.CreateServiceGroupRequestDomain) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.Domains = domains }
}

// WithCreateServiceGroupSoftLimit sets the concurrent-request count at which
// the load balancer wakes up another standby instance.
func WithCreateServiceGroupSoftLimit(n uint64) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.SoftLimit = new(n) }
}

// WithCreateServiceGroupHardLimit sets the maximum number of concurrent
// requests an assigned instance can handle.
func WithCreateServiceGroupHardLimit(n uint64) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.HardLimit = new(n) }
}

// WithCreateServiceGroupAutokill configures automatic delete-on-idle
// behavior for the service group.
func WithCreateServiceGroupAutokill(cfg gen.CreateServiceGroupRequestAutokill) CreateServiceGroupOption {
	return func(r *gen.CreateServiceGroupRequest) { r.Autokill = &cfg }
}

// ServiceGroupsOption configures a call to [Client.ServiceGroups].
type ServiceGroupsOption func(*gen.GetServiceGroupsOpts)

// WithServiceGroupDetails includes full details for each returned service group.
func WithServiceGroupDetails() ServiceGroupsOption {
	return func(o *gen.GetServiceGroupsOpts) { o.Details = new(true) }
}

// WithServiceGroupLimit caps the number of service groups returned.
func WithServiceGroupLimit(n uint32) ServiceGroupsOption {
	return func(o *gen.GetServiceGroupsOpts) { o.Count = new(n) }
}

// WithServiceGroupPage continues listing from the pagination cursor returned
// by a previous call.
func WithServiceGroupPage(from string) ServiceGroupsOption {
	return func(o *gen.GetServiceGroupsOpts) { o.From = new(from) }
}

// WithServiceGroupOrder sets the sort order for the returned service groups.
func WithServiceGroupOrder(order PaginationOrder) ServiceGroupsOption {
	return func(o *gen.GetServiceGroupsOpts) { o.Order = new(gen.PaginationOrder(order)) }
}

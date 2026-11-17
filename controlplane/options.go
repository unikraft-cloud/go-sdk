// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"unikraft.com/cloud/sdk/controlplane/gen"
)

// ProvisionNodeOption configures a call to [Client.ProvisionNode].
type ProvisionNodeOption func(*gen.ProvisionNodeRequest)

// WithProvisionName sets the node's name. If not set, a name is auto-generated.
func WithProvisionName(name string) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.Name = new(name) }
}

// WithProvisionCloudProvider sets the cloud provider where the node should
// be provisioned.
func WithProvisionCloudProvider(cp gen.CloudProvider) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.Cloudprovider = &cp }
}

// WithProvisionSSHKeys sets the SSH keys for accessing the node. At least
// one key is required.
func WithProvisionSSHKeys(keys ...gen.SSHKey) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.SshKeys = keys }
}

// WithProvisionTags associates tags with the node.
func WithProvisionTags(tags map[string]string) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.Tags = tags }
}

// WithProvisionCloudProviderConfig sets provider-specific configuration for
// advanced customization of the node.
func WithProvisionCloudProviderConfig(cfg gen.CloudProviderConfig) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.CloudproviderConfig = &cfg }
}

// WithProvisionMetro sets the Unikraft Cloud metro to associate the node with.
func WithProvisionMetro(metro string) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.Metro = new(metro) }
}

// WithProvisionPlatformConfig sets user overrides for platform configuration.
// Keys should be kebab-case (e.g., "vm-user-min-memory").
func WithProvisionPlatformConfig(cfg map[string]string) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.PlatformConfig = cfg }
}

// WithProvisionImagePulls pre-pulls the given image references on the
// provisioned node.
func WithProvisionImagePulls(refs ...string) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.ImagePulls = refs }
}

// WithProvisionNetCount overrides the network count for the node. Defaults to 1000.
func WithProvisionNetCount(n uint32) ProvisionNodeOption {
	return func(r *gen.ProvisionNodeRequest) { r.NetCount = new(n) }
}

// NodesOption configures a call to [Client.Nodes].
type NodesOption func(*gen.ListNodesOpts)

// WithNodesCloudProvider filters nodes to those provisioned on the given cloud provider.
func WithNodesCloudProvider(cp gen.CloudProvider) NodesOption {
	return func(o *gen.ListNodesOpts) { o.Cloudprovider = &cp }
}

// WithNodesState filters nodes to those currently in the given state.
func WithNodesState(state gen.NodeState) NodesOption {
	return func(o *gen.ListNodesOpts) { o.State = &state }
}

// WithNodesMetro filters nodes to those associated with the given metro.
func WithNodesMetro(metro string) NodesOption {
	return func(o *gen.ListNodesOpts) { o.Metro = new(metro) }
}

// WithNodesLimit caps the number of nodes returned.
func WithNodesLimit(n uint32) NodesOption {
	return func(o *gen.ListNodesOpts) { o.Limit = new(n) }
}

// WithNodesOffset skips the first n matching nodes.
func WithNodesOffset(n uint32) NodesOption {
	return func(o *gen.ListNodesOpts) { o.Offset = new(n) }
}

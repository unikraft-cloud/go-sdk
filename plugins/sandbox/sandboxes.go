// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"iter"
	"os"
	"slices"
	"time"

	plugin "unikraft.com/cloud/plugins/sandbox"
	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/cloud/sdk/platform/gen"
)

const (
	// DefaultImage is the instance image of a sandbox when none is given. It
	// has a shell and a main process that stays alive.
	DefaultImage = "debian-slim:latest"
	// DefaultPluginImage is the image of the sandbox plugin when none is
	// given.
	DefaultPluginImage = "plugins/sandbox:latest"
	// DefaultMemoryMB is the memory of a sandbox when none is given.
	DefaultMemoryMB = 2048
	// DefaultAutokill is how long a stopped sandbox stays before the
	// platform deletes it.
	DefaultAutokill = 5 * time.Minute
	// DefaultBootTimeout is how long Create waits for the instance to run.
	DefaultBootTimeout = 30 * time.Second
	// DefaultReadyTimeout is how long Create waits for the plugin to answer.
	DefaultReadyTimeout = time.Minute
)

// Sandboxes creates and finds sandboxes: instances with the sandbox plugin.
type Sandboxes struct {
	// Platform is the platform client of the metro of the sandboxes.
	Platform platform.Client
	// Client and Opts reach the sandbox plugin.
	Client *plugin.Client
	Opts   []plugin.Option
	// PluginName is the name of the plugin. Empty is PluginName.
	PluginName string
}

func (s Sandboxes) pluginName() string {
	if s.PluginName != "" {
		return s.PluginName
	}
	return PluginName
}

func (s Sandboxes) target(inst gen.Instance) Target {
	name := s.pluginName()
	return Target{
		Client:   s.Client,
		Instance: inst,
		Plugin:   name,
		Opts:     append(slices.Clone(s.Opts), plugin.WithPluginName(name)),
	}
}

// CreateOption configures a call to [Sandboxes.Create].
type CreateOption func(*createConfig)

type createConfig struct {
	instance     []platform.CreateInstanceOption
	pluginImage  gen.ImageSource
	bootTimeout  time.Duration
	readyTimeout time.Duration
	noReady      bool
}

// WithInstance applies opts to the instance of the sandbox. Use it to set
// the image, the memory, the name or a template.
func WithInstance(opts ...platform.CreateInstanceOption) CreateOption {
	return func(c *createConfig) { c.instance = append(c.instance, opts...) }
}

// WithPluginImage sets the image of the sandbox plugin.
func WithPluginImage(image gen.ImageSource) CreateOption {
	return func(c *createConfig) { c.pluginImage = image }
}

// WithBootTimeout sets how long Create waits for the instance to run.
func WithBootTimeout(d time.Duration) CreateOption {
	return func(c *createConfig) { c.bootTimeout = d }
}

// WithReadyTimeout sets how long Create waits for the plugin to answer.
func WithReadyTimeout(d time.Duration) CreateOption {
	return func(c *createConfig) { c.readyTimeout = d }
}

// WithoutReady makes Create return without waiting for the instance to run
// or for the plugin to answer.
func WithoutReady() CreateOption {
	return func(c *createConfig) { c.noReady = true }
}

// Create creates a sandbox and waits until its plugin answers.
//
// Without a template, branch or checkpoint, the sandbox gets DefaultImage
// (or $UKC_SANDBOX_IMAGE), DefaultMemoryMB and the sandbox plugin. A
// snapshot already has these, so they are then not set. If the plugin does
// not answer in time, Create deletes the sandbox. Use WithoutReady to keep a
// sandbox that fails to start.
func (s Sandboxes) Create(ctx context.Context, opts ...CreateOption) (Target, error) {
	cfg := createConfig{bootTimeout: DefaultBootTimeout, readyTimeout: DefaultReadyTimeout}
	for _, opt := range opts {
		opt(&cfg)
	}

	var pluginErr error
	defaults := func(r *gen.CreateInstanceRequest) {
		fromSnapshot := r.Template != nil || r.BranchFrom != nil || r.Checkpoint != nil
		if !fromSnapshot {
			if r.Image == nil {
				r.Image = gen.ImageReference(cmp.Or(os.Getenv("UKC_SANDBOX_IMAGE"), DefaultImage))
			}
			if r.MemoryMb == nil {
				r.MemoryMb = new(int64(DefaultMemoryMB))
			}
		}
		r.Plugins, pluginErr = s.withPlugin(r.Plugins, cfg.pluginImage, fromSnapshot)
		if r.Autokill == nil {
			r.Autokill = &gen.CreateInstanceRequestAutokill{TimeMs: new(uint64(DefaultAutokill / time.Millisecond))}
		}
		r.Autostart = new(true)
		if cfg.noReady {
			r.TimeoutS = nil
		} else {
			r.TimeoutS = new(int64(cfg.bootTimeout / time.Second))
		}
	}

	instances, err := s.Platform.CreateInstance(ctx, append(slices.Clone(cfg.instance), defaults)...)
	if pluginErr != nil {
		return Target{}, pluginErr
	}
	if err != nil {
		return Target{}, err
	}
	if len(instances) != 1 {
		return Target{}, fmt.Errorf("creating a sandbox made %d instances, not 1", len(instances))
	}
	t := s.target(instances[0].Instance)
	if cfg.noReady {
		return t, nil
	}

	if err := t.WaitReady(ctx, cfg.readyTimeout); err != nil {
		// A running instance with a silent plugin never stops, so autokill
		// does not remove it.
		derr := instances[0].Delete(context.WithoutCancel(ctx))
		return Target{}, errors.Join(err, derr)
	}
	return t, nil
}

// withPlugin adds the sandbox plugin to plugins, unless it is there.
func (s Sandboxes) withPlugin(plugins []gen.CreateInstanceRequestPlugin, image gen.ImageSource, fromSnapshot bool) ([]gen.CreateInstanceRequestPlugin, error) {
	name := s.pluginName()
	if slices.ContainsFunc(plugins, func(p gen.CreateInstanceRequestPlugin) bool { return p.Name == name }) {
		if image != nil {
			return nil, fmt.Errorf("the instance has a plugin named %q and WithPluginImage gives a second image for it", name)
		}
		return plugins, nil
	}
	// A snapshot has its plugins: add one only when an image is given.
	if fromSnapshot && image == nil {
		return plugins, nil
	}
	if image == nil {
		image = gen.ImageReference(DefaultPluginImage)
	}
	return append([]gen.CreateInstanceRequestPlugin{{Name: name, Image: image}}, plugins...), nil
}

// Get returns the sandbox that ref names.
func (s Sandboxes) Get(ctx context.Context, ref platform.Ref) (Target, error) {
	inst, err := s.Platform.Instance(ctx, ref)
	if err != nil {
		return Target{}, err
	}
	if !s.hasPlugin(inst.Instance) {
		return Target{}, fmt.Errorf("instance %s has no %q plugin", ref, s.pluginName())
	}
	return s.target(inst.Instance), nil
}

// List yields the instances that have the sandbox plugin.
func (s Sandboxes) List(ctx context.Context, opts ...platform.ListOption) iter.Seq2[Target, error] {
	return func(yield func(Target, error) bool) {
		for inst, err := range s.Platform.ListInstances(ctx, append(slices.Clone(opts), platform.WithListDetails())...) {
			if err != nil {
				yield(Target{}, err)
				return
			}
			if !s.hasPlugin(inst.Instance) {
				continue
			}
			if !yield(s.target(inst.Instance), nil) {
				return
			}
		}
	}
}

func (s Sandboxes) hasPlugin(inst gen.Instance) bool {
	name := s.pluginName()
	return slices.ContainsFunc(inst.Plugins, func(p gen.InstancePlugin) bool { return p.Name == name })
}

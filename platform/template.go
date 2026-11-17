// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"
	"errors"
	"fmt"
	"time"

	"unikraft.com/cloud/sdk/platform/gen"
)

const (
	prepareWait            = 5 * time.Minute
	preparePollInterval    = 500 * time.Millisecond
	preparePollMaxInterval = 3 * time.Second
)

// Template is a template instance: a snapshot that new instances clone. It
// embeds gen.Instance so its fields are accessed directly.
type Template struct {
	gen.Instance

	client gen.Client
}

func wrapTemplate(client gen.Client, data gen.Instance) *Template {
	return &Template{Instance: cloneInstance(data), client: client}
}

// Refresh re-fetches the template from the API.
func (t *Template) Refresh(ctx context.Context) error {
	resp, err := t.client.GetTemplateInstances(ctx, []gen.NameOrUUID{{Uuid: &t.Uuid}}, gen.GetTemplateInstancesOpts{Details: new(true)})
	if err != nil {
		return err
	}
	if len(resp.Data.Instances) == 0 {
		return fmt.Errorf("template %s not found", t.Uuid)
	}
	t.Instance = cloneInstance(resp.Data.Instances[0])
	return nil
}

// TemplateDeleteOption configures a call to Template.Delete.
type TemplateDeleteOption func(*templateDeleteConfig)

type templateDeleteConfig struct {
	missingOK bool
}

// WithTemplateDeleteMissingOK makes Delete return nil when the template does
// not exist.
func WithTemplateDeleteMissingOK() TemplateDeleteOption {
	return func(c *templateDeleteConfig) { c.missingOK = true }
}

// Delete deletes the template.
func (t *Template) Delete(ctx context.Context, opts ...TemplateDeleteOption) error {
	var cfg templateDeleteConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	_, err := t.client.DeleteTemplateInstanceByUUID(ctx, t.Uuid)
	if cfg.missingOK && gen.ErrorContains(err, gen.APIHTTPErrorNotFound) {
		return nil
	}
	return err
}

// Clone creates instances from the template. opts describe the new
// instances, and cannot set an image or another template.
func (t *Template) Clone(ctx context.Context, opts ...CreateInstanceOption) ([]*Instance, error) {
	var request gen.CreateInstanceRequest
	for _, opt := range opts {
		opt(&request)
	}
	if request.Image != nil || request.Template != nil {
		return nil, errors.New("a clone takes its image from the template: do not set an image or a template")
	}
	request.Template = &gen.CreateInstanceRequestTemplate{Uuid: &t.Uuid}
	resp, err := t.client.CreateInstance(ctx, request)
	if err != nil {
		return nil, err
	}
	instances := make([]*Instance, len(resp.Data.Instances))
	for i, data := range resp.Data.Instances {
		instances[i] = wrapInstance(t.client, data)
	}
	return instances, nil
}

// TemplatesOption configures a call to [Client.Templates].
type TemplatesOption func(*gen.GetTemplateInstancesOpts)

// WithTemplateDetails includes full details for each returned template.
func WithTemplateDetails() TemplatesOption {
	return func(o *gen.GetTemplateInstancesOpts) { o.Details = new(true) }
}

// WithTemplateLimit caps the number of templates returned.
func WithTemplateLimit(n uint32) TemplatesOption {
	return func(o *gen.GetTemplateInstancesOpts) { o.Count = new(n) }
}

// WithTemplatePage continues listing from the pagination cursor returned by
// a previous call.
func WithTemplatePage(from string) TemplatesOption {
	return func(o *gen.GetTemplateInstancesOpts) { o.From = new(from) }
}

// WithTemplateTags returns only templates that have all of tags.
func WithTemplateTags(tags ...string) TemplatesOption {
	return func(o *gen.GetTemplateInstancesOpts) { o.Tags = tags }
}

// WithTemplateOrder sets the sort order for the returned templates.
func WithTemplateOrder(order PaginationOrder) TemplatesOption {
	return func(o *gen.GetTemplateInstancesOpts) { o.Order = new(gen.PaginationOrder(order)) }
}

func (c *platformClient) Template(ctx context.Context, ref Ref) (*Template, error) {
	resp, err := c.raw.GetTemplateInstances(ctx, []gen.NameOrUUID{ref.toNameOrUUID()}, gen.GetTemplateInstancesOpts{Details: new(true)})
	if err != nil {
		return nil, err
	}
	if len(resp.Data.Instances) == 0 {
		return nil, &templateNotFoundError{ref: ref}
	}
	return wrapTemplate(c.raw, resp.Data.Instances[0]), nil
}

func (c *platformClient) Templates(ctx context.Context, refs []Ref, opts ...TemplatesOption) ([]*Template, error) {
	var request []gen.NameOrUUID
	for _, ref := range refs {
		request = append(request, ref.toNameOrUUID())
	}
	var genOpts gen.GetTemplateInstancesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetTemplateInstances(ctx, request, genOpts)
	if err != nil {
		return nil, err
	}
	templates := make([]*Template, len(resp.Data.Instances))
	for i, data := range resp.Data.Instances {
		templates[i] = wrapTemplate(c.raw, data)
	}
	return templates, nil
}

// PrepareTemplateOption configures a call to [Client.PrepareTemplate].
type PrepareTemplateOption func(*gen.CreateInstanceRequestTemplate)

// WithPrepareAutokill deletes the template when no instance has cloned it
// for d.
func WithPrepareAutokill(d time.Duration) PrepareTemplateOption {
	return func(t *gen.CreateInstanceRequestTemplate) {
		t.Autokill = &gen.TemplateAutokill{TimeMs: new(uint64(d / time.Millisecond))}
	}
}

// WithPrepareTimeout stops waiting for the template after d.
func WithPrepareTimeout(d time.Duration) PrepareTemplateOption {
	return func(t *gen.CreateInstanceRequestTemplate) {
		t.PrepareTimeoutS = new(int64(d / time.Second))
	}
}

// PrepareTemplate returns the template with name, and makes it from source
// if it does not exist.
//
// The platform makes a template when an instance clones it. So
// PrepareTemplate creates one stopped seed instance, waits until the template
// is listed, and then deletes the seed. If the template is never listed, the
// seed stays for inspection and PrepareTemplate returns an error. The guest
// of source must write 1 to /uk/libukp/template_instance when it is ready.
func (c *platformClient) PrepareTemplate(ctx context.Context, name string, source gen.Instance, opts ...PrepareTemplateOption) (*Template, error) {
	existing, err := c.Template(ctx, Name(name))
	if err == nil {
		return existing, nil
	}
	if !isTemplateMissing(err) {
		return nil, err
	}

	tmpl := gen.CreateInstanceRequestTemplate{
		Name:       &name,
		Prepare:    new(true),
		CreateArgs: &source,
	}
	for _, opt := range opts {
		opt(&tmpl)
	}
	resp, err := c.raw.CreateInstance(ctx, gen.CreateInstanceRequest{
		Template:  &tmpl,
		Autostart: new(false),
		Tags:      source.Tags,
	})
	if err != nil {
		return nil, err
	}

	wait := prepareWait
	if tmpl.PrepareTimeoutS != nil && *tmpl.PrepareTimeoutS > 0 {
		wait = time.Duration(*tmpl.PrepareTimeoutS) * time.Second
	}
	prepared, err := c.awaitTemplate(ctx, name, wait)
	if err != nil {
		return nil, err
	}

	for _, seed := range resp.Data.Instances {
		err = errors.Join(err, wrapInstance(c.raw, seed).Delete(ctx, WithDeleteMissingOK()))
	}
	return prepared, err
}

// awaitTemplate reads the template with name, and retries with backoff until
// it is listed or wait elapses.
func (c *platformClient) awaitTemplate(ctx context.Context, name string, wait time.Duration) (*Template, error) {
	deadline := time.Now().Add(wait)
	interval := preparePollInterval
	for {
		tmpl, err := c.Template(ctx, Name(name))
		if err == nil || !isTemplateMissing(err) {
			return tmpl, err
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return nil, fmt.Errorf("template %s was not prepared within %s: %w", name, wait, err)
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(min(interval, remaining)):
		}
		interval = min(interval*2, preparePollMaxInterval)
	}
}

func isTemplateMissing(err error) bool {
	var missing *templateNotFoundError
	return errors.As(err, &missing) || gen.ErrorContains(err, gen.APIHTTPErrorNotFound)
}

type templateNotFoundError struct {
	ref Ref
}

func (e *templateNotFoundError) Error() string {
	return fmt.Sprintf("template %s not found", e.ref)
}

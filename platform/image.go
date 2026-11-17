// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"context"

	"unikraft.com/cloud/sdk/platform/gen"
)

// ImagesOption configures a call to [Client.Images].
type ImagesOption func(*gen.GetImagesOpts)

// WithImageTag returns only images with tag.
func WithImageTag(tag string) ImagesOption {
	return func(o *gen.GetImagesOpts) { o.Tag = &tag }
}

// WithImageDigest returns only images with digest.
func WithImageDigest(digest string) ImagesOption {
	return func(o *gen.GetImagesOpts) { o.Digest = &digest }
}

func (c *platformClient) Images(ctx context.Context, opts ...ImagesOption) ([]gen.Image, error) {
	var genOpts gen.GetImagesOpts
	for _, opt := range opts {
		opt(&genOpts)
	}
	resp, err := c.raw.GetImages(ctx, nil, genOpts)
	if err != nil {
		return nil, err
	}
	return resp.Data.Images, nil
}

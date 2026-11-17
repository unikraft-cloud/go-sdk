// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package controlplane

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"unikraft.com/cloud/sdk/controlplane/gen"
)

// defaultNamespace is the namespace of an image reference without one.
const defaultNamespace = "official"

var (
	digestPattern = regexp.MustCompile(`(?i)^[a-z0-9]+(?:[.+_-][a-z0-9]+)*:[0-9a-f]+$`)
	tagPattern    = regexp.MustCompile(`^[A-Za-z0-9_][A-Za-z0-9_.-]{0,127}$`)
)

// RegistryImage is one tag of an image in the registry.
type RegistryImage struct {
	// Repository is the image repository, <namespace>/<name>.
	Repository string
	Tag        string
	Digest     string
	// Size is the size of the image in bytes.
	Size     uint64
	PushTime time.Time
}

// Reference returns the image reference, <repository>:<tag>.
func (i RegistryImage) Reference() string {
	return i.Repository + ":" + i.Tag
}

type imageRef struct {
	namespace  string
	repository string
	tag        string
	digest     string
}

// parseImageRef splits [<registry>/][<namespace>/]<name>[:<tag>][@<digest>].
// It drops a scheme and a registry host, puts a bare name in the official
// namespace, and uses the latest tag when there is no tag and no digest.
func parseImageRef(ref string) (imageRef, error) {
	rest := ref
	if _, after, ok := strings.Cut(rest, "://"); ok {
		rest = after
	}
	rest, digest, hasDigest := strings.Cut(rest, "@")

	head, name := "", rest
	if i := strings.LastIndex(rest, "/"); i >= 0 {
		head, name = rest[:i], rest[i+1:]
	}
	name, tag, hasTag := strings.Cut(name, ":")

	var segments []string
	for s := range strings.SplitSeq(head, "/") {
		if s != "" {
			segments = append(segments, s)
		}
	}
	// A registry host has a dot or a port, or is localhost. A namespace never
	// has one.
	if len(segments) > 0 && (strings.ContainsAny(segments[0], ".:") || segments[0] == "localhost") {
		segments = segments[1:]
	}

	if name == "" {
		return imageRef{}, fmt.Errorf("image reference %q has no image name", ref)
	}
	if hasTag && !tagPattern.MatchString(tag) {
		return imageRef{}, fmt.Errorf("image reference %q has an invalid tag %q", ref, tag)
	}
	if hasDigest && !digestPattern.MatchString(digest) {
		return imageRef{}, fmt.Errorf("image reference %q has an invalid digest %q", ref, digest)
	}
	if len(segments) == 0 {
		segments = []string{defaultNamespace}
	}
	if tag == "" && digest == "" {
		tag = "latest"
	}

	return imageRef{
		namespace:  segments[0],
		repository: strings.Join(append(segments, name), "/"),
		tag:        tag,
		digest:     strings.ToLower(digest),
	}, nil
}

// FindImage returns the registry record of the image that ref names, or nil
// if the registry does not have it. A found image is one that a node can
// pull. FindImage lists the full namespace of the image to find it.
func (c *controlplaneClient) FindImage(ctx context.Context, ref string) (*RegistryImage, error) {
	wanted, err := parseImageRef(ref)
	if err != nil {
		return nil, err
	}
	resp, err := c.raw.ListImages(ctx, listImagesOpts(wanted.namespace))
	if err != nil {
		return nil, err
	}
	for _, image := range resp.Data.Images {
		if image.Name != wanted.repository {
			continue
		}
		for _, tag := range image.Tags {
			if tag.Name == "" {
				continue
			}
			if wanted.tag != "" && tag.Name != wanted.tag {
				continue
			}
			if wanted.digest != "" && !strings.EqualFold(tag.Digest, wanted.digest) {
				continue
			}
			return &RegistryImage{
				Repository: image.Name,
				Tag:        tag.Name,
				Digest:     tag.Digest,
				Size:       tag.Size,
				PushTime:   tag.PushTime,
			}, nil
		}
	}
	return nil, nil
}

// ImageExists reports whether the registry has the image that ref names.
// See FindImage.
func (c *controlplaneClient) ImageExists(ctx context.Context, ref string) (bool, error) {
	image, err := c.FindImage(ctx, ref)
	return image != nil, err
}

func listImagesOpts(namespace string) gen.ListImagesOpts {
	return gen.ListImagesOpts{Details: new(true), Namespace: []string{namespace}}
}

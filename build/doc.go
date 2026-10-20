// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package build is the Go client of the Unikraft Cloud build API. The request
// and response types and the plain JSON operations are generated from
// build.yaml at the root of the repository; the constructors of the union
// types, the progress stream and the transport are written by hand next to
// them.
//
// A build runs on a builder in the caller's own account, and the caller's
// token is the only credential: the service validates it, finds the builder
// with it and pushes the image with it.
//
//	c, err := build.New("https://build.unikraft.cloud", build.WithToken(token))
//	if err != nil {
//		return err
//	}
//
//	src := build.Git("https://github.com/acme/app.git", "main").WithSubdir("app")
//	b, err := c.CreateBuild(ctx, build.CreateBuildRequest{
//		Kraftfile: build.KraftfileFrom(src),
//		Contexts:  build.BuildContexts{Context: src},
//		Output:    []build.BuildOutput{build.Image("index.unikraft.io/acme/app:1.2.3")},
//	})
//	if err != nil {
//		return err
//	}
//
//	logs := c.PlainLogs(ctx, b.ID)
//	defer logs.Close()
//	if _, err := io.Copy(os.Stderr, logs); err != nil {
//		return err // a failed build is an *Error here
//	}
package build

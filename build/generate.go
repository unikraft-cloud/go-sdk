// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

// The models and the JSON operations are rendered from ../build.yaml with the
// templates in templates/. `task generate` fetches the spec and runs this.
//
// The progress stream is not rendered: its schemas mirror buildkit's own
// client.SolveStatus, which Event embeds instead, so that anything that
// renders buildkit progress renders this stream as-is.

//go:generate go run unikraft.com/x/tools/openapi-gen@v0.0.0-20261020135908-5d596d5cad44 -i ../build.yaml -o . -t ./templates -v package=build -v skip=BuildContexts:BuildEvent:Vertex:VertexStatus:VertexLog:VertexWarning:ProgressGroup:SourceInfo:Range:Position

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"unikraft.com/x/kingkong"
)

type cli struct {
	Input   string            `short:"i" help:"Path to OpenAPI spec file" required:"" type:"existingfile"`
	Output  string            `short:"o" help:"Output directory for generated files" required:""`
	Package string            `short:"p" help:"Package name for generated code" required:""`
	TypeMap map[string]string `short:"t" help:"Override a generated Go type. Format: 'GoType=import/path.TargetType'. The key is the Go type that would normally be generated (e.g. 'time.Time') and the value is the fully-qualified replacement type including its import path (e.g. 'k8s.io/apimachinery/pkg/apis/meta/v1.Time')" mapsep:","`
}

func main() {
	var cli cli
	ctx := kong.Parse(&cli,
		kong.Name("codegen"),
		kong.Description("Generate Go SDK code from OpenAPI spec"),
		kong.Help(kingkong.HelpPrinter("")),
	)
	err := run(&cli)
	ctx.FatalIfErrorf(err)
}

func run(cli *cli) error {
	generator, err := NewGenerator(cli.Input, cli.Package, cli.TypeMap)
	if err != nil {
		return fmt.Errorf("error creating generator: %w", err)
	}

	if err := os.MkdirAll(cli.Output, 0o755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	var files []GeneratedFile

	files = append(files, generator.GenerateModels()...)
	files = append(files, generator.GenerateClient())

	files = append(files,
		generator.GenerateRequest(),
		generator.GenerateResponse(),
		generator.GenerateClientOptions(),
		generator.GenerateHTTPAPIErrors(),
	)

	for _, file := range files {
		if err := file.Generate(generator.templates, cli.Output); err != nil {
			return fmt.Errorf("error generating %s: %w", file.Basename, err)
		}
	}

	fmt.Println("Code generation completed successfully!")
	return nil
}

// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ettle/strcase"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// GeneratedFile represents a file to be generated from a template
type GeneratedFile struct {
	TemplateName string
	Data         any
	Basename     string
}

// Generate writes the generated file to the specified directory
func (f *GeneratedFile) Generate(templates *template.Template, outputDir string) error {
	tmpl := templates.Lookup(f.TemplateName)
	if tmpl == nil {
		return fmt.Errorf("template %s not found", f.TemplateName)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, f.Data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}
	if len(bytes.TrimSpace(buf.Bytes())) == 0 {
		// skip if template produced no output
		return nil
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: gofmt failed for %s: %v\n", f.TemplateName, err)
		fmt.Fprintf(os.Stderr, "Unformatted source:\n%s\n", buf.String())
		return fmt.Errorf("formatting code: %w", err)
	}

	filename := f.Basename
	if filename == "" {
		filename = strings.TrimSuffix(f.TemplateName, ".tmpl")
	}
	if !strings.Contains(filename, ".gen") {
		if name, ext, ok := strings.Cut(filename, "."); ok {
			filename = name + ".gen." + ext
		} else {
			filename += ".gen"
		}
	}

	if err := os.WriteFile(filepath.Join(outputDir, filename), formatted, 0o644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	fmt.Printf("Generated %s\n", filename)
	return nil
}

// Generator handles code generation from OpenAPI specs
type Generator struct {
	parser    *Parser
	templates *template.Template
}

// TypeMapping holds a parsed type override
type typeMapping struct {
	From       string
	To         string
	ImportPath string
}

// NewGenerator creates a new code generator
func NewGenerator(specPath, packageName string, rawTypeMaps map[string]string) (*Generator, error) {
	parser, err := NewParser(specPath, packageName)
	if err != nil {
		return nil, fmt.Errorf("creating parser: %w", err)
	}

	typeMaps, err := parseTypeMappings(rawTypeMaps)
	if err != nil {
		return nil, fmt.Errorf("parsing type maps: %w", err)
	}

	Preprocess(parser.doc, parser)

	tmpl, err := template.
		New("").
		Funcs(templateFuncs{parser: parser, typeMaps: typeMaps}.Funcs()).
		ParseFS(templatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("loading templates: %w", err)
	}

	return &Generator{parser: parser, templates: tmpl}, nil
}

func (g *Generator) GenerateModels() []GeneratedFile {
	modelFiles := g.parser.ParseModels()

	files := make([]GeneratedFile, 0, len(modelFiles))
	for _, mf := range modelFiles {
		files = append(files, GeneratedFile{
			TemplateName: "model.go.tmpl",
			Data: map[string]any{
				"PackageName": g.parser.packageName,
				"SchemaName":  mf.SchemaName,
				"Schema":      mf.Schema,
			},
			Basename: "model_" + strcase.ToSnake(mf.SchemaName) + ".gen.go",
		})
	}
	return files
}

func (g *Generator) GenerateClient() GeneratedFile {
	operations := g.parser.ParseOperations()

	data := g.packageData()
	data["Operations"] = operations
	return GeneratedFile{TemplateName: "client.go.tmpl", Data: data}
}

func (g *Generator) GenerateRequest() GeneratedFile {
	return GeneratedFile{TemplateName: "request.go.tmpl", Data: g.packageData()}
}

func (g *Generator) GenerateResponse() GeneratedFile {
	return GeneratedFile{TemplateName: "response.go.tmpl", Data: g.packageData()}
}

func (g *Generator) GenerateClientOptions() GeneratedFile {
	return GeneratedFile{TemplateName: "client_options.go.tmpl", Data: g.packageData()}
}

func (g *Generator) GenerateHTTPAPIErrors() GeneratedFile {
	return GeneratedFile{TemplateName: "http_api_errors.go.tmpl", Data: g.packageData()}
}

func (g *Generator) packageData() map[string]any {
	return map[string]any{"PackageName": g.parser.packageName}
}

// parseTypeMappings parses a raw flag map such as
//
//	"time.Time" -> "k8s.io/apimachinery/pkg/apis/meta/v1.Time"
//
// into TypeMapping values.  The value must be a fully-qualified type of the
// form "<import/path>.<TypeName>", where the last dot separates the import
// path from the identifier.  The generated Go expression becomes
// "<pkg>.<TypeName>" where <pkg> is the last path segment of the import path.
func parseTypeMappings(raw map[string]string) ([]typeMapping, error) {
	mappings := make([]typeMapping, 0, len(raw))
	for from, to := range raw {
		dot := strings.LastIndex(to, ".")
		if dot < 0 {
			return nil, fmt.Errorf("type-map value %q must be in the form 'import/path.TypeName'", to)
		}

		importPath := to[:dot]
		typeName := to[dot+1:]

		// Derive the package qualifier from the last segment of the import path.
		slash := strings.LastIndex(importPath, "/")
		pkgName := importPath[slash+1:]
		mappings = append(mappings, typeMapping{
			From:       from,
			To:         pkgName + "." + typeName,
			ImportPath: importPath,
		})
	}
	return mappings, nil
}

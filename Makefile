# SPDX-License-Identifier: BSD-3-Clause
# Copyright (c) 2025, Unikraft GmbH.
# Licensed under the BSD-3-Clause License (the "License").
# You may not use this file except in compliance with the License.

# Prelude
WORKDIR             ?= $(CURDIR)
Q                   ?= @
CHANNEL             ?= prod-stable

# Tools
GO                  ?= go
CURL                ?= curl
GOIMPORTS           ?= $(GO) run golang.org/x/tools/cmd/goimports@latest

.PHONY: all
all: generate fmt

.PHONY: generate
# sandbox is deliberately excluded until sandbox.yaml is published on the
# release channels; until then it 404s and would fail CI.  Run `make sandbox`
# explicitly to regenerate it, and add it back here once the spec lands.
generate: platform controlplane

.PHONY: platform
platform: platform.yaml
	$(Q)rm -f $(WORKDIR)/platform/*.gen.go
	$(GO) generate ./platform
	ls $(WORKDIR)/platform/*.gen.go | xargs $(GOIMPORTS) -l -w

platform.yaml:
	$(Q)$(CURL) -f -o $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/platform.yaml

.PHONY: controlplane
controlplane: controlplane.yaml
	$(Q)rm -f $(WORKDIR)/controlplane/*.gen.go
	$(GO) generate ./controlplane
	ls $(WORKDIR)/controlplane/*.gen.go | xargs $(GOIMPORTS) -l -w

controlplane.yaml:
	$(Q)$(CURL) -f -o $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/controlplane.yaml

.PHONY: sandbox
sandbox: sandbox.yaml
	$(Q)rm -f $(WORKDIR)/sandbox/*.gen.go
	$(GO) generate ./sandbox
	ls $(WORKDIR)/sandbox/*.gen.go | xargs $(GOIMPORTS) -l -w

sandbox.yaml:
	$(Q)$(CURL) -f -o $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/sandbox.yaml

.PHONY: fmt
fmt:
	ls $(WORKDIR)/platform/*.go | xargs $(GOIMPORTS) -l -w
	ls $(WORKDIR)/controlplane/*.go | xargs $(GOIMPORTS) -l -w
	ls $(WORKDIR)/sandbox/*.go | xargs $(GOIMPORTS) -l -w

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	$(GO) test ./...

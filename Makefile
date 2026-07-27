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

.PHONY: fmt
fmt:
	ls $(WORKDIR)/platform/*.go | xargs $(GOIMPORTS) -l -w
	ls $(WORKDIR)/controlplane/*.go | xargs $(GOIMPORTS) -l -w

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	$(GO) test ./...

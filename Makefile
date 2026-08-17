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

.PHONY: all
all: generate fmt

.PHONY: generate
generate: platform controlplane

.PHONY: platform
platform: platform.yaml
	$(Q)rm -f $(WORKDIR)/platform/*.gen.go
	$(GO) generate ./platform

platform.yaml:
	$(Q)$(CURL) -f -o $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/platform.yaml

.PHONY: controlplane
controlplane: controlplane.yaml
	$(Q)rm -f $(WORKDIR)/controlplane/*.gen.go
	$(GO) generate ./controlplane

controlplane.yaml:
	$(Q)$(CURL) -f -o $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/controlplane.yaml

.PHONY: fmt
fmt:
	go fmt $(WORKDIR)/platform/*.go
	go fmt $(WORKDIR)/controlplane/*.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	$(GO) test ./...

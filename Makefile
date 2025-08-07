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
WGET                ?= wget
DOCKER              ?= docker
OPENAPI_GEN_VERSION ?= v6.4.0
GOIMPORTS           ?= $(GO) run golang.org/x/tools/cmd/goimports@latest

.PHONY: all
all: generate

.PHONY: generate
generate: platform

.PHONY: platform
platform:
	$(DOCKER) run \
		--rm \
		--volume "$(WORKDIR):/local" \
		--user="$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli:$(OPENAPI_GEN_VERSION) generate \
				--generator-name go \
				--engine         "handlebars" \
				--input-spec     /local/platform.yaml \
				--config         /local/platform/config.yaml \
				--template-dir   /local/templates \
				--output         /local/platform \
				--git-repo-id    unikraft-cloud \
				--git-user-id    go-sdk \
				$(OPENAPI_GENERATOR_EXTRA_OPTIONS)

platform.yaml:
	$(Q)$(WGET) -O $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/platform.yaml

.PHONY: controlplane
controlplane: controlplane.yaml
	$(DOCKER) run \
		--rm \
		--volume "$(WORKDIR):/local" \
		--user="$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli:$(OPENAPI_GEN_VERSION) generate \
				--generator-name go \
				--engine         "handlebars" \
				--input-spec     /local/controlplane.yaml \
				--config         /local/controlplane/config.yaml \
				--template-dir   /local/templates \
				--output         /local/controlplane \
				--git-repo-id    unikraft-cloud \
				--git-user-id    go-sdk \
				$(OPENAPI_GENERATOR_EXTRA_OPTIONS)

controlplane.yaml:
	$(Q)$(WGET) -O $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/controlplane.yaml

.PHONY: fmt
fmt:
	ls $(WORKDIR)/platform/*.go | xargs $(GOIMPORTS) -l -w
	ls $(WORKDIR)/controlplane/*.go | xargs $(GOIMPORTS) -l -w

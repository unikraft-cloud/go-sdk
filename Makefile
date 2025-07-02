# SPDX-License-Identifier: BSD-3-Clause
# Copyright (c) 2025, Unikraft GmbH.
# Licensed under the BSD-3-Clause License (the "License").
# You may not use this file except in compliance with the License.

# Tools
GO      ?= go
WGET    ?= wget
OGEN    ?= $(GO) run github.com/ogen-go/ogen/cmd/ogen

# Source
CHANNEL ?= prod-stable

# Misc
Q       ?= @

.PHONY: all
all: generate

.PHONY: generate
generate: platform controlplane

.PHONY: platform
platform: platform.yaml
	$(Q)$(OGEN) -package platform -target platform -clean platform.yaml

platform.yaml:
	$(Q)$(WGET) -O $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/platform.yaml

.PHONY: controlplane
controlplane: controlplane.yaml
	$(Q)$(OGEN) -package controlplane -target controlplane -clean controlplane.yaml

controlplane.yaml:
	$(Q)$(WGET) -O $@ https://raw.githubusercontent.com/unikraft-cloud/openapi/$(CHANNEL)/controlplane.yaml

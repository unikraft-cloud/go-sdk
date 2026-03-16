// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"fmt"
	"strings"
)

const (
	// DefaultMetro is set to a default node based in Frankfurt.
	DefaultMetro = "fra"
	// DefaultMetroEndpoint is the default location of the default node API.
	DefaultMetroEndpoint = "https://api." + DefaultMetro + ".unikraft.cloud"

	// DefaultUserAgent is the default user agent used for API requests.
	DefaultUserAgent = "unikraft-cloud-go-sdk/"
)

func EndpointForMetro(metro string) string {
	if metro == "" {
		return ""
	}
	if strings.Contains(metro, "://") {
		return metro
	}
	return fmt.Sprintf("https://api.%s.unikraft.cloud", metro)
}

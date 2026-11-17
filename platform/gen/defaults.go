// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package gen

import (
	"fmt"
	"os"
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

// tokenEnvVars lists the environment variables that hold the API token, in
// order of precedence.
var tokenEnvVars = []string{"UKC_TOKEN", "UNIKRAFT_CLOUD_TOKEN", "KRAFTCLOUD_TOKEN"}

// MetroEnvVar is the environment variable that holds the default metro.
const MetroEnvVar = "UKC_METRO"

// TokenFromEnv returns the API token from the first set environment variable
// of UKC_TOKEN, UNIKRAFT_CLOUD_TOKEN and KRAFTCLOUD_TOKEN, or "" if none is set.
func TokenFromEnv() string {
	for _, key := range tokenEnvVars {
		if v := os.Getenv(key); v != "" {
			return v
		}
	}
	return ""
}

// MetroFromEnv returns the default metro from UKC_METRO, or "" if it is not set.
func MetroFromEnv() string {
	return os.Getenv(MetroEnvVar)
}

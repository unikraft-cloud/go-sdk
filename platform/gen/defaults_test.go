// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package gen

import "testing"

func TestTokenFromEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want string
	}{
		{name: "none", want: ""},
		{name: "kraftcloud", env: map[string]string{"KRAFTCLOUD_TOKEN": "c"}, want: "c"},
		{name: "unikraft cloud over kraftcloud", env: map[string]string{"UNIKRAFT_CLOUD_TOKEN": "b", "KRAFTCLOUD_TOKEN": "c"}, want: "b"},
		{name: "ukc over all", env: map[string]string{"UKC_TOKEN": "a", "UNIKRAFT_CLOUD_TOKEN": "b", "KRAFTCLOUD_TOKEN": "c"}, want: "a"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range tokenEnvVars {
				t.Setenv(key, tc.env[key])
			}
			if got := TokenFromEnv(); got != tc.want {
				t.Errorf("TokenFromEnv() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestMetroFromEnv(t *testing.T) {
	t.Setenv(MetroEnvVar, "was")
	if got := MetroFromEnv(); got != "was" {
		t.Errorf("MetroFromEnv() = %q, want %q", got, "was")
	}
}

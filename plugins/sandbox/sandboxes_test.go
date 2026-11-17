// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package sandbox

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	plugin "unikraft.com/cloud/plugins/sandbox"
	"unikraft.com/cloud/sdk/platform"
	"unikraft.com/cloud/sdk/platform/gen"
)

// fakeCloud serves the platform API and the sandbox plugin of instance i1.
type fakeCloud struct {
	t          *testing.T
	pluginUp   bool
	instances  []map[string]any
	mu         sync.Mutex
	created    map[string]any
	deletedIDs []string
}

func (f *fakeCloud) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	f.mu.Lock()
	defer f.mu.Unlock()

	reply := func(status int, data any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		require.NoError(f.t, json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": data}))
	}

	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/v1/instances":
		require.NoError(f.t, json.Unmarshal(body, &f.created))
		reply(http.StatusOK, map[string]any{"instances": []map[string]any{{"uuid": "i1", "name": "sbx"}}})
	case r.Method == http.MethodGet && r.URL.Path == "/v1/instances":
		reply(http.StatusOK, map[string]any{"instances": f.instances})
	case r.Method == http.MethodDelete && r.URL.Path == "/v1/instances/i1":
		f.deletedIDs = append(f.deletedIDs, "i1")
		reply(http.StatusOK, map[string]any{})
	case r.URL.Path == pluginPrefixFor("i1")+"/commands":
		if !f.pluginUp {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		reply(http.StatusOK, map[string]any{"commands": []string{}})
	default:
		f.t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	}
}

func pluginPrefixFor(uuid string) string {
	return "/v1/instances/" + uuid + "/plugins/sandbox"
}

func newSandboxes(t *testing.T, f *fakeCloud) Sandboxes {
	t.Helper()
	f.t = t
	srv := httptest.NewServer(f)
	t.Cleanup(srv.Close)
	return Sandboxes{
		Platform: platform.NewClient(platform.WithToken("token"), platform.WithDefaultMetro(srv.URL)),
		Client:   plugin.NewClient(),
		Opts: []plugin.Option{
			plugin.WithEndpoint(srv.URL),
			plugin.WithToken("token"),
			plugin.WithHTTPClient(srv.Client()),
		},
	}
}

func TestSandboxesCreate(t *testing.T) {
	t.Setenv("UKC_SANDBOX_IMAGE", "")

	tests := []struct {
		name string
		opts []CreateOption
		want map[string]any
	}{
		{
			name: "defaults",
			want: map[string]any{
				"image":     DefaultImage,
				"memory_mb": float64(DefaultMemoryMB),
				"plugins":   []any{map[string]any{"name": "sandbox", "image": DefaultPluginImage}},
				"autokill":  map[string]any{"time_ms": float64(300000)},
				"autostart": true,
				"timeout_s": float64(30),
			},
		},
		{
			name: "from a template",
			opts: []CreateOption{
				WithInstance(platform.WithCreateTemplate(gen.CreateInstanceRequestTemplate{Name: new("tpl")})),
				WithBootTimeout(10 * time.Second),
			},
			want: map[string]any{
				"template":  map[string]any{"name": "tpl"},
				"autokill":  map[string]any{"time_ms": float64(300000)},
				"autostart": true,
				"timeout_s": float64(10),
			},
		},
		{
			name: "own image and plugin image",
			opts: []CreateOption{
				WithInstance(platform.WithCreateImage(gen.ImageReference("alpine")), platform.WithCreateMemoryMB(512)),
				WithPluginImage(gen.ImageReference("plugins/sandbox:v2")),
			},
			want: map[string]any{
				"image":     "alpine",
				"memory_mb": float64(512),
				"plugins":   []any{map[string]any{"name": "sandbox", "image": "plugins/sandbox:v2"}},
				"autokill":  map[string]any{"time_ms": float64(300000)},
				"autostart": true,
				"timeout_s": float64(30),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fakeCloud{pluginUp: true}
			target, err := newSandboxes(t, f).Create(t.Context(), tt.opts...)
			require.NoError(t, err)
			assert.Equal(t, "i1", target.Instance.Uuid)
			assert.Equal(t, "sandbox", target.Plugin)
			assert.Equal(t, tt.want, f.created)
			assert.Empty(t, f.deletedIDs)
		})
	}
}

func TestSandboxesCreateDeletesWhenPluginIsSilent(t *testing.T) {
	t.Parallel()

	f := &fakeCloud{}
	_, err := newSandboxes(t, f).Create(t.Context(), WithReadyTimeout(10*time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, []string{"i1"}, f.deletedIDs)
}

func TestSandboxesCreateWithoutReady(t *testing.T) {
	t.Parallel()

	f := &fakeCloud{}
	_, err := newSandboxes(t, f).Create(t.Context(), WithoutReady())
	require.NoError(t, err)
	assert.NotContains(t, f.created, "timeout_s")
	assert.Empty(t, f.deletedIDs)
}

func TestSandboxesCreateRefusesSecondPluginImage(t *testing.T) {
	t.Parallel()

	f := &fakeCloud{pluginUp: true}
	_, err := newSandboxes(t, f).Create(t.Context(),
		WithInstance(platform.WithCreatePlugins(gen.CreateInstanceRequestPlugin{Name: "sandbox"})),
		WithPluginImage(gen.ImageReference("plugins/sandbox:v2")),
	)
	assert.ErrorContains(t, err, "second image")
}

func TestSandboxesGetAndList(t *testing.T) {
	t.Parallel()

	f := &fakeCloud{instances: []map[string]any{
		{"uuid": "i1", "name": "sbx", "plugins": []map[string]any{{"name": "sandbox", "image": "x"}}},
		{"uuid": "i2", "name": "web"},
	}}
	s := newSandboxes(t, f)

	var uuids []string
	for target, err := range s.List(t.Context()) {
		require.NoError(t, err)
		uuids = append(uuids, target.Instance.Uuid)
	}
	assert.Equal(t, []string{"i1"}, uuids)

	target, err := s.Get(t.Context(), platform.Name("sbx"))
	require.NoError(t, err)
	assert.Equal(t, "i1", target.Instance.Uuid)
}

func TestSandboxesGetRefusesInstanceWithoutPlugin(t *testing.T) {
	t.Parallel()

	f := &fakeCloud{instances: []map[string]any{{"uuid": "i2", "name": "web"}}}
	_, err := newSandboxes(t, f).Get(t.Context(), platform.Name("web"))
	assert.ErrorContains(t, err, `no "sandbox" plugin`)
}

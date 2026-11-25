// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH and The Unikraft CLI Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package sandbox

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	plugin "unikraft.com/cloud/plugins/sandbox"
)

// TestStatLocalFile pins what a local source is accepted as: a symlink is
// followed to whatever it points at, a hard link is an ordinary file, and a
// symlink that points nowhere names both itself and its target.
func TestStatLocalFile(t *testing.T) {
	dir := t.TempDir()

	regular := filepath.Join(dir, "regular.txt")
	require.NoError(t, os.WriteFile(regular, []byte("contents\n"), 0o644))

	subdir := filepath.Join(dir, "subdir")
	require.NoError(t, os.Mkdir(subdir, 0o755))

	big := filepath.Join(dir, "big.bin")
	require.NoError(t, os.WriteFile(big, nil, 0o644))
	require.NoError(t, os.Truncate(big, maxFileSize+1))

	hard := filepath.Join(dir, "hard.txt")
	require.NoError(t, os.Link(regular, hard))

	symlink := func(name, target string) string {
		t.Helper()
		link := filepath.Join(dir, name)
		require.NoError(t, os.Symlink(target, link))
		return link
	}

	for _, tt := range []struct {
		name  string
		local string
		wants []string
	}{
		{name: "regular", local: regular},
		{name: "symlink-to-file", local: symlink("to-file", regular)},
		{name: "hard-link", local: hard},
		{name: "directory", local: subdir, wants: []string{"is a directory", "written"}},
		{name: "symlink-to-directory", local: symlink("to-dir", subdir), wants: []string{"is a directory"}},
		{name: "dangling-symlink", local: symlink("dangling", filepath.Join(dir, "missing.txt")), wants: []string{"dangling", "missing.txt", "does not exist"}},
		{name: "over-the-limit", local: big, wants: []string{"over the", "limit"}},
		{name: "symlink-over-the-limit", local: symlink("to-big", big), wants: []string{"over the", "limit"}},
		{name: "missing", local: filepath.Join(dir, "nope.txt"), wants: []string{"reading local file", "nope.txt"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := StatLocalFile(tt.local, "written")
			if len(tt.wants) == 0 {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			for _, want := range tt.wants {
				assert.ErrorContains(t, err, want)
			}
		})
	}
}

// TestReadFile pins that a read gives the plugin its limit, and that the
// plugin's refusal of a file over it is an error that names the file.
func TestReadFile(t *testing.T) {
	const contents = "contents\n"
	target := newTargetServing(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req plugin.ReadFileRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		w.Header().Set("Content-Type", "application/json")
		if uint64(len(contents)) > *req.MaxSize {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			fmt.Fprint(w, `{"status":"error","message":"file is over max_size","errors":[{"status":413}]}`)
			return
		}
		data := plugin.ReadFileData{Contents: base64.StdEncoding.EncodeToString([]byte(contents))}
		_ = json.NewEncoder(w).Encode(map[string]any{"status": "success", "data": data})
	}))

	data, err := target.ReadFile(t.Context(), "/file", int64(len(contents)))
	require.NoError(t, err)
	assert.Equal(t, contents, string(data))

	_, err = target.ReadFile(t.Context(), "/file", int64(len(contents))-1)
	assert.ErrorContains(t, err, `"/file" is over the`)
}

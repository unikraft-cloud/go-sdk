// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package build

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Git returns a source that is the git repository at url, checked out at ref:
// a branch, a tag or a commit. An empty ref is the default branch of the
// repository. The repository is cloned anonymously by the builder.
func Git(url, ref string) Source {
	s := Source{Git: new(url)}
	if ref != "" {
		s.Ref = new(ref)
	}

	return s
}

// URL returns a source that is the tarball at url, which the builder fetches
// and unpacks.
func URL(url string) Source {
	return Source{URL: new(url)}
}

// WithSubdir returns the source narrowed to dir, a directory inside the
// repository or tarball.
func (s Source) WithSubdir(dir string) Source {
	s.Subdir = new(dir)
	return s
}

// WithFilename returns the source pointed at the file name, relative to its
// subdirectory. It is only meaningful for a Kraftfile source, where the file
// defaults to Kraftfile, and for the dockerfile context, where it defaults to
// the name of the Dockerfile that the Kraftfile builds with.
func (s Source) WithFilename(name string) Source {
	s.Filename = new(name)
	return s
}

// KraftfileFrom returns a Kraftfile read out of src by the service, from the
// file named by its Filename, which defaults to Kraftfile. Relative paths in
// the file resolve inside the build context, not inside src.
func KraftfileFrom(src Source) Kraftfile {
	return src
}

// KraftfileObjectOf returns the Kraftfile that is v in its parsed form, for a
// value which encodes to a JSON object with the keys of a Kraftfile, such as
// a parsed *kraftfile.Kraftfile of unikraft.com/x/kraftfile. Its spec is not
// carried, and the object form does not need one. A value that encodes to
// anything but an object, or to an object with a session, git or url key,
// which would make it a source, is refused.
func KraftfileObjectOf(v any) (KraftfileObject, error) {
	if v == nil {
		return nil, errors.New("a Kraftfile object is required")
	}

	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("encoding Kraftfile: %w", err)
	}

	var obj KraftfileObject
	if err := json.Unmarshal(b, &obj); err != nil {
		return nil, fmt.Errorf("a Kraftfile is a JSON object: %w", err)
	}
	if obj == nil {
		return nil, errors.New("a Kraftfile is a JSON object, not null")
	}
	for _, key := range []string{"session", "git", "url"} {
		if _, ok := obj[key]; ok {
			return nil, fmt.Errorf("a Kraftfile object cannot have a %q key, which makes it a source", key)
		}
	}

	return obj, nil
}

// Image returns the output that pushes the built image to ref, including its
// tag, with the caller's token.
func Image(ref string) BuildOutput {
	return BuildOutput{Image: new(ref)}
}

// Volume returns the output that imports the build into the volume name; not supported yet.
func Volume(name string) BuildOutput {
	return BuildOutput{Volume: new(name)}
}

// SessionSource returns the directory that the caller's session serves under name.
func SessionSource(name string) Source {
	return Source{Session: new(name)}
}

// SecretValue returns a secret whose value is sent in the request.
func SecretValue(value string) Secret {
	return Secret{Value: new(value)}
}

// SessionSecret returns the secret that the caller's session serves under id.
func SessionSecret(id string) Secret {
	return Secret{Session: new(id)}
}

// Credentials returns a registry login with username and password.
func Credentials(username, password string) Registry {
	return Registry{Username: new(username), Password: new(password)}
}

// SessionRegistry returns a registry login that the caller's session answers for.
func SessionRegistry() Registry {
	return Registry{Session: new(true)}
}

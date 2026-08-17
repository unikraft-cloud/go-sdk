// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"fmt"

	"unikraft.com/cloud/sdk/platform/gen"
)

// Ref identifies an instance by name or by UUID. Use Name or UUID to
// construct one; a bare string cannot be passed where a Ref is expected,
// which removes the ambiguity of "is this a name or a UUID."
type Ref interface {
	fmt.Stringer

	toNameOrUUID() gen.NameOrUUID
}

type nameRef string

func (r nameRef) String() string { return string(r) }

func (r nameRef) toNameOrUUID() gen.NameOrUUID {
	s := string(r)
	return gen.NameOrUUID{Name: &s}
}

type uuidRef string

func (r uuidRef) String() string { return string(r) }

func (r uuidRef) toNameOrUUID() gen.NameOrUUID {
	s := string(r)
	return gen.NameOrUUID{Uuid: &s}
}

// Name creates a Ref that identifies an instance by its name.
func Name(name string) Ref { return nameRef(name) }

// UUID creates a Ref that identifies an instance by its UUID.
func UUID(uuid string) Ref { return uuidRef(uuid) }

// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"fmt"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// ImageSource is a closed union: exactly one of ImageReference, ImageSpec.
//
// The zero value (nil) means "not set" and is omitted when marshalling.
type ImageSource interface {
	isImageSource()
}

// The image as a plain image reference.
type ImageReference string

func (ImageReference) isImageSource() {}

func (ImageSpec) isImageSource() {}

// UnmarshalImageSource decodes v into whichever ImageSource variant matches the
// kind of the JSON value.
func UnmarshalImageSource(v jsontext.Value) (ImageSource, error) {
	// The kind of the JSON value narrows the candidates but does not prove a
	// match: a failed decode falls through to the variants that share the kind,
	// and its error is only reported once none of them matched either.
	var err error

	switch v.Kind() {
	case '"':
		var out ImageReference
		if err = json.Unmarshal(v, &out); err == nil {
			return out, nil
		}
	case '{':
		var out ImageSpec
		if err = json.Unmarshal(v, &out); err == nil {
			return out, nil
		}
	}

	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("cannot unmarshal %v into ImageSource", v.Kind())
}

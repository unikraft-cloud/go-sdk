// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import "context"

type bearerToken struct {
	token string
}

// Compile-time assertion to ensure that bearerToken implements SecuritySource.
var _ SecuritySource = (*bearerToken)(nil)

func (token *bearerToken) BearerToken(ctx context.Context, operationName OperationName) (BearerToken, error) {
	return BearerToken{
		Token: token.token,
	}, nil
}

// Add a bearer token to the client configuration.  This token will be used for
// authentication when making requests to the platform API.
func Token(token string) SecuritySource {
	return &bearerToken{
		token: token,
	}
}

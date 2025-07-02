# Unikraft Cloud Go SDK

[![](https://pkg.go.dev/badge/github.com/unikraft-cloud/go-sdk.svg)](https://pkg.go.dev/github.com/unikraft-cloud/go-sdk)

This repository contains an auto-generated Go SDK which interfaces with
[Unikraft Cloud](https://unikraft.cloud) based on the public
[OpenAPI](https://github.com/unikraft-cloud/openapi) specification.

> **Get started with Unikraft Cloud Today**
>
> Sign up at https://console.unikraft.cloud/signup


## Quickstart

```go
package main

import (
	"context"
	"fmt"

	ukp "github.com/unikraft-cloud/go-sdk/platform"
)

func main() {
	ctx := context.Background()

	client, err := ukp.NewClient(
		"https://api.fra0.kraft.cloud",
		ukp.Token("..."),
	)
	if err != nil {
		panic(err)
	}

	inst, err := client.GetInstanceByUUID(ctx, platform.GetInstanceByUUIDParams{
		UUID: "ee5611f9-2d83-4ee4-8d76-c882ab0f2c22",
	})
	if err != nil {
		panic(err)
	}

	data, _ := inst.GetData().Get()

	fmt.Sprintf("%#v\n", data)
}
```

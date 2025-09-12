# Unikraft Cloud Go SDK

[![](https://pkg.go.dev/badge/unikraft.com/cloud/sdk.svg)](https://pkg.go.dev/unikraft.com/cloud/sdk)

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

	ukp "unikraft.com/cloud/sdk/platform"
)

func main() {
	ctx := context.Background()

	client := ukp.NewClient(
		ukp.WithDefaultMetro("fra0"),
		ukp.WithToken("...."),
	)

	instResp, err := client.GetInstances(ctx, nil, true)
	if err != nil {
		panic(err)
	}

	for _, instance := range instResp.Data.Instances {
		fmt.Printf("UUID:  %s\n", *instance.Uuid)
		fmt.Printf("Name:  %s\n", *instance.Name)
		fmt.Printf("Image: %s\n", *instance.Image)
		fmt.Printf("Args:  %v\n", instance.Args)
		fmt.Printf("State: %s\n", *instance.State)
		fmt.Printf("-----------------------------------\n")
	}

	fmt.Sprintf("%#v\n", data)
}
```

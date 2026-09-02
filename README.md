# Unikraft Cloud Go SDK

[![](https://pkg.go.dev/badge/unikraft.com/cloud/sdk.svg)](https://pkg.go.dev/unikraft.com/cloud/sdk)

This repository contains an auto-generated Go SDK which interfaces with
[Unikraft Cloud](https://unikraft.cloud) based on the public
[OpenAPI](https://github.com/unikraft-cloud/openapi) specification.

> **Get started with Unikraft Cloud today!**
>
> Sign up at <https://console.unikraft.cloud/signup>.


## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"os"

	"unikraft.com/cloud/sdk/platform"
)

func main() {
	ctx := context.Background()

	// Create a new client. The token and default metro are automatically
	// read from environment variables (UKC_TOKEN and UKC_METRO).
	client := platform.NewClientFromEnv()

	// List all instances
	resp, err := client.GetInstances(ctx, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Check for API-level errors
	if resp.Status != platform.ResponseStatusSuccess {
		fmt.Fprintf(os.Stderr, "error: %s\n", resp.Message)
		os.Exit(1)
	}

	for _, inst := range resp.Data.Instances {
		fmt.Printf("Name:  %s\n", inst.Name)
		fmt.Printf("UUID:  %s\n", inst.Uuid)
		fmt.Printf("State: %s\n", inst.State)
		fmt.Printf("Image: %s\n", inst.Image)
		fmt.Println("---")
	}
}
```

See the [examples/platform-list](examples/platform-list/main.go) for a complete example.

## Configuration

`NewClientFromEnv()` reads configuration from environment variables:

- `UKC_TOKEN` - Your Unikraft Cloud API token
- `UKC_METRO` - The default metro/region to use (e.g., `fra`/`sfo`/`dal`/etc)

Options passed to it take precedence over the environment:

```go
client := platform.NewClientFromEnv(
	platform.WithDefaultMetro("fra"),
)
```

`NewClient()` reads no environment variables at all, so the client is
configured entirely by the options it is given:

```go
client := platform.NewClient(
	platform.WithToken("your-api-token"),
	platform.WithDefaultMetro("fra"),
)
```

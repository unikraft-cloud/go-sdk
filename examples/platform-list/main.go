// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Example program demonstrating how to list instances using the Unikraft Cloud SDK.
package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"unikraft.com/cloud/sdk/platform"
)

func main() {
	ctx := context.Background()

	// Create a new client. The token + default metro is automatically read
	// from environment variables (UKC_TOKEN + UKC_METRO).
	client := platform.NewClient()

	// List all instances with full details
	resp, err := client.GetInstances(ctx, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Check for API-level errors
	if resp.Status != "success" {
		fmt.Fprintf(os.Stderr, "error: %s\n", resp.Message)
		os.Exit(1)
	}
	instances := resp.Data.Instances
	if len(instances) == 0 {
		fmt.Println("No instances found.")
		return
	}

	// Create a tabwriter for aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tUUID\tSTATE\tIMAGE\tMEMORY")

	for _, inst := range instances {
		name := "-"
		if inst.Name != nil {
			name = *inst.Name
		}

		uuid := "-"
		if inst.Uuid != nil {
			uuid = *inst.Uuid
		}

		state := "-"
		if inst.State != nil {
			state = string(*inst.State)
		}

		image := "-"
		if inst.Image != nil {
			image = *inst.Image
		}

		memory := "-"
		if inst.MemoryMb != nil {
			memory = fmt.Sprintf("%dMiB", *inst.MemoryMb)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", name, uuid, state, image, memory)
	}

	w.Flush()
}

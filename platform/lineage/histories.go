// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2026, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package lineage

import (
	"cmp"
	"context"
	"errors"
	"slices"

	"unikraft.com/cloud/sdk/platform"
)

// FetchResult is the result of FetchHistories.
type FetchResult struct {
	// Instances are the histories that were read.
	Instances []platform.GetCheckpointHistoryResponseInstanceHistory
	// Missing are the UUIDs of the instances whose history could not be read.
	Missing []string
	// Errs are the errors of the batches that failed.
	Errs []error
}

// FetchHistories reads the checkpoint histories of instances, at most
// platform.MaxOperations per request. It skips templates and retained
// instances, as the platform gives no history for them. A failed request only
// adds its instances to Missing and its error to Errs. FetchHistories returns
// an error only when every request fails, in which case only Errs is set.
func FetchHistories(ctx context.Context, client platform.Client, instances []platform.Instance) (FetchResult, error) {
	uuids := make([]string, 0, len(instances))
	for _, instance := range instances {
		if instance.State != platform.InstanceStateTemplate && instance.RetainedUntil == nil {
			uuids = append(uuids, instance.Uuid)
		}
	}
	return batchHistories(uuids, func(refs []platform.NameOrUUID) (*platform.Response[platform.GetCheckpointHistoryResponseData], error) {
		return client.GetInstanceHistory(ctx, refs, platform.GetInstanceHistoryOpts{})
	})
}

// batchHistories calls get for uuids, at most platform.MaxOperations at a time.
func batchHistories(uuids []string, get func([]platform.NameOrUUID) (*platform.Response[platform.GetCheckpointHistoryResponseData], error)) (FetchResult, error) {
	var (
		read    FetchResult
		batches int
	)
	for batch := range slices.Chunk(uuids, platform.MaxOperations) {
		batches++
		refs := make([]platform.NameOrUUID, len(batch))
		for i := range batch {
			refs[i].Uuid = &batch[i]
		}
		histories, missing, err := splitHistories(get(refs))
		if err != nil {
			read.Errs = append(read.Errs, err)
			read.Missing = append(read.Missing, batch...)
			continue
		}
		read.Instances = append(read.Instances, histories...)
		read.Missing = append(read.Missing, missing...)
	}
	if len(read.Errs) > 0 && len(read.Errs) == batches {
		return FetchResult{Errs: read.Errs}, errors.Join(read.Errs...)
	}
	return read, nil
}

// splitHistories returns the histories that were read and the UUIDs of the
// instances whose history failed. The client returns an error on a partial
// success, but still decodes the response.
func splitHistories(resp *platform.Response[platform.GetCheckpointHistoryResponseData], err error) (histories []platform.GetCheckpointHistoryResponseInstanceHistory, missing []string, _ error) {
	if resp == nil || resp.Data == nil {
		return nil, nil, cmp.Or(err, errors.New("empty instance history response"))
	}
	if err != nil && resp.Status != platform.ResponseStatusPartialSuccess {
		return nil, nil, err
	}
	for _, h := range resp.Data.Instances {
		if h.Status != nil && *h.Status != platform.ResponseStatusSuccess {
			missing = append(missing, h.Uuid)
			continue
		}
		histories = append(histories, h)
	}
	return histories, missing, nil
}

// MissesHistory reports whether the missing histories could change the tree
// that contains current in t, as Scope returns it before trimming it to a
// limit. t is every node, as Build returns them. An instance in the tree whose
// history is missing may have checkpoints that a limit would keep, and a
// placeholder for a deleted instance may stand in for an instance whose
// history is missing.
func (t Tree) MissesHistory(current string, missing []string) bool {
	if len(missing) == 0 {
		return false
	}
	tree, _ := t.Scope(current, 0)
	nodes := make(map[string]bool, len(tree))
	for _, n := range tree {
		if n.Deleted {
			return true
		}
		nodes[n.UUID] = true
	}
	return slices.ContainsFunc(missing, func(id string) bool { return nodes[id] })
}

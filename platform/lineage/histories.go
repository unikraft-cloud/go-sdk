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
	"strings"

	"unikraft.com/cloud/sdk/platform"
)

// Histories is the result of FetchHistories.
type Histories struct {
	// Instances are the histories that were read.
	Instances []platform.GetCheckpointHistoryResponseInstanceHistory
	// Missing are the UUIDs of the instances whose history could not be read.
	Missing []string
	// Errs are the errors of the batches that failed.
	Errs []error
}

// FetchHistories reads the checkpoint histories of instances, at most
// platform.MaxOperations per request. A failed request only adds its
// instances to Missing and its error to Errs. FetchHistories returns an error
// only when every request fails, in which case only Errs is set.
func FetchHistories(ctx context.Context, client platform.Client, instances []platform.Instance) (Histories, error) {
	uuids := make([]string, 0, len(instances))
	for _, instance := range instances {
		uuids = append(uuids, instance.Uuid)
	}
	return batchHistories(uuids, func(refs []platform.NameOrUUID) (*platform.Response[platform.GetCheckpointHistoryResponseData], error) {
		return client.GetInstanceHistory(ctx, refs, platform.GetInstanceHistoryOpts{})
	})
}

// batchHistories calls get for uuids, at most platform.MaxOperations at a time.
func batchHistories(uuids []string, get func([]platform.NameOrUUID) (*platform.Response[platform.GetCheckpointHistoryResponseData], error)) (Histories, error) {
	var (
		read    Histories
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
		return Histories{Errs: read.Errs}, errors.Join(read.Errs...)
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

// MissesHistory reports whether the missing histories could change t. An
// instance in t whose history is missing may lose checkpoints, and a
// checkpoint whose owner's history is missing is attributed to a placeholder
// for a deleted instance.
func (t Tree) MissesHistory(missing []string) bool {
	if len(missing) == 0 {
		return false
	}
	nodes := make(map[string]bool, len(t))
	for _, n := range t {
		// t may hold a placeholder's checkpoints but not the placeholder.
		if n.Deleted || strings.HasPrefix(n.Source, DeletedPrefix) {
			return true
		}
		nodes[n.UUID] = true
	}
	return slices.ContainsFunc(missing, func(id string) bool { return nodes[id] })
}

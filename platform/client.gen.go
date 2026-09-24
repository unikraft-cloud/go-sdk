// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"unikraft.com/cloud/sdk/pkg/httpclient"
)

type Client interface {
	// Subscribe to audit events as they are raised.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/audit
	//
	// See: https://unikraft.com/docs/api/platform/v1/audit#subscribe-audit-events
	SubscribeAuditEvents(ctx context.Context, opts SubscribeAuditEventsOpts) (<-chan *Response[AuditEventData], error)
	// Create an autoscale configuration for a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/services/{uuid}/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#create-autoscale-configuration-by-service-group-uuid
	CreateAutoscaleConfigurationByServiceGroupUUID(ctx context.Context, uuid string, request CreateAutoscaleConfigurationByServiceGroupUUIDRequest) (*Response[CreateAutoscaleConfigurationsResponseData], error)
	// Add a new autoscale policy to a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/services/{uuid}/autoscale/policies
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#create-autoscale-configuration-policy
	CreateAutoscaleConfigurationPolicy(ctx context.Context, uuid string, request CreateAutoscaleConfigurationPolicyRequest) (*Response[CreateAutoscaleConfigurationPolicyResponseData], error)
	// Create one or more autoscale configurations for the specified service groups.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/services/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#create-autoscale-configurations
	CreateAutoscaleConfigurations(ctx context.Context, request []CreateAutoscaleConfigurationsRequestConfiguration) (*Response[CreateAutoscaleConfigurationsResponseData], error)
	// Delete autoscale policies for a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/services/{uuid}/autoscale/policies
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#delete-autoscale-configuration-policies
	DeleteAutoscaleConfigurationPolicies(ctx context.Context, uuid string, request DeletePolicyRequest) (*Response[DeleteAutoscaleConfigurationPolicyResponseData], error)
	// Delete an autoscale policy by name for a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `name`
	//
	// Performs: DELETE /v1/services/{uuid}/autoscale/policies/{name}
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#delete-autoscale-configuration-policy-by-name
	DeleteAutoscaleConfigurationPolicyByName(ctx context.Context, uuid string, name string) (*Response[DeleteAutoscaleConfigurationPolicyResponseData], error)
	// Delete autoscale configurations for the specified service groups.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/services/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#delete-autoscale-configurations
	DeleteAutoscaleConfigurations(ctx context.Context, request []NameOrUUID) (*Response[DeleteAutoscaleConfigurationsResponseData], error)
	// Delete the autoscale configuration for a service group by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/services/{uuid}/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#delete-autoscale-configurations-by-service-group-uuid
	DeleteAutoscaleConfigurationsByServiceGroupUUID(ctx context.Context, uuid string) (*Response[DeleteAutoscaleConfigurationsResponseData], error)
	// List autoscale policies for a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: GET /v1/services/{uuid}/autoscale/policies
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#get-autoscale-configuration-policies
	GetAutoscaleConfigurationPolicies(ctx context.Context, uuid string, request GetAutoscaleConfigurationPolicyRequest) (*Response[GetAutoscaleConfigurationPolicyResponseData], error)
	// Get an autoscale policy by name for a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `name`
	//
	// Performs: GET /v1/services/{uuid}/autoscale/policies/{name}
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#get-autoscale-configuration-policy-by-name
	GetAutoscaleConfigurationPolicyByName(ctx context.Context, uuid string, name string) (*Response[GetAutoscaleConfigurationPolicyResponseData], error)
	// List autoscale configurations for the specified service groups.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/services/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#get-autoscale-configurations
	GetAutoscaleConfigurations(ctx context.Context, request []NameOrUUID, opts GetAutoscaleConfigurationsOpts) (*Response[GetAutoscaleConfigurationsResponseData], error)
	// Get autoscale configurations for a service group by UUID.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/services/{uuid}/autoscale
	//
	// See: https://unikraft.com/docs/api/platform/v1/autoscale#get-autoscale-configurations-by-service-group-uuid
	GetAutoscaleConfigurationsByServiceGroupUUID(ctx context.Context, uuid string) (*Response[GetAutoscaleConfigurationsResponseData], error)
	// Upload a new certificate.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/certificates
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#create-certificate
	CreateCertificate(ctx context.Context, request CreateCertificateRequest) (*Response[CreateCertificateResponseData], error)
	// Delete a certificate by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/certificates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#delete-certificate-by-uuid
	DeleteCertificateByUUID(ctx context.Context, uuid string) (*Response[DeleteCertificatesResponseData], error)
	// Delete certificates by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/certificates
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#delete-certificates
	DeleteCertificates(ctx context.Context, request []NameOrUUID) (*Response[DeleteCertificatesResponseData], error)
	// Get a certificate by UUID.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/certificates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#get-certificate-by-uuid
	GetCertificateByUUID(ctx context.Context, uuid string) (*Response[GetCertificatesResponseData], error)
	// List certificates.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/certificates
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#get-certificates
	GetCertificates(ctx context.Context, request []NameOrUUID, opts GetCertificatesOpts) (*Response[GetCertificatesResponseData], error)
	// Update a certificate by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/certificates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#update-certificate-by-uuid
	UpdateCertificateByUUID(ctx context.Context, uuid string, request UpdateCertificateByUUIDRequestBody) (*Response[UpdateCertificatesResponseData], error)
	// Update certificates.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/certificates
	//
	// See: https://unikraft.com/docs/api/platform/v1/certificates#update-certificates
	UpdateCertificates(ctx context.Context, request []UpdateCertificatesRequestItem) (*Response[UpdateCertificatesResponseData], error)
	// Retrieve all images in store.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/image-store
	//
	// See: https://unikraft.com/docs/api/platform/v1/images#get-image-store
	GetImageStore(ctx context.Context, request []GetImagesRequestTagOrDigest, opts GetImageStoreOpts) (*Response[GetImagesResponseData], error)
	// Retrieve all images.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/images
	//
	// See: https://unikraft.com/docs/api/platform/v1/images#get-images
	GetImages(ctx context.Context, request []GetImagesRequestTagOrDigest, opts GetImagesOpts) (*Response[GetImagesResponseData], error)
	// Pull and pin one or more images so they stay cached and are never
	// evicted, without relying on an on-demand pull to succeed at instance
	// start. If a pull fails, the agent's error is returned in the entry's
	// `message` field.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/images
	//
	// See: https://unikraft.com/docs/api/platform/v1/images#pin-images
	PinImages(ctx context.Context, request []PinImageRequestItem) (*Response[PinImagesResponseData], error)
	// Unpin one or more images by UUID, making them eligible for normal
	// cache eviction again. This does not delete the image.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/images
	//
	// See: https://unikraft.com/docs/api/platform/v1/images#unpin-images
	UnpinImages(ctx context.Context, request []UnpinImageRequestItem) (*Response[UnpinImagesResponseData], error)
	// Create a checkpoint from an existing instance. A checkpoint captures the
	// state of an instance at a specific point in time. Checkpoints can be
	// created from running, stopped, or standby instances.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/instances/checkpoints
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#create-checkpoint-instances
	CreateCheckpointInstances(ctx context.Context, request []CreateCheckpointInstancesRequestItem) (*Response[CreateCheckpointInstancesResponseData], error)
	// Create an instance in Unikraft Cloud.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/instances
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#create-instance
	CreateInstance(ctx context.Context, request CreateInstanceRequest) (*Response[CreateInstanceResponseData], error)
	// Convert instances to template instances.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/instances/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#create-template-instances
	CreateTemplateInstances(ctx context.Context, request []CreateTemplateInstancesRequestItem) (*Response[CreateTemplateInstancesResponseData], error)
	// Delete a specified checkpoint instance by its UUID. After this call the
	// UUID of the checkpoint instance is no longer valid.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/instances/checkpoints/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-checkpoint-instance-by-uuid
	DeleteCheckpointInstanceByUUID(ctx context.Context, uuid string) (*Response[DeleteCheckpointInstancesResponseData], error)
	// Delete the specified checkpoint instance(s) by ID(s) (name or UUID).
	// After this call the IDs of the checkpoint instances are no longer valid.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/instances/checkpoints
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-checkpoint-instances
	DeleteCheckpointInstances(ctx context.Context, request []NameOrUUID) (*Response[DeleteCheckpointInstancesResponseData], error)
	// Delete instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/instances/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-instance-by-uuid
	DeleteInstanceByUUID(ctx context.Context, uuid string, request DeleteInstanceByUUIDRequestBody) (*Response[DeleteInstancesResponseData], error)
	// Delete instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/instances
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-instances
	DeleteInstances(ctx context.Context, request []DeleteInstanceRequestItem) (*Response[DeleteInstancesResponseData], error)
	// Delete a template instance by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/instances/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-template-instance-by-uuid
	DeleteTemplateInstanceByUUID(ctx context.Context, uuid string) (*Response[DeleteTemplateInstancesResponseData], error)
	// Delete template instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/instances/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#delete-template-instances
	DeleteTemplateInstances(ctx context.Context, request []NameOrUUID) (*Response[DeleteTemplateInstancesResponseData], error)
	// Get the checkpoint history of one or more checkpoint instances.
	// Returns the ordered list of checkpoints in each checkpoint's history.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/checkpoints/history
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-checkpoint-history
	GetCheckpointHistory(ctx context.Context, request []NameOrUUID, opts GetCheckpointHistoryOpts) (*Response[GetCheckpointHistoryResponseData], error)
	// Get the checkpoint history of a checkpoint instance by its UUID.
	// Returns the ordered list of checkpoints in the checkpoint's history.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/instances/checkpoints/{uuid}/history
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-checkpoint-history-by-uuid
	GetCheckpointHistoryByUUID(ctx context.Context, uuid string) (*Response[GetCheckpointHistoryResponseData], error)
	// Get a single checkpoint instance by its UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/checkpoints/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-checkpoint-instance-by-uuid
	GetCheckpointInstanceByUUID(ctx context.Context, uuid string, opts GetCheckpointInstanceByUUIDOpts) (*Response[GetCheckpointInstancesResponseData], error)
	// Get one or more checkpoint instances by their UUID(s) or name(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/checkpoints
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-checkpoint-instances
	GetCheckpointInstances(ctx context.Context, request []NameOrUUID, opts GetCheckpointInstancesOpts) (*Response[GetCheckpointInstancesResponseData], error)
	// Get a single instance by UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-by-uuid
	GetInstanceByUUID(ctx context.Context, uuid string, opts GetInstanceByUUIDOpts) (*Response[GetInstancesResponseData], error)
	// Get the checkpoint history of one or more instances.
	// Returns the ordered list of checkpoints associated with each instance.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/history
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-history
	GetInstanceHistory(ctx context.Context, request []NameOrUUID, opts GetInstanceHistoryOpts) (*Response[GetCheckpointHistoryResponseData], error)
	// Get the checkpoint history of an instance by its UUID.
	// Returns the ordered list of checkpoints associated with the instance.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/instances/{uuid}/history
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-history-by-uuid
	GetInstanceHistoryByUUID(ctx context.Context, uuid string) (*Response[GetCheckpointHistoryResponseData], error)
	// Get instances logs.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/log
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-logs
	GetInstanceLogs(ctx context.Context, request []GetInstancesLogsRequestItem, opts GetInstanceLogsOpts) (*Response[GetInstancesLogsResponseData], error)
	// Get instance logs by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: GET /v1/instances/{uuid}/log
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-logs-by-uuid
	GetInstanceLogsByUUID(ctx context.Context, uuid string, request GetInstanceLogsByUUIDRequestBody) (*Response[GetInstancesLogsResponseData], error)
	// Get instances metrics.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/metrics
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-metrics
	GetInstanceMetrics(ctx context.Context, request []NameOrUUID, opts GetInstanceMetricsOpts) (*Response[GetInstancesMetricsResponseData], error)
	// Get instance metrics by UUID.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/instances/{uuid}/metrics
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instance-metrics-by-uuid
	GetInstanceMetricsByUUID(ctx context.Context, uuid string) (*Response[GetInstancesMetricsResponseData], error)
	// Get one or many instances with their current status and configuration.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-instances
	GetInstances(ctx context.Context, request []NameOrUUID, opts GetInstancesOpts) (*Response[GetInstancesResponseData], error)
	// Get a single template instance by UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-template-instance-by-uuid
	GetTemplateInstanceByUUID(ctx context.Context, uuid string, opts GetTemplateInstanceByUUIDOpts) (*Response[GetTemplateInstancesResponseData], error)
	// Get one or more template instances.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#get-template-instances
	GetTemplateInstances(ctx context.Context, request []NameOrUUID, opts GetTemplateInstancesOpts) (*Response[GetTemplateInstancesResponseData], error)
	// Start instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/{uuid}/start
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#start-instance-by-uuid
	StartInstanceByUUID(ctx context.Context, uuid string, request StartInstanceByUUIDRequestBody) (*Response[StartInstancesResponseData], error)
	// Start instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/start
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#start-instances
	StartInstances(ctx context.Context, request []StartInstancesRequestItem) (*Response[StartInstancesResponseData], error)
	// Stop instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/{uuid}/stop
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#stop-instance-by-uuid
	StopInstanceByUUID(ctx context.Context, uuid string, request StopInstanceByUUIDRequestBody) (*Response[StopInstancesResponseData], error)
	// Stop instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/stop
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#stop-instances
	StopInstances(ctx context.Context, request []StopInstancesRequestItem) (*Response[StopInstancesResponseData], error)
	// Suspend instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/{uuid}/suspend
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#suspend-instance-by-uuid
	SuspendInstanceByUUID(ctx context.Context, uuid string, request SuspendInstanceByUUIDRequestBody) (*Response[SuspendInstancesResponseData], error)
	// Suspend instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/instances/suspend
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#suspend-instances
	SuspendInstances(ctx context.Context, request []SuspendInstancesRequestItem) (*Response[SuspendInstancesResponseData], error)
	// Update (modify) a checkpoint instance by its UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances/checkpoints/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-checkpoint-instance-by-uuid
	UpdateCheckpointInstanceByUUID(ctx context.Context, uuid string, request UpdateCheckpointInstanceByUUIDRequestBody) (*Response[UpdateCheckpointInstancesResponseData], error)
	// Update (modify) one or more checkpoint instances by ID(s) (name or UUID).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances/checkpoints
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-checkpoint-instances
	UpdateCheckpointInstances(ctx context.Context, request []UpdateCheckpointInstancesRequestItem) (*Response[UpdateCheckpointInstancesResponseData], error)
	// Update instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-instance-by-uuid
	UpdateInstanceByUUID(ctx context.Context, uuid string, request UpdateInstanceByUUIDRequestBody) (*Response[UpdateInstancesResponseData], error)
	// Update instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-instances
	UpdateInstances(ctx context.Context, request []UpdateInstancesRequestItem) (*Response[UpdateInstancesResponseData], error)
	// Update a template instance by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-template-instance-by-uuid
	UpdateTemplateInstanceByUUID(ctx context.Context, uuid string, request UpdateTemplateInstanceByUUIDRequestBody) (*Response[UpdateTemplateInstancesResponseData], error)
	// Update template instances by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/instances/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#update-template-instances
	UpdateTemplateInstances(ctx context.Context, request []UpdateTemplateInstancesRequestItem) (*Response[UpdateTemplateInstancesResponseData], error)
	// Wait for instance state by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: GET /v1/instances/{uuid}/wait
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#wait-instance-by-uuid
	WaitInstanceByUUID(ctx context.Context, uuid string, request WaitInstanceByUUIDRequestBody) (*Response[WaitInstancesResponseData], error)
	// Wait for instances to reach states.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/instances/wait
	//
	// See: https://unikraft.com/docs/api/platform/v1/instances#wait-instances
	WaitInstances(ctx context.Context, request []WaitInstancesRequestItem, opts WaitInstancesOpts) (*Response[WaitInstancesResponseData], error)
	// Return the status of a full-system health check.
	//
	// Performs: GET /v1/healthz
	//
	// See: https://unikraft.com/docs/api/platform/v1/node#healthz
	Healthz(ctx context.Context) (*Response[HealthzResponseData], error)
	// Create a new service group.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/services
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#create-service-group
	CreateServiceGroup(ctx context.Context, request CreateServiceGroupRequest) (*Response[CreateServiceGroupResponseData], error)
	// Delete a service group by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/services/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#delete-service-group-by-uuid
	DeleteServiceGroupByUUID(ctx context.Context, uuid string) (*Response[DeleteServiceGroupsResponseData], error)
	// Delete service groups by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/services
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#delete-service-groups
	DeleteServiceGroups(ctx context.Context, request []NameOrUUID) (*Response[DeleteServiceGroupsResponseData], error)
	// Get a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/services/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#get-service-group-by-uuid
	GetServiceGroupByUUID(ctx context.Context, uuid string, opts GetServiceGroupByUUIDOpts) (*Response[GetServiceGroupsResponseData], error)
	// List service groups.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/services
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#get-service-groups
	GetServiceGroups(ctx context.Context, request []NameOrUUID, opts GetServiceGroupsOpts) (*Response[GetServiceGroupsResponseData], error)
	// Update a service group by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/services/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#update-service-group-by-uuid
	UpdateServiceGroupByUUID(ctx context.Context, uuid string, request UpdateServiceGroupByUUIDRequestBody) (*Response[UpdateServiceGroupsResponseData], error)
	// Update service groups.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/services
	//
	// See: https://unikraft.com/docs/api/platform/v1/services#update-service-groups
	UpdateServiceGroups(ctx context.Context, request []UpdateServiceGroupsRequestItem) (*Response[UpdateServiceGroupsResponseData], error)
	// List quota usage and limits of your user account.
	//
	// Performs: GET /v1/users/quotas
	//
	// See: https://unikraft.com/docs/api/platform/v1/users#get-user
	GetUser(ctx context.Context) (*Response[QuotasResponseData], error)
	// List quota usage and limits of a user account by UUID.
	//
	// @param `uuid`
	//
	// Performs: GET /v1/users/{uuid}/quotas
	//
	// See: https://unikraft.com/docs/api/platform/v1/users#get-user-by-uuid
	GetUserByUUID(ctx context.Context, uuid string) (*Response[QuotasResponseData], error)
	// Attach a volume by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/volumes/{uuid}/attach
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#attach-volume-by-uuid
	AttachVolumeByUUID(ctx context.Context, uuid string, request AttachVolumeByUUIDRequestBody) (*Response[AttachVolumesResponseData], error)
	// Attach volumes to instances.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/volumes/attach
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#attach-volumes
	AttachVolumes(ctx context.Context, request []AttachVolumesRequestItem) (*Response[AttachVolumesResponseData], error)
	// Clone a volume by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/volumes/{uuid}/clone
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#clone-volume-by-uuid
	CloneVolumeByUUID(ctx context.Context, uuid string, request CloneVolumeByUUIDRequestBody) (*Response[CloneVolumesResponseData], error)
	// Clone volumes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/volumes/clone
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#clone-volumes
	CloneVolumes(ctx context.Context, request []CloneVolumesRequestItem) (*Response[CloneVolumesResponseData], error)
	// Create a template volume.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/volumes/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#create-template-volume
	CreateTemplateVolume(ctx context.Context, request []NameOrUUID) (*Response[CreateTemplateVolumesResponseData], error)
	// Create a new volume.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: POST /v1/volumes
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#create-volume
	CreateVolume(ctx context.Context, request CreateVolumeRequest) (*Response[CreateVolumeResponseData], error)
	// Delete a template volume by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/volumes/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#delete-template-volume-by-uuid
	DeleteTemplateVolumeByUUID(ctx context.Context, uuid string) (*Response[DeleteTemplateVolumesResponseData], error)
	// Delete template volumes by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/volumes/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#delete-template-volumes
	DeleteTemplateVolumes(ctx context.Context, request []NameOrUUID) (*Response[DeleteTemplateVolumesResponseData], error)
	// Delete a volume by UUID.
	//
	// @param `uuid`
	//
	// Performs: DELETE /v1/volumes/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#delete-volume-by-uuid
	DeleteVolumeByUUID(ctx context.Context, uuid string) (*Response[DeleteVolumesResponseData], error)
	// Delete volumes by ID(s).
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: DELETE /v1/volumes
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#delete-volumes
	DeleteVolumes(ctx context.Context, request []NameOrUUID) (*Response[DeleteVolumesResponseData], error)
	// Detach a volume by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/volumes/{uuid}/detach
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#detach-volume-by-uuid
	DetachVolumeByUUID(ctx context.Context, uuid string, request DetachVolumeByUUIDRequestBody) (*Response[DetachVolumesResponseData], error)
	// Detach volumes from instances.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PUT /v1/volumes/detach
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#detach-volumes
	DetachVolumes(ctx context.Context, request []DetachVolumesRequestItem) (*Response[DetachVolumesResponseData], error)
	// Get a template volume by UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/volumes/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#get-template-volume-by-uuid
	GetTemplateVolumeByUUID(ctx context.Context, uuid string, opts GetTemplateVolumeByUUIDOpts) (*Response[GetTemplateVolumesResponseData], error)
	// List template volumes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/volumes/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#get-template-volumes
	GetTemplateVolumes(ctx context.Context, request []NameOrUUID, opts GetTemplateVolumesOpts) (*Response[GetTemplateVolumesResponseData], error)
	// Get a volume by UUID.
	//
	// @param `uuid`
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/volumes/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#get-volume-by-uuid
	GetVolumeByUUID(ctx context.Context, uuid string, opts GetVolumeByUUIDOpts) (*Response[GetVolumesResponseData], error)
	// List volumes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// @param `opts`
	// 	Optional query parameters for this operation.
	//
	// Performs: GET /v1/volumes
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#get-volumes
	GetVolumes(ctx context.Context, request []NameOrUUID, opts GetVolumesOpts) (*Response[GetVolumesResponseData], error)
	// Update a template volume by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/volumes/templates/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#update-template-volume-by-uuid
	UpdateTemplateVolumeByUUID(ctx context.Context, uuid string, request UpdateTemplateVolumeByUUIDRequestBody) (*Response[UpdateTemplateVolumesResponseData], error)
	// Update template volumes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/volumes/templates
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#update-template-volumes
	UpdateTemplateVolumes(ctx context.Context, request []UpdateTemplateVolumesRequestItem) (*Response[UpdateTemplateVolumesResponseData], error)
	// Update a volume by UUID.
	//
	// @param `uuid`
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/volumes/{uuid}
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#update-volume-by-uuid
	UpdateVolumeByUUID(ctx context.Context, uuid string, request UpdateVolumeByUUIDRequestBody) (*Response[UpdateVolumesResponseData], error)
	// Update volumes.
	//
	// @param `request`
	// 	The request body for this operation.
	//
	// Performs: PATCH /v1/volumes
	//
	// See: https://unikraft.com/docs/api/platform/v1/volumes#update-volumes
	UpdateVolumes(ctx context.Context, request []UpdateVolumesRequestItem) (*Response[UpdateVolumesResponseData], error)
	// WithMetro sets the metro to use when connecting to the API.
	WithMetro(string) Client
	// WithTimeout sets the timeout when making the request.
	WithTimeout(time.Duration) Client
	// WithHTTPClient overwrites the base HTTP client.
	WithHTTPClient(httpclient.HTTPClient) Client
}

// NewClient creates a new client for the API, configured only by the given
// options. Use [NewClientFromEnv] to also take configuration from the
// environment.
func NewClient(copts ...ClientOption) Client {
	options := ClientOptions{}

	for _, opt := range copts {
		opt(&options)
	}

	return newClient(&options)
}

// NewClientFromEnv creates a new client for the API, filling in whatever the
// given options leave unset from the environment: the token from UKC_TOKEN,
// UNIKRAFT_CLOUD_TOKEN or KRAFTCLOUD_TOKEN, and the metro from UKC_METRO.
func NewClientFromEnv(copts ...ClientOption) Client {
	options := ClientOptions{}

	for _, opt := range copts {
		opt(&options)
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UKC_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UNIKRAFT_CLOUD_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("KRAFTCLOUD_TOKEN"))
	}

	if options.DefaultEndpoint() == "" {
		options.SetDefaultMetro(os.Getenv("UKC_METRO"))
	}

	return newClient(&options)
}

// newClient applies the defaults for anything options leaves unset, and
// returns a client backed by them.
func newClient(options *ClientOptions) Client {
	if options.DefaultEndpoint() == "" {
		options.SetDefaultMetro(DefaultMetro)
	}

	if options.AllowInsecure() && options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewHTTPClient(httpclient.WithInsecure()))
	}

	if options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewHTTPClient())
	}

	return &client{
		request: &Request{
			copts: options,
		},
	}
}

type client struct {
	request *Request
}

// WithMetro sets the metro to use when connecting to the API.
func (c *client) WithMetro(m string) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *client) WithHTTPClient(hc httpclient.HTTPClient) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *client) WithTimeout(to time.Duration) Client {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *client) clone() *client {
	ccpy := *c
	return &ccpy
}

func (c *client) SubscribeAuditEvents(ctx context.Context, opts SubscribeAuditEventsOpts) (<-chan *Response[AuditEventData], error) {
	requestPath := "/v1/audit"

	query := make(url.Values)
	for _, v := range opts.Events {
		query.Add("events", string(v))
	}
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	resp := &Response[AuditEventData]{}
	if err := doRequest[AuditEventData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}
	return resp.Events()
}

func (c *client) CreateAutoscaleConfigurationByServiceGroupUUID(ctx context.Context, uuid string, request CreateAutoscaleConfigurationByServiceGroupUUIDRequest) (*Response[CreateAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateAutoscaleConfigurationsResponseData]{}
	if err := doRequest[CreateAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateAutoscaleConfigurationPolicy(ctx context.Context, uuid string, request CreateAutoscaleConfigurationPolicyRequest) (*Response[CreateAutoscaleConfigurationPolicyResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale/policies"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateAutoscaleConfigurationPolicyResponseData]{}
	if err := doRequest[CreateAutoscaleConfigurationPolicyResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateAutoscaleConfigurations(ctx context.Context, request []CreateAutoscaleConfigurationsRequestConfiguration) (*Response[CreateAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/autoscale"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[CreateAutoscaleConfigurationsResponseData]{}
	if err := doRequest[CreateAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteAutoscaleConfigurationPolicies(ctx context.Context, uuid string, request DeletePolicyRequest) (*Response[DeleteAutoscaleConfigurationPolicyResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale/policies"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[DeleteAutoscaleConfigurationPolicyResponseData]{}
	if err := doRequest[DeleteAutoscaleConfigurationPolicyResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteAutoscaleConfigurationPolicyByName(ctx context.Context, uuid string, name string) (*Response[DeleteAutoscaleConfigurationPolicyResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale/policies/{name}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))
	requestPath = strings.ReplaceAll(requestPath, "{name}", url.PathEscape(string(name)))

	resp := &Response[DeleteAutoscaleConfigurationPolicyResponseData]{}
	if err := doRequest[DeleteAutoscaleConfigurationPolicyResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteAutoscaleConfigurations(ctx context.Context, request []NameOrUUID) (*Response[DeleteAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/autoscale"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteAutoscaleConfigurationsResponseData]{}
	if err := doRequest[DeleteAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteAutoscaleConfigurationsByServiceGroupUUID(ctx context.Context, uuid string) (*Response[DeleteAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteAutoscaleConfigurationsResponseData]{}
	if err := doRequest[DeleteAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetAutoscaleConfigurationPolicies(ctx context.Context, uuid string, request GetAutoscaleConfigurationPolicyRequest) (*Response[GetAutoscaleConfigurationPolicyResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale/policies"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[GetAutoscaleConfigurationPolicyResponseData]{}
	if err := doRequest[GetAutoscaleConfigurationPolicyResponseData](ctx, c.request, http.MethodGet, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetAutoscaleConfigurationPolicyByName(ctx context.Context, uuid string, name string) (*Response[GetAutoscaleConfigurationPolicyResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale/policies/{name}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))
	requestPath = strings.ReplaceAll(requestPath, "{name}", url.PathEscape(string(name)))

	resp := &Response[GetAutoscaleConfigurationPolicyResponseData]{}
	if err := doRequest[GetAutoscaleConfigurationPolicyResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetAutoscaleConfigurations(ctx context.Context, request []NameOrUUID, opts GetAutoscaleConfigurationsOpts) (*Response[GetAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/autoscale"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetAutoscaleConfigurationsResponseData]{}
	if err := doRequest[GetAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetAutoscaleConfigurationsByServiceGroupUUID(ctx context.Context, uuid string) (*Response[GetAutoscaleConfigurationsResponseData], error) {
	requestPath := "/v1/services/{uuid}/autoscale"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[GetAutoscaleConfigurationsResponseData]{}
	if err := doRequest[GetAutoscaleConfigurationsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (*Response[CreateCertificateResponseData], error) {
	requestPath := "/v1/certificates"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateCertificateResponseData]{}
	if err := doRequest[CreateCertificateResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteCertificateByUUID(ctx context.Context, uuid string) (*Response[DeleteCertificatesResponseData], error) {
	requestPath := "/v1/certificates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteCertificatesResponseData]{}
	if err := doRequest[DeleteCertificatesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteCertificates(ctx context.Context, request []NameOrUUID) (*Response[DeleteCertificatesResponseData], error) {
	requestPath := "/v1/certificates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteCertificatesResponseData]{}
	if err := doRequest[DeleteCertificatesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCertificateByUUID(ctx context.Context, uuid string) (*Response[GetCertificatesResponseData], error) {
	requestPath := "/v1/certificates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[GetCertificatesResponseData]{}
	if err := doRequest[GetCertificatesResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCertificates(ctx context.Context, request []NameOrUUID, opts GetCertificatesOpts) (*Response[GetCertificatesResponseData], error) {
	requestPath := "/v1/certificates"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetCertificatesResponseData]{}
	if err := doRequest[GetCertificatesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateCertificateByUUID(ctx context.Context, uuid string, request UpdateCertificateByUUIDRequestBody) (*Response[UpdateCertificatesResponseData], error) {
	requestPath := "/v1/certificates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateCertificatesResponseData]{}
	if err := doRequest[UpdateCertificatesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateCertificates(ctx context.Context, request []UpdateCertificatesRequestItem) (*Response[UpdateCertificatesResponseData], error) {
	requestPath := "/v1/certificates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateCertificatesResponseData]{}
	if err := doRequest[UpdateCertificatesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetImageStore(ctx context.Context, request []GetImagesRequestTagOrDigest, opts GetImageStoreOpts) (*Response[GetImagesResponseData], error) {
	requestPath := "/v1/image-store"

	query := make(url.Values)
	if opts.Digest != nil {
		query.Add("digest", string(*opts.Digest))
	}
	if opts.Tag != nil {
		query.Add("tag", string(*opts.Tag))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetImagesResponseData]{}
	if err := doRequest[GetImagesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetImages(ctx context.Context, request []GetImagesRequestTagOrDigest, opts GetImagesOpts) (*Response[GetImagesResponseData], error) {
	requestPath := "/v1/images"

	query := make(url.Values)
	if opts.Digest != nil {
		query.Add("digest", string(*opts.Digest))
	}
	if opts.Tag != nil {
		query.Add("tag", string(*opts.Tag))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetImagesResponseData]{}
	if err := doRequest[GetImagesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) PinImages(ctx context.Context, request []PinImageRequestItem) (*Response[PinImagesResponseData], error) {
	requestPath := "/v1/images"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[PinImagesResponseData]{}
	if err := doRequest[PinImagesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UnpinImages(ctx context.Context, request []UnpinImageRequestItem) (*Response[UnpinImagesResponseData], error) {
	requestPath := "/v1/images"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UnpinImagesResponseData]{}
	if err := doRequest[UnpinImagesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateCheckpointInstances(ctx context.Context, request []CreateCheckpointInstancesRequestItem) (*Response[CreateCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[CreateCheckpointInstancesResponseData]{}
	if err := doRequest[CreateCheckpointInstancesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateInstance(ctx context.Context, request CreateInstanceRequest) (*Response[CreateInstanceResponseData], error) {
	requestPath := "/v1/instances"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateInstanceResponseData]{}
	if err := doRequest[CreateInstanceResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateTemplateInstances(ctx context.Context, request []CreateTemplateInstancesRequestItem) (*Response[CreateTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[CreateTemplateInstancesResponseData]{}
	if err := doRequest[CreateTemplateInstancesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteCheckpointInstanceByUUID(ctx context.Context, uuid string) (*Response[DeleteCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteCheckpointInstancesResponseData]{}
	if err := doRequest[DeleteCheckpointInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteCheckpointInstances(ctx context.Context, request []NameOrUUID) (*Response[DeleteCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteCheckpointInstancesResponseData]{}
	if err := doRequest[DeleteCheckpointInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteInstanceByUUID(ctx context.Context, uuid string, request DeleteInstanceByUUIDRequestBody) (*Response[DeleteInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[DeleteInstancesResponseData]{}
	if err := doRequest[DeleteInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteInstances(ctx context.Context, request []DeleteInstanceRequestItem) (*Response[DeleteInstancesResponseData], error) {
	requestPath := "/v1/instances"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteInstancesResponseData]{}
	if err := doRequest[DeleteInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteTemplateInstanceByUUID(ctx context.Context, uuid string) (*Response[DeleteTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteTemplateInstancesResponseData]{}
	if err := doRequest[DeleteTemplateInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteTemplateInstances(ctx context.Context, request []NameOrUUID) (*Response[DeleteTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteTemplateInstancesResponseData]{}
	if err := doRequest[DeleteTemplateInstancesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCheckpointHistory(ctx context.Context, request []NameOrUUID, opts GetCheckpointHistoryOpts) (*Response[GetCheckpointHistoryResponseData], error) {
	requestPath := "/v1/instances/checkpoints/history"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetCheckpointHistoryResponseData]{}
	if err := doRequest[GetCheckpointHistoryResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCheckpointHistoryByUUID(ctx context.Context, uuid string) (*Response[GetCheckpointHistoryResponseData], error) {
	requestPath := "/v1/instances/checkpoints/{uuid}/history"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[GetCheckpointHistoryResponseData]{}
	if err := doRequest[GetCheckpointHistoryResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCheckpointInstanceByUUID(ctx context.Context, uuid string, opts GetCheckpointInstanceByUUIDOpts) (*Response[GetCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetCheckpointInstancesResponseData]{}
	if err := doRequest[GetCheckpointInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetCheckpointInstances(ctx context.Context, request []NameOrUUID, opts GetCheckpointInstancesOpts) (*Response[GetCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetCheckpointInstancesResponseData]{}
	if err := doRequest[GetCheckpointInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceByUUID(ctx context.Context, uuid string, opts GetInstanceByUUIDOpts) (*Response[GetInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetInstancesResponseData]{}
	if err := doRequest[GetInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceHistory(ctx context.Context, request []NameOrUUID, opts GetInstanceHistoryOpts) (*Response[GetCheckpointHistoryResponseData], error) {
	requestPath := "/v1/instances/history"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetCheckpointHistoryResponseData]{}
	if err := doRequest[GetCheckpointHistoryResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceHistoryByUUID(ctx context.Context, uuid string) (*Response[GetCheckpointHistoryResponseData], error) {
	requestPath := "/v1/instances/{uuid}/history"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[GetCheckpointHistoryResponseData]{}
	if err := doRequest[GetCheckpointHistoryResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceLogs(ctx context.Context, request []GetInstancesLogsRequestItem, opts GetInstanceLogsOpts) (*Response[GetInstancesLogsResponseData], error) {
	requestPath := "/v1/instances/log"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	for _, v := range opts.Offset {
		query.Add("offset", fmt.Sprintf("%d", v))
	}
	for _, v := range opts.Limit {
		query.Add("limit", fmt.Sprintf("%d", v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetInstancesLogsResponseData]{}
	if err := doRequest[GetInstancesLogsResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceLogsByUUID(ctx context.Context, uuid string, request GetInstanceLogsByUUIDRequestBody) (*Response[GetInstancesLogsResponseData], error) {
	requestPath := "/v1/instances/{uuid}/log"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[GetInstancesLogsResponseData]{}
	if err := doRequest[GetInstancesLogsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceMetrics(ctx context.Context, request []NameOrUUID, opts GetInstanceMetricsOpts) (*Response[GetInstancesMetricsResponseData], error) {
	requestPath := "/v1/instances/metrics"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetInstancesMetricsResponseData]{}
	if err := doRequest[GetInstancesMetricsResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstanceMetricsByUUID(ctx context.Context, uuid string) (*Response[GetInstancesMetricsResponseData], error) {
	requestPath := "/v1/instances/{uuid}/metrics"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[GetInstancesMetricsResponseData]{}
	if err := doRequest[GetInstancesMetricsResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetInstances(ctx context.Context, request []NameOrUUID, opts GetInstancesOpts) (*Response[GetInstancesResponseData], error) {
	requestPath := "/v1/instances"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetInstancesResponseData]{}
	if err := doRequest[GetInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetTemplateInstanceByUUID(ctx context.Context, uuid string, opts GetTemplateInstanceByUUIDOpts) (*Response[GetTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetTemplateInstancesResponseData]{}
	if err := doRequest[GetTemplateInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetTemplateInstances(ctx context.Context, request []NameOrUUID, opts GetTemplateInstancesOpts) (*Response[GetTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetTemplateInstancesResponseData]{}
	if err := doRequest[GetTemplateInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) StartInstanceByUUID(ctx context.Context, uuid string, request StartInstanceByUUIDRequestBody) (*Response[StartInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}/start"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[StartInstancesResponseData]{}
	if err := doRequest[StartInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) StartInstances(ctx context.Context, request []StartInstancesRequestItem) (*Response[StartInstancesResponseData], error) {
	requestPath := "/v1/instances/start"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[StartInstancesResponseData]{}
	if err := doRequest[StartInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) StopInstanceByUUID(ctx context.Context, uuid string, request StopInstanceByUUIDRequestBody) (*Response[StopInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}/stop"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[StopInstancesResponseData]{}
	if err := doRequest[StopInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) StopInstances(ctx context.Context, request []StopInstancesRequestItem) (*Response[StopInstancesResponseData], error) {
	requestPath := "/v1/instances/stop"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[StopInstancesResponseData]{}
	if err := doRequest[StopInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) SuspendInstanceByUUID(ctx context.Context, uuid string, request SuspendInstanceByUUIDRequestBody) (*Response[SuspendInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}/suspend"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[SuspendInstancesResponseData]{}
	if err := doRequest[SuspendInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) SuspendInstances(ctx context.Context, request []SuspendInstancesRequestItem) (*Response[SuspendInstancesResponseData], error) {
	requestPath := "/v1/instances/suspend"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[SuspendInstancesResponseData]{}
	if err := doRequest[SuspendInstancesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateCheckpointInstanceByUUID(ctx context.Context, uuid string, request UpdateCheckpointInstanceByUUIDRequestBody) (*Response[UpdateCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateCheckpointInstancesResponseData]{}
	if err := doRequest[UpdateCheckpointInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateCheckpointInstances(ctx context.Context, request []UpdateCheckpointInstancesRequestItem) (*Response[UpdateCheckpointInstancesResponseData], error) {
	requestPath := "/v1/instances/checkpoints"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateCheckpointInstancesResponseData]{}
	if err := doRequest[UpdateCheckpointInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateInstanceByUUID(ctx context.Context, uuid string, request UpdateInstanceByUUIDRequestBody) (*Response[UpdateInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateInstancesResponseData]{}
	if err := doRequest[UpdateInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateInstances(ctx context.Context, request []UpdateInstancesRequestItem) (*Response[UpdateInstancesResponseData], error) {
	requestPath := "/v1/instances"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateInstancesResponseData]{}
	if err := doRequest[UpdateInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateTemplateInstanceByUUID(ctx context.Context, uuid string, request UpdateTemplateInstanceByUUIDRequestBody) (*Response[UpdateTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateTemplateInstancesResponseData]{}
	if err := doRequest[UpdateTemplateInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateTemplateInstances(ctx context.Context, request []UpdateTemplateInstancesRequestItem) (*Response[UpdateTemplateInstancesResponseData], error) {
	requestPath := "/v1/instances/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateTemplateInstancesResponseData]{}
	if err := doRequest[UpdateTemplateInstancesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) WaitInstanceByUUID(ctx context.Context, uuid string, request WaitInstanceByUUIDRequestBody) (*Response[WaitInstancesResponseData], error) {
	requestPath := "/v1/instances/{uuid}/wait"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[WaitInstancesResponseData]{}
	if err := doRequest[WaitInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) WaitInstances(ctx context.Context, request []WaitInstancesRequestItem, opts WaitInstancesOpts) (*Response[WaitInstancesResponseData], error) {
	requestPath := "/v1/instances/wait"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	for _, v := range opts.State {
		query.Add("state", string(v))
	}
	for _, v := range opts.TimeoutMs {
		query.Add("timeout_ms", fmt.Sprintf("%d", v))
	}
	for _, v := range opts.TimeoutS {
		query.Add("timeout_s", fmt.Sprintf("%d", v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[WaitInstancesResponseData]{}
	if err := doRequest[WaitInstancesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) Healthz(ctx context.Context) (*Response[HealthzResponseData], error) {
	requestPath := "/v1/healthz"

	resp := &Response[HealthzResponseData]{}
	if err := doRequest[HealthzResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateServiceGroup(ctx context.Context, request CreateServiceGroupRequest) (*Response[CreateServiceGroupResponseData], error) {
	requestPath := "/v1/services"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateServiceGroupResponseData]{}
	if err := doRequest[CreateServiceGroupResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteServiceGroupByUUID(ctx context.Context, uuid string) (*Response[DeleteServiceGroupsResponseData], error) {
	requestPath := "/v1/services/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteServiceGroupsResponseData]{}
	if err := doRequest[DeleteServiceGroupsResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteServiceGroups(ctx context.Context, request []NameOrUUID) (*Response[DeleteServiceGroupsResponseData], error) {
	requestPath := "/v1/services"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteServiceGroupsResponseData]{}
	if err := doRequest[DeleteServiceGroupsResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetServiceGroupByUUID(ctx context.Context, uuid string, opts GetServiceGroupByUUIDOpts) (*Response[GetServiceGroupsResponseData], error) {
	requestPath := "/v1/services/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetServiceGroupsResponseData]{}
	if err := doRequest[GetServiceGroupsResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetServiceGroups(ctx context.Context, request []NameOrUUID, opts GetServiceGroupsOpts) (*Response[GetServiceGroupsResponseData], error) {
	requestPath := "/v1/services"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetServiceGroupsResponseData]{}
	if err := doRequest[GetServiceGroupsResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateServiceGroupByUUID(ctx context.Context, uuid string, request UpdateServiceGroupByUUIDRequestBody) (*Response[UpdateServiceGroupsResponseData], error) {
	requestPath := "/v1/services/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateServiceGroupsResponseData]{}
	if err := doRequest[UpdateServiceGroupsResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateServiceGroups(ctx context.Context, request []UpdateServiceGroupsRequestItem) (*Response[UpdateServiceGroupsResponseData], error) {
	requestPath := "/v1/services"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateServiceGroupsResponseData]{}
	if err := doRequest[UpdateServiceGroupsResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetUser(ctx context.Context) (*Response[QuotasResponseData], error) {
	requestPath := "/v1/users/quotas"

	resp := &Response[QuotasResponseData]{}
	if err := doRequest[QuotasResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetUserByUUID(ctx context.Context, uuid string) (*Response[QuotasResponseData], error) {
	requestPath := "/v1/users/{uuid}/quotas"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[QuotasResponseData]{}
	if err := doRequest[QuotasResponseData](ctx, c.request, http.MethodGet, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) AttachVolumeByUUID(ctx context.Context, uuid string, request AttachVolumeByUUIDRequestBody) (*Response[AttachVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}/attach"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[AttachVolumesResponseData]{}
	if err := doRequest[AttachVolumesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) AttachVolumes(ctx context.Context, request []AttachVolumesRequestItem) (*Response[AttachVolumesResponseData], error) {
	requestPath := "/v1/volumes/attach"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[AttachVolumesResponseData]{}
	if err := doRequest[AttachVolumesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CloneVolumeByUUID(ctx context.Context, uuid string, request CloneVolumeByUUIDRequestBody) (*Response[CloneVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}/clone"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CloneVolumesResponseData]{}
	if err := doRequest[CloneVolumesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CloneVolumes(ctx context.Context, request []CloneVolumesRequestItem) (*Response[CloneVolumesResponseData], error) {
	requestPath := "/v1/volumes/clone"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[CloneVolumesResponseData]{}
	if err := doRequest[CloneVolumesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateTemplateVolume(ctx context.Context, request []NameOrUUID) (*Response[CreateTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[CreateTemplateVolumesResponseData]{}
	if err := doRequest[CreateTemplateVolumesResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) CreateVolume(ctx context.Context, request CreateVolumeRequest) (*Response[CreateVolumeResponseData], error) {
	requestPath := "/v1/volumes"

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[CreateVolumeResponseData]{}
	if err := doRequest[CreateVolumeResponseData](ctx, c.request, http.MethodPost, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteTemplateVolumeByUUID(ctx context.Context, uuid string) (*Response[DeleteTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteTemplateVolumesResponseData]{}
	if err := doRequest[DeleteTemplateVolumesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteTemplateVolumes(ctx context.Context, request []NameOrUUID) (*Response[DeleteTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteTemplateVolumesResponseData]{}
	if err := doRequest[DeleteTemplateVolumesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteVolumeByUUID(ctx context.Context, uuid string) (*Response[DeleteVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	resp := &Response[DeleteVolumesResponseData]{}
	if err := doRequest[DeleteVolumesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DeleteVolumes(ctx context.Context, request []NameOrUUID) (*Response[DeleteVolumesResponseData], error) {
	requestPath := "/v1/volumes"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DeleteVolumesResponseData]{}
	if err := doRequest[DeleteVolumesResponseData](ctx, c.request, http.MethodDelete, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DetachVolumeByUUID(ctx context.Context, uuid string, request DetachVolumeByUUIDRequestBody) (*Response[DetachVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}/detach"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[DetachVolumesResponseData]{}
	if err := doRequest[DetachVolumesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) DetachVolumes(ctx context.Context, request []DetachVolumesRequestItem) (*Response[DetachVolumesResponseData], error) {
	requestPath := "/v1/volumes/detach"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[DetachVolumesResponseData]{}
	if err := doRequest[DetachVolumesResponseData](ctx, c.request, http.MethodPut, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetTemplateVolumeByUUID(ctx context.Context, uuid string, opts GetTemplateVolumeByUUIDOpts) (*Response[GetTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetTemplateVolumesResponseData]{}
	if err := doRequest[GetTemplateVolumesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetTemplateVolumes(ctx context.Context, request []NameOrUUID, opts GetTemplateVolumesOpts) (*Response[GetTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetTemplateVolumesResponseData]{}
	if err := doRequest[GetTemplateVolumesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetVolumeByUUID(ctx context.Context, uuid string, opts GetVolumeByUUIDOpts) (*Response[GetVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	query := make(url.Values)
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}

	resp := &Response[GetVolumesResponseData]{}
	if err := doRequest[GetVolumesResponseData](ctx, c.request, http.MethodGet, requestPath, query, nil, resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) GetVolumes(ctx context.Context, request []NameOrUUID, opts GetVolumesOpts) (*Response[GetVolumesResponseData], error) {
	requestPath := "/v1/volumes"

	query := make(url.Values)
	for _, v := range opts.Uuid {
		query.Add("uuid", string(v))
	}
	for _, v := range opts.Name {
		query.Add("name", string(v))
	}
	if opts.Details != nil {
		query.Add("details", fmt.Sprintf("%t", *opts.Details))
	}
	if opts.Count != nil {
		query.Add("count", fmt.Sprintf("%d", *opts.Count))
	}
	if opts.From != nil {
		query.Add("from", string(*opts.From))
	}
	if opts.Order != nil {
		query.Add("order", string(*opts.Order))
	}
	if opts.Sortby != nil {
		query.Add("sortby", string(*opts.Sortby))
	}
	for _, v := range opts.Tags {
		query.Add("tags", string(v))
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[GetVolumesResponseData]{}
	if err := doRequest[GetVolumesResponseData](ctx, c.request, http.MethodGet, requestPath, query, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateTemplateVolumeByUUID(ctx context.Context, uuid string, request UpdateTemplateVolumeByUUIDRequestBody) (*Response[UpdateTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateTemplateVolumesResponseData]{}
	if err := doRequest[UpdateTemplateVolumesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateTemplateVolumes(ctx context.Context, request []UpdateTemplateVolumesRequestItem) (*Response[UpdateTemplateVolumesResponseData], error) {
	requestPath := "/v1/volumes/templates"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateTemplateVolumesResponseData]{}
	if err := doRequest[UpdateTemplateVolumesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateVolumeByUUID(ctx context.Context, uuid string, request UpdateVolumeByUUIDRequestBody) (*Response[UpdateVolumesResponseData], error) {
	requestPath := "/v1/volumes/{uuid}"
	requestPath = strings.ReplaceAll(requestPath, "{uuid}", url.PathEscape(string(uuid)))

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &Response[UpdateVolumesResponseData]{}
	if err := doRequest[UpdateVolumesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *client) UpdateVolumes(ctx context.Context, request []UpdateVolumesRequestItem) (*Response[UpdateVolumesResponseData], error) {
	requestPath := "/v1/volumes"

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}

	resp := &Response[UpdateVolumesResponseData]{}
	if err := doRequest[UpdateVolumesResponseData](ctx, c.request, http.MethodPatch, requestPath, nil, bytes.NewReader(body), resp); err != nil {
		return resp, err
	}
	return resp, nil
}

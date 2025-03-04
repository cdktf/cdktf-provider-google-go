// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshRouteDestinations struct {
	// Required.
	//
	// The clusters where the Gateway API HTTPRoute resource will be deployed to. Valid entries include the associated entities IDs configured in the Target resource and "@self" to include the Target cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddeploy_delivery_pipeline#destination_ids ClouddeployDeliveryPipeline#destination_ids}
	DestinationIds *[]*string `field:"required" json:"destinationIds" yaml:"destinationIds"`
	// Optional.
	//
	// Whether to propagate the Kubernetes Service to the route destination clusters. The Service will always be deployed to the Target cluster even if the HTTPRoute is not. This option may be used to facilitiate successful DNS lookup in the route destination clusters. Can only be set to true if destinations are specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddeploy_delivery_pipeline#propagate_service ClouddeployDeliveryPipeline#propagate_service}
	PropagateService interface{} `field:"optional" json:"propagateService" yaml:"propagateService"`
}


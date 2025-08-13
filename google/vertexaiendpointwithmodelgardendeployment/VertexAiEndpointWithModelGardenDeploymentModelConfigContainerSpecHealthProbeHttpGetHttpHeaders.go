// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeHttpGetHttpHeaders struct {
	// The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#name VertexAiEndpointWithModelGardenDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#value VertexAiEndpointWithModelGardenDeployment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


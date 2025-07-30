// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentDeployConfig struct {
	// dedicated_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#dedicated_resources VertexAiEndpointWithModelGardenDeployment#dedicated_resources}
	DedicatedResources *VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources `field:"optional" json:"dedicatedResources" yaml:"dedicatedResources"`
	// If true, enable the QMT fast tryout feature for this model if possible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#fast_tryout_enabled VertexAiEndpointWithModelGardenDeployment#fast_tryout_enabled}
	FastTryoutEnabled interface{} `field:"optional" json:"fastTryoutEnabled" yaml:"fastTryoutEnabled"`
	// System labels for Model Garden deployments. These labels are managed by Google and for tracking purposes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#system_labels VertexAiEndpointWithModelGardenDeployment#system_labels}
	SystemLabels *map[string]*string `field:"optional" json:"systemLabels" yaml:"systemLabels"`
}


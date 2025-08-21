// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfig struct {
	// Whether the user accepts the End User License Agreement (EULA) for the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#accept_eula VertexAiEndpointWithModelGardenDeployment#accept_eula}
	AcceptEula interface{} `field:"optional" json:"acceptEula" yaml:"acceptEula"`
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#container_spec VertexAiEndpointWithModelGardenDeployment#container_spec}
	ContainerSpec *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// The Hugging Face read access token used to access the model artifacts of gated models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#hugging_face_access_token VertexAiEndpointWithModelGardenDeployment#hugging_face_access_token}
	HuggingFaceAccessToken *string `field:"optional" json:"huggingFaceAccessToken" yaml:"huggingFaceAccessToken"`
	// If true, the model will deploy with a cached version instead of directly downloading the model artifacts from Hugging Face.
	//
	// This is suitable for
	// VPC-SC users with limited internet access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#hugging_face_cache_enabled VertexAiEndpointWithModelGardenDeployment#hugging_face_cache_enabled}
	HuggingFaceCacheEnabled interface{} `field:"optional" json:"huggingFaceCacheEnabled" yaml:"huggingFaceCacheEnabled"`
	// The user-specified display name of the uploaded model. If not set, a default name will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#model_display_name VertexAiEndpointWithModelGardenDeployment#model_display_name}
	ModelDisplayName *string `field:"optional" json:"modelDisplayName" yaml:"modelDisplayName"`
}


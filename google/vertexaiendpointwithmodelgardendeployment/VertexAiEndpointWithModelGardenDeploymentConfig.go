// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiEndpointWithModelGardenDeploymentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Resource ID segment making up resource 'location'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#location VertexAiEndpointWithModelGardenDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// deploy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#deploy_config VertexAiEndpointWithModelGardenDeployment#deploy_config}
	DeployConfig *VertexAiEndpointWithModelGardenDeploymentDeployConfig `field:"optional" json:"deployConfig" yaml:"deployConfig"`
	// endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#endpoint_config VertexAiEndpointWithModelGardenDeployment#endpoint_config}
	EndpointConfig *VertexAiEndpointWithModelGardenDeploymentEndpointConfig `field:"optional" json:"endpointConfig" yaml:"endpointConfig"`
	// The Hugging Face model to deploy. Format: Hugging Face model ID like 'google/gemma-2-2b-it'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#hugging_face_model_id VertexAiEndpointWithModelGardenDeployment#hugging_face_model_id}
	HuggingFaceModelId *string `field:"optional" json:"huggingFaceModelId" yaml:"huggingFaceModelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#id VertexAiEndpointWithModelGardenDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#model_config VertexAiEndpointWithModelGardenDeployment#model_config}
	ModelConfig *VertexAiEndpointWithModelGardenDeploymentModelConfig `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#project VertexAiEndpointWithModelGardenDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Model Garden model to deploy. Format: 'publishers/{publisher}/models/{publisher_model}@{version_id}', or 'publishers/hf-{hugging-face-author}/models/{hugging-face-model-name}@001'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#publisher_model_name VertexAiEndpointWithModelGardenDeployment#publisher_model_name}
	PublisherModelName *string `field:"optional" json:"publisherModelName" yaml:"publisherModelName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#timeouts VertexAiEndpointWithModelGardenDeployment#timeouts}
	Timeouts *VertexAiEndpointWithModelGardenDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxToolConfig struct {
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
	// High level description of the Tool and its usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#description DialogflowCxTool#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The human-readable name of the tool, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#display_name DialogflowCxTool#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// data_store_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#data_store_spec DialogflowCxTool#data_store_spec}
	DataStoreSpec *DialogflowCxToolDataStoreSpec `field:"optional" json:"dataStoreSpec" yaml:"dataStoreSpec"`
	// function_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#function_spec DialogflowCxTool#function_spec}
	FunctionSpec *DialogflowCxToolFunctionSpec `field:"optional" json:"functionSpec" yaml:"functionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#id DialogflowCxTool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// open_api_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#open_api_spec DialogflowCxTool#open_api_spec}
	OpenApiSpec *DialogflowCxToolOpenApiSpec `field:"optional" json:"openApiSpec" yaml:"openApiSpec"`
	// The agent to create a Tool for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#parent DialogflowCxTool#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_tool#timeouts DialogflowCxTool#timeouts}
	Timeouts *DialogflowCxToolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


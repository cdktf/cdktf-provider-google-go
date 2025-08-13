// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxGeneratorConfig struct {
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
	// The human-readable name of the generator, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#display_name DialogflowCxGenerator#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// prompt_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#prompt_text DialogflowCxGenerator#prompt_text}
	PromptText *DialogflowCxGeneratorPromptText `field:"required" json:"promptText" yaml:"promptText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#id DialogflowCxGenerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The language to create generators for the following fields: * Generator.prompt_text.text If not specified, the agent's default language is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#language_code DialogflowCxGenerator#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// llm_model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#llm_model_settings DialogflowCxGenerator#llm_model_settings}
	LlmModelSettings *DialogflowCxGeneratorLlmModelSettings `field:"optional" json:"llmModelSettings" yaml:"llmModelSettings"`
	// model_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#model_parameter DialogflowCxGenerator#model_parameter}
	ModelParameter *DialogflowCxGeneratorModelParameter `field:"optional" json:"modelParameter" yaml:"modelParameter"`
	// The agent to create a Generator for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#parent DialogflowCxGenerator#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// placeholders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#placeholders DialogflowCxGenerator#placeholders}
	Placeholders interface{} `field:"optional" json:"placeholders" yaml:"placeholders"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_generator#timeouts DialogflowCxGenerator#timeouts}
	Timeouts *DialogflowCxGeneratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


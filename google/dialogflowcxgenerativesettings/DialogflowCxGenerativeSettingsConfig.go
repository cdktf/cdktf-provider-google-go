// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxGenerativeSettingsConfig struct {
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
	// Language for this settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#language_code DialogflowCxGenerativeSettings#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// fallback_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#fallback_settings DialogflowCxGenerativeSettings#fallback_settings}
	FallbackSettings *DialogflowCxGenerativeSettingsFallbackSettings `field:"optional" json:"fallbackSettings" yaml:"fallbackSettings"`
	// generative_safety_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#generative_safety_settings DialogflowCxGenerativeSettings#generative_safety_settings}
	GenerativeSafetySettings *DialogflowCxGenerativeSettingsGenerativeSafetySettings `field:"optional" json:"generativeSafetySettings" yaml:"generativeSafetySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#id DialogflowCxGenerativeSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// knowledge_connector_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#knowledge_connector_settings DialogflowCxGenerativeSettings#knowledge_connector_settings}
	KnowledgeConnectorSettings *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings `field:"optional" json:"knowledgeConnectorSettings" yaml:"knowledgeConnectorSettings"`
	// llm_model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#llm_model_settings DialogflowCxGenerativeSettings#llm_model_settings}
	LlmModelSettings *DialogflowCxGenerativeSettingsLlmModelSettings `field:"optional" json:"llmModelSettings" yaml:"llmModelSettings"`
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#parent DialogflowCxGenerativeSettings#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_generative_settings#timeouts DialogflowCxGenerativeSettings#timeouts}
	Timeouts *DialogflowCxGenerativeSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


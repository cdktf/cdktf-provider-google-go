// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsFallbackSettings struct {
	// prompt_templates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_generative_settings#prompt_templates DialogflowCxGenerativeSettings#prompt_templates}
	PromptTemplates interface{} `field:"optional" json:"promptTemplates" yaml:"promptTemplates"`
	// Display name of the selected prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_generative_settings#selected_prompt DialogflowCxGenerativeSettings#selected_prompt}
	SelectedPrompt *string `field:"optional" json:"selectedPrompt" yaml:"selectedPrompt"`
}


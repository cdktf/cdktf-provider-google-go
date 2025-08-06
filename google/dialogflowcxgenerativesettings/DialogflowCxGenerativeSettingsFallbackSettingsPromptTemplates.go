// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsFallbackSettingsPromptTemplates struct {
	// Prompt name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_generative_settings#display_name DialogflowCxGenerativeSettings#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// If the flag is true, the prompt is frozen and cannot be modified by users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_generative_settings#frozen DialogflowCxGenerativeSettings#frozen}
	Frozen interface{} `field:"optional" json:"frozen" yaml:"frozen"`
	// Prompt text that is sent to a LLM on no-match default, placeholders are filled downstream.
	//
	// For example: "Here is a conversation $conversation, a response is: "
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_generative_settings#prompt_text DialogflowCxGenerativeSettings#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}


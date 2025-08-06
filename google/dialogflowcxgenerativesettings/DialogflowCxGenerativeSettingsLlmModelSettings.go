// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsLlmModelSettings struct {
	// The selected LLM model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_generative_settings#model DialogflowCxGenerativeSettings#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The custom prompt to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_generative_settings#prompt_text DialogflowCxGenerativeSettings#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}


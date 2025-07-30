// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrases struct {
	// Language code of the phrase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_generative_settings#language_code DialogflowCxGenerativeSettings#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Text input which can be used for prompt or banned phrases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_generative_settings#text DialogflowCxGenerativeSettings#text}
	Text *string `field:"required" json:"text" yaml:"text"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contactcenterinsightsanalysisrule


type ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig struct {
	// Resource name of the Dialogflow conversation profile. Format: projects/{project}/locations/{location}/conversationProfiles/{conversation_profile}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/contact_center_insights_analysis_rule#conversation_profile ContactCenterInsightsAnalysisRule#conversation_profile}
	ConversationProfile *string `field:"optional" json:"conversationProfile" yaml:"conversationProfile"`
	// Default summarization model to be used. Possible values: SUMMARIZATION_MODEL_UNSPECIFIED BASELINE_MODEL BASELINE_MODEL_V2_0 Possible values: ["BASELINE_MODEL", "BASELINE_MODEL_V2_0"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/contact_center_insights_analysis_rule#summarization_model ContactCenterInsightsAnalysisRule#summarization_model}
	SummarizationModel *string `field:"optional" json:"summarizationModel" yaml:"summarizationModel"`
}


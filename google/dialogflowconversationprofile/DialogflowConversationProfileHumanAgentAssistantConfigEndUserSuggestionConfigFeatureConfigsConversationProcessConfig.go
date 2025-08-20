// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig struct {
	// Number of recent non-small-talk sentences to use as context for article and FAQ suggestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#recent_sentences_count DialogflowConversationProfile#recent_sentences_count}
	RecentSentencesCount *float64 `field:"optional" json:"recentSentencesCount" yaml:"recentSentencesCount"`
}


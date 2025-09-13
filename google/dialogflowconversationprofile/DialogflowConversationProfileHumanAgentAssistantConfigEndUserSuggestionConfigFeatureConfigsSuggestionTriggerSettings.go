// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings struct {
	// Do not trigger if last utterance is small talk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#no_small_talk DialogflowConversationProfile#no_small_talk}
	NoSmallTalk interface{} `field:"optional" json:"noSmallTalk" yaml:"noSmallTalk"`
	// Only trigger suggestion if participant role of last utterance is END_USER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#only_end_user DialogflowConversationProfile#only_end_user}
	OnlyEndUser interface{} `field:"optional" json:"onlyEndUser" yaml:"onlyEndUser"`
}


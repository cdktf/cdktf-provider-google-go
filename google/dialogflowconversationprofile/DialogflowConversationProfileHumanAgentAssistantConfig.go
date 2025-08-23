// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfig struct {
	// end_user_suggestion_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#end_user_suggestion_config DialogflowConversationProfile#end_user_suggestion_config}
	EndUserSuggestionConfig *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig `field:"optional" json:"endUserSuggestionConfig" yaml:"endUserSuggestionConfig"`
	// human_agent_suggestion_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#human_agent_suggestion_config DialogflowConversationProfile#human_agent_suggestion_config}
	HumanAgentSuggestionConfig *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig `field:"optional" json:"humanAgentSuggestionConfig" yaml:"humanAgentSuggestionConfig"`
	// message_analysis_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#message_analysis_config DialogflowConversationProfile#message_analysis_config}
	MessageAnalysisConfig *DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfig `field:"optional" json:"messageAnalysisConfig" yaml:"messageAnalysisConfig"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#notification_config DialogflowConversationProfile#notification_config}
	NotificationConfig *DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigs struct {
	// conversation_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#conversation_model_config DialogflowConversationProfile#conversation_model_config}
	ConversationModelConfig *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig `field:"optional" json:"conversationModelConfig" yaml:"conversationModelConfig"`
	// conversation_process_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#conversation_process_config DialogflowConversationProfile#conversation_process_config}
	ConversationProcessConfig *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig `field:"optional" json:"conversationProcessConfig" yaml:"conversationProcessConfig"`
	// Disable the logging of search queries sent by human agents.
	//
	// It can prevent those queries from being stored at answer records.
	// This feature is only supported for types: KNOWLEDGE_SEARCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#disable_agent_query_logging DialogflowConversationProfile#disable_agent_query_logging}
	DisableAgentQueryLogging interface{} `field:"optional" json:"disableAgentQueryLogging" yaml:"disableAgentQueryLogging"`
	// Enable including conversation context during query answer generation. This feature is only supported for types: KNOWLEDGE_SEARCH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#enable_conversation_augmented_query DialogflowConversationProfile#enable_conversation_augmented_query}
	EnableConversationAugmentedQuery interface{} `field:"optional" json:"enableConversationAugmentedQuery" yaml:"enableConversationAugmentedQuery"`
	// Automatically iterates all participants and tries to compile suggestions. This feature is only supported for types: ARTICLE_SUGGESTION, FAQ, DIALOGFLOW_ASSIST, KNOWLEDGE_ASSIST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#enable_event_based_suggestion DialogflowConversationProfile#enable_event_based_suggestion}
	EnableEventBasedSuggestion interface{} `field:"optional" json:"enableEventBasedSuggestion" yaml:"enableEventBasedSuggestion"`
	// Enable query suggestion only. This feature is only supported for types: KNOWLEDGE_ASSIST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#enable_query_suggestion_only DialogflowConversationProfile#enable_query_suggestion_only}
	EnableQuerySuggestionOnly interface{} `field:"optional" json:"enableQuerySuggestionOnly" yaml:"enableQuerySuggestionOnly"`
	// Enable query suggestion even if we can't find its answer.
	//
	// By default, queries are suggested only if we find its answer.
	// This feature is only supported for types: KNOWLEDGE_ASSIST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#enable_query_suggestion_when_no_answer DialogflowConversationProfile#enable_query_suggestion_when_no_answer}
	EnableQuerySuggestionWhenNoAnswer interface{} `field:"optional" json:"enableQuerySuggestionWhenNoAnswer" yaml:"enableQuerySuggestionWhenNoAnswer"`
	// query_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#query_config DialogflowConversationProfile#query_config}
	QueryConfig *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig `field:"optional" json:"queryConfig" yaml:"queryConfig"`
	// suggestion_feature block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#suggestion_feature DialogflowConversationProfile#suggestion_feature}
	SuggestionFeature *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature `field:"optional" json:"suggestionFeature" yaml:"suggestionFeature"`
	// suggestion_trigger_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#suggestion_trigger_settings DialogflowConversationProfile#suggestion_trigger_settings}
	SuggestionTriggerSettings *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings `field:"optional" json:"suggestionTriggerSettings" yaml:"suggestionTriggerSettings"`
}


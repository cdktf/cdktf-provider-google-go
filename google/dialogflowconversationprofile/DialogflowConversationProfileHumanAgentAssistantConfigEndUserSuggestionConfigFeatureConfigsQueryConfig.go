// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig struct {
	// Confidence threshold of query result. This feature is only supported for types: ARTICLE_SUGGESTION, FAQ, SMART_REPLY, SMART_COMPOSE, KNOWLEDGE_SEARCH, KNOWLEDGE_ASSIST, ENTITY_EXTRACTION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#confidence_threshold DialogflowConversationProfile#confidence_threshold}
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// context_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#context_filter_settings DialogflowConversationProfile#context_filter_settings}
	ContextFilterSettings *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings `field:"optional" json:"contextFilterSettings" yaml:"contextFilterSettings"`
	// dialogflow_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#dialogflow_query_source DialogflowConversationProfile#dialogflow_query_source}
	DialogflowQuerySource *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource `field:"optional" json:"dialogflowQuerySource" yaml:"dialogflowQuerySource"`
	// document_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#document_query_source DialogflowConversationProfile#document_query_source}
	DocumentQuerySource *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySource `field:"optional" json:"documentQuerySource" yaml:"documentQuerySource"`
	// knowledge_base_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#knowledge_base_query_source DialogflowConversationProfile#knowledge_base_query_source}
	KnowledgeBaseQuerySource *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySource `field:"optional" json:"knowledgeBaseQuerySource" yaml:"knowledgeBaseQuerySource"`
	// Maximum number of results to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#max_results DialogflowConversationProfile#max_results}
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// sections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#sections DialogflowConversationProfile#sections}
	Sections *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSections `field:"optional" json:"sections" yaml:"sections"`
}


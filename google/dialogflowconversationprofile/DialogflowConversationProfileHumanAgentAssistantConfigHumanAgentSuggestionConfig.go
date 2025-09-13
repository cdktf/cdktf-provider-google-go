// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig struct {
	// When disableHighLatencyFeaturesSyncDelivery is true and using the AnalyzeContent API, we will not deliver the responses from high latency features in the API response.
	//
	// The humanAgentAssistantConfig.notification_config must be configured and enableEventBasedSuggestion must be set to true to receive the responses from high latency features in Pub/Sub. High latency feature(s): KNOWLEDGE_ASSIST
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#disable_high_latency_features_sync_delivery DialogflowConversationProfile#disable_high_latency_features_sync_delivery}
	DisableHighLatencyFeaturesSyncDelivery interface{} `field:"optional" json:"disableHighLatencyFeaturesSyncDelivery" yaml:"disableHighLatencyFeaturesSyncDelivery"`
	// feature_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#feature_configs DialogflowConversationProfile#feature_configs}
	FeatureConfigs interface{} `field:"optional" json:"featureConfigs" yaml:"featureConfigs"`
	// List of various generator resource names used in the conversation profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#generators DialogflowConversationProfile#generators}
	Generators *[]*string `field:"optional" json:"generators" yaml:"generators"`
	// If groupSuggestionResponses is false, and there are multiple featureConfigs in event based suggestion or StreamingAnalyzeContent, we will try to deliver suggestions to customers as soon as we get new suggestion.
	//
	// Different type of suggestions based on the same context will be in separate Pub/Sub event or StreamingAnalyzeContentResponse.
	//
	// If groupSuggestionResponses set to true. All the suggestions to the same participant based on the same context will be grouped into a single Pub/Sub event or StreamingAnalyzeContentResponse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#group_suggestion_responses DialogflowConversationProfile#group_suggestion_responses}
	GroupSuggestionResponses interface{} `field:"optional" json:"groupSuggestionResponses" yaml:"groupSuggestionResponses"`
}


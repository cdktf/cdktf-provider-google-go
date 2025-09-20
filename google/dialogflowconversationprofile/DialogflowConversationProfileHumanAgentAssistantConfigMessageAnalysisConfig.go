// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfig struct {
	// Enable entity extraction in conversation messages on agent assist stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#enable_entity_extraction DialogflowConversationProfile#enable_entity_extraction}
	EnableEntityExtraction interface{} `field:"optional" json:"enableEntityExtraction" yaml:"enableEntityExtraction"`
	// Enable sentiment analysis in conversation messages on agent assist stage.
	//
	// Sentiment analysis inspects user input and identifies the prevailing subjective opinion, especially to determine a user's attitude as positive, negative, or neutral.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#enable_sentiment_analysis DialogflowConversationProfile#enable_sentiment_analysis}
	EnableSentimentAnalysis interface{} `field:"optional" json:"enableSentimentAnalysis" yaml:"enableSentimentAnalysis"`
}


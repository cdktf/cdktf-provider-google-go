// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurnsUserInput struct {
	// Whether sentiment analysis is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_test_case#enable_sentiment_analysis DialogflowCxTestCase#enable_sentiment_analysis}
	EnableSentimentAnalysis interface{} `field:"optional" json:"enableSentimentAnalysis" yaml:"enableSentimentAnalysis"`
	// Parameters that need to be injected into the conversation during intent detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_test_case#injected_parameters DialogflowCxTestCase#injected_parameters}
	InjectedParameters *string `field:"optional" json:"injectedParameters" yaml:"injectedParameters"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_test_case#input DialogflowCxTestCase#input}
	Input *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput `field:"optional" json:"input" yaml:"input"`
	// If webhooks should be allowed to trigger in response to the user utterance.
	//
	// Often if parameters are injected, webhooks should not be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_test_case#is_webhook_enabled DialogflowCxTestCase#is_webhook_enabled}
	IsWebhookEnabled interface{} `field:"optional" json:"isWebhookEnabled" yaml:"isWebhookEnabled"`
}


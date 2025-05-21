// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment struct {
	// advanced_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#advanced_settings DialogflowCxFlow#advanced_settings}
	AdvancedSettings *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// conditional_cases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#conditional_cases DialogflowCxFlow#conditional_cases}
	ConditionalCases interface{} `field:"optional" json:"conditionalCases" yaml:"conditionalCases"`
	// If the flag is true, the agent will utilize LLM to generate a text response.
	//
	// If LLM generation fails, the defined responses in the fulfillment will be respected. This flag is only useful for fulfillments associated with no-match event handlers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#enable_generative_fallback DialogflowCxFlow#enable_generative_fallback}
	EnableGenerativeFallback interface{} `field:"optional" json:"enableGenerativeFallback" yaml:"enableGenerativeFallback"`
	// messages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#messages DialogflowCxFlow#messages}
	Messages interface{} `field:"optional" json:"messages" yaml:"messages"`
	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs.
	//
	// If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#return_partial_responses DialogflowCxFlow#return_partial_responses}
	ReturnPartialResponses interface{} `field:"optional" json:"returnPartialResponses" yaml:"returnPartialResponses"`
	// set_parameter_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#set_parameter_actions DialogflowCxFlow#set_parameter_actions}
	SetParameterActions interface{} `field:"optional" json:"setParameterActions" yaml:"setParameterActions"`
	// The tag used by the webhook to identify which fulfillment is being called.
	//
	// This field is required if webhook is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#tag DialogflowCxFlow#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/dialogflow_cx_flow#webhook DialogflowCxFlow#webhook}
	Webhook *string `field:"optional" json:"webhook" yaml:"webhook"`
}


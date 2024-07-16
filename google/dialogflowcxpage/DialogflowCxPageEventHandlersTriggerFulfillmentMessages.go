// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageEventHandlersTriggerFulfillmentMessages struct {
	// The channel which the response is associated with.
	//
	// Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#channel DialogflowCxPage#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// conversation_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#conversation_success DialogflowCxPage#conversation_success}
	ConversationSuccess *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess `field:"optional" json:"conversationSuccess" yaml:"conversationSuccess"`
	// live_agent_handoff block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#live_agent_handoff DialogflowCxPage#live_agent_handoff}
	LiveAgentHandoff *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff `field:"optional" json:"liveAgentHandoff" yaml:"liveAgentHandoff"`
	// output_audio_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#output_audio_text DialogflowCxPage#output_audio_text}
	OutputAudioText *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText `field:"optional" json:"outputAudioText" yaml:"outputAudioText"`
	// A custom, platform-specific payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#payload DialogflowCxPage#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// play_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#play_audio DialogflowCxPage#play_audio}
	PlayAudio *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio `field:"optional" json:"playAudio" yaml:"playAudio"`
	// telephony_transfer_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#telephony_transfer_call DialogflowCxPage#telephony_transfer_call}
	TelephonyTransferCall *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall `field:"optional" json:"telephonyTransferCall" yaml:"telephonyTransferCall"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/dialogflow_cx_page#text DialogflowCxPage#text}
	Text *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText `field:"optional" json:"text" yaml:"text"`
}


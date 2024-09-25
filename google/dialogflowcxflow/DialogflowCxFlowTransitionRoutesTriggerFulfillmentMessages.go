// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessages struct {
	// The channel which the response is associated with.
	//
	// Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#channel DialogflowCxFlow#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// conversation_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#conversation_success DialogflowCxFlow#conversation_success}
	ConversationSuccess *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess `field:"optional" json:"conversationSuccess" yaml:"conversationSuccess"`
	// live_agent_handoff block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#live_agent_handoff DialogflowCxFlow#live_agent_handoff}
	LiveAgentHandoff *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff `field:"optional" json:"liveAgentHandoff" yaml:"liveAgentHandoff"`
	// output_audio_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#output_audio_text DialogflowCxFlow#output_audio_text}
	OutputAudioText *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText `field:"optional" json:"outputAudioText" yaml:"outputAudioText"`
	// A custom, platform-specific payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#payload DialogflowCxFlow#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// play_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#play_audio DialogflowCxFlow#play_audio}
	PlayAudio *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio `field:"optional" json:"playAudio" yaml:"playAudio"`
	// telephony_transfer_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#telephony_transfer_call DialogflowCxFlow#telephony_transfer_call}
	TelephonyTransferCall *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall `field:"optional" json:"telephonyTransferCall" yaml:"telephonyTransferCall"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/dialogflow_cx_flow#text DialogflowCxFlow#text}
	Text *DialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText `field:"optional" json:"text" yaml:"text"`
}


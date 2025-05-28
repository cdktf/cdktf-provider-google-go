// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessages struct {
	// The channel which the response is associated with.
	//
	// Clients can specify the channel via QueryParameters.channel, and only associated channel response will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#channel DialogflowCxFlow#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// conversation_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#conversation_success DialogflowCxFlow#conversation_success}
	ConversationSuccess *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess `field:"optional" json:"conversationSuccess" yaml:"conversationSuccess"`
	// knowledge_info_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#knowledge_info_card DialogflowCxFlow#knowledge_info_card}
	KnowledgeInfoCard *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard `field:"optional" json:"knowledgeInfoCard" yaml:"knowledgeInfoCard"`
	// live_agent_handoff block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#live_agent_handoff DialogflowCxFlow#live_agent_handoff}
	LiveAgentHandoff *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff `field:"optional" json:"liveAgentHandoff" yaml:"liveAgentHandoff"`
	// output_audio_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#output_audio_text DialogflowCxFlow#output_audio_text}
	OutputAudioText *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText `field:"optional" json:"outputAudioText" yaml:"outputAudioText"`
	// Returns a response containing a custom, platform-specific payload.
	//
	// This field is part of a union field 'message': Only one of 'text', 'payload', 'conversationSuccess', 'outputAudioText', 'liveAgentHandoff', 'endInteraction', 'playAudio', 'mixedAudio', 'telephonyTransferCall', or 'knowledgeInfoCard' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#payload DialogflowCxFlow#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// play_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#play_audio DialogflowCxFlow#play_audio}
	PlayAudio *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio `field:"optional" json:"playAudio" yaml:"playAudio"`
	// telephony_transfer_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#telephony_transfer_call DialogflowCxFlow#telephony_transfer_call}
	TelephonyTransferCall *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall `field:"optional" json:"telephonyTransferCall" yaml:"telephonyTransferCall"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#text DialogflowCxFlow#text}
	Text *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText `field:"optional" json:"text" yaml:"text"`
}


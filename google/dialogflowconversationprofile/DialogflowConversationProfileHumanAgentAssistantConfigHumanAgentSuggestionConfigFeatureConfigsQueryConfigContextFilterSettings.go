// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings struct {
	// If set to true, the last message from virtual agent (hand off message) and the message before it (trigger message of hand off) are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#drop_handoff_messages DialogflowConversationProfile#drop_handoff_messages}
	DropHandoffMessages interface{} `field:"optional" json:"dropHandoffMessages" yaml:"dropHandoffMessages"`
	// If set to true, all messages from ivr stage are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#drop_ivr_messages DialogflowConversationProfile#drop_ivr_messages}
	DropIvrMessages interface{} `field:"optional" json:"dropIvrMessages" yaml:"dropIvrMessages"`
	// If set to true, all messages from virtual agent are dropped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_conversation_profile#drop_virtual_agent_messages DialogflowConversationProfile#drop_virtual_agent_messages}
	DropVirtualAgentMessages interface{} `field:"optional" json:"dropVirtualAgentMessages" yaml:"dropVirtualAgentMessages"`
}


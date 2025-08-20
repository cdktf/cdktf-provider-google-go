// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentHandoffConfig struct {
	// live_person_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_conversation_profile#live_person_config DialogflowConversationProfile#live_person_config}
	LivePersonConfig *DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfig `field:"optional" json:"livePersonConfig" yaml:"livePersonConfig"`
}


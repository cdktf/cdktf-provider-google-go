// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig struct {
	// The name of a dialogflow virtual agent used for intent detection and suggestion triggered by human agent.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_conversation_profile#agent DialogflowConversationProfile#agent}
	Agent *string `field:"optional" json:"agent" yaml:"agent"`
}


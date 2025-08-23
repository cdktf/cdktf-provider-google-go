// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile


type DialogflowConversationProfileAutomatedAgentConfig struct {
	// ID of the Dialogflow agent environment to use. Expects the format "projects/<Project ID>/locations/<Location ID>/agent/environments/<EnvironmentID>".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#agent DialogflowConversationProfile#agent}
	Agent *string `field:"required" json:"agent" yaml:"agent"`
	// Configure lifetime of the Dialogflow session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_conversation_profile#session_ttl DialogflowConversationProfile#session_ttl}
	SessionTtl *string `field:"optional" json:"sessionTtl" yaml:"sessionTtl"`
}


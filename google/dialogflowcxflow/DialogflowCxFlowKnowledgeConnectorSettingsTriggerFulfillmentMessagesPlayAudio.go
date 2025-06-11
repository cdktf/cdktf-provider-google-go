// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio struct {
	// URI of the audio clip.
	//
	// Dialogflow does not impose any validation on this value. It is specific to the client that reads it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dialogflow_cx_flow#audio_uri DialogflowCxFlow#audio_uri}
	AudioUri *string `field:"required" json:"audioUri" yaml:"audioUri"`
}


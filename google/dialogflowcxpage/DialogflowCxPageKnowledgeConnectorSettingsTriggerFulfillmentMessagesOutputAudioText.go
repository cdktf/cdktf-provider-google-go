// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText struct {
	// The SSML text to be synthesized.
	//
	// For more information, see SSML.
	// This field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_page#ssml DialogflowCxPage#ssml}
	Ssml *string `field:"optional" json:"ssml" yaml:"ssml"`
	// The raw text to be synthesized.
	//
	// This field is part of a union field 'source': Only one of 'text' or 'ssml' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_cx_page#text DialogflowCxPage#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}


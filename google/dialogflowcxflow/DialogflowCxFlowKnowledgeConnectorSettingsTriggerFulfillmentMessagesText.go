// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText struct {
	// A collection of text response variants.
	//
	// If multiple variants are defined, only one text response variant is returned at runtime.
	// required: true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dialogflow_cx_flow#text DialogflowCxFlow#text}
	Text *[]*string `field:"optional" json:"text" yaml:"text"`
}


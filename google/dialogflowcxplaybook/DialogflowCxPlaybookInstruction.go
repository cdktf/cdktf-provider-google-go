// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxplaybook


type DialogflowCxPlaybookInstruction struct {
	// General guidelines for the playbook.
	//
	// These are unstructured instructions that are not directly part of the goal, e.g. "Always be polite". It's valid for this text to be long and used instead of steps altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_playbook#guidelines DialogflowCxPlaybook#guidelines}
	Guidelines *string `field:"optional" json:"guidelines" yaml:"guidelines"`
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_playbook#steps DialogflowCxPlaybook#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
}


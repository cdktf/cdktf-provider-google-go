// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxplaybook


type DialogflowCxPlaybookInstructionSteps struct {
	// Sub-processing needed to execute the current step.
	//
	// This field uses JSON data as a string. The value provided must be a valid JSON representation documented in [Step](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.playbooks#step).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_playbook#steps DialogflowCxPlaybook#steps}
	Steps *string `field:"optional" json:"steps" yaml:"steps"`
	// Step instruction in text format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dialogflow_cx_playbook#text DialogflowCxPlaybook#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}


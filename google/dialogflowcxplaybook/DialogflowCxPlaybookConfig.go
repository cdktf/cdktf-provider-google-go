// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxplaybook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxPlaybookConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The human-readable name of the playbook, unique within an agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#display_name DialogflowCxPlaybook#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// High level description of the goal the playbook intend to accomplish.
	//
	// A goal should be concise since it's visible to other playbooks that may reference this playbook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#goal DialogflowCxPlaybook#goal}
	Goal *string `field:"required" json:"goal" yaml:"goal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#id DialogflowCxPlaybook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instruction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#instruction DialogflowCxPlaybook#instruction}
	Instruction *DialogflowCxPlaybookInstruction `field:"optional" json:"instruction" yaml:"instruction"`
	// llm_model_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#llm_model_settings DialogflowCxPlaybook#llm_model_settings}
	LlmModelSettings *DialogflowCxPlaybookLlmModelSettings `field:"optional" json:"llmModelSettings" yaml:"llmModelSettings"`
	// The agent to create a Playbook for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#parent DialogflowCxPlaybook#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Type of the playbook. Possible values: ["PLAYBOOK_TYPE_UNSPECIFIED", "TASK", "ROUTINE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#playbook_type DialogflowCxPlaybook#playbook_type}
	PlaybookType *string `field:"optional" json:"playbookType" yaml:"playbookType"`
	// The resource name of tools referenced by the current playbook in the instructions.
	//
	// If not provided explicitly, they are will be implied using the tool being referenced in goal and steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#referenced_tools DialogflowCxPlaybook#referenced_tools}
	ReferencedTools *[]*string `field:"optional" json:"referencedTools" yaml:"referencedTools"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dialogflow_cx_playbook#timeouts DialogflowCxPlaybook#timeouts}
	Timeouts *DialogflowCxPlaybookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


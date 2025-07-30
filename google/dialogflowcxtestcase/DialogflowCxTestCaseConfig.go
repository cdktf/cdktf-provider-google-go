// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxTestCaseConfig struct {
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
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#display_name DialogflowCxTestCase#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#id DialogflowCxTestCase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Additional freeform notes about the test case. Limit of 400 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#notes DialogflowCxTestCase#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#parent DialogflowCxTestCase#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes.
	//
	// Each tag should start with "#" and has a limit of 30 characters
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#tags DialogflowCxTestCase#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// test_case_conversation_turns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#test_case_conversation_turns DialogflowCxTestCase#test_case_conversation_turns}
	TestCaseConversationTurns interface{} `field:"optional" json:"testCaseConversationTurns" yaml:"testCaseConversationTurns"`
	// test_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#test_config DialogflowCxTestCase#test_config}
	TestConfig *DialogflowCxTestCaseTestConfig `field:"optional" json:"testConfig" yaml:"testConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_test_case#timeouts DialogflowCxTestCase#timeouts}
	Timeouts *DialogflowCxTestCaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


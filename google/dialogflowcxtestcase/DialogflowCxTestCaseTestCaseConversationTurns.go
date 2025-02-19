// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurns struct {
	// user_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/dialogflow_cx_test_case#user_input DialogflowCxTestCase#user_input}
	UserInput *DialogflowCxTestCaseTestCaseConversationTurnsUserInput `field:"optional" json:"userInput" yaml:"userInput"`
	// virtual_agent_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/dialogflow_cx_test_case#virtual_agent_output DialogflowCxTestCase#virtual_agent_output}
	VirtualAgentOutput *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput `field:"optional" json:"virtualAgentOutput" yaml:"virtualAgentOutput"`
}


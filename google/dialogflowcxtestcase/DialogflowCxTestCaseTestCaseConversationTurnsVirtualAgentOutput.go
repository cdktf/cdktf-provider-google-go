// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput struct {
	// current_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/dialogflow_cx_test_case#current_page DialogflowCxTestCase#current_page}
	CurrentPage *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage `field:"optional" json:"currentPage" yaml:"currentPage"`
	// The session parameters available to the bot at this point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/dialogflow_cx_test_case#session_parameters DialogflowCxTestCase#session_parameters}
	SessionParameters *string `field:"optional" json:"sessionParameters" yaml:"sessionParameters"`
	// text_responses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/dialogflow_cx_test_case#text_responses DialogflowCxTestCase#text_responses}
	TextResponses interface{} `field:"optional" json:"textResponses" yaml:"textResponses"`
	// triggered_intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/dialogflow_cx_test_case#triggered_intent DialogflowCxTestCase#triggered_intent}
	TriggeredIntent *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent `field:"optional" json:"triggeredIntent" yaml:"triggeredIntent"`
}


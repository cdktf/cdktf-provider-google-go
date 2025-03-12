// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText struct {
	// The natural language text to be processed. Text length must not exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/dialogflow_cx_test_case#text DialogflowCxTestCase#text}
	Text *string `field:"required" json:"text" yaml:"text"`
}


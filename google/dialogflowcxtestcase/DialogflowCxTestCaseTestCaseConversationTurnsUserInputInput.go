// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput struct {
	// dtmf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/dialogflow_cx_test_case#dtmf DialogflowCxTestCase#dtmf}
	Dtmf *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf `field:"optional" json:"dtmf" yaml:"dtmf"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/dialogflow_cx_test_case#event DialogflowCxTestCase#event}
	Event *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent `field:"optional" json:"event" yaml:"event"`
	// The language of the input.
	//
	// See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes.
	// Note that queries in the same session do not necessarily need to specify the same language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/dialogflow_cx_test_case#language_code DialogflowCxTestCase#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/dialogflow_cx_test_case#text DialogflowCxTestCase#text}
	Text *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText `field:"optional" json:"text" yaml:"text"`
}


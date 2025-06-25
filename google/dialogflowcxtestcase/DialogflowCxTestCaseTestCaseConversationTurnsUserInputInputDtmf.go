// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf struct {
	// The dtmf digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dialogflow_cx_test_case#digits DialogflowCxTestCase#digits}
	Digits *string `field:"optional" json:"digits" yaml:"digits"`
	// The finish digit (if any).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/dialogflow_cx_test_case#finish_digit DialogflowCxTestCase#finish_digit}
	FinishDigit *string `field:"optional" json:"finishDigit" yaml:"finishDigit"`
}

